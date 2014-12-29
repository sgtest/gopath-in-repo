package mypkg

import (
	"mycompany.example.com/foo"
	"othercompany.example.com/otherpkg"
)

func Foo() {
	foo.Foo()
	otherpkg.Other()
}
