package t

import (
	"testing"
)

func TestT(testing *testing.T) {
	T(true == true)
}

func TestF(*testing.T) {
	F(false == true)
}

func TestEQ(*testing.T) {
	EQ(true, true)
	EQ("hi", "hi")
	EQ(1,1)
	EQ(3.4, 3.4)
}

func TestNEQ(*testing.T) {
	NEQ(false, true)
	NEQ("ih", "hi")
	NEQ(1,2)
	NEQ(3.4, 44)
}

