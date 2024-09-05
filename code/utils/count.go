package utils

// counts occurrences of all strings in a slice by converting into map.
func Count_string_slice_occurrences(slice []string) map[string]int {
	count_map := make(map[string]int) // make a map with string key (item) and int as value (count)
	for _, string := range slice {
		count_map[string]++
	}
	return count_map
}
