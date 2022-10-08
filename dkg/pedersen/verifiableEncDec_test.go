package pedersen

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"go.dedis.ch/dela/dkg"
	"go.dedis.ch/dela/dkg/pedersen/types"
	"go.dedis.ch/dela/mino"

	"go.dedis.ch/dela/mino/minoch"
	"go.dedis.ch/dela/mino/minogrpc"
	"go.dedis.ch/dela/mino/router/tree"

	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/xof/keccak"
)

func init() {
	rand.Seed(0)
}

func Test_verifiableEncDec_minoch(t *testing.T) {
	// setup up the dkg
	n := 18
	threshold := n
	batchSize := 10

	minos := make([]mino.Mino, n)
	dkgs := make([]dkg.DKG, n)
	addrs := make([]mino.Address, n)

	// Create GBar. We need a generator in order to follow the encryption and
	// decryption protocol of https://arxiv.org/pdf/2205.08529.pdf / we take an
	// agreed data among the participants and embed it as a point. The result is
	// the generator that we are seeking.
	agreedData := make([]byte, 32)
	_, err := rand.Read(agreedData)
	require.NoError(t, err)
	GBar := suite.Point().Embed(agreedData, keccak.New(agreedData))

	gBarbuff, err := GBar.MarshalBinary()
	require.NoError(t, err)
	fmt.Println(hex.EncodeToString(gBarbuff))

	minoManager := minoch.NewManager()

	for i := 0; i < n; i++ {
		minoch := minoch.MustCreate(minoManager, fmt.Sprintf("addr %d", i))
		minos[i] = minoch
		addrs[i] = minoch.GetAddress()
	}

	pubkeys := make([]kyber.Point, len(minos))

	for i, mino := range minos {
		dkg, pubkey := NewPedersen(mino)
		dkgs[i] = dkg
		pubkeys[i] = pubkey
	}

	fakeAuthority := NewAuthority(addrs, pubkeys)

	actors := make([]dkg.Actor, n)
	for i := 0; i < n; i++ {
		actor, err := dkgs[i].Listen()
		require.NoError(t, err)
		actors[i] = actor
	}

	t.Log("setting up the dkg ...")

	_, err = actors[0].Setup(fakeAuthority, threshold)
	require.NoError(t, err)

	t.Log("generating the message and encrypting it ...")

	// generating random messages in batch and encrypt them
	keys := make([][29]byte, batchSize)

	var ciphertexts []types.Ciphertext
	for i := 0; i < batchSize; i++ {
		_, err = rand.Read(keys[i][:])
		require.NoError(t, err)

		ciphertext, remainder, err := actors[0].VerifiableEncrypt(keys[i][:], GBar)
		require.NoError(t, err)
		require.Len(t, remainder, 0)

		ciphertexts = append(ciphertexts, ciphertext)
	}

	// decrypting the batch ciphertext message
	fmt.Println("decrypting the ciphertext ...")
	decrypted, _, _, err := actors[0].VerifiableDecrypt(ciphertexts)
	require.NoError(t, err)
	for i := 0; i < batchSize; i++ {
		require.Equal(t, keys[i][:], decrypted[i])
	}
}

func Test_verifiableEncDec_minogrpc(t *testing.T) {
	// we want to time the decryption for different batch sizes with different
	// number of nodes

	//numWorkersSlice := []int{8}
	batchSizeSlice := []int{32}

	// setting up the dkg
	n := 128
	threshold := n

	minos := make([]mino.Mino, n)
	dkgs := make([]dkg.DKG, n)
	addrs := make([]mino.Address, n)

	// Create GBar. We need a generator in order to follow the encryption and
	// decryption protocol of https://arxiv.org/pdf/2205.08529.pdf. We take an
	// agreed data among the participants and embed it as a point. The result is
	// the generator that we are seeking.
	agreedData := make([]byte, 32)
	_, err := rand.Read(agreedData)
	require.NoError(t, err)
	GBar := suite.Point().Embed(agreedData, keccak.New(agreedData))

	t.Log("initiating the dkg nodes ...")

	for i := 0; i < n; i++ {
		addr := minogrpc.ParseAddress("127.0.0.1", 0)

		minogrpc, err := minogrpc.NewMinogrpc(addr, nil, tree.NewRouter(minogrpc.NewAddressFactory()))
		require.NoError(t, err)

		defer minogrpc.GracefulStop()

		minos[i] = minogrpc
		addrs[i] = minogrpc.GetAddress()
	}

	pubkeys := make([]kyber.Point, len(minos))

	for i, mino := range minos {
		for _, m := range minos {
			mino.(*minogrpc.Minogrpc).GetCertificateStore().Store(m.GetAddress(),
				m.(*minogrpc.Minogrpc).GetCertificateChain())
		}
		dkg, pubkey := NewPedersen(mino.(*minogrpc.Minogrpc))
		dkgs[i] = dkg
		pubkeys[i] = pubkey
	}

	fakeAuthority := NewAuthority(addrs, pubkeys)

	actors := make([]dkg.Actor, n)
	for i := 0; i < n; i++ {
		actor, err := dkgs[i].Listen()
		require.NoError(t, err)
		actors[i] = actor
	}

	t.Log("setting up the dkg ...")

	start := time.Now()
	_, err = actors[0].Setup(fakeAuthority, threshold)
	require.NoError(t, err)
	setupTime := time.Since(start)

	// generating random messages in batch and encrypt them
	for _, batchSize := range batchSizeSlice {
		t.Logf("=== starting the process with batch size = %d === \n", batchSize)

		//workerNum = numWorkersSlice[i]

		keys := make([][29]byte, batchSize)
		var ciphertexts []types.Ciphertext
		for i := 0; i < batchSize; i++ {
			_, err = rand.Read(keys[i][:])
			require.NoError(t, err)

			ciphertext, remainder, err := actors[0].VerifiableEncrypt(keys[i][:], GBar)
			require.NoError(t, err)
			require.Len(t, remainder, 0)

			ciphertexts = append(ciphertexts, ciphertext)
		}

		t.Log("decrypting the batch ...")

		start = time.Now()
		decrypted, _, _, err := actors[0].VerifiableDecrypt(ciphertexts)
		decryptionTime := time.Since(start)
		require.NoError(t, err)

		for i := 0; i < batchSize; i++ {
			require.Equal(t, keys[i][:], decrypted[i])
		}

		// initiating the log file for writing the delay and throughput data
		f, err := os.OpenFile("../logs/test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		require.NoError(t, err)
		defer f.Close()
		wrt := io.MultiWriter(os.Stdout, f)
		log.SetOutput(wrt)

		log.Printf("n = %d , batchSize = %d  ,decryption time = %v s, throughput =  %v tx/s , dkg setup time = %v s",
			n, batchSize, decryptionTime.Seconds(), float32(batchSize)/float32(decryptionTime.Seconds()), float32(setupTime.Seconds()))

		t.Logf("n=%d, batchSize=%d, workerNum=%d, decryption time=%s, "+
			"throughput=%v[tx/s], dkg setup time=%s", n, batchSize, workerNum,
			decryptionTime, float64(batchSize)/decryptionTime.Seconds(), setupTime)
	}
}
