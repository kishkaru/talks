package main
import ( "fmt";  "time"; "math/rand"; "sync")
var buf []int; var bufMu = new(sync.Mutex) // HL
func slowIO() int {
  for { // HL
    time.Sleep(200 * time.Millisecond)
    bufMu.Lock(); buf = append(buf, rand.Int() % 100); bufMu.Unlock() // HL
  } // HL
}
func stupidLoop() string { for i := 0; i < 4e8; i++ { i += 1 }; return "Hey" }
func main() {
  start := time.Now()
  go slowIO() // Think of it as a light-weight pthread // HL
  for i := 0; i < 5; i++ {
    str := stupidLoop() // HL
    for {                               // HL
      bufMu.Lock(); if len(buf) >= 2 * (i + 1) { break } // break if data is ready // HL
      bufMu.Unlock(); time.Sleep(100 * time.Millisecond) // HL
    }                                   // HL
    fmt.Printf("%s: %d, %d\n", str, buf[2 * i], buf[2 * i + 1]); bufMu.Unlock() // HL
  }
  fmt.Println(time.Since(start))
}
