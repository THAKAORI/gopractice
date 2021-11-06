package main		// mainパッケージであることを宣言

// import (
// 	"fmt";
// 	"context"
// )

func main()  {
	waitForever := make(chan interface{})
	go func ()  {
		panic("test panic")
	}()
	<-waitForever
}

