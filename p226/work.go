package work

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "Type=Foo"
import "github.com/cheekybits/genny/generic"

type Type generic.Type

func doWork(strings <-chan string) <-chan Type {
	completed := make(chan Type)
	go func() {
		defer fmt.Println("doWork exited.")
		defer close(completed)
		for s := range strings {
			fmt.Println(s)
		}
	}()
	return completed
}