package main

import (
	"fmt"
)

var cheapest int

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {

	//visited := make(map[int]bool)
	//routeMap := make(map[int][]*Nbr)
	//
	//for _,route := range flights {
	//	if nbrs,ok := routeMap[route[0]]; ok {
	//		all := append(nbrs, &Nbr{Name: route[1], Price: route[2]})
	//		routeMap[route[0]] = all
	//	} else {
	//		routeMap[route[0]] = []*Nbr{{Name: route[1], Price: route[2]}}
	//	}
	//}

	//possiblePaths := make(map[string]Route)
	//cheapest = -1
	//getAllPaths(routeMap, visited, src, dst, possiblePaths, 0, 0, fmt.Sprintf("%d", src), k)
	//for _, route := range possiblePaths {
	//	fmt.Println(route.Price, route.Stops)
	//	if cheapest == -1 {
	//		cheapest = route.Price
	//	} else if route.Price < cheapest {
	//		cheapest = route.Price
	//	}
	//}
	//
	//return cheapest

	// optimized solution
	return cheapestPriceBellmanFord(n, flights, src, dst, k)
}

func cheapestPriceBellmanFord(n int, flights [][]int, src, dest, k int) int {

	prices := make([]int, 0)
	for i:=0; i<n; i++ {
		prices = append(prices, MaxInt)
	}
	prices[src] = 0

	for i:=0; i<k+1; i++ {
		tmpPrices := make([]int, len(prices))
		copy(tmpPrices, prices)
		for _,flight := range flights {
			if prices[flight[0]] == MaxInt {
				continue
			}
			if prices[flight[0]] + flight[2] < tmpPrices[flight[1]] {
				tmpPrices[flight[1]] = prices[flight[0]] + flight[2]
			}
		}
		prices = tmpPrices
	}

	if prices[dest] == MaxInt {
		return -1
	}
	return prices[dest]
}

func getAllPaths(routeMap map[int][]*Nbr, visited map[int]bool, src, dest int, possiblePaths map[string]Route, priceSoFar, stopsSoFar int, pathSoFar string, k int) {

	if src == dest {
		cheapest = priceSoFar
		possiblePaths[pathSoFar] = Route{Price: priceSoFar, Stops: stopsSoFar - 1}
		return
	}

	if stopsSoFar - 1 >= k || cheapest != -1 && cheapest < priceSoFar {
		return
	}

	visited[src] = true

	for _,nbr := range routeMap[src] {
		if visited[nbr.Name] {
			continue
		}
		totalPrice := priceSoFar + nbr.Price
		getAllPaths(routeMap, visited, nbr.Name, dest, possiblePaths, totalPrice, stopsSoFar + 1, fmt.Sprintf("%s%d", pathSoFar, nbr.Name), k)
	}

	visited[src] = false
}

type Route struct {
	Price int
	Stops int
}

type Nbr struct {
	Name int
	Price int
}
