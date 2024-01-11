# Codingame Framework

[![codecov](https://codecov.io/gh/mrsombre/codingame-framework/graph/badge.svg?token=I8RYIUSN6Q)](https://codecov.io/gh/mrsombre/codingame-framework)

A set of algorithms and data structures for solving puzzles.

Feel free to follow me on [Codingame Profile](https://www.codingame.com/profile/9dd9f9f38412d78eaf21718bf6e87ca0626964).

### Golang Version

Currently, the framework is written in Golang 1.19.x as it is closer to 1.18.x available on Codingame.

### Tests

```shell
go test -cover ./...
```

### Benchmarks

```shell
go test -bench=. -benchmem -run=^$ > bench.out
```
