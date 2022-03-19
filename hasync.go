package hasync

const F_OK = 1
const F_ERR = 2

type Future[T any] struct {
	defVal T
	result T
	c      chan int
	closed bool
	erorr  string
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
	f.erorr = s
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

func (f *Future[T]) TryGet() (T, bool) {
	if f.erorr != "" {
		return f.defVal, false
	}
	if !f.closed {
		return f.defVal, false
	}
	return f.result, true
}

func (f *Future[T]) Get() (T, bool) {
	if f.closed {
		return f.result, true
	}
	res, ok := <-f.c
	if !ok {
		return f.defVal, false
	}
	if res == F_OK {
		return f.result, true
	}
	return f.defVal, false
}

func Await[T any](defVal T, f func(*Future[T])) *Future[T] {
	future := NewFuture(defVal)
	go f(future)
	return future
}

func (f *Future[T]) GetError() string {
	return f.erorr
}
