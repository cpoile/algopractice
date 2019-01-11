package poker

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

type card struct {
	val  int
	suit rune
}

type cardsAndOrig struct {
	cards []card
	orig  string
	rank  int
}

func BestHand(in []string) ([]string, error) {
	var hands []cardsAndOrig
	for _, h := range in {
		curHand, ok := parseHand(h)
		if !ok {
			return nil, errors.New("invalid hand")
		}
		sortHand(curHand)
		hands = append(hands, cardsAndOrig{curHand, h, rank(curHand)})
	}

	sort.Slice(hands, func(a, b int) bool {
		return hands[a].rank < hands[b].rank ||
			hands[a].rank == hands[b].rank &&
				compare(hands[a].cards, hands[b].cards) > 0
	})

	winners := []string{hands[0].orig}
	for i := 1; i < len(in); i++ {
		if compare(hands[i-1].cards, hands[i].cards) == 0 {
			winners = append(winners, hands[i].orig)
		}
	}

	return winners, nil
}

func sortHand(hand []card) {
	counts := make(map[int]int)
	for _, c := range hand {
		counts[c.val]++
	}
	sort.Slice(hand, func(a, b int) bool {
		return counts[hand[a].val] > counts[hand[b].val] ||
			counts[hand[a].val] == counts[hand[b].val] && hand[a].val > hand[b].val
	})
}

func rank(hand []card) int {
	if isStraight(hand) && isFlush(hand) {
		return 1
	}
	if isFourKind(hand) {
		return 2
	}
	if isFullHouse(hand) {
		return 3
	}
	if isFlush(hand) {
		return 4
	}
	// special case of ace low in straights
	if isStraight(hand) && is5High(hand) {
		return 6
	}
	if isStraight(hand) {
		return 5
	}
	if isThreeKind(hand) {
		return 7
	}
	if isTwoPair(hand) {
		return 8
	}
	if isOnePair(hand) {
		return 9
	}
	return 10
}

func isFlush(hand []card) bool {
	suit := hand[0].suit
	for _, c := range hand {
		if c.suit != suit {
			return false
		}
	}
	return true
}

func isStraight(hand []card) bool {
	for i := 1; i < len(hand); i++ {
		if hand[i].val+1 != hand[i-1].val {
			// might be special case of ace low in straights
			return is5High(hand)
		}
	}
	return true
}

func is5High(hand []card) bool {
	if hand[0].val == 14 {
		tempCards := make([]card, 5)
		copy(tempCards, hand)
		tempCards[0].val = 1
		sortHand(tempCards)
		return isStraight(tempCards)
	}
	return false
}

func isFullHouse(hand []card) bool {
	return hand[0].val == hand[2].val && hand[3].val == hand[4].val
}

func isFourKind(hand []card) bool {
	return hand[0].val == hand[3].val
}

func isThreeKind(hand []card) bool {
	return hand[0].val == hand[2].val
}

func isTwoPair(hand []card) bool {
	return hand[0].val == hand[1].val && hand[2].val == hand[3].val
}

func isOnePair(hand []card) bool {
	return hand[0].val == hand[1].val
}

func compare(a, b []card) int {
	for i := 0; i < 5; i++ {
		cmp := a[i].val - b[i].val
		if cmp != 0 {
			return cmp
		}
	}
	return 0
}

func parseCard(c string) (card, bool) {
	var suit rune
	rcard := []rune(c)
	switch rcard[len(rcard)-1] {
	case '♧', '♤', '♡', '♢':
		suit = rcard[len(rcard)-1]
	default:
		return card{}, false
	}

	var value int
	var err error
	switch string(rcard[:len(rcard)-1]) {
	case "J":
		value = 11
	case "Q":
		value = 12
	case "K":
		value = 13
	case "A":
		value = 14
	default:
		value, err = strconv.Atoi(string(rcard[:len(rcard)-1]))
		if err != nil || value < 2 || value > 10 {
			return card{}, false
		}
	}

	return card{value, suit}, true
}

func parseHand(hand string) ([]card, bool) {
	tokens := strings.Split(hand, " ")
	if len(tokens) != 5 {
		return nil, false
	}

	var curHand []card
	for _, t := range tokens {
		c, ok := parseCard(t)
		if !ok {
			return nil, false
		}
		curHand = append(curHand, c)
	}

	return curHand, true
}
