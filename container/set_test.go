package container

import (
	"testing"
)

func TestSet(t *testing.T) {
	// Test NewSet
	set1 := NewSet[int]()
	if set1 == nil {
		t.Errorf("NewSet failed to create a set")
	}

	// Test Add
	set1.Add(1)
	set1.Add(2)
	set1.Add(3)
	if set1.Len() != 3 {
		t.Errorf("Add failed to add elements to the set")
	}

	// Test Contains
	if !set1.Contains(2) {
		t.Errorf("Contains failed to check if an element exists in the set")
	}

	// Test Remove
	set1.Remove(2)
	if set1.Contains(2) {
		t.Errorf("Remove failed to remove an element from the set")
	}

	// Test ToSlice
	slice1 := set1.ToSlice()
	if len(slice1) != 2 || slice1[0] != 1 || slice1[1] != 3 {
		t.Errorf("ToSlice failed to convert the set to a slice")
	}

	// Test Union
	set2 := NewSetFromSlice([]int{2, 3, 4})
	unionSet := set1.Union(set2)
	if unionSet.Len() != 4 {
		t.Errorf("Union failed to combine two sets")
	}

	// Test Intersection
	intersectionSet := set1.Intersection(set2)
	if intersectionSet.Len() != 1 {
		t.Errorf("Intersection failed to find the intersection of two sets")
	}

	// Test Difference
	differenceSet := set1.Difference(set2)
	if differenceSet.Len() != 1 {
		t.Errorf("Difference failed to find the difference between two sets")
	}

	// Test ForEach
	var sum int
	set1.ForEach(func(item int) {
		sum += item
	})
	if sum != 4 {
		t.Errorf("ForEach failed to apply a function to each element")
	}

	// Test Map
	mappedSet := set1.Map(func(item int) int {
		return item * 2
	})
	if mappedSet.Contains(2) || !mappedSet.Contains(4) || !mappedSet.Contains(6) {
		t.Errorf("Map failed to apply a function to each element and create a new set")
	}

	// Test Filter
	filteredSet := set1.Filter(func(item int) bool {
		return item%2 == 0
	})
	if filteredSet.Contains(1) || !filteredSet.Contains(3) {
		t.Errorf("Filter failed to filter elements based on a condition")
	}

	// Test Reduce
	reducedValue := set1.Reduce(func(a int, b int) int {
		return a + b
	}, 0)
	if reducedValue != 4 {
		t.Errorf("Reduce failed to reduce elements to a single value")
	}
}
