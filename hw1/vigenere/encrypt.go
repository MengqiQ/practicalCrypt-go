package vigenere

func Encipher(key string, p string) string {
	cipherBytes := []byte(p)
	for i := 0; i < len(p); i++ {
		if 97 <= cipherBytes[i] && cipherBytes[i] < 123 {
			cipherBytes[i] -= 32
		}
		if 65 <= cipherBytes[i] && cipherBytes[i] < 91 {
			cipherBytes[i] = (p[i]+key[i%len(key)]-130)%26 + 65
		}
	}
	c := string(cipherBytes)
	return c
}
