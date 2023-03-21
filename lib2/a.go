package lib2

import (
	"github.com/cuiweixie/ref1/lib1"
)

func Lib2() string {
	return "ref1" + lib1.Lib1()
}
