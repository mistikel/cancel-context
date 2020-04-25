package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	talker(ctx, 5*time.Second, "hello")
}

func talker(ctx context.Context, timeout time.Duration, teks string) {
	select {
	case <-time.After(timeout):
		fmt.Println(teks)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}
}
