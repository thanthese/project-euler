package main

import "fmt"

func main() {
	total := 5000
	coins := []int{200, 100, 50, 20, 10, 5, 2, 1}
	ways := make_change(total, coins)
	fmt.Println(ways)
}

// find ways to make change for the target_amount using coins of the given
// denominations
func make_change(target_amount int, coins []int) int {

	// this func merely wraps and supports (with memoization) the central
	// recursive alg

	type args struct {
		total_so_far, coin_index int
	}

	memo := map[args]int{}

	var recur func(total_so_far, coin_index int) int
	recur = func(total_so_far, coin_index int) int {

		// memoization shortcut
		if v, ok := memo[args{total_so_far, coin_index}]; ok {
			return v
		}

		// the alg proper
		coin_val := coins[coin_index]
		if total_so_far == target_amount || coin_val == 1 {
			return 1
		}
		will_fit := (target_amount - total_so_far) / coin_val
		sum := 0
		for n := 0; n < will_fit+1; n++ {
			sum += recur(total_so_far+coin_val*n, coin_index+1)
		}

		// save result and return
		memo[args{total_so_far, coin_index}] = sum
		return sum
	}
	return recur(0, 0)
}
