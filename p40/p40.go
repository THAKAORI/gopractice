package main		// mainパッケージであることを宣言

import "fmt"		// fmtモジュールをインポート
import "sync"

var wg sync.WaitGroup

func sayHello()  {
	defer wg.Done()
	fmt.Println("hello")
}

func main(){
	wg.Add(1)
	go sayHello()
	wg.Wait()
}