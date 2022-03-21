package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	now := int64(1642651327)
	// now := int64(1542651327)
	// now := 0
	fmt.Printf("flag = %v\n", IsRecentlyItem(now))
	// fmt.Println(time.Hour)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "aaa", "bbb")
	fmt.Printf("%v", ctx)
	fmt.Println(ctx.Value("aaa"))
}

const (
	LimitDays = 5.0
)

func IsRecentlyItem(createTime int64) bool {
	duration := LimitDays * 24 * time.Hour
	createAt := time.Unix(createTime, 0)
	leftBound := time.Now().Add(-duration)
	return createAt.After(leftBound)
}
