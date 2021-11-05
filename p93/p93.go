package main		// mainパッケージであることを宣言

import (
	"fmt";
)



func main() {		// 最初に実行されるmain()関数を定義
	doWork := func (strings <- chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func ()  {
			defer fmt.Println("doWork exited")
			defer close(completed)
			for s := range strings {
				fmt.Println(s)
			}
		}()
		return completed
	}

	completed := doWork(nil)

	<-completed

	fmt.Println("Done.")
}