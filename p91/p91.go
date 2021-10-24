package main		// mainパッケージであることを宣言

import (

)



func main() {		// 最初に実行されるmain()関数を定義
	for _, s := range []string{"a", "b", "c"} {
		select {
		case <-done:
			return
		case stringStream <- s:

		}
	}
}