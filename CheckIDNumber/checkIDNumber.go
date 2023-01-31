package checkidnumber

import (
	"math"
)

func CheckIDNumber(id string) bool {
	idNumber := id[:17]
	idCheck := []byte(id[len(id)-1 : ])
	sum := 0
	for index := 0; index < 17; index++ {
		sum += int(math.Pow(2, float64(17-index))) % 11 * int(idNumber[index]-'0')
	}
	num := (12 - (sum % 11)) % 11
	if num < 10 {
		return num == int(idCheck[0]-'0')
	} else {
		return idCheck[0] == 'X'
	}

}
