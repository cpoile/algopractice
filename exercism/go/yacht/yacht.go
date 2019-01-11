package yacht

import "sort"

var targets = map[string]int{
	"ones": 1, "twos": 2, "threes": 3, "fours": 4, "fives": 5, "sixes": 6,
}

func Score(dice []int, cat string) int {
	switch cat {
	case "yacht":
		return yacht(dice)
	case "ones", "twos", "threes", "fours", "fives", "sixes":
		return singles(dice, targets[cat])
	case "full house":
		return fullhouse(dice)
	case "four of a kind":
		return four(dice)
	case "little straight":
		return straight(dice, false)
	case "big straight":
		return straight(dice, true)
	case "choice":
		return choice(dice)
	}
	return 0
}

func yacht(dice []int) int {
	sortDice(dice)
	if dice[0] == dice[4] {
		return 50
	}
	return 0
}

func four(dice []int) int {
	sortDice(dice)
	if dice[0] == dice[3] {
		return dice[0] * 4
	}
	return 0
}

func singles(dice []int, target int) int {
	var total int
	for _, d := range dice {
		if d == target {
			total += target
		}
	}
	return total
}

func straight(dice []int, isBigStraight bool) int {
	sort.Ints(dice)
	expected := 1
	if isBigStraight {
		expected = 2
	}
	for i := 0; i < 5; i, expected = i+1, expected+1 {
		if dice[i] != expected {
			return 0
		}
	}
	return 30
}

func fullhouse(dice []int) int {
	sortDice(dice)
	if dice[0] == dice[2] && dice[3] == dice[4] && dice[0] != dice[3] {
		return dice[0]*3 + dice[3]*2
	}
	return 0
}

func choice(dice []int) int {
	var sum int
	for _, d := range dice {
		sum += d
	}
	return sum
}

// sort by frequency then by value
func sortDice(dice []int) {
	counts := make(map[int]int)
	for _, d := range dice {
		counts[d]++
	}
	sort.Slice(dice, func(a, b int) bool {
		return counts[dice[a]] > counts[dice[b]] ||
			counts[dice[a]] == counts[dice[b]] && dice[a] < dice[b]
	})
}
