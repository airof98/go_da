package prime

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPrime(a int) bool {
	if a <= 1 {
		return false
	}

	aa := int(math.Sqrt(float64(a)))
	for i := 2; i <= aa; i++ {
		if a%i == 0 {
			return false
		}
	}

	return true
}

func makeNumbers(coms, others string, numberSet map[int]bool) {
	if len(coms) != 0 {
		a, _ := strconv.Atoi(coms)
		numberSet[a] = true
	}

	for i, c := range others {
		makeNumbers(fmt.Sprintf("%s%c", coms, c), others[0:i]+others[i+1:], numberSet)
	}
}

func TestPrimeNumbers(t *testing.T) {
	assert.False(t, isPrime(91))
	assert.True(t, isPrime(71))

	others := "173"
	numberSet := make(map[int]bool, 0)
	makeNumbers("", others, numberSet)
	for k, _ := range numberSet {
		if isPrime(k) {
			fmt.Println(k)
		}
	}
}
