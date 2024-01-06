## 01 - Package Vs Module

a package is a folder with a collection of go files.

a module is a colleection of packages

**Package:**
```plaintext
folder_1
    |_file1.go
    |_file2.go
    |_file3.go
```

**Module:**
```plaintext
module_1
    |_folder_1
    |   |_file1.go
    |   |_file2.go
    |   |_file3.go
    |
    |_folder_2
        |_file1.go
        |_file2.go
        |_file3.go
```

## 02 - Start a new module

go mod init ...

usually you would name the module after the github location path:

```bash
go mod init github.com/hoejsagerc/go_tutorials
```

every file is part of a package and should be declared in the top of the file

```go
package main
```

the main package is a special package which should contain the main entry point of the program.
so the compiler will look for a main method whithin the main package


```go
package main


func main() {

}
```


## 03 - Importing packages

you can import a package / library in the top of a file. You are not allowed to import
libraries which are not used in you code

```go
package main
import "fmt"


func main(){
    fmt.Println("Hello World");
}
```


## 04 - Building / Running go programs

```bash
# Building the go file
go build cmd/tutorial_1/main.go
# Then run the program maing.exe
./main.ext


# Building and running the go file
go run cmd/tutorial_1/main.go
```


## 05 - Constants, Variables & Data types

variables are declared with "var" - "variable name" - "data type"
example:
```go
var intNum int
```
all variables declared needs to be used for the compiler to run

### Integers
int8 = 127 (the reason for the 128 bits - 1 => is that the 1 bit is used for the + or - sign)
int16 = 32.767
int32 = 2.147.483.648
int64 = 9.223.372.036.854.775.807

if you forexample try to store 32.768 in an int16 the compiler will give you an overflow error.

**Important note!:** when running the code you will not get an error. Only when compiling. =>
instead you will experience some wierd results when running your code. so be sure the handle the
types correctly.


### Unsigned Integers
like int(integers) we have uint(unsigned integers) which will only store the positive value meaning you will
have twice as large numbers. The unsigned integers will not store a bit for the sign.

uint8 = 255
uint16 = 65.535
uint32 = 4.294.967.295
uint64 = 18.446.744.073.709.551.615


### Floats

we have either a float32 or the float64. Just like the integers it bit representation indicates
how precise the decimal will be.

a quick note on the precesion

In the example below you could have a float32 and when you try to print it you see that
you won't get the exact number you stored in the variable. This is due to how the float
is stored on the machine.
```go
package main
import "fmt"

func main() {
    var floatNum float32 = 12345678.9
    fmt.Println(floatNum) >> 123456789.000000
}
```

where as in the example below you see that the exact decimal will be printed
```go
package main
import "fmt"

func main() {
    var floatNum float64 = 12345678.9
    fmt.Println(floatNum) >> 12345678.900000
}
```

be sure to think about which data types to use instead of just choosing the most precice everytime.


### Data type operations
you cannot do operations between to different data types.

so you cannot do float32 * int32 and you cannot do float32 + float64.

so if you wanted to do something like this you would need to cast the type to a different type:

#### Casting

```go
func main() {
    var flaotNum float32 = 10.1
    var intNum int32 = 2
    var result float32 = floatNum + float32(intNum)
    fmt.Println(result) >> 12.1
}
```

#### Integer division
Integer division results in a division and the result will be rounded down

```go
func main() {
    var intNum1 int = 3
    var intNum2 int = 2
    fmt.Println(intNum1/intNum2) >> 1

    // To get the remainder use the modulus
    fmt.Println(intNum1%intNum2) >> 1
}
```

### String and Runes
you can store strings within the string type and be expressed with either " or with a `


```go
func main() {
    var myString string = "this is my fist string"
    var myOtherString string = `This
    is my other string`
}
```

when using "" => you cannot format the string on multiple lines. You can use \n to indicate a line break

When using `` => you can format the string on multiple lines


You can conatinate string with the + sign

```go
func main() {
    var myString string = "string 1" + "string 2"
}
```

You cand use len methos to get the length. But it is not the the characters in the string.
it will be the bytes of the string.


So if you want the actual character length of a string you need to import a module

