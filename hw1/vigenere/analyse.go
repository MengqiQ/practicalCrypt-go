package vigenere

func Analyse(ciphertext string, keyLen int) string {
	cipherbytes := []byte(ciphertext)
	res := make([]byte, keyLen)
	for i := 0; i < keyLen; i++ {
		differ := byte(65)
		count := make([]int, 26)
		max := 0
		for j := i; j <= len(cipherbytes)-keyLen; j += keyLen {
			tmp := cipherbytes[j] - 65
			count[tmp]++
			if count[tmp] > max {
				max = count[tmp]
				differ = cipherbytes[j]
			}
		}
		//最多的，应为plaintext的e : 69
		if differ >= 69 {
			res[i] = differ - 4
		} else {
			res[i] = differ + 22
		}
	}

	return string(res)
}

//func getMax(cipherbytes []byte) byte{
//	count := make([]int, 26)
//	for i := 0; i < len(cipherbytes); i++ {
//		count[cipherbytes[i] - 65]++
//	}
//	max := 0
//	maxByte := byte(0)
//	for i := 0; i < 26; i++ {
//		if max < count[i] {
//			maxByte = byte(i + 65)
//			max = count[i]
//		}
//	}
//	return maxByte
//}
