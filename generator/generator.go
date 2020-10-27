package generator

import (
	"math/rand"
	"time"
)

func Generate(length ...int) string {
	var arr string

	for i := 0; i < length[0]; i++ {

		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		char := string(r1.Intn(126-33) + 33) // ASCI CHAR NUMBER 33-126
		arr = arr + char
		time.Sleep(time.Millisecond)
	}

	return arr
}
