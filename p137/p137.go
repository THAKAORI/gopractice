package main		// mainパッケージであることを宣言

import (
	"fmt"
)

func main() {
	var wg sync.WaitGroup
	done := make(chan interface{})
	defer close(done)

	wg.Add(1)
	go func ()  {
		defer wg.Done()
		if err := printGreeting(done); err != nil {
			fmt.Printf("%v", err)
			return
		}
	}()

	wg.Add(1)
	go func ()  {
		defer wg.Done()
		if err := printFarewell(done); err != nil {
			fmt.Printf("%v", err)
			return
		}
	}()

	wg.Wait()
}

func printGreeting(done <-chan interface{}) error {
	greeting, err := genGreeting(done)
	if err != nil {
		return nil
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}

func printFarewell(done <-chan interface{}) error {
	farewell, err := genFarewell(done)
	if err != nil {
		return nil
	}
	
}