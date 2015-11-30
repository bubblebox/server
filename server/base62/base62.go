package base62

import (
	"bytes"
	"math"
)

var (
	chars = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	base  = uint64(62)
)

// Encode takes an integer and returns the Base62 encoded string representation.
func Encode(value uint64) string {
	var buffer []byte

	val, remainder := value, uint64(0)

	for val >= base {
		val, remainder = divmod(val, base)
		buffer = append([]byte{chars[remainder]}, buffer...)
	}
	buffer = append([]byte{chars[val]}, buffer...)

	return string(buffer)
}

// Decode takes a Base62 encoded string and returns the integer representation.
func Decode(code string) uint64 {
	value := uint64(0)

	// Convert and reverse code
	bcode := []byte(code)
	for i, j := 0, len(bcode)-1; i < j; i, j = i+1, j-1 {
		bcode[i], bcode[j] = bcode[j], bcode[i]
	}

	// Calculate integer value
	for i, b := range bcode {
		mult := uint64(math.Pow(float64(base), float64(i)))
		value += uint64(bytes.IndexByte(chars, b)) * mult
	}

	return value
}

func divmod(val uint64, base uint64) (uint64, uint64) {
	return val / base, val % base
}
