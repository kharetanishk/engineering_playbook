# Go

Runnable Go examples with Hinglish comments. Each folder is one topic.

| Folder | Topic |
|---|---|
| `basics/` | variables, loops, slices, maps, defer |
| `structs/` | structs, methods, interfaces, type switch |
| `errorsdemo/` | error wrapping, `errors.Is/As`, panic/recover |
| `goroutines/` | goroutines, channels, select, timeout |
| `workerpool/` | worker pool pattern |
| `mutex/` | Mutex, atomic, race detector |
| `ctxdemo/` | context timeout / cancel |
| `generics/` | type params, constraints, generic Stack |
| `lru/` | LRU cache + table-driven test + benchmark |
| `httpserver/` | net/http JSON API + httptest |
| `cmd/server/` | runs the http server |

## Commands

```
go run ./basics          # koi bhi ek topic chalao
go test -race ./...      # sab tests, data race check ke saath
go test -bench=. ./lru   # benchmark
go vet ./...             # static check
gofmt -l .               # formatting check
go run ./cmd/server      # server :8080
```
