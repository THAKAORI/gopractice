package main		// mainパッケージであることを宣言

import "fmt"		// fmtモジュールをインポート
import "sync"
// import "time"
// import "text/tabwriter"
// import "os"
// import "math"

var wg sync.WaitGroup

func main() {		// 最初に実行されるmain()関数を定義
	var count int

	increment := func ()  {
		count++
	}

	var once sync.Once

	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func ()  {
			defer increments.Done()
			once.Do(increment)
			// increment()
		}()
	}

	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}