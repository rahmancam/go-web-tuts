package main

import (
	"context"
	"fmt"

	"github.com/rahmancam/go-web-tuts/099_Level2/context/03_session_withValue/session"
)

func isUserLoggedIn(ctx context.Context) bool {
	userID := session.GetUser(ctx)
	fmt.Println(*userID)
	return userID != nil
}

func main() {
	ctx := session.WithUser(context.Background(), 10001)
	if ok := isUserLoggedIn(ctx); ok {
		fmt.Println("User Logged in")
	} else {
		fmt.Println("Not authorized")
	}
}
