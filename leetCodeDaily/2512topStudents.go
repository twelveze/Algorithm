package main

import (
	"fmt"
	"sort"
	"strings"
)

func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
	source := make([]int, len(student_id))

	positiveMap := make(map[string]struct{}, 0)
	negativeMap := make(map[string]struct{}, 0)
	type pair struct {
		source int
		id     int
	}
	sortSource := make([]pair, len(student_id))
	//初始化positive map
	for _, positive := range positive_feedback {
		positiveMap[positive] = struct{}{}
	}
	//初始化negative map
	for _, negative := range negative_feedback {
		negativeMap[negative] = struct{}{}
	}
	for index, studentReport := range report {
		reportArray := strings.Split(studentReport, " ")
		for _, word := range reportArray {
			if _, exist := positiveMap[word]; exist {
				source[index] += 3
			}
			if _, exist := negativeMap[word]; exist {
				source[index] -= 1
			}
		}
		sortSource[index] = pair{source[index], student_id[index]}
	}
	sort.Slice(sortSource, func(i, j int) bool {
		return sortSource[i].source > sortSource[j].source || sortSource[i].source == sortSource[j].source && sortSource[i].id < sortSource[j].id
	})
	var res []int
	for _, p := range sortSource[:k] {
		res = append(res, p.id)
	}
	return res
}

func main() {
	positiveFeedback := []string{"smart", "brilliant", "studious"}
	negativeFeedback := []string{"not"}
	report := []string{"this student is studious", "the student is smart"}
	studentId := []int{1, 2}
	k := 2
	res := topStudents(positiveFeedback, negativeFeedback, report, studentId, k)
	fmt.Println(res)
}
