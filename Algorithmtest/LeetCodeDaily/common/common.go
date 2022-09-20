package common

func Abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
