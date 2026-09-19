Errors aur testing

- Error ek value hai: `if err != nil { return err }`.
- Context add karo: `fmt.Errorf("op: %w", err)`.
- `errors.Is` sentinel ke liye, `errors.As` custom type ke liye.
- panic sirf programmer bug ke liye. recover sirf defer ke andar kaam karta hai.
- Tests: file `x_test.go`, function `TestXxx(t *testing.T)`. Table-driven tests use karo.
- Benchmark: `BenchmarkXxx(b *testing.B)`, loop `b.N` tak.
- HTTP handler test: `httptest.NewRecorder()`, server chalane ki zaroorat nahi.

Code: errorsdemo/, lru/lru_test.go, httpserver/server_test.go
