package percentagerounds

import "strconv"

func GetPercentageRounds(percentage int) string {
	if percentage == 0 {
		return "[=>         ]  " + strconv.Itoa(percentage)+"%"
	} else if percentage <= 10 {
		return "[#=>        ]  " + strconv.Itoa(percentage)+"%"
	} else if percentage <= 20 {
		return "[##=>       ]  " + strconv.Itoa(percentage)+"%"
	} else if percentage <= 30 {
		return "[###=>      ]  " + strconv.Itoa(percentage)+"%"
	} else if percentage <= 40 {
		return "[####=>     ]  " + strconv.Itoa(percentage)+"%"
	} else if percentage <= 50 {
		return "[#####=>    ]  " + strconv.Itoa(percentage)+"%"
	} else if percentage <= 60 {
		return "[######=>   ]  " + strconv.Itoa(percentage)+"%"
	} else if percentage <= 70 {
		return "[#######=>  ]  " + strconv.Itoa(percentage)+"%"
	} else if percentage <= 80 {
		return "[########=> ]  " + strconv.Itoa(percentage)+"%"
	} else if percentage <= 90 {
		return "[#########=>]  " + strconv.Itoa(percentage)+"%"
	} else if percentage < 100 {
		return "[##########=]  " + strconv.Itoa(percentage)+"%"
	} else {
		return "[###########]  " + strconv.Itoa(percentage)+"%"
	}
}

func GetPercentageRoundsSubstring(percentage int) string {
	// symbols:="★★★★★★★★★★☆☆☆☆☆☆☆☆☆☆"
	symbols:="##########          "
	offset:=(100-percentage)/10
	return symbols[offset:offset+(percentage/10)]
}