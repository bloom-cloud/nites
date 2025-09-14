
----------

# 🚀 Start Project

```bash
# Init project
go mod init name-of-project

# Create main.go

```

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world. This is start")
	time.Sleep(5 * time.Second)
}

```

```bash
# Run
go run ./main.go

# Build
go build ./main.go

```

----------

> Programming is not a knowledge-based endeavor.  
> It's a craft. Practice like swimming.  
> Fail. Persevere. Acquire the skill.

----------

# 🧠 Exercise

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

----------

# 📝 Declaring Variables

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
}

```

----------

# 📦 Basic Data Types

> _Local variables can't have the same name._  
> _Global + local can._

----------

### 🔢 Integer

```go
// Types: int, int8, int16, int32, int64
//        uint, uint8, uint16, uint32, uint64
//        byte (alias for uint8)
//        rune (alias for int32)
//        uintptr - memory address

var x int = 10
var p *int = &x
var u uintptr = uintptr(unsafe.Pointer(p))

var number int = 4

```

----------

### 🔣 Floating Point

```go
// Types: float32, float64

var realNum float32 = 1.46
var wholeNum float32 = 1    // valid too

```

----------

### 🎯 Rune, Boolean, String

```go
var isCool bool = true

var letter1 rune = 'A'
var letter2 rune = A        // invalid - missing quotes
var letter3 rune = '🦁'

var text string = "Some text"

```
