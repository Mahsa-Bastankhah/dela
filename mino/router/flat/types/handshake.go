// The file contains the implementation of a handshake message that is sent to
// create the routing table.
//

package types

import (
	"go.dedis.ch/dela/mino"
	"go.dedis.ch/dela/mino/router"
	"go.dedis.ch/dela/serde"
	"go.dedis.ch/dela/serde/registry"
	"golang.org/x/xerrors"
)

var hsFormats = registry.NewSimpleRegistry()

// Handshake is a message to send the initial parameters of a routing table.
//
// - implements serde.Message
type Handshake struct {
	known []mino.Address
}

// NewHandshake returns a new handshake message.
func NewHandshake(known ...mino.Address) Handshake {
	return Handshake{
		known: known,
	}
}

// GetKnown returns the known addresses
func (h Handshake) GetKnown() []mino.Address {
	return h.known
}

// Serialize implements serde.Message. It returns the serialized data for the
// handshake.
func (h Handshake) Serialize(ctx serde.Context) ([]byte, error) {
	format := hsFormats.Get(ctx.GetFormat())

	data, err := format.Encode(ctx, h)
	if err != nil {
		return nil, xerrors.Errorf("encode: %v", err)
	}

	return data, nil
}

// HandshakeFactory is a factory to serialize and deserialize handshake
// messages.
//
// - implements router.HandshakeFactory
type HandshakeFactory struct {
	addrFac mino.AddressFactory
}

// NewHandshakeFactory creates a new factory.
func NewHandshakeFactory(addrFac mino.AddressFactory) HandshakeFactory {
	return HandshakeFactory{
		addrFac: addrFac,
	}
}

// Deserialize implements serde.Factory. It populates the handshake if
// appropriate, otherwise it returns an error.
func (fac HandshakeFactory) Deserialize(ctx serde.Context, data []byte) (serde.Message, error) {
	return fac.HandshakeOf(ctx, data)
}

// HandshakeOf implements router.HandshakeFactory. It populates the handshake if
// appropriate, otherwise it returns an error.
func (fac HandshakeFactory) HandshakeOf(ctx serde.Context, data []byte) (router.Handshake, error) {
	format := hsFormats.Get(ctx.GetFormat())

	ctx = serde.WithFactory(ctx, AddrKey{}, fac.addrFac)

	msg, err := format.Decode(ctx, data)
	if err != nil {
		return nil, xerrors.Errorf("decode: %v", err)
	}

	hs, ok := msg.(Handshake)
	if !ok {
		return nil, xerrors.Errorf("invalid handshake '%T'", msg)
	}

	return hs, nil
}
