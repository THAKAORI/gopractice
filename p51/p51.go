package main		// mainパッケージであることを宣言

import "fmt"		// fmtモジュールをインポート
import "sync"
import "time"
import "text/tabwriter"
import "os"
import "math"

var wg sync.WaitGroup

func main() {		// 最初に実行されるmain()関数を定義
	producer := func(wg *sync.WaitGroup, l sync.Locker){
		defer wg.Done()
		for i := 5; i > 0; i--{
			l.Lock()
			l.Unlock()
			time.Sleep(1)
		}
	}

	observer := func (wg *sync.WaitGroup, l sync.Locker)  {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func (count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count+1)
		beginTestTime := time.Now()
		go producer(&wg, mutex)
		for i := count; i > 0; i--{
			go observer(&wg, rwMutex)
		}

		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutex\tMutex\n")
	for i := 0; i < 20; i++{
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m),
		)
	}
}