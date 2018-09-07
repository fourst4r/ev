# ev
Simple event handling in Go, without magic strings.

##Example
```go
func onMessage(args ev.Args) {
    println(args.String(0)) // cast arg 0 to string and print it!
}

var message ev.Ent
message.On(onMessage)
message.Invoke("Hello World!")
message.Off(onMessage) // we're done
```