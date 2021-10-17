package main		// mainパッケージであることを宣言

// import "fmt"		// fmtモジュールをインポート
import "sync"
// import "time"
// import "text/tabwriter"
// import "os"
// import "math"

var wg sync.WaitGroup

func main() {		// 最初に実行されるmain()関数を定義
	var onceA, onceB sync.Once
	var initB func()
	initA := func ()  { onceB.Do(initB) }
	initB = func ()  { onceA.Do(initA) }
	onceA.Do(initA)
}