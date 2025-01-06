package ti

import (
	"fmt"
	"sync"
	"time"
)

var g1 int

func Ti() {

	g1++
	fmt.Printf("结果：g = %d\n", g1)
}

var lck sync.Mutex

func Foo() {
	fmt.Println(11110)
	lck.Lock()

	fmt.Println(11111)
	time.Sleep(1 * time.Second)
	defer lck.Unlock()
	// ...
}
