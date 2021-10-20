package main		// mainパッケージであることを宣言

import (
	"fmt";
	"sync";
	"bytes"
)



func main() {		// 最初に実行されるmain()関数を定義
	printData := func (wg *sync.WaitGroup, data []byte)  {
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])

	wg.Wait()
}