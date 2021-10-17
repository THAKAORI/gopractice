package main		// mainパッケージであることを宣言

import "fmt"		// fmtモジュールをインポート
import "sync"
import "time"

var wg sync.WaitGroup

func main() {		// 最初に実行されるmain()関数を定義
	wg.Add(1)
	go func(){
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
	}()

	wg.Add(1)
	go func ()  {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()

	wg.Wait()
	fmt.Println("All goroutines complete.")
}