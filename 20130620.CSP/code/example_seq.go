package main

import ( "fmt";  "time"; "math/rand")

func slowIO() int {
  time.Sleep(200 * time.Millisecond)
  return rand.Int() % 100
}

func stupidLoop() string { // Think of this as some Processor-intensive Task
  for i := 0; i < 4e8; i++ { i += 1 };
  return "Hey"
}

func main() {
  start := time.Now()
  for i := 0; i < 5; i++ {
    v1, v2 := slowIO(), slowIO()
    fmt.Printf("%s: %d, %d\n", stupidLoop(), v1, v2)
  }
  fmt.Println(time.Since(start))
}
