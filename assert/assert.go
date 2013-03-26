package assert

import "github.com/jprichardson/goatee-go"

//screw Google and their encouragement of verbosity
func True (expr bool) {
	t.T(expr)
}

func False (expr bool) {
	t.F(expr)
}

func Equals (v1 interface{}, v2 interface{}) {
	t.EQ(v1, v2)
}

func Equal (v1 interface{}, v2 interface{}) {
	Equals(v1, v2)
}

func NotEquals (v1 interface{}, v2 interface{}) {
	t.NEQ(v1, v2)
}

func NotEqual (v1 interface{}, v2 interface{}) {
	NotEquals(v1, v2)	
}