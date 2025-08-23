package main

// Euclidean Algorithm - to calculates Greatest Common Divisor(GCD) of two integers
func gcd(num1, num2 int) int {
	if num2 == 0 {
		// Base case: if the second number is 0, the GCD is num1
		return num1
	} else {
		// Recursive case: GCD of (num2, num1 % num2)
		return gcd(num2, num1%num2)
	}
}

// gcdOfStrings finds the largest string that can divide both str1 and str2.
func gcdOfStrings(str1 string, str2 string) string {
	// If concatenating str1+str2 is not equal to str2+str1,
	// it means they don't have a common divisor string.
	if str1+str2 != str2+str1 {
		return ""
	} else {
		// Find the GCD of the lengths of str1 and str2
		gcdIndex := gcd(len(str1), len(str2))

		// The GCD string will be the prefix of str1 with length gcdIndex
		return str1[0:gcdIndex]
	}
}

/*
func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))   // Output: "ABC"
	fmt.Println(gcdOfStrings("ABABAB", "AB"))    // Output: "AB"
	fmt.Println(gcdOfStrings("LEET", "CODE"))    // Output: ""
}
*/
