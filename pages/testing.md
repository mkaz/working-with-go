---
title: Testing
layout: page
order: 24
---

# Testing

The [Go testing package](https://golang.org/pkg/testing/) supports creating unit test in go. Your test code is written in plain go, there is not a new assert language or syntax to learn. Write go code and use `t.Fail` or `t.Error` to signal the test did not pass.

To write a new test create a test file, ending in `_test.go` that includes functions with the following signature: `func TestXxx(*testing.T)`

Note: The function name must start with `Test` and include a capitalized second part to identify the test routing.

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

## Test Table

Use a table structure to setup and test multiple inputs and expected values.

```go
func TestAdd(t *testing.T) {
	var testData = []struct {
		param1 int
		param2 int
		expect int
	}{
		{ 1, 2, 3 },
		{ 0, 1, 1 },
		{ 1, 0, 1 },
		{ -1, 1, 0 },
		{ -1, -1, -2 },
	}

	for _, td := range testData {
		result := Add(td.param1, td.param2)
		if td.expect != result {
			t.Errorf("Expected %s but received %s", td.expect, result)
		}
	}
}
```