```go
import "unicode/utf8"

func main() {
    var myString string = "test"
    fmt.Println(utf8.RuneCountInString(myString)) >> 4
}
```

A Rune is the same as a Char in C# and you can use the data type like this:

```go
import "unicode/utf8"

func main() {
    var myChar rune = "t"
}
```

If you try to print a rune you will get a wierd output since it is not handled like a string

### Booleans

Boolean is like a boolean: true or false

```go
func main() {
    var myBool bool
    fmt.Println(myBool) >> false
}
```


### Default values

for: 
- uint
- uint8
- uint16
- uint32
- uint64
- int 
- int8
- int16
- int32
- int64
- float32
- float64
- runes

The default value is **0**

for:
- string

the default value is **""**

for:
- boolean

the default value is **false**


### Inferred variable types and variable declarations

You can skip the type definition when assigning a variable if the type is inferred

```go
func main() {
    var myString = "mystring"
}
```

You can also use the shorthand method:

```go
func main() {
    myString := "mystring"
}
```

You can assign multiple variables in a single lines with all of the above methods

```go
func main() {
    var var1, var2 int = 1, 2
    var var1, var2 = 1, 2
    var1, var2 := 1, 2
}
```

### Constants

Everything learned up untill now so everything for the variables and declarations
applies to constants aswell, except that you cannot change the value of a const
after it has been declared.

```go
func main() {
    const myConst string = "My constant string"
}
```

You will have to initialize a constant with a value when declaring it.


## 06 - Functions and control structures

Functions starting with a lowercase letter is set as private functions and cannot
be imported into other files whereas functions starting with a capital letter is
public functions and can be imported.

Defining and calling a function

```go
package main
import "fmt"

func main() {
    const myString string = "Hello world"
    printMe(myString)
}

func printMe(printMeValue string) {
    fmt.Println(printMeValue)
}
```


if you need to return a value from a function you will need to define the return type

```go
package main

func intDivision(numerator int, denominator int) int {
    var result int = numerator/denominator
    return result
}
```

you can return multiple values aswell. Note here that i am using Prinf to print with 
string interpolation.

```go
package main
import "fmt"

func main() {
    var numerator int = 11
    var denominator int = 2
    var result, remainder int = intDivision(numerator, denominator)
    fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
}

func intDivision(numerator int, denominator int) (int, int) {
    var result int = numerator/denominator
    var remainder int = numerator%denominator
    return result, remainder
}
```

Errors will be handled as types or values is go meaning that you can do error handling and
returning the errors.

```go
package main
import "fmt"

func main() {
    var numerator int = 11
    var denominator int = 2
    var result, remainder, err = intDivision(numerator, denominator)
    if err != nill {
        fmt.Printf(err.Error())
    } else if remainde r== 0 {
        fmt.Prinf("The result of the integer division id %v", result)
    } else {
        fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
    }
}

func intDivision(numerator int, denominator int) (int, int, error) {
    var err error 
    if denominator==0{
        err = errors.New("cannot divide by zero")
        return 0,0,err
    }
    var result int = numerator/denominator
    var remainder int = numerator%denominator
    return result, remainder, err
}
```

### Operators

== equals
!= not equals
&& and operator
|| or operator


### Switch statement

instead of using the if else statements from earlier we can use the switch statement
```go
package main
import "fmt"

func main() {
    var numerator int = 11
    var denominator int = 2
    var result, remainder, err = intDivision(numerator, denominator)

    switch {
        case err != nill {
            fmt.Printf(err.Error())
        }
        case remainder == 0 {

           fmt.Prinf("The result of the integer division id %v", result)
        }
        default {
            fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
        }
    }
}

func intDivision(numerator int, denominator int) (int, int, error) {
    var err error 
    if denominator==0{
        err = errors.New("cannot divide by zero")
        return 0,0,err
    }
    var result int = numerator/denominator
    var remainder int = numerator%denominator
    return result, remainder, err
}
```

important note here is that the break is implicit and we do not need to declare it after
the switch case statement.


we can also write conditional switch statements. Taken with context from above example we can say:

