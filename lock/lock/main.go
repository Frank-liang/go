package main

import (
	"fmt"
	"sync"
	"time"
)

type IncType struct {
	i     int64
	ILock sync.Mutex
}

var Invar IncType

func Inc(inc *IncType) {
	counter := 0
	for {
		counter += 1
		inc.ILock.Lock()
		inc.i += 1
		inc.ILock.Unlock()
		if counter > 10000000 {
			break
		}
	}
	fmt.Println("finish")
	return
}

func main() {
	go Inc(&Invar)
	go Inc(&Invar)
	go Inc(&Invar)

	time.Sleep(time.Second * 5)
	fmt.Println(Invar.i)
}
