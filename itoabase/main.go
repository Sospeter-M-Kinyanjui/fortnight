package main

import "fmt"

func main() {
	  // Test case 1: Base 16 (Hexadecimal)
    fmt.Println(ItoaBase(255, 16))  // Expected: "FF"
    fmt.Println(ItoaBase(-255, 16)) // Expected: "-FF"
    
    // Test case 2: Base 2 (Binary)
    fmt.Println(ItoaBase(10, 2))    // Expected: "1010"
    fmt.Println(ItoaBase(-10, 2))   // Expected: "-1010"
    
    // Test case 3: Base 10 (Decimal)
    fmt.Println(ItoaBase(0, 10))    // Expected: "0"
    fmt.Println(ItoaBase(12345, 10)) // Expected: "12345"
    
    // Test case 4: Base 8 (Octal)
    fmt.Println(ItoaBase(63, 8))    // Expected: "77"
    fmt.Println(ItoaBase(-63, 8))   // Expected: "-77"
    
    // Test case 5: Base 3
    fmt.Println(ItoaBase(10, 3))    // Expected: "101"
    fmt.Println(ItoaBase(-10, 3))   // Expected: "-101"
    
    // Test case 6: Base 16 (Large number)
    fmt.Println(ItoaBase(123456789, 16))  // Expected: "75BCD15"
    fmt.Println(ItoaBase(-123456789, 16)) // Expected: "-75BCD15"
}

func ItoaBase(value int, base int) string {
	chars := "0123456789ABCDEF"
	if value == 0 {
		return "0"
	}
	isNeg := false
	if value < 0 {
		isNeg = true
		value = -value
	}

	var result []rune

	for value > 0 {
		remainder := value % base
		result = append([]rune{rune(chars[remainder])}, result...)
		value /= base
	}

	if isNeg {
		result = append([]rune{'-'}, result...)
	}
	return string(result)
}
