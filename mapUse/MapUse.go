package main

import "fmt"

func main() {
	tokens := []string{"abc", "cde", "23243423", "werwer", "sdfjsidfjisdjf"}
	ip := []string{"192.168.0.1", "192.168.0.2", "192.168.0.3"}
	serverIPMap := make(map[string][]string)
	for i, token := range tokens {
		ipValue := ip[i%3]
		serverIPMap[ipValue] = append(serverIPMap[ipValue], token)
	}

	fmt.Println(serverIPMap)

	for key, iptokens := range serverIPMap {

		fmt.Println("key", key)
		fmt.Println("value", iptokens)

	}
}
