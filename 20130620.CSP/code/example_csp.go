package main

import ( "fmt";  "time"; "math/rand")

var ch = make(chan int, 4) // HL
func slowIO() int {
  for {
    time.Sleep(200 * time.Millisecond)
    ch <- rand.Int() % 100 // HL
  }
}

func stupidLoop() string { for i := 0; i < 4e8; i++ { i += 1 }; return "Hey" }

func main() {
  start := time.Now()
  go slowIO()
  for i := 0; i < 5; i++ {
    fmt.Printf("%s: %d, %d\n", stupidLoop(), <-ch, <-ch) // HL
  }
  fmt.Println(time.Since(start))
}
