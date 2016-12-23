package main

func main() {

}

// subtract a number by one
func subtractByOne(number int) int {
	return number - 1
}

// add strings to an array
func addStringsToArray(args ...string) []string {

	// make the array to append strings too
	var arr = make([]string, len(args))

	// for each argument, add the string
	for index, value := range args {
		arr[index] = value
	}

	// return the array
	return arr
}
