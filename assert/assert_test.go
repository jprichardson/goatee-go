package assert

import "testing"

func TestTrue(testing *testing.T) {
	True(true == true)
}

func TestFalse(*testing.T) {
	False(false == true)
}

func TestEqual(*testing.T) {
	Equal(true, true)
	Equal("hi", "hi")
	Equal(1,1)
	Equal(3.4, 3.4)
}

func TestNotEqual(*testing.T) {
	NotEqual(false, true)
	NotEqual("ih", "hi")
	NotEqual(1,2)
	NotEqual(3.4, 44)
}

