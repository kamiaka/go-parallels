# go-parallels

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
parallels.Do(func(i int) error {
  fmt.Println(i)
  return nil
}, 10, parallels.Concurrent(2))
```

## License

MIT