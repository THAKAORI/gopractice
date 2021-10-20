package main		// mainパッケージであることを宣言

import "fmt"		// fmtモジュールをインポート
// import "sync"
// import "time"
// import "text/tabwriter"
// import "os"
// import "math"

func main() {		// 最初に実行されるmain()関数を定義
	stringStream := make(chan string)
	go func ()  {
		stringStream <- "Hello channels"
	}()
	fmt.Println(<-stringStream)
}