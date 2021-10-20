package main		// mainパッケージであることを宣言

import (
	"fmt"
)



func main() {		// 最初に実行されるmain()関数を定義
	chanOwner := func () <-chan int {
		results := make(chan int, 5)
		go func ()  {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		} ()
		return results
	}
	
	consumer := func (results <-chan int)  {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner()
	consumer(results)
}