```go
func main() {
    var numerator int = 11
    var denominator int = 2
    var result, remainder, err = intDivision(numerator, denominator)

    switch {
        case err != nill {
            fmt.Printf(err.Error())
        }
        case remainder == 0 {

           fmt.Prinf("The result of the integer division id %v", result)
        }
        default {
            fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
        }
    }
}

switch remainder {
    case 0: // if 0
        fmt.Printf("The division was exact")
    case 1,2 // if either 1 or 2
        fmt.Printf("The division was close")
    default: // otherwise run
        fmt.Printf("The division was not close")
}
```


## 07 - Arrays, Slices, Maps and Loops

### Arrays

- Fixed length
- Same type
- indexable
- Contigues in memory (right next to each other in a single unbroken block of memory)

```go
func main() {
    var intArray [3]int32
}
```

- In the example above we state the the lenght of the array is 3 meaning it can hold 3 elements.
- The 3 elements will be of type int32

the array will hold the default value of the types used in the array. so for the int32 type
the default value is 0 => therefore at this moment the array will be [0,0,0]

you can index the first element like this: intArray[0]

and you can index elemnent 1 and 2 like this: intArray [1:3]

You can change the an elemeent by using indexing: intArray[1] = 123


You can initialize the array with the following syntaxes:


```go
func main() {
    var intArray [3]int32 = [3]int32{1,2,3}

    // or...
    intArray := [3]int32{1,2,3}

    // or...
    int Array := [...]int32{1,2,3}
}
```


### Slices

slices are just arrays under the hood with some more functionality. They can be compared
to List in c#

```go
func main() {
    var intSlice []int32 = []int32{,4,5,6}
}
```

you can then add a new value to the slice using append

```go
func main() {
    var intSlice []int32 = []int32{,4,5,6}
    intSlice = append(intSlice, 7)
}
```

what actually happens under the hood is that once the intSlice array was created
with a specific placement in memory. When we then try to append to the array it will check
to see if the array has enough room to add the new value.
In the example above the array first created will not have enough room, and a new array with
room for the new value will be created instead.

you can use the builtin methods len and cap to check how it works:

```go
immport "fmt"

func main() {
    var intSlice []int32 = []int32{4,5,6}
    fmt.Printf("The length is: %v, with a capacity of: %v", len(intSlice), cap(intSlice))
    // >> The length is 3 with capacity of: 3

    intSlice = append(intSlice, 7)
    fmt.Printf("\nThe length is: %v, with a capacity of: %v", len(intSlice), cap(intSlice))
    // >> The length is 4 with capacity of: 6
}
```


You can append multiple values to the slice like this:


```go
func main() {
    var intSlice []int32 = []int32{4,5,6}

    var intSlice 2 []int32 = []int32{7,8}
    intSlice = append(intSlice, intSlice2...) // using the spread operator
}
```

You can also initialize a slice with a given capacity by using the make function

```go
func main() {
    // we specify the lenght of the array with 3 and optionally the capacity with 8
    var intSlice []int32 = make(int32[], 3, 8) // by default the capacity will be the lenght
}
```

You can iterate over a slice and array by using:

```go
func main() {
    var intSlice []int32 = []int32{4,5,6}

    for i, v := range intSlice {
        fmt.Printf("Index: %v and Value: %v", i, v)
    }
}
```

### Maps

a set of key value pairs

you can define a map with the keys of type string and the values of type uint8

```go
func main() {
    var myMap map[string]uint8 = make(map[string]uint8)

    var myMap2 = map[string]uint8{"Adam": 32, "Sarah": 45}
}
```

And you can get the value by indexing the key

```go
func main() {
    var myMap2 = map[string]uint8{"Adam": 32, "Sarah": 45}
    fmt.Prinln(myMap2["Adam"]) >> 32
}
```

if i try to index a key which does not exist i will not get an error, instead i will
get the default value of the "value" for a key within this map. Meaning that here i would
get 0 since it is the default value of an uint8

by default when indexing the key it will return 2 values. the Value and a boolean.
The boolean defines if the key exists in the map or not

```go
func main() {
    var myMap2 = map[string]uint8{"Adam": 32, "Sarah": 45}
    var age, ok = myMap2["Adadm"]
    fmt.Println{ok} >> true
}
```

