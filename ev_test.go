package ev

import (
	"fmt"
	"testing"
)

func TestEnt(t *testing.T) {
	var e Ent

	f1 := func(args Args) {
		fmt.Print("f1")
		if args.String(0) != "invoked!" {
			t.Fail()
		}
	}
	e.On(f1)

	f2 := func(args Args) {
		fmt.Print("f2")
		if args.Int(1) != 1 {
			t.Fail()
		}
	}
	e.On(f2)

	f3 := func(args Args) {
		fmt.Print("f3 (once)")
		if args.String(0) != "invoked!" && args.Int(1) != 1 {
			t.Fail()
		}
	}

	e.Invoke("invoked!", 1) // f1 f2
	fmt.Println()
	e.Off(f2)
	e.Invoke("invoked!", 1) // f1
	fmt.Println()
	e.Off(f1)
	e.On(f2)
	e.Invoke("invoked!", 1) // f2
	fmt.Println()
	e.Once(f3)
	e.Invoke("invoked!", 1) // f2 f3
	fmt.Println()
	e.Invoke("invoked!", 1) // f2
	print("\r\ndone!\r\n")
	e.Off(f2)
	e.Invoke("invoked!", 1) // <nothing>

	// Output:
	// f1f2
	// f1
	// f2
	// f2f3 (once)
	// f
}
