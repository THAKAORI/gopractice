package main		// mainパッケージであることを宣言

import "fmt"		// fmtモジュールをインポート
import "sync"

var wg sync.WaitGroup

func main(){
	for _, salutation := range []string{"hello", "greeting", "good day"}{
		wg.Add(1)
		go func(salutation string){
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}