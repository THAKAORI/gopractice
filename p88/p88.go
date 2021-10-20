package main		// mainパッケージであることを宣言

import (
	"fmt"
)

func main() {		// 最初に実行されるmain()関数を定義
	data := make([]int, 4)

	loopData := func (handleData chan<- int)  {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}