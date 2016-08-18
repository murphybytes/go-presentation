package main

import (
  "fmt"
  "crypto/rand"
  "sync"
)

var primes = []byte{
    2, 3, 5,     7,     11,     13,     17,     19,     23,     29,
31,     37,     41,     43,     47,     53,     59,     61,     67,     71,
73,     79,     83,     89,     97,    101,    103,    107,    109,    113,
127,    131,    137,    139,    149,    151,    157,    163,    167,    173,
179,    181,    191,    193,    197,    199,    211,    223,    227,    229,
233,    239,    241,    251,
}


// generates a bunch of numbers and feeds them to a channel
func produceSample(c int ) <-chan byte {
  out := make(chan byte)

  buffer := make([]byte, c )
  rand.Read(buffer)

  go func() {
    for _, b := range buffer {
      out <- b
    }
    close(out)
  }()

  return out
}
// Finds numbers that are prime
func producePrimes(primes []byte, in <-chan byte) <-chan byte {
  out := make(chan byte)
  go func(){
    var wg sync.WaitGroup
    wg.Add(10)
    for i := 0; i < 10; i++ {

      go func(wg *sync.WaitGroup) {
        defer wg.Done()
        for b := range in {
          for _, p := range primes {
            if b == p {
              out <- b
            }
          }
        }
      }(&wg)

    }
    wg.Wait()
    close(out)
  }()

  return out
}
// extracts last digit of number
func getLastDigits(in <-chan byte) <-chan byte {
  out := make(chan byte)
  go func(){
    for b := range in {
      out <- b % 10
    }
    close(out)
  }()

  return out

}


func main() {
  s := produceSample(100000)
  p := producePrimes(primes, s)
  d := getLastDigits(p)

  ones := 0
  twos := 0
  threes := 0
  fives := 0
  sevens := 0
  nines := 0
  for b := range d {
    switch b {
    case 1 :
      ones++
    case 2 :
      twos++
    case 3 :
      threes++
    case 5 :
      fives++
    case 7 :
      sevens++
    case 9:
      nines++
    }
  }

  fmt.Printf("Ones %d\n", ones)
  fmt.Printf("Twos %d\n", twos)
  fmt.Printf("Threes %d\n", threes)
  fmt.Printf("Fives %d\n", fives)
  fmt.Printf("Sevens %d\n", sevens)
  fmt.Printf("Nines %d\n", nines)
}
