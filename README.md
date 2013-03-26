golang: goatee
=============

[![Build Status](https://travis-ci.org/jprichardson/goatee-go.png)](https://travis-ci.org/jprichardson/goatee-go)

Screw Google's verbose Go testing package. "go t", or "goatee" will help make your tests more concise and readable.


Why?
----

Because, this is too verbose:

```go
myvar := "hello"
if myvar != "hell" {
  t.Error("%s != hell\n", myvar)
}
```

more concise:

```go
myvar := "hello"
assert.True(myvar, "hell")

//even better:
t.T (myvar, "hell")
```



Installing
----------

    go get github.com/jprichardson/goatee-go


Usage
-----

**mypackage_test.go**:

```go
import (
    "github.com/jprichardson/goatee-go"
    "github.com/jprichardson/goatee-go/assert" //import the assert package if you like the assert style
    "testing" //optionally use Google's testing package
)

//traditionally, this is (t *testing.T), but you'll have a namespace conflict
func TestMyMethod (tu *testing.T) { 
    t.EQ (1, 1)
    t.EQ ("hi", "hi")
    t.NEQ ("hi", "bye")
    t.T (true == true)
    t.F (false == true)
}

func TestMyMethod2 (*testing.T) {
    assert.True(true == true)
    assert.False(true != false)
    assert.Equal(1, 1)
    assert.NotEqual("hi", "bye")
}

```



License (New BSD)
-------------
Copyright (c) 2013 JP Richardson

