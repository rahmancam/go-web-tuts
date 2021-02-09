package session

import "context"

type key int

const UserKey key = 0

func WithUser(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, UserKey, userID)
}

func GetUser(ctx context.Context) *int {
	userID, ok := ctx.Value(UserKey).(int)
	if !ok {
		return nil
	}
	return &userID
}