you can delete a key by using the method delete

```go
func main() {
    var myMap2 = map[string]uint8{"Adam": 32, "Sarah": 45}
    delete(myMap2, "Adam")
}
```

We can iterate over maps by using:

```go
func main() {
    var myMap2 = map[string]uint8{"Adam": 32, "Sarah": 45}

    for name, value := range myMap2 {
        fmt.Printf("Name: %v, Age: %v", name, age)
    }
}
```

### Loops

Looping over a Slice

```go
func main() {
    var intSlice []int32 = []int32{4,5,6}

    for i, v := range intSlice {
        fmt.Printf("Index: %v and Value: %v", i, v)
    }
}
```

Looping over a Map

```go
func main() {
    var myMap2 = map[string]uint8{"Adam": 32, "Sarah": 45}

    for name, value := range myMap2 {
        fmt.Printf("Name: %v, Age: %v", name, age)
    }
}
```

a while loop in go would look similar to: 

```go
func main() {
    var i int = 0
    for i < 10 {
        fmt.Println(i)
        i = i + 1
    }
}
```

We can also omit the loop by using the break

```go
func main() {
    var i int = 0
    for i < 10 {
        if >= 10 {
            break
        }
        fmt.Println(i)
        i = i + 1
    }
}
```

and the above example can be done with a different syntax:

```go
func main() {
    for i:=0; i<10; i++ {
        fmt.Println(i)
    }
}
```

for the operator in the "post" part => i++
there are multiple operators that can be used:

- i--        => subtract
- i++        => add
- i+=10      => add by 10
- i-=10      => decrement by 10
- i*=10      => multiply by 10
- i/=10      => divide by 10


## 8 - Strings, Runes and bytes

string building can be done with the "strings" module

```go
import (
    "fmt"
    "strings"
)

func main() {
    var stringSlice = []string {"t","e","s","t","i","n","g"}
    var strBuilder = strings.builder
    for i := range stringSlice {
        strBuilder.WriteString(stringSlice[i])
    }

    var catStr string = strBuilder.String()
    fmt.Printf("\n%v", catStr)
}
```

## 9 - Structs and Interfaces

### Structs

a struct is similar to your own type and can be defined like below:

```go
package main

type gasEngine struct{
    mpg uint8
    gallons uint8
    ownerInfo owner // you could also just have owner this would allow you to access with myEngine.name <= syntax from subfield
}

type owner struct {
    name string
}

func main() {
    var myEngine gasEngine {mpg: 25, 15, owner{"Max"}}
    gasEngine.gallons = 20
    // printing the owner name
    fmt.Println(myEngine.ownerInfo.name)

    // anonymous struct
    var myEngine2 = struct {
        mpg uint,
        gallons uint
    }{25, 15} // this cannot be reused, you would have the define it everytime
}
```


You can define methods to a struct like the below example:

```go
package main

type gasEngine struct{
    mpg uint8
    gallons uint8
}

func (e gasEngine) milesLeft() uint8 {
    return e.gallons * e.mpg
}

func main() {
    var myEngine gasEngine {25, 15}
    fmt.Println(myEngine.milesleft()) >> result of 25*15
}
```

### Interfaces

lets say we have the following example:

```go
package main

type electricEngine struct {
    mpkwh uint8
    kwh uint8
}

func (e electricEngine) milesLeft() uint8 {
    return e.kwh * e.mpkwh
}

type gasEngine struct {
    mpg uint8
    gallons uint8
}

func (e gasEngine) milesLeft() uint8 {
    return e.gallons * e.mpg
}

// i define an interface the the function specification
type engine interface {
    milesleft() uint8
}

// the function then takes the interface engine as e => meaning i can pass an engine into the function
func canMakeIt(e engine, miles uint8) {
    if miles <= e.mileslef() {
        fmt.Printf("Can make it there!")
    }else{
        fmt.Println("Need to fuel!")
    }
}


func main() {
    var myEngine electricEngine = electricEngine{25, 15}
    canMakeIt(myEngine, 50)
}
```


## 10 - Pointers

pointers are a special type which stores memory locations instead of actual values

