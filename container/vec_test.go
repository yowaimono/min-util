package container

import (
    "testing"
    "reflect"
)

func TestVec(t *testing.T) {
    // Test NewVec
    vec := NewVec([]int{1, 2, 3})
    if vec.Len()!= 3 {
        t.Errorf("Expected length 3, got %d", vec.Len())
    }

    // Test Map
    mappedVec := vec.Map(func(i int) int { return i * 2 })
    expectedMappedData := []int{2, 4, 6}
    if!reflect.DeepEqual(mappedVec.Container(), expectedMappedData) {
        t.Errorf("Expected mapped data %v, got %v", expectedMappedData, mappedVec.Container())
    }

    // Test Filter
    filteredVec := vec.Filter(func(i int) bool { return i%2 == 0 })
    expectedFilteredData := []int{2}
    if!reflect.DeepEqual(filteredVec.Container(), expectedFilteredData) {
        t.Errorf("Expected filtered data %v, got %v", expectedFilteredData, filteredVec.Container())
    }

    // Test ForEach
    forEachCount := 0
    vec.ForEach(func(i int) { forEachCount++ })
    if forEachCount!= 3 {
        t.Errorf("Expected ForEach to be called 3 times, got %d", forEachCount)
    }

    // Test Contains
    if!vec.Contains(2) {
        t.Errorf("Expected vec to contain 2")
    }

    // Test Len
    if vec.Len()!= 3 {
        t.Errorf("Expected length 3, got %d", vec.Len())
    }

    // Test Container
    if!reflect.DeepEqual(vec.Container(), []int{1, 2, 3}) {
        t.Errorf("Expected container data %v, got %v", []int{1, 2, 3}, vec.Container())
    }

    // Test Reduce
    reducedValue := vec.Reduce(func(a int, b int) int { return a + b }, 0)
    expectedReducedValue := 6
    if reducedValue!= expectedReducedValue {
        t.Errorf("Expected reduced value %d, got %d", expectedReducedValue, reducedValue)
    }

    // Test FlatMap
    flatMappedVec := vec.FlatMap(func(i int) []int { return []int{i, i * 2} })
    expectedFlatMappedData := []int{1, 2, 2, 4, 3, 6}
    if!reflect.DeepEqual(flatMappedVec.Container(), expectedFlatMappedData) {
        t.Errorf("Expected flat mapped data %v, got %v", expectedFlatMappedData, flatMappedVec.Container())
    }

    // Test GroupBy
    groupedMap := vec.GroupBy(func(i int) string { return "Group" })
    expectedGroupedMap := map[string][]int{"Group": {1, 2, 3}}
    if!reflect.DeepEqual(groupedMap, expectedGroupedMap) {
        t.Errorf("Expected grouped map %v, got %v", expectedGroupedMap, groupedMap)
    }

    // Test Distinct
    distinctVec := vec.Distinct()
    expectedDistinctData := []int{1, 2, 3}
    if!reflect.DeepEqual(distinctVec.Container(), expectedDistinctData) {
        t.Errorf("Expected distinct data %v, got %v", expectedDistinctData, distinctVec.Container())
    }

    // Test Take
    takenVec := vec.Take(2)
    expectedTakenData := []int{1, 2}
    if!reflect.DeepEqual(takenVec.Container(), expectedTakenData) {
        t.Errorf("Expected taken data %v, got %v", expectedTakenData, takenVec.Container())
    }

    // Test Skip
    skippedVec := vec.Skip(1)
    expectedSkippedData := []int{2, 3}
    if!reflect.DeepEqual(skippedVec.Container(), expectedSkippedData) {
        t.Errorf("Expected skipped data %v, got %v", expectedSkippedData, skippedVec.Container())
    }

    // Test Reverse
    reversedVec := vec.Reverse()
    expectedReversedData := []int{3, 2, 1}
    if!reflect.DeepEqual(reversedVec.Container(), expectedReversedData) {
        t.Errorf("Expected reversed data %v, got %v", expectedReversedData, reversedVec.Container())
    }
}
