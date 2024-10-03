package main

func main() {

}

func minimumOperations(nums []int) int {
	nonZeros := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		if v != 0 {
			nonZeros[v] = struct{}{}
		}
	}

	return len(nonZeros)
}
