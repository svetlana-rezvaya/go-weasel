package main

func initialize(length int) string {
	result := ""
	i := 0
	for i < length {
		character := '#'
		result = result + string(character)
		i = i + 1
	}

	return result
}
