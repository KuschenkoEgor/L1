package main
//Имеется последовательность строк - (cat, cat, dog, cat, tree)
//создать для нее собственное множество.
import "fmt"

func multiplicity(s []string) []string {
	var result []string
	IndividualMap := make(map[string]bool, len(s))

	for _, val := range s {
		IndividualMap[val] = true
	}

	for k := range IndividualMap {
		result = append(result, k)
	}

	return result
}

func main() {
	A := []string{"cat", "cat", "dog", "cat", "tree"}

	ChangeArray := multiplicity(A)

	fmt.Printf("Default array: %v\nChange Array: %v\n", A, ChangeArray)
}