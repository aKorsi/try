# try

Simple and idiomatic Go try package.

### Usage

Just call `try.Do` with the function you want to try in the event of an error or panic:

```go
var value string
err := try.Do(func() error {
    var err error
    value, err = SomeFunction()
    return err
})
if err != nil {
    log.Fatalln("error:", err)
}
```

#### Retrying

Try supports retrying in the event of an error or panic.

```go
count:= 5
var value string
err := try.ReDoByCounter(func () error {
    var err error
    value, err = SomeFunction()
    return err
}, count)
if err != nil {
    log.Fatalln("error:", err)
}
```

#### Delay between retries

To introduce a delay between retries, just make a `try.ReDoByCounterWithDelay` call.

```go
count:= 5
var value string
err := try.ReDoByCounterWithDelay(func () error {
    var err error
    value, err = SomeFunction()
    return err
}, count, time.Second)
if err != nil {
    log.Fatalln("error:", err)
}
```