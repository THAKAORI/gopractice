package main		// mainパッケージであることを宣言

import (
	"fmt";
	"time"
)



func main() {		// 最初に実行されるmain()関数を定義
	doWork := func (done <-chan interface{}, strings <- chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func ()  {
			defer fmt.Println("doWork exited")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func ()  {
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		time.Sleep(1 * time.Second)
		close(done)
	}()

	<-terminated
	fmt.Println("Done .")
}