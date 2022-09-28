package tl_test

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dgrr/tl"
)

func ExampleMap() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	numbersToString := tl.Map(numbers, func(number int) string {
		return strconv.Itoa(number)
	})

	fmt.Println(numbers)
	fmt.Println(numbersToString)
}

func ExampleFilter() {
	filterThis := []string{
		"Each type parameter has a type",
		"constraint that acts as a kind of",
		"meta-type for the type parameter",
		"Each type constraint specifies the permissible",
		"type arguments that calling",
		"code can use for the respective",
		"type parameter",
	}

	// filter out the strings containing `type`.
	afterFilter := tl.Filter(filterThis, func(x string) bool {
		return !strings.Contains(x, "type")
	})

	fmt.Println(strings.Join(filterThis, "\n"))
	fmt.Println("--- after ----")
	fmt.Println(strings.Join(afterFilter, "\n"))
}

func ExampleFilterInPlace() {
	filterThis := []string{
		"Each type parameter has a type",
		"constraint that acts as a kind of",
		"meta-type for the type parameter",
		"Each type constraint specifies the permissible",
		"type arguments that calling",
		"code can use for the respective",
		"type parameter",
	}

	// filter out the strings containing `type`.
	filterThis = tl.FilterInPlace(filterThis, func(x string) bool {
		return !strings.Contains(x, "type")
	})

	fmt.Println(strings.Join(filterThis, "\n"))
}

func ExampleJoin() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	b := []int{5, 7, 9, 11, 13}

	joined := tl.Join(a, b)

	fmt.Println("All", joined)
}

func ExampleAntiJoin() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	b := []int{5, 7, 9, 11, 13}

	antijoined := tl.AntiJoin(a, b)

	fmt.Println("All", antijoined)
}

func ExampleMerge() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	b := []int{5, 7, 9, 11, 13}

	merged := tl.Merge(a, b)

	fmt.Println("All", merged)
}
