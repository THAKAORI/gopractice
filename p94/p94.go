package main		// mainパッケージであることを宣言

import (
	"fmt";
	"math/rand";
	"time"
)



func main() {		// 最初に実行されるmain()関数を定義
	newRandStream := func () <-chan int {
		randStream := make(chan int)
		go func ()  {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				randStream <- rand.Int()
			}
			fmt.Println("generated")
		}()

		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}

	time.Sleep(1 * time.Second)
}