<p align="center">
  <h1 align="center">Fibonacci</h1>
  <p align="center">A Fibonacci sequence generator.</p>
  <p align="center">
    <a href="https://travis-ci.org/andrewrynhard/fibonacci"><img alt="Travis" src="https://img.shields.io/travis/andrewrynhard/fibonacci.svg?style=flat-square"></a>
    <a href="https://codecov.io/gh/andrewrynhard/fibonacci"><img alt="Codecov" src="https://img.shields.io/codecov/c/github/andrewrynhard/fibonacci.svg?style=flat-square"></a>
    <a href="https://goreportcard.com/report/github.com/andrewrynhard/fibonacci"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/andrewrynhard/fibonacci?style=flat-square"></a>
    <a href="https://github.com/andrewrynhard/fibonacci/releases/latest"><img alt="Release" src="https://img.shields.io/github/release/andrewrynhard/fibonacci.svg?style=flat-square"></a>
    <a href="https://github.com/andrewrynhard/fibonacci/releases/latest"><img alt="GitHub (pre-)release" src="https://img.shields.io/github/release/andrewrynhard/fibonacci/all.svg?style=flat-square"></a>
  </p>
</p>

---

Fibonacci is a simple RESTful service that generates a sequence of Fibonacci
numbers. For example:

```bash
$ curl fibonacci.local/v1/sequence/5
{"sequence":["0","1","1","2","3"]}
```

For more information see the [documentation](https://andrewrynhard.github.io/fibonacci).
