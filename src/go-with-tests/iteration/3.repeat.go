package iteration

func Repeat(text string) (output string) {
	for i := 0; i < 5; i++ {
		output += text
	}
	return
}
