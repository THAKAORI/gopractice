package main		// mainパッケージであることを宣言

import (
	"fmt";
)

func main()  {
	var data int
	go func ()  {
		data++
	}()
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}

