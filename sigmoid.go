package main

import "fmt"
import "math"

func main() {
  fmt.Println(sigmoid32(3.2))
}

func sigmoid64 (n float64) float64 {
	return 1/(1 + math.Exp(-n))
}
func sigmoid32 (n float32) float32 {
	return float32(1/(1 + math.Exp(-(float64(n)))))
}
