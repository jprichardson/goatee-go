package t

import (
	"fmt"
	"os"
	"runtime/debug"
)

//screw Google and their encouragement of verbosity
func T (expr bool) {
	if !expr {
		fmt.Fprintf(os.Stderr, "NOT TRUE")
		debug.PrintStack()
		os.Exit(1)
	}
}

func F (expr bool) {
	if expr {
		fmt.Fprintf(os.Stderr, "TRUE")
		debug.PrintStack()
		os.Exit(1)
	}
}

func EQ (v1 interface{}, v2 interface{}) {
	if v1 != v2 {
		fmt.Fprintf(os.Stderr, "%v != %v\n", v1,  v2)
		debug.PrintStack()
		os.Exit(1)
	}
}

func NEQ (v1 interface{}, v2 interface{}) {
	if v1 == v2 {
		fmt.Fprintf(os.Stderr, "%v == %v\n", v1,  v2)
		debug.PrintStack()
		os.Exit(1)
	}
}