# DKGCLI

DKGCLI is a CLI tool for using the DKG protocol. Here is a complete scenario:

## Dkg
```sh
# Install the CLI
go install .

# Run 3 nodes. Do that in 3 different sessions
LLVL=info dkgcli --config /tmp/node1 start --routing tree --listen tcp://127.0.0.1:2001
LLVL=info dkgcli --config /tmp/node2 start --routing tree --listen tcp://127.0.0.1:2002
LLVL=info dkgcli --config /tmp/node3 start --routing tree --listen tcp://127.0.0.1:2003

# Exchange certificates
dkgcli --config /tmp/node2 minogrpc join --address //127.0.0.1:2001 $(dkgcli --config /tmp/node1 minogrpc token)
dkgcli --config /tmp/node3 minogrpc join --address //127.0.0.1:2001 $(dkgcli --config /tmp/node1 minogrpc token)

# Initialize DKG on each node. Do that in a 4th session.
dkgcli --config /tmp/node1 dkg listen
dkgcli --config /tmp/node2 dkg listen
dkgcli --config /tmp/node3 dkg listen

# Do the setup in one of the node:
dkgcli --config /tmp/node1 dkg setup \
    --authority $(cat /tmp/node1/dkgauthority) \
    --authority $(cat /tmp/node2/dkgauthority) \
    --authority $(cat /tmp/node3/dkgauthority) --threshold 3

# Encrypt a message:
dkgcli --config /tmp/node2 dkg encrypt --message deadbeef

# Decrypt a message
dkgcli --config /tmp/node3 dkg decrypt --encrypted <...>
```

## verifiable encryption and decryption (F3B)
```sh
# Encrypt a message with verifiable encryption function 
# Receive the ciphertext as well as the encryption proofs
# GBar is the second generator of the group.
dkgcli --config /tmp/node2 dkg verifiableEncrypt --message deadbeef --GBar 1d0194fdc2fa2ffcc041d3ff12045b73c86e4ff95ff662a5eee82abdf44a53c7

# Verify the encryption proof and decrypt the ciphertext
# Ciphertext should be in the form of <hex(K)>:<hex(C)>:<hex(Ubar)>:<hex(E)>:<hex(F)>
# In the case of batch decryption you can append as many as ciphertexts you want using ":" as the separator 
# GBar should be the same that we used for encryption
dkgcli --config /tmp/node3 dkg verifiableDecrypt --ciphertexts <...> --GBar 1d0194fdc2fa2ffcc041d3ff12045b73c86e4ff95ff662a5eee82abdf44a53c7
```

## Resharing
```sh
# Adding a new node
LLVL=info dkgcli --config /tmp/node4 start --routing tree --listen tcp://127.0.0.1:2004

# Exchange certificates
dkgcli --config /tmp/node4 minogrpc join --address //127.0.0.1:2001 $(dkgcli --config /tmp/node1 minogrpc token)

# Initialize DKG on the new node. Do that in a 4th session.
dkgcli --config /tmp/node4 dkg listen

# Reshare the secret from node1-node2-node3 committee to node1-node2-node4 committee
dkgcli --config /tmp/node1 dkg reshare \
    --authority $(cat /tmp/node1/dkgauthority) \
    --authority $(cat /tmp/node2/dkgauthority) \
    --authority $(cat /tmp/node4/dkgauthority) --thresholdNew 3


# You should be able to decrypt the same ciphertext with the new committee
dkgcli --config /tmp/node4 dkg verifiableDecrypt --ciphertexts <...> --GBar 1d0194fdc2fa2ffcc041d3ff12045b73c86e4ff95ff662a5eee82abdf44a53c7
```
## Docker image
```sh
# Now for creating the docker image of the dkgcli you should do the following:
# These commands would build a docker image of the package and write it in docker file dela/dkg:latest
cd ../../..
docker build -t dela/dkg:latest -f dkg/pedersen/dkgcli/dockerfile .
docker run --rm -e LLVL=info dela/dkg

# Then you need to change the tag of your docker image in order to push it to your docker hub repository
docker tag dela/dkg:latest YOUR_DOCKERHUB_NAME/firstimage

# Push the docker image to docker hub
docker push YOUR_DOCKERHUB_NAME/firstimage
