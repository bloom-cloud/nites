# Start Project

```bash
# Initialize project
go mod init name-of-project

# Install library
go get -u github.com/inancgumus/screen

# Create main.go
```

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world.", "New line element", "Even newer")
	time.Sleep(5 * time.Second)
}
```

```bash
# Run
go run ./main.go

# Build
go build ./main.go
```

---

> Programming is not knowledge-based.
> It’s a craft. Practice it. Fail, persevere, and master the skill.

> It’s inconvenient to know multiple ways of doing the same thing.

---

# String Formatting

```go
package main

import "fmt"

func main() {
	// Newlines and Raw Strings
	fmt.Println("This is multiline\nThis is my code") // \n creates a new line
	fmt.Println(`This is another
statement
called a raw string literal`) // Backticks preserve formatting

	// Basic Formatting
	var name = "T"
	var age = 33
	var initial rune = 'P' // rune prints as int by default
	fmt.Printf("String: %s | Number: %d | Char: %c\n", name, age, initial)

	// Auto Type Detection
	var anything = "t"
	fmt.Printf("Go infers type: %v\n", anything)

	// Floats
	var fl = 10.1111
	fmt.Printf("Two decimals: %.2f\n", fl)
	fmt.Printf("Fixed width (10, 2 decimals): %10.2f\n", fl)

	// Formatted String (store instead of print)
	var formattedOutput = fmt.Sprintf("Formatted %s", "word")
	println(formattedOutput)
}
```

**Quick Reference**

| Verb | Meaning            | Example               |
| ---- | ------------------ | --------------------- |
| `%v` | default format     | `fmt.Printf("%v", x)` |
| `%T` | type of value      | `fmt.Printf("%T", x)` |
| `%s` | string             | `"text"`              |
| `%d` | integer (base 10)  | `42`                  |
| `%f` | float              | `3.14`                |
| `%c` | character (rune)   | `'A'`                 |
| `%t` | boolean            | `true`                |
| `%x` | hex representation | `255 -> ff`           |

* `%.nf` → control decimal places
* `%<width>.<prec>f` → alignment and precision

---

# Exercise

```go
package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"time"
)

func main() {
	fmt.Println("For loop")
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

	fmt.Println("Power")
	fmt.Println(math.Pow(2, 11))

	fmt.Println("Random")
	fmt.Println(rand.IntN(7))

	fmt.Println("Time")
	fmt.Println(time.Now().Date())
}
```

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	year, month, day := t.Date()
	fmt.Println(year, month, day)
}
```

---

# Declaring Variables

```go
package main

func main() {
	var text1 string
	text1 = "Text 1"

	var text2 = "Text 2"   // Type inferred
	text3 := "Text 3"      // Idiomatic Go

	println(text1)
	println(text2)
	println(text3)

	// Constants don't change, don't require type, and must be primitive
	const pi = 3.14
}
```

---

# Basic Data Types

> Local variables cannot share names.
> Global and local variables can.

---

### Integer

```go
// Types: int, int8, int16, int32, int64
//        uint, uint8, uint16, uint32, uint64
//        byte (alias for uint8)
//        rune (alias for int32)
//        uintptr (for memory addresses)

var x int = 10
var p *int = &x
var u uintptr = uintptr(unsafe.Pointer(p))

var number int = 4
```

---

### Floating Point

```go
// Types: float32, float64

var realNum float32 = 1.46
var wholeNum float32 = 1
```

---

### Rune, Boolean, String

```go
var isCool bool = true

var letter1 rune = 'A'
var letter2 rune = A        // invalid - missing quotes
var letter3 rune = '🦁'

var text string = "Some text"
```

---

### Conversion

```go
package main

import (
	"strconv"
)

func main() {
	// String to int
	numAsStr := "522"
	num, err := strconv.Atoi(numAsStr)

	// Int to float
	numAsFloat := float32(num)

	// Int to larger int
	numAsBiggerInt := int64(num)

	// Int to smaller int (may overflow)
	numAsSmallerInt := int8(num)

	println(numAsStr, num, err, numAsFloat, numAsBiggerInt, numAsSmallerInt)
}
```

---