### Creating a pointer
in the example below we create the variable p which stores the memory address for an int32 value
this means that the pointer can point to the actual data created and not a copy of that data:
```go
package main

func main() {
    var p *int32
}
```

### Dereferencing the pointer
If you want to print out the memory location you can use dereferencing the pointer

```go
package main
import "fmt"

func main() {
    var p *int32

    fmt.Printf("The value p points to is: %v", *p)
}
```


### Setting and changing the value of the pointer

```go
package main
import "fmt"

func main() {
    var p *int32 = new(int32) // setting the value to default => 0
    fmt.Printf("The value p points to is: %v", *p)

    *p = 10 // changing the value

}
```
in the example above we see that the * syntaxt is using two seperate roles.
In the first line of the main function we are using the * syntax to _initialize_ a _pointer_.
But on line 2 and 4 of the main function we are using it to _referencing_ the _value_ of the pointer


### Experiencing the **_Nil Pointer Exception_**

in the example below we don't initialize the pointer with a value
and when we then try to dereferencing the value on line 2 of the main function
we will get at nil pointer exception.
```go
package main
import "fmt"

func main() {
    var p *int32 // = new(int32) <= in this example we are not initializing the pointer
    fmt.Printf("The value p points to is: %v", *p) // <= then this line will throw the nil pointer exception 
}
```

**so make sure the pointers are not nil before trying to assing values to them.**

### Creating a pointer from the address of another variable

You can create a pointer from the address of another variable by using the & sign.


In the example below we initialize the pointer p and then we declare the variable
i as an int32. We then set p to refer to the memory address of i.

This means that both p and i reference the same in32 value.

and if we then use the * notation to change the value of p => we then also change
the value of i
```go
package main

func main() {
    var p *int32 = new(int32)
    var i int32

    p = &i // p is now referencing the memory address of i

    *p = 10 // since p is referencing the same memory address as i, both i and p will be changed by this.
}
```

This behavior is not the same as if you try to assign one variable to another 
variable like the example below. Here we are just copying the the value of k
into i. Whereas with pointer we are changing the original value.
```go
package main

func main() {
    var i int32
    var k int32 = 2

    i = k
}
```

_The only place where this copy logic is not working like the example above is with_
_slices. So if you cope a slice into another slice and then changing one value in the_
_new slice. Then the original slice will also be changed. The reason for this is that_
_under the hood Slices are just using pointers._



### Pointer Example with Functions

In the example below we see that when we pass thing1 to the square function we get a result.
But we also see that the value of thing1 is actually copied into the square function instead
of the value actually being passed into the function. We see this by looking at the memory
addresses. thing1 address is different from thing2 address. And when we print thing1 values
at the end of the main function we see that the actual values of thing1 has not been changed.

This makes our program less efficient since we are using alot more memory because of the copying
instead of just working with the actual values.
```go
package main
import "fmt"

func main() {
    var thing1 = [5]float64{1,2,3,4,5}
    fmt.Printf("\nThe memory location of thing1 is: %v", &thing1) // >> memory location: 0x1400001e240
    var result [5]float64 = square(thing1)
    fmt.Printf("\nThe result is: %v", result) // >> result = [1 4 9 16 25]
    fmt.Printf("\nThe value of thing1: ", thing1) // >> thing1 vvalue = [1 2 3 4 5]
}

func square(thing2 [5]float64) [5]float64 {
    fmt.Printf("\nThe memory location of thing2 is: %v", &thing2) // >> memory location: 0x1400001e2a0
    for i := range thing2 {
        thing2[i] = thing2[i] * thing2[i]
    }
    return thing2
}
```

We can change this behavior and make our program more efficient by just changing the
sqaure function to accepts a pointer to an array as parameter by adding the * in the
parameters for the Square function. And then instead of passing thing1 we pass &thing1

We now see that the the memory locations of thing1 and thing1 is the same. But we also see
that the values of thing1 has been changed.

