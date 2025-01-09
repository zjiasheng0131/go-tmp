package main

import (
	"fmt"
	"regexp"

	"github.com/james-0131-2/gocustom2"
	"github.com/james-0131-2/jamesgo1"
	"github.com/james-0131-2/agpl3"
)

func main() {
	re := regexp.MustCompile(".com")
	fmt.Println(re.FindString("(cainiaojc4com)"))
	fmt.Println(re.FindString("abc.org"))
	fmt.Println(re.FindString("fb.com"))

	re1 := regexp.MustCompile("secret=123")
	fmt.Println(re1.FindString("const secret = 123"))
	a1, e1 := gocustom2.ShanghaiToUTC("01:00")
	fmt.Println("gocus", a1, e1)
	a2, e2 := jamesgo1.ShanghaiToUTC("01:00")
	fmt.Println("gocus", a2, e2)
        a3, e3 := agpl3.ShanghaiToUTC("01:00")
        fmt.Println("gocus", a3,e3)
}
