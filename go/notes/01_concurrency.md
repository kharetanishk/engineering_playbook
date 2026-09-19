Go concurrency cheat sheet

- goroutine: `go f()`, bahut halka, thousands chal sakte hain.
- channel: goroutines ke beech data. unbuffered = sync, buffered = queue. Sender close karta hai.
- select: multiple channels + `time.After` se timeout.
- WaitGroup: sabke khatam hone ka wait. `Add` goroutine start karne se PEHLE.
- Mutex: shared state ka lock, `defer Unlock()`. Simple counter ke liye `atomic`.
- context: cancel/timeout propagate karo, hamesha `defer cancel()`.
- Worker pool: fixed workers + jobs channel = concurrency limit.
- Rule: "Share memory by communicating", jahan ho sake channel use karo, warna mutex.
- Hamesha `go test -race` chalao.

Common interview gotchas
- Loop variable capture: parameter me pass karo (Go 1.22+ me per-iteration copy hai, phir bhi samajhna zaroori).
- Nil channel pe send/receive hamesha block karta hai.
- Closed channel pe send = panic, receive = zero value.
- Goroutine leak: koi goroutine channel pe atka reh gaya to kabhi khatam nahi hoga.

Code: goroutines/, workerpool/, mutex/, ctxdemo/
