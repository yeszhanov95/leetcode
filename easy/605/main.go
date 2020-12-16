package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; n > 0 && i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			i++
			continue
		}
		if i + 1 == len(flowerbed) || flowerbed[i+1] == 0 {
			i++
			n--
		}
	}
	return n == 0
}