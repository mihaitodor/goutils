package noop

import "sync"

type Tracer struct{}

func (t *Tracer) Trace(name string, fn func(), wg *sync.WaitGroup) error {
	if wg == nil {
		fn()
	} else {
		wg.Add(1)

		go func() {
			defer wg.Done()

			fn()
		}()
	}

	return nil
}