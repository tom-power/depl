package depl

import "sort"

func sorted(input []string) []string {
	sorted := make([]string, len(input))
	copy(sorted, input)
	sort.Strings(sorted)
	return sorted
}

func keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Contains(args []string, target string) bool {
	for _, arg := range args {
		if arg == target {
			return true
		}
	}
	return false
}
