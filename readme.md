# MultiGroup 

The `multigroup` Go module provides a utility function for grouping slices of any type based on multiple key selectors. It's particularly useful when you need to categorize data into nested groups dynamically.

## Installation

```sh
go get github.com/davidroman0O/multigroup
```

## Usage

To use the `MultiGroup` function, you need to provide a slice of elements and a set of key selectors. Each selector function should take an element as input and return a key name and its corresponding value as a string.

Here is a simple example of how to use `MultiGroup`:

```go
package main

import (
	"fmt"
	"strconv"
	"github.com/davidroman0O/multigroup"
)

// Define your data structure.
type Person struct {
	Name    string
	Country string
	Age     int
}

func main() {
	// Prepare your data.
	people := []Person{
		{Name: "Alice", Country: "USA", Age: 30},
		{Name: "Bob", Country: "USA", Age: 30},
		{Name: "Charlie", Country: "USA", Age: 25},
		{Name: "Diana", Country: "UK", Age: 30},
	}

	// Define key selectors.
	countrySelector := func(p Person) (string, string) { return "Country", p.Country }
	ageSelector := func(p Person) (string, string) { return "Age", strconv.Itoa(p.Age) }

	// Group the data.
	groupedPeople := multigroup.By(people, countrySelector, ageSelector)

	// Output the results.
	for _, group := range groupedPeople {
		fmt.Printf("Group: %+v\n", group.Keys)
		for _, person := range group.Items {
			fmt.Printf("  %+v\n", person)
		}
	}
}
```

## License

`multigroup` is open-source software licensed under the MIT license. See the `LICENSE` file for more details.
