package main

import (
	"fmt"
)

func main() {

	//Variable declaration
	var sTring string = "Hello World"
	var a, b, c, d int64 = 42, 30, 28, 76
	fmt.Printf(" a: %v#\n b: %v\n c: %T\n d: %v\n", a, b, c, d)
	fmt.Println(sTring)
	fmt.Println("Hello World!")

	//Format specifiers
	var i = 3.141

	fmt.Printf("%e\n", i)
	fmt.Printf("%f\n", i)
	fmt.Printf("%.2f\n", i)
	fmt.Printf("%6.2f\n", i)
	fmt.Printf("%g\n", i)

	//t specifier for boolean evaluation
	var t = true
	var j = false

	fmt.Printf("%t\n", t)
	fmt.Printf("%t\n", j)

	var txt = "Hello"

	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)

	var k = 15

	fmt.Printf("%b\n", k)
	fmt.Printf("%d\n", k)
	fmt.Printf("%+d\n", k)
	fmt.Printf("%o\n", k)
	fmt.Printf("%O\n", k)
	fmt.Printf("%x\n", k)
	fmt.Printf("%X\n", k)
	fmt.Printf("%#x\n", k)
	fmt.Printf("%4d\n", k)
	fmt.Printf("%-4d\n", k)
	fmt.Printf("%04d\n", k)
	/*
		Go has three general data types: (bool, numeric and string value)
	*/
	//Here are ways to declare boolean variables
	var b1 bool = true
	var b2 = true
	var b3 bool
	b4 := true

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)

	//Integers
	//Unsigned (positive values only)
	//Signed integers take both negative and positive values.
	var x uint = 500
	var y uint = 4500
	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v\n", y, y)

	//Array initialization
	//Recall, that all variables in golang are initialised
	arr1 := [5]int{1: 10, 2: 40, 0: 1400}
	fmt.Printf("Type: %T Value: %v length: %v\n", arr1, arr1, len(arr1))

	/*Go slices
	Unlike arrays, slices can grow or shrink.
	They store data of the same type just like arrays
	Here are ways to create a slice:
	1. Using the []datatype{values} format
	2. Create a slice from an array
	3. Using the make() function
	*/

	//1. Using the []datatype{values} format
	myslice := []int{1, 2, 3} //has a capacity of 3
	/*
		In golang, there are two function that can be used to return the capacity of a slice:
		1. len()function - returns the length of the slice
		2. cap()function - returns the capacity of the slice(the
		number of elements the slice can grow or shrink to)
	*/
	fmt.Println(myslice)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))

	/*
		2. Create a slice by slicing an array
		syntax:
		1. var myarray = [length]datatype{values}
		myslice := myarray[start:end]
	*/
	arr3 := [6]int{10, 11, 12, 13, 14, 15}
	myslice2 := arr3[2:4]

	fmt.Printf("myslice = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))

	/*
		3. Create a slice with the make() function
		Syntax:
		slice_name := make([]type, length, capacity)
		If the capacity parameter is not defined, it will be equal to length.
	*/
	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	//with omitted capacity
	myslice3 := make([]int, 5)
	fmt.Printf("myslice1 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))

	/*
		To add an element in a list use the index then assign the new value.
		append() function is used to add a value at the end of a slice
		syntax:
		slice3 = append(slice1, slice2...)
		The '...' after slice2 is necessary when appending the elements of one lice to another.

		You can alter the size of a slice by slicing the slice, or by appending new elements to it.
	*/

	//How to append one slice to another slice:
	my_slice1 := []int{1, 2, 3}
	my_slice2 := []int{4, 5, 6}
	my_slice3 := append(my_slice1, my_slice2...)

	fmt.Printf("myslice3=%v\n", my_slice3)
	fmt.Printf("length= %d\n", len(my_slice3))
	fmt.Printf("capacity=%d\n", cap(my_slice3))

	/*
		Memory Efficiency:
		If an array is large enough, and you do not intend to use all the elements, it is best practice to make a copy of a relatively smaller program to minimise the amount of memory used by the program.
		The copy function is used to create a new underlying array with only the required elements for the slice.
		copy syntax
		copy(destination, source) and returns the number of elements copied
	*/
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))

	//--Operators In Golang--
	/*
		1. Arithmetic opearators
		2. Assignment operators
		3. Comparison operators
		4. Logical operators
		5. Bitwise operators

		Go has the following conditional statements
		if, else, else if, switch
		note that else {}should be on the same line
		i.e: } else {
			//not any other way
		}
		//Switch statement;
		switch expression {
		case x:
			//code block
		case y:
			//code block
		case z:
			...
		default:
			//code block
		}
	*/

	//Single-Case Switch Example
	day := 4

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}

	//The default keyword is included as the last statement to specify wht code should run if there is no case match
	/*
		The Multiple-case switch statement
		Syntax:
		switch expression {
		case x, y:
			//code block if expression is evaluated to x or y
		case v, w:
			//code block if expression evaluates to v or w
		case z:
			...
		default:
			//code block if expression is not found in any cases
		}
	*/
	day1 := 5

	switch day1 {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")

	}

	/*
		The for loop can take up to three statements:
		for statement1;statement2; statement3 {
			//code to be executed for each iteration
		}
	*/

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for j := 0; j <= 100; j += 10 {
		fmt.Println(j)
	}

	//The continue statement is used to skip one or more iterations in the loop, It then continues with the next iteration in the loop.
	for k := 0; k < 5; k++ {
		if k == 3 {
			continue
		}
		fmt.Println(i)
	}

	//The break statement breaks out of the loop when i is equal to 3:

	for l := 0; l < 5; l++ {
		if l == 3 {
			break
		}
		fmt.Println(l)
	}
	//Note that break and continue statements are always used with condition.

	//Nested loops
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for t := 0; t < len(adj); t++ {
		for w := 0; w < len(fruits); w++ {
			fmt.Println(adj[t], fruits[w])
		}
	}
	/*The Range Keyword is used to easily iterate over an array, slice or map. It returns both the index and the value.
	The range keyword is used like this:
	for index, value := array|slice|map {
		//code to be executed for each iteration
	}*/

	fruits2 := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits2 {
		fmt.Printf("%v\t%v\n", idx, val)
	}
	//We want to omit the indices(idx stores the index, val stores the value)
	fruits3 := [3]string{"apple", "orange", "banana"}
	for _, val := range fruits3 {
		fmt.Printf("%v\n", val)
	} //The same will happen when we wish to omit the value and print indices

	/*
		####--- GO FUNCTIONS ---####
		A function is a block of statements that can be used repeatedly in a program.
		A function will not execute automatically when a page loads
		A  function will be executed by a call to the function.

		--Create a Function--
		Use the func keyword
		specify a name for the function followed by parentheses()
		Add code that defines what the function should do, inside the curly braces {}
		Basic syntax:
		func FunctionName() {
			//code to be executed
		}

		Naming Rules for Go Functions
		1. Must start with a letter
		2. Can only contain only alpha-numeric characters and underscores
		3. They are case sensitive
		4. It cannot contain spaces
		5. Use multi-word variable naming if function name consists of multiple words.
		6. Give the function a name that regflects what the function does.
	*/

	return
}
