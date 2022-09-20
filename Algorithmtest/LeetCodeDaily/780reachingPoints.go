package main

import "fmt"

func reachingPoints(sx, sy, tx, ty int) bool {
	for tx > sx && ty > sy && tx != ty {
		if tx > ty {
			tx %= ty
		} else {
			ty %= tx
		}
	}
	switch {
	case tx == sx && ty == sy:
		return true
	case tx == sx:
		return ty > sy && (ty - sy) % tx == 0
	case ty == sy:
		return tx > sx && (tx - sx) % ty == 0
	default:
		return false
	}
}

func main() {
	res := reachingPoints(1,1,3,5)
	fmt.Println(res)
}
