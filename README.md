# strpad

String padding for Go.

## Usage

```go
import (
    "fmt"
    "github.com/bepisdev/strpad"
)

func main() {
    myPaddedString := strpad.Pad("Hello, world.")
    fmt.Println(myPaddedString)
    // "Hello, world." becomes " Hello, world. "
}
```

The `Pad()` function can be swapped out for `PadLeft()` or `PadRight()` using
the identical API to pad the left or right side of a string without the other.

> There are also other functions for padding that follow this naming convention.

- `PadSymbol` and `PadNSymbol` functions, take an extra param as a string, and use that to pad the input string.

```go

import (
    "fmt"
    "github.com/bepisdev/strpad"
)

func main() {
    myPaddedString := strpad.PadN("Hello, world.", 2)
    fmt.Println(myPaddedString)
    // "Hello, world." becomes "  Hello, world.  "

    symbolPaddedStr := strpad.PadSymbol("Hello, world.", "**")
    fmt.Println(symbolPaddedStr)
    // "Hello, world." becomes "**Hello, world.**"
}

```

The complete API documentation can be found on
[pkg.go.dev](https://pkg.go.dev/github.com/bepisdev/strpad)

## Installation

```
go get github.com/bepisdev/strpad
```

## License

This software is open-sourced under the MIT license.
