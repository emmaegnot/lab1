package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello from goroutine", i)
}

//func main() {
//	for i := 0; i < 20; i++ {
//		fmt.Println("Hello World")
//	}
//}

func main() {
	for i := 0; i < 5; i++ {
		go hello(i)
	}
	time.Sleep(1 * time.Second)
}

//why does sleep line solve the problem of not printing anything?
