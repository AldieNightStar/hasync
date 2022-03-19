# H-Async - Await-async lib with Futures

# Code samples
* Run async
```go
res, err := Await[T](INIT_VAL, func(f Future[T]) {
    f.Ok(T)
    f.Error("Bad!")
})
```
* Get results
```go
// returns (T, error).
// [!] Could be used ONCE. Second time will be: ERR_CHANERR
// Will block until future will be completed
//   ERR_CHANERR   - when channel is already closed
f.Get()
// retunrs (T, error).
// Will NOT block and will return result.
//   ERR_NOTCLOSED - if future is not completed yet
f.TryGet() 
```


# Example
```go
// Initial value
// First value before result is ended
// It could be any type
const INITAL_VAL = 0

// Run async function
res, err := Await[int](INITAL_VAL, func(f *Future[int]) {
    // Some processing. Let it be Sleep(...) function
    time.Sleep(5 * time.Millisecond)
    // In case of ok, return result
    // Later result could be used with Get() or TryGet()
    f.Ok(122)
    // For error
    // Later could be used with 
    f.Error("Something is bad here")
}).Get()

if err != nil {
    fmt.Println("Error happened: ", err.Error())
    return
}
fmt.Println("Resul is: ", res)
```