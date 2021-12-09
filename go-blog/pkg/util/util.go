package util

import "strconv"

func Convertoi(s string) (i int, err error) {
	i, err = strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return

}
