package core

import "fmt"

func Contains[T comparable](slice []T, value T) bool {
	for _, element := range slice {
		if element == value {
			return true
		}
	}

	return false
}

func CategoryFromString(value string) (Category, error) {
	switch value {
	case string(SPORT):
		return SPORT, nil
	case string(FINANCE):
		return FINANCE, nil
	case string(MOVIES):
		return MOVIES, nil
	default:
		return Category(""), fmt.Errorf("unknown %s", value)
	}
}
