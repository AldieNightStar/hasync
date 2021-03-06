package hasync

import (
	"errors"
)

const F_OK = 1
const F_ERR = 2

const ERR_NOTCLOSED = "Not closed yet!"
const ERR_CHANERR = "Channel error!"

type Future[T any] struct {
	defVal T
	result T
	c      chan int
	closed bool
	error  string
}

func NewFuture[T any](defVal T) *Future[T] {
	return &Future[T]{defVal, defVal, make(chan int, 1), false, ""}
}

func (f *Future[T]) Error(s string) bool {
	if f.closed {
		return false
	}
	f.c <- F_ERR
	close(f.c)
	f.closed = true
	f.error = s
	return true
}

func (f *Future[T]) Ok(result T) bool {
	if f.closed {
		return false
	}
	f.result = result
	f.closed = true
	f.c <- F_OK
	close(f.c)
	return true
}

func (f *Future[T]) TryGet() (T, error) {
	if f.error != "" {
		return f.defVal, errors.New(f.error)
	}
	if !f.closed {
		return f.defVal, errors.New(ERR_NOTCLOSED)
	}
	return f.result, nil
}

func (f *Future[T]) Get() (T, error) {
	if f.closed {
		return f.result, errors.New(ERR_CHANERR)
	}
	res, ok := <-f.c
	if !ok {
		return f.defVal, errors.New(ERR_CHANERR)
	}
	if res == F_OK {
		return f.result, nil
	}
	return f.defVal, errors.New(f.error)
}

func Await[T any](defVal T, f func(*Future[T])) FutureResult[T] {
	future := NewFuture(defVal)
	go f(future)
	return future
}

type FutureResult[T any] interface {
	Get() (T, error)
	TryGet() (T, error)
}
