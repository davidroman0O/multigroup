package multigroup

import (
	"reflect"
	"strconv"
	"testing"
)

// Person is a simple struct for testing purposes.
type Person struct {
	Name    string
	Country string
	Age     int
}

// TestMultiGroupBy tests the MultiGroupBy function with a slice of Person.
func TestMultiGroupBy(t *testing.T) {
	// Prepare the data.
	people := []Person{
		{Name: "Alice", Country: "USA", Age: 30},
		{Name: "Bob", Country: "USA", Age: 30},
		{Name: "Charlie", Country: "USA", Age: 25},
		{Name: "Diana", Country: "UK", Age: 30},
	}

	// Define the selectors for grouping.
	countrySelector := func(p Person) (string, string) { return "Country", p.Country }
	ageSelector := func(p Person) (string, string) { return "Age", strconv.Itoa(p.Age) }

	// Perform the grouping.
	groups := By(people, countrySelector, ageSelector)

	// Define the expected groups, making sure to convert the Age to a string.
	expectedGroups := []Group[string, Person]{
		{
			Keys: []KeyValue[string]{
				{Key: "Country", Value: "USA"},
				{Key: "Age", Value: "30"},
			},
			Items: []Person{people[0], people[1]}, // Alice and Bob
		},
		{
			Keys: []KeyValue[string]{
				{Key: "Country", Value: "USA"},
				{Key: "Age", Value: "25"},
			},
			Items: []Person{people[2]}, // Charlie
		},
		{
			Keys: []KeyValue[string]{
				{Key: "Country", Value: "UK"},
				{Key: "Age", Value: "30"},
			},
			Items: []Person{people[3]}, // Diana
		},
	}

	// Check if the groups match the expected groups.
	if !reflect.DeepEqual(groups, expectedGroups) {
		t.Errorf("MultiGroupBy() got = %v, want %v", groups, expectedGroups)
	}
}
