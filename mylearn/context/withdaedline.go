package main

import (
	"fmt"
	"context"
	"time"
)

func withdeadline() {
	d := time.Now().Add(5 * time.Second)

	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	for {
		select {
		case <- ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}

func main() {
	withdeadline()
}
