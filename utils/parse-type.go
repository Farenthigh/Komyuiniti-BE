package utils

import "fmt"

func StringToUint(s string) (uint, error) {
	var i uint
	_, err := fmt.Sscanf(s, "%d", &i)
	if err != nil {
		return 0, err
	}
	return i, nil
}
