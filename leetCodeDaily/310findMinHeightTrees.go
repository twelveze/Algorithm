package main

import "fmt"

//从边缘开始，先找到所有出度为1的节点，然后把所有出度为1的节点进队列，然后不断地bfs
//最后找到的就是两边同时向中间靠近的节点，那么这个中间节点就相当于把整个距离二分了
//那么它当然就是到两边距离最小的点啦，也就是到其他叶子节点最近的节点了
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	var res []int
	nodeMap := make(map[int][]int, n) //相邻节点的映射
	degree := make([]int, n)          //存储节点的度
	for _, e := range edges {
		nodeMap[e[0]] = append(nodeMap[e[0]], e[1])
		nodeMap[e[1]] = append(nodeMap[e[1]], e[0])
		degree[e[0]]++
		degree[e[1]]++
	}

	queue := make([]int, 0)
	for i := range degree { //把所有度为1的节点(即叶子节点)加入到队列
		if degree[i] == 1 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 { //
		n := len(queue)
		res = make([]int, n)
		copy(res, queue)
		for i := 0; i < n; i++ { //遍历叶子节点
			cur := queue[0]
			queue = queue[1:]
			for _, next := range nodeMap[cur] { //遍历与叶子节点的所有邻接节点，将度减1
				degree[next]--
				if degree[next] == 1 { //若是新的叶子节点，则入队
					queue = append(queue, next)
				}
			}
		}
	}
	return res //返回最后一轮的叶子节点，就是最小高度树的根节点
}

func main() {
	n := 6
	edges := [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}
	res := findMinHeightTrees(n, edges)
	fmt.Println(res)
}
