package main		// mainパッケージであることを宣言

import "fmt"		// fmtモジュールをインポート
import "time"

func main() {		// 最初に実行されるmain()関数を定義
	stream := make(chan interface{})
    go func() {
        defer fmt.Println("close")
        defer close(stream)
        for i := 0; i < 3; i++ {
            stream <- i
			fmt.Println("range開始:",stream)
        }
    }()

	takeStream := make(chan interface{})
	defer close(takeStream)
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 1)
		takeStream <- stream
		fmt.Println("range開始:",takeStream)
	}
	time.Sleep(time.Second * 5)
}