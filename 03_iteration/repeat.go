package iteration

// const repeatCount = 5

func Repeat(character string, repetition int) string {
	var repeated string
	for i := 0; i < repetition; i++ {
		repeated += character
	}
	return repeated
}