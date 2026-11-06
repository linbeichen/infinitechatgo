package utils

import (
	"fmt"
	"math"
	"math/rand"
)

func RandamNumber(len int) string {
	max := int(math.Pow10(len))
	return fmt.Sprintf("%0*d", len, rand.Intn(max))
}
