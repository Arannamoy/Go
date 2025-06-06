# Go
## Table of contents:
- [Init mode file](#init-mod-file)
- [Sample Code](#sample-code)
- [Run Code](#run-code) 
- [Escape sequence and comment](#escape-sequence-and-comment)
- [Variable](#variable)
- [Formatting output using %](#formatting-output-using-)
- [Take input](#take-input)
- [Integer formatting number conversion for decimal to binary,octal,hexadecimal](#integer-formatting--number-conversion-for-decimal-to-binaryoctalhexadecimal)
- [Arithmatic operation](#arithmatic-operation)
- [Conditions: If, else-if, else](#ifelse-ifelse)
- [Odd Even](#odd-even)
- [Switch statement](#switch-statement)
- [Multiple case in switch statement](#multiple-case-in-switch-statement)
- [For Loop Only One Loop in Golang](#for-loop)
- [Break and continue statement](#break-and-continue-statement)
- [Structure](#structure)
- [Array](#array)
- [Dynamic-array | slice](#dynamic-array--slice)
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
- [Topic](#table-of-contents)
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

- [Topic](#table-of-contents)
## Variable
- Convention and rule is same as others.

```go
package main

import "fmt"

func main() {
	var type1 string = "Hello world"
	const type2 int = 32
	var type3 bool = true
	var type4, type5, type6 string
	type4 = " "
	type5 = " "
	type6 = " "
	type7:="No Var using colon after variable"
	fmt.Println("Output: ", type1, type2, type3, type4, type5, type6,"\n",type7)
}

```
- [Topic](#table-of-contents)
## Formatting output using %:

```go
package main

import "fmt"

func main() {
	var type1 string = "Hello world"
	fmt.Printf("%s", type1)
}
```

- [Topic](#table-of-contents)
## Take input:

```go
package main

import "fmt"

func main() {
	// For input three method Scanf,Scanln,Scan
	var type1, type2, type3 string
	fmt.Println("type1:")
	fmt.Scanln(&type1)
	fmt.Println("type2:")
	fmt.Scan(&type2)
	type3 = type2
	fmt.Println(type1, type2, type3)
	fmt.Scan(&type1, &type2, &type3)
	fmt.Println(type1, type2, type3)
}
```
- [Topic](#table-of-contents)
## Integer Formatting & number conversion for decimal to (binary,octal,hexadecimal):

```go
package main

import "fmt"

func main() {
	var decimal int
	fmt.Print("Enter a decimal number:")
	fmt.Scan(&decimal)
	fmt.Printf("Binary number: %b , Octal number: %o, Hex number: %x", decimal, decimal, decimal)
	type2 := 2.141166525662
	fmt.Printf("\n%.2f", type2)
}

```
- [Topic](#table-of-contents)
## Arithmatic operation:

```go
package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("Enter two number:")
	fmt.Scan(&num1, &num2)
	fmt.Println("For modulus(%):", num1%num2)

	// Others operations are same
}
```
- [Topic](#table-of-contents)
## If,else if,else

> Syntax is case sensitive for Golang.

```go
package main

import ("fmt")

func main() {
	var num1, num2 int
	fmt.Print("Enter two number:")
	fmt.Scan(&num1)

	// Must follow this sytax for if,else starting position
	if(num1>num2){
		fmt.Print("Number 1 is max.")
	}else if(num1<num2){
		fmt.Print("Num2 is max")
	}else{
		fmt.Print("Same value.")
	}
}
```
- [Topic](#table-of-contents)
## Odd even
```go
package main

import ("fmt")

func main() {
	var num1 int
	fmt.Print("Enter a number:")
	fmt.Scan(&num1)

	// Must follow this sytax for if,else starting position
	if(num1%2 ==0){
		fmt.Println("Number is even.")
	}else{
		fmt.Println("Number is odd.")
	}
}
```
- [Topic](#table-of-contents)
## Switch statement

```go
package main

import ("fmt")

func main() {
	var num1 int
	fmt.Print("Enter a number:")
	fmt.Scan(&num1)

	switch num1{
	case 1:fmt.Println("Sunday")
	case 2:fmt.Println("Monday")
	case 3:fmt.Println("Tuesday")
	case 4:fmt.Println("Wednesday")
	case 5:fmt.Println("Thursday")
	case 6:fmt.Println("Friday")
	case 7:fmt.Println("Saturday")
	default:fmt.Println("Invalid input")
	}
}
```
- [Topic](#table-of-contents)
## Multiple case in switch statement

```go
package main

import ("fmt")

func main() {
	var num1 int
	fmt.Print("Enter a number:")
	fmt.Scan(&num1)

	switch num1{
	case 6,7:fmt.Println("Weekend")
	case 1,2,3,4,5:fmt.Println("Working day")
	default:fmt.Println("Invalid input")
	}
}
```
- [Topic](#table-of-contents)
## For loop

> Only one loop in golang

```go
package main

import (
	"fmt"
)

func main() {
	var num1 int
	fmt.Print("Enter a number:")
	fmt.Scan(&num1)

	// syntax
	for i := 0; i < num1; i++ {
		fmt.Println(i)
	}

	for i := num1; i > 0; i -= 2 {
		fmt.Println(i)
	}

    x:=10

    for x>=0{
        fmt.Println(x)
		x--
    }
}
```

- [Topic](#table-of-contents)
## Break and continue statement

```go
package main 

import "fmt"

func main(){
	for i:=10;;i+=3{
		if(i%2 ==0){
			fmt.Println("Even")
			continue
		}else if(i>101){
			break
		}
	}
}
```
- [Topic](#table-of-contents)
## Function

```go
package main 

import "fmt"


// return type function 
func add(x int,y int)int{  
	return x+y;
}

// no return type function
func sum(x int,y int){  
	fmt.Println(x+y)
}

func main(){
	fmt.Println(add(10,15))
	sum(10,15)
}
```
- [Topic](#table-of-contents)
## Structure

```go
package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	std1 := Student{"A", 32}
	var std2 Student
	std2.name = "B"
	std2.age = 20
	fmt.Println(std1, "\n", std2)
}
```
- [Topic](#table-of-contents)
## Array

```go
package main

import "fmt"

func main() {
	var array1 [50]int
	for i := 1; i < 50; i++ {
		array1[i] = i
	}

	fmt.Println(array1)
}

```

- [Topic](#table-of-contents)
## Dynamic array | slice:

```go
package main
import "fmt"

func main(){
	var slice [] int;
	for i:=0;i<10;i++{
		slice = append(slice, i)
	}
	fmt.Println(slice)
}
```