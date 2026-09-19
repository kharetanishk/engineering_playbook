// context: cancel, timeout aur deadline ko goroutines tak pahunchane ke liye.
package main

import (
	"context"
	"fmt"
	"time"
)

func slowWork(ctx context.Context) error {
	select {
	case <-time.After(2 * time.Second): // kaam poora hone ka time
		return nil
	case <-ctx.Done(): // upar se cancel ya timeout aaya
		return ctx.Err()
	}
}

func main() {
	// 100ms ke baad apne aap cancel
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel() // hamesha call karo, warna resource leak

	err := slowWork(ctx)
	fmt.Println("result:", err) // context deadline exceeded

	// rule: ctx function ka pehla parameter hota hai, struct me store mat karo
}
