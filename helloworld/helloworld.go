package main		// mainパッケージであることを宣言

import "fmt"		// fmtモジュールをインポート

func main() {		// 最初に実行されるmain()関数を定義
    a := []int{}
	for i := 0; i < 10; i++ {
		a = append(a, i)
		fmt.Println(len(a), cap(a))
	}

}