```go
package main
import "fmt"

func main() {
    var thing1 = [5]float64{1,2,3,4,5}
    fmt.Printf("\nThe memory location of thing1 is: %v", &thing1) // >> memory location: 0x1400001e240
    var result [5]float64 = square(&thing1)
    fmt.Printf("\nThe result is: %v", result) // >> result = [1 4 9 16 25]
    fmt.Printf("\nThe value of thing1: ", thing1) // >> thing1 vvalue = [1 4 9 16 25] <= value has changed
}

func square(thing2 *[5]float64) [5]float64 {
    fmt.Printf("\nThe memory location of thing2 is: %v", &thing2) // >> memory location: 0x1400001e240 <= same as thing1
    for i := range thing2 {
        thing2[i] = thing2[i] * thing2[i]
    }
    return thing2
}
```


This i very usefull when passing large amount of data into functions so you don't have
to make copies wasting memory.


## 11 - Goroutines

we can make our tasks in our program run concurrent which may not be the same as
parallel execution.

So if we have a single cpu core available and we have two main tasks with a couple of sub
tasks, we can make our program handle the two main tasks concurrently. But in this case
it will only handle one sub task at the time. It mights switch between the sub tasks in the two
main tasks, so if it is waiting for an external call in Main task 1, it might handle a sub task
from Main task 2 i the mean time.

If we then have 2 cpu cores we can handle our program concurrently but execute our main tasks
in parallel execution. This means that we can make each cpu core handle each Main task on the
same time.



In the example below we are simulating a db call and each call the db will need
to finish before we can continue to the next call.
```go
package main
import (
    "fmt"
    "math/rand"
    "time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
    t0 := time.Now()
    for i := 0; i < len(dbData); i++ {
        dbCall(i)
    }
    fmt.Printf("\nTotal execution time: %v", time.Since(t0))
}

func dbCall(i int) {
    var delay int32 = rand.Float32()*2000
    time.Sleep(time.Duration(delay) * time.Milliseconds)
    fmt.Println("\nThe result from the database is:", dbData[i])
}
```


We can modify the example above to handle the executions concurrently
So basicly we are creating a counter and adding 1 to the counter foreach time
we start a task. When the task is complete it will decrement the counter.
And then we are just waiting for the counter to reach 0, so we know that
all the tasks has completed.
```go
package main
import (
    "fmt"
    "math/rand"
    "time"
    "sync"
)

var wg = sync.WaitGroup{} // we are creating a new WaitGroup
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
    t0 := time.Now()
    for i := 0; i < len(dbData); i++ {
        wg.Add(1) // here we are incrementing the WaitGroup counter
        go dbCall(i) // we are adding "go" to utilize goroutines
    }
    wg.Wait() // and here we are simply waiting for the counter to reach 0
    fmt.Printf("\nTotal execution time: %v", time.Since(t0))
}

func dbCall(i int) {
    var delay int32 = rand.Float32()*2000
    time.Sleep(time.Duration(delay) * time.Milliseconds)
    fmt.Println("\nThe result from the database is:", dbData[i])
    wg.Done() // here we are decrementing the WaitGroup counter
}
```

Now if we wanted to send the result of the function back to the main function instead
of just printing them out. One thing to be aware of is that when you are working with data
concurrently you might have two tasks trying to modify the same memory location which will
result in corruption of the data. So to avoid this we can use Mutex to make the data handling
safe in a concurrent program.

The placement of the Lock is very important for the concurrency of your program.

```go
package main
import (
    "fmt"
    "math/rand"
    "time"
    "sync"
)

var m = sync.Mutex{} // here we are adding the Mutex
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{} // we are declaring a result variable to append out data to

func main() {
    t0 := time.Now()
    for i := 0; i < len(dbData); i++ {
        wg.Add(1)
        go dbCall(i)
    }
    wg.Wait()
    fmt.Printf("\nTotal execution time: %v", time.Since(t0))
}

func dbCall(i int) {
    var delay int32 = rand.Float32()*2000
    time.Sleep(time.Duration(delay) * time.Milliseconds)
    fmt.Println("\nThe result from the database is:", dbData[i])
    m.Lock() // the lock will check to see if a lock has been added by a diffrent routine. If so it will wait here until the lock is gone
    results = append(results, dbData[i]) // we are appending the dbData to the results variable
    m.UnLock() // once the data has been appended then memory location will be unlocked.
    wg.Done()
}
```

