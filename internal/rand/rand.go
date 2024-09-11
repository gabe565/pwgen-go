package rand

func IntN(n int) int {
	return globalRand.IntN(n)
}

func ShuffleSlice[T any](s []T) []T {
	globalRand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	return s
}

func Numeric(n int) string {
	return FromString("0123456789", n)
}

func Alpha(n int) string {
	return FromString("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", n)
}

func AlphaNum(n int) string {
	return FromString("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", n)
}

func ASCII(n int) string {
	const minByte, maxByte = ' ', '~'
	result := make([]byte, 0, n)
	for range n {
		result = append(result, byte(globalRand.IntN(maxByte+1-minByte)+minByte))
	}
	return string(result)
}

func FromString(letters string, n int) string {
	result := make([]byte, 0, n)
	for range n {
		result = append(result, letters[globalRand.IntN(len(letters))])
	}
	return string(result)
}
