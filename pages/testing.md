---
title: Testing
parent: 1465
template: page-tut.php
order: 24
---

# Testing

The [Go testing package](https://golang.org/pkg/testing/) supports creating unit test in go. Your test code is written in plain go, there is not a new assert language or syntax to learn. You write go code, and then Fail or Error out to signal the test did not pass.

To write a new test created a test file, ending in `_test.go` which includes functions with the following signature: `func TestXxx(*testing.T)`

Note: The function name must start with `Test` and include a capitalized second part which is used to identify the test routing.

For example, if you want to test an `Add` function in your project:

Create `add_test.go`:

```go
package main

import "testing"

func TestAdd(t *testing.T) {
    sum := Add(1, 1)
    if sum != 2 {
       t.Errorf("Expected %d => Received %d", 2, sum)
    }
}
```

Run test using: `go test`

