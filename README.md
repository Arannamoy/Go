# Go
## Table of contents:
- [Init mode file](#init-mod-file)
- [Sample Code](#sample-code)
- [Run Code](#run-code) 
- [Escape sequence and comment](#escape-sequence-and-comment)
- [Variable](#variable)

- [Topic](#table-of-contents)
## Init mod file:

```
go mod init App-Name
```
- [Topic](#table-of-contents)
## Sample Code:

```go
package main

import "fmt"

func main() {
	fmt.Print("Hello\n")
}

```

- [Topic](#table-of-contents)
## Run code:

```
go run filename
```

## Escape sequence and comment


`Escape sequence: \n,\t`

```go
package main

import "fmt"

func main() {
	/*fmt.Print("Hello\n")
	fmt.Print(("Hello"))
	fmt.Print("Hello")*/
	fmt.Println("Only programming is real.\tAlso problem solving real.\nWithout problem coding is like a desktop without cpu.\\\"")
}
```

`Single line comment`

```go
package main

import "fmt"

func main() {
	// fmt.Print("Hello\n")
	// fmt.Print(("Hello"))
	// fmt.Print("Hello")
	fmt.Println("Only programming is real.")
}

```

`Multi line comment`

```go
package main

import "fmt"

func main() {
	/*fmt.Print("Hello\n")
	fmt.Print(("Hello"))
	fmt.Print("Hello")*/
	fmt.Println("Only programming is real.")
}
```


## Variable
- Convention and rule is same as others.

```go
package main

import "fmt"

func main() {
	var type1 string = "Hello world"
	const type2 int = 32
	var type3 bool = true
	var type4,type5,type6 string;
	type4=" "
	type5=" "
	type6=" "
	fmt.Println("Output: ",type1, type2, type3,type4,type5,type6)
}
```