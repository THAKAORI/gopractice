package main		// mainパッケージであることを宣言

import (
	"fmt";
)



func main() {		// 最初に実行されるmain()関数を定義
	repeat := func (
		done <-chan interface{},
		values ...interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func ()  {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	take := func (
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func ()  {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <- valueStream:
				}
			}
		}()
		return takeStream
	}

	// toInt := func (
	// 	done <-chan interface{},
	// 	valueStream <-chan interface{},
	// ) <-chan int {
	// 	intStream := make(chan int)
	// 	go func ()  {
	// 		defer close(intStream)
	// 		for v := range valueStream {
	// 			select {
	// 			case <-done:
	// 				return
	// 			case intStream <- v.(int):
	// 			}
	// 		}
	// 	}()
	// 	return intStream
	// }

	// primeFinder := func(done <-chan interface{}, intStream <-chan int) <-chan interface{} {
	// 	primeStream := make(chan interface{})
	// 	go func() {
	// 		defer close(primeStream)
	// 		for integer := range intStream {
	// 			integer -= 1
	// 			prime := true
	// 			for divisor := integer - 1; divisor > 1; divisor-- {
	// 				if integer%divisor == 0 {
	// 					prime = false
	// 					break
	// 				}
	// 			}

	// 			if prime {
	// 				select {
	// 				case <-done:
	// 					return
	// 				case primeStream <- integer:
	// 				}
	// 			}
	// 		}
	// 	}()
	// 	return primeStream
	// }

	// fanIn := func (
	// 	done <-chan interface{},
	// 	channels ...<-chan interface{},
	// ) <-chan interface{} {
	// 	var wg sync.WaitGroup
	// 	multiplexedStream := make(chan interface{})

	// 	multiplex := func(c <-chan interface{}) {
	// 		defer wg.Done()
	// 		for i := range c {
	// 			select {
	// 			case <-done:
	// 				return
	// 			case multiplexedStream <- i:
	// 			}
	// 		}
	// 	}

	// 	wg.Add(len(channels))
	// 	for _, c := range channels {
	// 		go multiplex(c)
	// 	}

	// 	go func ()  {
	// 		wg.Wait()
	// 		close(multiplexedStream)
	// 	}()

	// 	return multiplexedStream
	// }

	orDone := func (done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func ()  {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if ok == false {
						return
					}
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}

	tee := func (
		done <-chan interface{},
		in <-chan interface{},
	) (_, _ <-chan interface{}) {
		out1 := make(chan interface{})
		out2 := make(chan interface{})
		go func ()  {
			defer close(out1)
			defer close(out2)
			for val := range orDone(done, in) {
				var out1, out2 = out1, out2
				for i := 0; i < 2; i++ {
					select {
					case out1 <- val:
						out1 = nil
					case out2 <- val:
						out2 = nil
					}
				}
			}
		}()
		return out1, out2
	}

	done := make(chan interface{})
	defer close(done)

	out1, out2 := tee(done, take(done, repeat(done, 1, 2), 4))

	for val1 := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
	}

}