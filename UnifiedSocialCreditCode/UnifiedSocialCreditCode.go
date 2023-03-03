package unifiedSocialCreditCode

import "math"

func CheckUnifiedSocialCreditCode(code string) bool {
	cc := code[:17]
	check := []byte(code[len(code)-1:])
	sum := 0
	for index := 0; index < 17; index++ {
		sum += int(cc[index]-'0') * (int(math.Pow(3, float64(index))) % 31)
	}
	c18 := 31 - (sum % 31)
	if c18 == int(check[0]-'0') {
		return true
	} else {
		return false
	}
}
