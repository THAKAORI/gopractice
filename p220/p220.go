package main		// mainパッケージであることを宣言

import (
	"log";
	"os";
	"runtime/trace";
	"context";
	"time"
)

func main()  {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal("failed to create trace output file: %v", err)
	}

	defer func ()  {
		if err := f.Close(); err != nil {
			log.Fatal("failed to close trace output file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		panic(err)
	}
	defer trace.Stop()

	ctx := context.Background()
	ctx, task := trace.NewTask(ctx, "makeCoffee")
	defer task.End()
	trace.Log(ctx, "orderID", "1")

	coffee := make(chan bool)

	go func ()  {
		trace.WithRegion(ctx, "extractCoffee", extractCoffee)
		coffee <- true
	}()
	<-coffee
}

func extractCoffee()  {
	time.Sleep(1*time.MicroSecond)
	return
}