package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Server struct {
	Name string
}
type LoadBalance struct {
	serverCluster []*Server
	size          int32
}

func (c *Server) Call() {
	fmt.Println("call")
}
func (m *LoadBalance) getServer(n int32) *Server {
	rand.Seed(time.Now().UnixNano())
	x := rand.Int31n(n)
	return m.serverCluster[x%m.size]
}
func NewLoadServer(size int32) *LoadBalance {
	loadBalance := &LoadBalance{serverCluster: make([]*Server, size), size: size}
	return loadBalance
}
func main() {
	loadBalance := NewLoadServer(4)
	loadBalance.getServer(4).Call()
}
