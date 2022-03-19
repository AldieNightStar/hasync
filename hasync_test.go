package hasync

import (
	"testing"
	"time"
)

func TestHasync(t *testing.T) {
	res, _ := Await(0, func(f *Future[int]) {
		time.Sleep(5 * time.Millisecond)
		f.Ok(122)
		f.Ok(44)
	}).Get()
	if res != 122 {
		t.Fatal("Final result is bad!")
	}
}

func TestHasyncBad(t *testing.T) {
	res, ok := Await(0, func(f *Future[int]) {
		f.Error("That is bad!")
	}).Get()
	if ok || res != 0 {
		t.Fatal("Final result is bad!")
	}
}

func TestFutureClosed(t *testing.T) {
	f := NewFuture(0)
	_, nonError := f.TryGet()
	if nonError {
		t.Fatal("First result should have fail, but has success")
	}
	f.Ok(3000)
	n, ok := f.TryGet()
	if !ok || n != 3000 {
		t.Fatal("Final result is bad!")
	}
}

func TestFutureClosedWithError(t *testing.T) {
	f := NewFuture(0)
	_, nonError := f.TryGet()
	if nonError {
		t.Fatal("First result should have fail, but has success")
	}
	f.Error("Teddy")
	_, nonError = f.TryGet()
	if nonError {
		t.Fatal("Should be error here")
	}
	if f.GetError() != "Teddy" {
		t.Fatal("Error result is bad!")
	}
}
