package main

func getKeysFromMap[T comparable, U comparable](m map[T]U) []T {
	var result []T
	for key, _ := range m {
		result = append(result, key)
	}
	return result
}
