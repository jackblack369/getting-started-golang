package main

import (
	"math"
	"runtime"
)

type Vector []float64

// Define the Op function as a method of the Vector type
func (v Vector) Op(x float64) float64 {
	// Example operation: square root of the element
	return math.Sqrt(x)
}

// Apply the operation to v[i], v[i+1] ... up to v[n-1].
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}
	c <- 1 // signal that this piece is done
}

var numCPU = runtime.NumCPU() // number of CPU cores

func (v Vector) DoAll(u Vector) {
	c := make(chan int, numCPU) // Buffering optional but sensible.
	for i := 0; i < numCPU; i++ {
		go v.DoSome(i*len(v)/numCPU, (i+1)*len(v)/numCPU, u, c)
	}
	// Drain the channel.
	for i := 0; i < numCPU; i++ {
		<-c // wait for one task to complete
	}
	// All done.
}
