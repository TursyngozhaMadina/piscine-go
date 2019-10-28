package piscine

func ConcatParams(args []string) string {

	var a string
	var count int = 0

	for range args {
		count++
	}
	for let := range args {
		a += args[let]

		if let != count-1 {
			a += "\n"
		}
	}
	return a
}
