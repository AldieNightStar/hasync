# H-Async - Await-async lib with Futures

```go
// Initial value
const INITAL_VAL = 0

res, err := Await(INITAL_VAL, func(f *Future[int]) {
    // Some processing. Let it be Sleep(...) function
    time.Sleep(5 * time.Millisecond)
    // In case of ok, return result
    // Later result could be used with Get() or TryGet()
    f.Ok(122)
    // For error
    // Later could be used with 
    //   f.Error("Something is bad here")
}).Get()

if err != nil {
    fmt.Println("Error happened: ", err.Error())
    return
}
fmt.Println("Resul is: ", res)
```