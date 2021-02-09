package main

import (
	"context"
	"fmt"
)

type userInfo struct {
	Name string
	ID   int
}

type key string

// UserKey of the system
const UserKey key = "UserKey"

func getRestrictedData(ctx context.Context) (string, error) {
	usr, ok := ctx.Value(UserKey).(userInfo)
	if !ok {
		return "", fmt.Errorf("Not Logged In")
	}
	return fmt.Sprintf("'%s' got the restricted data", usr.Name), nil
}

func main() {

	usr := userInfo{
		"Abdul Rahman", 10001,
	}
	ctx := context.WithValue(context.Background(), UserKey, usr)
	// ctx := context.WithValue(context.Background(), "test", usr)
	data, err := getRestrictedData(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

}
