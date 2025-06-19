package main

import "fmt"

func main() {
	flowerbed := []int{0, 0, 1, 0, 0}
	fmt.Println(canPlaceFlowers(flowerbed, 1))

}

func canPlaceFlowers(flowerbed []int, n int) bool {
	flowersPlanted := 0

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			prevElementEmpty := i == 0 || flowerbed[i-1] == 0
			nextElementEmpty := i == len(flowerbed)-1 || flowerbed[i+1] == 0
			if prevElementEmpty && nextElementEmpty {
				flowerbed[i] = 1
				flowersPlanted++

			}
		}

	}
	return flowersPlanted >= n

}
