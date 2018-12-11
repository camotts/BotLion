package core

import (
	"time"
)

func Retry5Times(f func() error, delay time.Duration) error {
	var e error
	for i := 0; i < 5; i++ {
		e = f()
		if e == nil {
			return nil
		}
	}
	return e
}
