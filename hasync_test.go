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
	_, err := Await(0, func(f *Future[int]) {
		f.Error("That is bad!")
	}).Get()
	if err == nil {
		t.Fatal("No error, but should be!")
	}
	if err.Error() != "That is bad!" {
		t.Fatal("Final error message is not correct!", err.Error())
	}
}

func TestFutureClosed(t *testing.T) {
	f := NewFuture(0)
	_, err := f.TryGet()
	if err == nil {
		t.Fatal("First result should have error, but has success")
	}
	f.Ok(3000)
	n, err := f.TryGet()
	if err != nil || n != 3000 {
		t.Fatal("Final result is bad!")
	}
}

func TestFutureClosedWithError(t *testing.T) {
	f := NewFuture(0)
	_, err := f.TryGet()
	if err == nil {
		t.Fatal("First result should have fail, but has success")
	}
	f.Error("Teddy")
	_, err = f.TryGet()
	if err == nil {
		t.Fatal("Should be error here")
	}
	if err.Error() != "Teddy" {
		t.Fatal("Error result is bad!")
	}
}

func TestChanErr(t *testing.T) {
	future := Await(0, func(f *Future[int]) {
		f.Ok(1000)
	})
	res, _ := future.Get()
	if res != 1000 {
		t.Fatal("Result is not correct")
	}
	_, err := future.Get()
	if err == nil {
		t.Fatal("Here should be error!")
	}
	if err.Error() != ERR_CHANERR {
		t.Fatal("Should be ERR_CHANERR error: ", err.Error())
	}
}
