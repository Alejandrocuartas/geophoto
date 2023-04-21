package helpers

import "strconv"

func ParseStringToFloat(str string) (float64, error) {
	value, e := strconv.ParseFloat(str, 64)
	if e != nil {
		return 0, e
	}
	return value, nil
}
