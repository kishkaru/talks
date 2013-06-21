package main
import ( "fmt";  "time"; "math/rand"; "sync" )

var buf []int; var bufMu = new(sync.Mutex);
var signal = make(chan string, 4) // Think a thread-safe queue that blocks on dequeue // HL
func slowIO() int {
  for {
    time.Sleep(200 * time.Millisecond)
    bufMu.Lock(); buf = append(buf, rand.Int() % 100); bufMu.Unlock(); signal <- "done" // HL
  }
}

func stupidLoop() string { for i := 0; i < 4e8; i++ { i += 1 }; return "Hey" }

func main() {
  start := time.Now()
  go slowIO()
  for i := 0; i < 5; i++ {
    str := stupidLoop()
    <-signal; <-signal // Wait for 2 signals // HL
    bufMu.Lock(); fmt.Printf("%s: %d, %d\n", str, buf[2*i], buf[2*i+1]); bufMu.Unlock() // HL
  }
  fmt.Println(time.Since(start))
}
