package rand

import "strconv"

func ShuffleSlice[T any](s []T) []T {
	Rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	return s
}

func Numeric(n int) string {
	var result string
	for range n {
		result += strconv.Itoa(Rand.IntN(10))
	}
	return result
}

func Alpha(n int) string {
	const bytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	result := make([]byte, 0, n)
	for range n {
		result = append(result, bytes[Rand.IntN(len(bytes))])
	}
	return string(result)
}

func AlphaNum(n int) string {
	const bytes = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	result := make([]byte, 0, n)
	for range n {
		result = append(result, bytes[Rand.IntN(len(bytes))])
	}
	return string(result)
}

func ASCII(n int) string {
	result := make([]byte, 0, n)
	for range n {
		// Generate random bytes between 32 (space) and 126 (~)
		result = append(result, byte(Rand.IntN(127-32)+32))
	}
	return string(result)
}
