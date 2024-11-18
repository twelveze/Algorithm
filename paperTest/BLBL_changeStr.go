package main

import "fmt"

func changeStr(s1, s2 string) string {
	if len(s1) < len(s2) {
		return ""
	}
	smap := make(map[byte]int, 0)
	for i := 0; i < len(s2); i++ {
		smap[s2[i]] = i
	}
	s1map := make(map[byte]int, 0)
	for i := 0; i < len(s1); i++ {
		s1map[s1[i]] = i
	}
	sByte := []byte(s1)
	for i := 0; i < len(s2); i++ {
		if _, ok := smap[s2[i]]; ok {
			if _, ok := s1map[s2[i]]; !ok {
				continue
			}
			sByte[i], sByte[s1map[s2[i]]] = sByte[s1map[s2[i]]], sByte[i]
			delete(s1map, s2[i])
		}
	}
	res := string(sByte)
	return res
}
func main() {
	var s1 string
	var s2 string
	fmt.Scanln(&s1)
	fmt.Scanln(&s2)
	res := changeStr(s1, s2)
	fmt.Println(res)
}
