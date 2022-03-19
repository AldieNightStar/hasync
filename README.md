# H-Async - Await-async lib with Futures

```go
// Initial value
const INITAL_VAL = 0

future := Await(INITAL_VAL, func(f *Future[int]) {
    // Some processing. Let it be Sleep(...) function
    time.Sleep(5 * time.Millisecond)
    // In case of ok, return result
    // Later result could be used with Get() or TryGet()
    f.Ok(122)
    // For error
    // Later could be used with 
    //   f.Error("Something is bad here")
})
res, ok := future.Get()

if !ok {
    fmt.Println("Error happened: ", future.GetError())
    return
}
fmt.Println("Resul is: ", res)
```