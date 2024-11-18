package main

import (
	"fmt"
	"sort"
)

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	var restaurantsFilter [][]int
	var restaurantId []int
	for _, restaurant := range restaurants {
		if ((restaurant[2] == 1 && veganFriendly == 1) || veganFriendly == 0) && restaurant[3] <= maxPrice && restaurant[4] <= maxDistance {
			restaurantsFilter = append(restaurantsFilter, restaurant)
		}
	}
	sort.Slice(restaurantsFilter, func(i, j int) bool {
		return restaurantsFilter[i][1] > restaurantsFilter[j][1] || (restaurantsFilter[i][1] == restaurantsFilter[j][1] && restaurantsFilter[i][0] > restaurantsFilter[j][0])
	})
	for _, restaurant := range restaurantsFilter {
		restaurantId = append(restaurantId, restaurant[0])
	}
	return restaurantId
}

func main() {
	restaurants := [][]int{{1, 4, 1, 40, 10}, {2, 8, 0, 50, 5}, {3, 8, 1, 30, 4}, {4, 10, 0, 10, 3}, {5, 1, 1, 15, 1}}
	veganFriendly := 1
	maxPrice := 50
	maxDistance := 10
	res := filterRestaurants(restaurants, veganFriendly, maxPrice, maxDistance)
	fmt.Print(res)
}
