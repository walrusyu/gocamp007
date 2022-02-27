package main

import (
	"fmt"
	"github.com/walrusyu/gocamp007/week5/counter"
	"time"
)

func main() {
	fmt.Printf("start at: %d\n", time.Now().Unix())
	c := counter.CreateCounter(10)
	for i := 0; i < 500; i++ {
		c.Add()
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Printf("stop at: %d\n", time.Now().Unix())
	fmt.Println(c.ToString())
}
