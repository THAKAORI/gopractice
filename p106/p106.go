package main		// mainパッケージであることを宣言

import (
	"fmt";
)



func main() {		// 最初に実行されるmain()関数を定義
	generator := func (done <-chan interface{}, intergers ...int) <-chan int {
		intStream := make(chan int, len(intergers))
		go func ()  {
			defer close(intStream)
			for _, i := range intergers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()
		return intStream
	}

	multiply := func (
		done <-chan interface{},
		intStream <-chan int,
		multiplier int,
	) <-chan int {
		multipliedStream := make(chan int)
		go func ()  {
			defer close(multipliedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- i*multiplier:
				}
			}
		}()
		return multipliedStream
	}

	add := func (
		done <-chan interface{},
		intStream <-chan int,
		additive int,
	) <-chan int {
		addedStream := make(chan int)
		go func ()  {
			defer close(addedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case addedStream <- i+additive:
				}
			}
		}()
		return addedStream
	}

	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

	for v := range pipeline {
		fmt.Println(v)
	}
}