package vigenere

import (
	"math"
)

func FindKeyLength(ciphertext string) int {
	cipherbytes := []byte(ciphertext)
	differ := 1.0
	res := 0
	for i := 1; i < 21; i++ {
		ic := 0.0
		count := 0
		for j := 0; j < i; j++ {
			each := make([]int, 26)
			count = 0
			for k := j; k <= len(cipherbytes)-i; k += i {
				each[cipherbytes[k]-65]++
				count++
			}
			ic += icEach(each, count)
		}
		ic /= float64(i)
		//get closed length
		curr := math.Abs(ic - 0.065)
		if curr < differ && math.Abs(curr-differ) > 0.001 {
			differ = curr
			res = i
		}
	}

	return res
}

func icEach(each []int, len int) float64 {
	count := 0.0
	for i := 0; i < 26; i++ {
		count += float64(each[i] * (each[i] - 1))
	}
	count /= float64(len * (len - 1))
	return count
}
