package main

func initialize(length int) string {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ "
	result := ""
	i := 0
	for i < length {
		character := alphabet[5]
		result = result + string(character)
		i = i + 1
	}

	return result
}
