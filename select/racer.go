package _select

import (
	"net/http"
	"time"
)

type RacerErr string

const ErrRequestTimedOut = RacerErr("request timed out")

const defaultResponseTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, defaultResponseTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", ErrRequestTimedOut
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func (e RacerErr) Error() string {
	return string(e)
}
