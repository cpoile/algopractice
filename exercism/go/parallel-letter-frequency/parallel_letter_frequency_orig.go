// Package letter provides a parallelized letter frequency counter.
package letter

import (
	"runtime"
	"strings"
	"sync"
)

// ConcurrentFrequency uses go routines to parallelize the Frequency function.
// The input string is split into chunks equal to the number of CPUs on the host
// machine, and each chunk is processed by a go routine.
func ConcurrentFrequencyOrig(strs []string) FreqMap {
	s := strings.Join(strs, "")
	perCPU := len(s)/runtime.NumCPU() + 1
	ch := make(chan FreqMap)
	var wg sync.WaitGroup

	for i := 0; i < len(s); i += perCPU {
		wg.Add(1)
		start, end := i, i+perCPU
		if end > len(s) {
			end = len(s)
		}
		go func() {
			ch <- Frequency(s[start:end])
			wg.Done()
		}()
	}

	// closer
	go func() {
		wg.Wait()
		close(ch)
	}()

	// merge the parallelized frequency maps
	ret := make(FreqMap)
	for m := range ch {
		for k := range m {
			ret[k] += m[k]
		}
	}

	return ret
}
