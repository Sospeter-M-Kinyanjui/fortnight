package main

func main() {
}

func Chunky(arr []int, size int) [][]int {
	if len(arr) < 1 {
		return nil
	}

	big := [][]int{}
	small := []int{}

	counter := 0
	for _, v := range arr {
		if counter == size {
			big = append(big, small...)
			small = nil
			counter = 0
		} else {
			small = append(small, v)
			counter++
		}
	}
	return big
}
