package main		// mainパッケージであることを宣言

import (
	"fmt";
	"context"
)

func main()  {
	ProcessRequest("jane", "abc123")
}

func ProcessRequest(userID, authToken string)  {
	ctx := context.WithValue(context.Background(), "userID", userID)
	ctx = context.WithValue(ctx, "authToken", authToken)
	HandleResponse(ctx)
}

func HandleResponse(ctx context.Context)  {
	fmt.Printf(
		"handling response for %v (%v)",
		ctx.Value("userID"),
		ctx.Value("authToken"),
	)
}