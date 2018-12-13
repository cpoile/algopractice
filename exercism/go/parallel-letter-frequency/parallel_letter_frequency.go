// Package letter provides a parallelized letter frequency counter.
package letter

// ConcurrentFrequency uses go routines to parallelize the Frequency function.
func ConcurrentFrequency(strs []string) FreqMap {
	ch := make(chan FreqMap, len(strs))

	for _, s := range strs {
		go func(str string) {
			ch <- Frequency(str)
		}(s)
	}

	// merge the parallelized frequency maps
	ret := make(FreqMap)
	for range strs {
		m := <-ch
		for k, v := range m {
			ret[k] += v
		}
	}

	return ret
}
