package main 

import (
  "fmt"
  "os"
  "strconv"
)

func main() {

  // Take user input
  if len(os.Args) < 2 {
    fmt.Println("Must provide a max integer")
    os.Exit(1)
  }

  // Argv is provided as a string, convert to integer
  s_max := os.Args[1]
  max, err := strconv.Atoi(s_max)
  if err != nil {
    fmt.Println("Must provide an integer")
    os.Exit(1)
  }

  // Initialize boolean slice
  sieve := make([]bool, max)

  // Begin the Sieve of Eratosthenes algorithm
  for i := 2; i < max; i++ {
    if !sieve[i] {
      // Current index isn't marked, mark all multiples of this index
      for multi := 2; i*multi < max; multi++ {
        sieve[i*multi] = true
      }
    }
  }

  for i := 2; i < max; i++ {
    if !sieve[i] {
      fmt.Printf("%v\n", i)
    }
  }
}