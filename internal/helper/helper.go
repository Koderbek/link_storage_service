package helper

import "strings"

const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base     = uint(len(alphabet))
)

func IdToCode(id uint) string {
	if id == 0 {
		return string(alphabet[0])
	}

	var sb strings.Builder
	for id > 0 {
		rem := id % base
		sb.WriteByte(alphabet[rem])
		id = id / base
	}

	return reverse(sb.String())
}

func CodeToId(code string) uint {
	var id uint
	for _, char := range code {
		index := strings.IndexRune(alphabet, char)
		id = id*base + uint(index)
	}
	return id
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
