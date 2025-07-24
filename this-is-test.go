package main

import "fmt"

func fail() {
    // This code has an obvious error - trying to divide by zero
    x := 10
    y := 0
    result := x / y  // Division by zero error!
    
    fmt.Printf("The result is: %d\n", result)
    
    // Another error - using an undefined variable
    fmt.Println("Value of z is:", z)  // z is not declared
}
