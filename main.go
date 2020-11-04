package main

func initialize(length int) string {
	result := ""
	i := 0
	for i < length {
		result = result + "#"
		i = i + 1
	}

	return result
}
