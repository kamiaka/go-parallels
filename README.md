# go-parallels
[![GoDoc](https://godoc.org/github.com/kamiaka/go-parallels?status.svg)](https://godoc.org/github.com/kamiaka/go-parallels)

Go package for parallel execute function.

## How to use

### Parallel execute all

```go
parallels.Do(func(i int) error {
  fmt.Println(i)
  return nil
}, 10)
```

### Specify number of concurrent executions

```go
err := parallels.Do(func(i int) error {
  fmt.Println(i)
  return nil
}, 10, parallels.Concurrent(2))
if err != nil {
  // handling error
}
```

## License

MIT
