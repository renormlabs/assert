# assert

`assert` provides a minimal, zero‑dependency set of test assertions that make writing tests easier and more clear.

## Installation

```go get github.com/renormlabs/assert@latest```

## Example

```go
package example_test

import (
    "testing"
    "github.com/renormlabs/assert"
)

func TestAdd(t *testing.T) {
    got := add(40, 2)
    assert.Equal(t, 42, got)
}
```
