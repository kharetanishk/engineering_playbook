// Goroutines aur channels: Go ki asli power.
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// go keyword se function alag lightweight thread (goroutine) me chalta hai
	var wg sync.WaitGroup // WaitGroup = sabke khatam hone ka wait
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) { // id ko parameter me pass kiya taaki har goroutine ki apni copy ho
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
			fmt.Println("goroutine", id, "done")
		}(i)
	}
	wg.Wait()

	// channel = goroutines ke beech data bhejne ka pipe
	ch := make(chan int) // unbuffered: sender tab tak ruka rahega jab tak receiver na le
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch) // sender hi close karta hai, receiver ko pata chal jata hai ki ab kuch nahi aayega
	}()
	for v := range ch { // close hone tak padhta rahega
		fmt.Println("mila:", v)
	}

	// select: kai channels me se jo pehle ready ho. timeout ke liye bhi use hota hai
	slow := make(chan string)
	select {
	case m := <-slow:
		fmt.Println(m)
	case <-time.After(50 * time.Millisecond):
		fmt.Println("timeout ho gaya")
	}
}
