package teststuff

import "crypto/rand"

func Add(a, b int) int {
	return a + b
}

func AddUint64(a, b uint64) uint64 {
	return a + b
}

func RandomBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}