Like the Lock we also have RLock and RUnlock these handles that multiple goroutines
may be able to read from the same memory location. So if the data is only locked by
a RLock, then other routines will be able to read from the memory location whereas with a full lock
other routines will not be able to read the data aswell.

To use RLocks you need to use the RWMutex{} instead.


## 12 - Channels

Channels are a way to pass around data between your goroutines
Channels holde data, are thread safe and can listen for data.

we can create a channel and add data to it like below:
```go
func main() {
    var c = make(chan int) // we create channel c and initialize it to hand ints
    c <- 1 // we then use this special syntax to add data to the channel
}
```

We can think of channels as an underlying array c: [1] in the example above

So we can take the value from the channel and store it in a variable like in the example below.
```go
func main() {
    var c = make(chan int)
    c <- 1

    var i  = <- c // here we pop out the data from the "array" and pass it into the variable i
}
```

Now the two examples above is not showcasing the correct way of using channels and if you run
the code the compiler will throw a deadlock exception.

To utilize the channels correctly we can look at the example below.

```go
import "fmt"

func main() {
    var c = make(chan int)
    go process(c) // executing the function within a goroutine
    fmt.Printf(<-c) // popping the value out of the channel and printin directly

}

// we create a function which adds 123 to the channel
func process(c chan int) {
    c <- 123
}
```

Now we can also utilize loops within channels

```go
import "fmt"

func main() {
    var c = make(chan int)
    go process(c) 
    // because of the range c it will wait until a value has been added to the channel, and then run a loop.
    for i := range c {
        fmt.Printf(<-c)
    }
}

func process(c chan int) {
    // The defer is simply just to ensure that the close(c) will always be called, like a "finally"
    defer close(c) // we need to close the channel to let other routines know. Otherwise we would have the deadlock
    for i := 0; i < 5; i++ {
        c <- i
    }
}
```

In all of the example above we are working with unbuffered channels, meaning that the channel
can only hold a single value. We can also create buffered channels which can hold multiple values.

In the example below. If we were not creating a buffered channel, because we make the main function
sleep for 1 second, we actually also make the process function wait for 1 second, since it cannot add
a new value to the channel before the value has been processed by the main function. But by making the
channel buffered we can now let the process function work independant of the main function.
```go
import (
    "fmt"
    "time"
)

func main() {
    var c = make(chan int 5) // by adding 5 we are creating a buffered channel
    go process(c) 
    for i := range c {
        fmt.Printf(<-c)
        time.Sleep(time.Seconds*1) // here we making the main function wait 1 second.
    }
}

func process(c chan int) {
    defer close(c)
    for i := 0; i < 5; i++ {
        c <- i
    }
}
```

## 13 - Generics

You can use generics in your functions like the example below:

when defining a function which accepts a Generic you need to define which types could be used
ahnd then you seperate them with |. you also need to define the return type as T
```go
func main() {
    var intSlice = []int{1,2,3,4,5}
    fmt.Println(sumSlice[int](intSlice)) // calling the sumSlice function and defining that we are using int type

    var float32Slice = []float32{1,2,3,4,5}
    fmt.Println(sumSlice[float32](float32Slice)) // calling the sumSlice function and defining that we are using int type
}

func sumSlice[T int | float32 | float64](slice[] T) T {
    var sum T
    for _,v := range slice { // the _ serves as a blank identifier for a value which should be discarded
        sum += v
    }
    return sum
}
```

You can also use Generics within structs. We are defining the struct
to accept either a gasEngine or an electricEngine. And when we are declaring the 
structs we are then definining withing [] which type we are using.


```go
type gasEngine struct {
    gallons float32
    mpg float32
}

type electricEngine struct {
    kwh float32
    mpkwh float32
}

type car[T gasEngine | electricEngine] struct {
    carMake string
    carModel string
    engine T
}

func main() {
    var gasCar = car[gasEngine] {
        carMake = "Honda"
        carModel = "Civic"
        engine = gasEngine {
            gallons = 12.4
            mpg = 40
        }
    }

    var electricCar = car[electricEngine] {
        carMake = "Tesla"
        carModel = "Model 3"
        engine = electricEngine {
            kwh = 57.5
            mpkwh = 4.17
        }
    }
}
```




