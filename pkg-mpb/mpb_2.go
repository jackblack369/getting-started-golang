package pkgmpb

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
)

func test2(t *testing.T) {

	var wg sync.WaitGroup
	// passed wg will be accounted at p.Wait() call
	p := mpb.New(mpb.WithWaitGroup(&wg))
	total, numBars := 100, 3
	wg.Add(numBars)

	for i := 0; i < numBars; i++ {
		name := fmt.Sprintf("Bar#%d:", i)
		bar := p.AddBar(int64(total),
			mpb.PrependDecorators(
				// simple name decorator
				decor.Name(name),
				// decor.DSyncWidth bit enables column width synchronization
				decor.Percentage(decor.WCSyncSpace),
			),
			mpb.AppendDecorators(
				// replace ETA decorator with "done" message, OnComplete event
				decor.OnComplete(
					// ETA decorator with ewma age of 30
					decor.EwmaETA(decor.ET_STYLE_GO, 30, decor.WCSyncWidth), "done",
				),
			),
		)
		// simulating some work
		go func() {
			defer wg.Done()
			rng := rand.New(rand.NewSource(time.Now().UnixNano()))
			max := 100 * time.Millisecond
			for i := 0; i < total; i++ {
				// start variable is solely for EWMA calculation
				// EWMA's unit of measure is an iteration's duration
				start := time.Now()
				time.Sleep(time.Duration(rng.Intn(10)+1) * max / 10)
				// we need to call EwmaIncrement to fulfill ewma decorator's contract
				bar.EwmaIncrement(time.Since(start))
			}
		}()
	}
	// wait for passed wg and for all bars to complete and flush
	p.Wait()
}
