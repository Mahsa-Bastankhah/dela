# Controller

For adding a new command first you should add the command to the `SetCommands` function of the `mod.go` file.
In the following you see an example of how to add a new command.
```go
sub = cmd.SetSubCommand("encrypt")
	sub.SetDescription("encrypt a message. Outputs <hex(K)>:<hex(C)>:<hex(remainder)>")
	sub.SetFlags(
		cli.StringFlag{
			Name:  "message",
			Usage: "the message to encrypt, encoded in hex",
		},
	)
	sub.SetAction(builder.MakeAction(encryptAction{}))
```

For specifying the action of each command you should add a new struct and new function to the `action.go` file.
```go
type encryptAction struct{}
func (a encryptAction) Execute(ctx node.Context) error {

}
```

In the function first you should get the actor.
```go
var actor dkg.Actor

	err := ctx.Injector.Resolve(&actor)
	if err != nil {
		return xerrors.Errorf("failed to resolve actor, did you call listen?: %v", err)
	}
```

Then convert the flags to bytes.
```go
message, err := hex.DecodeString(ctx.Flags.String("message"))
	if err != nil {
		return xerrors.Errorf("failed to decode message: %v", err)
	}
```

Then call the corresponding `dkg` function and handle the error if any.
```go
k, c, remainder, err := actor.Encrypt(message)
	if err != nil {
		return xerrors.Errorf("failed to encrypt: %v", err)
	}
```

And finally convert the results to bytes and write them.
```go
kbuff, err := k.MarshalBinary()
if err != nil {
	return "", xerrors.Errorf("failed to marshal k: %v", err)
}

cbuff, err := c.MarshalBinary()
if err != nil {
	return "", xerrors.Errorf("failed to marshal c: %v", err)
}

encoded := hex.EncodeToString(kbuff) + separator +
	hex.EncodeToString(cbuff) + separator +
	hex.EncodeToString(remainder)

fmt.Fprint(ctx.Out, encoded)
```




