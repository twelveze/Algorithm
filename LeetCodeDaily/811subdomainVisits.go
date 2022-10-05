package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	domainMap := make(map[string]int)
	for _, s := range cpdomains {
		spaceIndex := strings.Index(s, " ")
		count, _ := strconv.Atoi(s[:spaceIndex]) //次数
		subDomain := s[spaceIndex+1:]            //域名str
		domainMap[subDomain] += count
		for {
			spot := strings.Index(subDomain, ".")
			if spot < 0 {
				break
			}
			subDomain = subDomain[spot+1:]
			domainMap[subDomain] += count
		}
	}
	ans := make([]string, 0, len(domainMap))
	for str, count := range domainMap {
		ans = append(ans, strconv.Itoa(count)+" "+str)
	}
	return ans
}

func main() {
	cpdomains := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	res := subdomainVisits(cpdomains)
	fmt.Println(res)
}
