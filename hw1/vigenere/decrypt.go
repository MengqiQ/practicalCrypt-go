package vigenere

func Decipher(key string, c string) string {
	cipherBytes := []byte(c)
	for i := 0; i < len(c); i++ {
		if 97 <= cipherBytes[i] && cipherBytes[i] < 123 {
			cipherBytes[i] -= 32
		}
		if 65 <= cipherBytes[i] && cipherBytes[i] < 91 {
			cipherBytes[i] = (cipherBytes[i]+26-key[i%len(key)])%26 + 65
			//cipherBytes[i] = cipherBytes[i] + 26 - key[i % len(key)] + 65
		}
	}
	p := string(cipherBytes)
	return p
}
