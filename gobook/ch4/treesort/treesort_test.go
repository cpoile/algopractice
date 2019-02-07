package treesort

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestTreesort(t *testing.T) {
	cases := [][]string{
		strings.Split("I am the very model of a modern major general. I am information, vegetable, animal and mineral.", " "),
		strings.Split("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed mollis justo sit amet diam tempus ultricies. Etiam porttitor diam et turpis pretium commodo. Quisque sit amet laoreet felis, quis finibus dui. Integer posuere nulla ac venenatis facilisis. Donec arcu ipsum, ornare posuere leo at, finibus faucibus leo. Nunc imperdiet hendrerit metus a lobortis. Maecenas vel vehicula leo, non sollicitudin mi. Donec eget nisi libero. Nam vestibulum lacus sed turpis maximus, consectetur pretium nulla lacinia. Donec varius quam leo, ac commodo eros molestie a. Maecenas a tortor in risus pretium consequat. Nulla venenatis, enim id scelerisque condimentum, nisi mauris euismod ante, nec tempus risus purus sodales orci. Duis imperdiet facilisis lectus, id vestibulum magna pretium et. Maecenas ultricies nisi et libero sodales tristique.", " ")}

	for _, c := range cases {
		orig := make([]string, len(c))
		copy(orig, c)
		sort.Strings(c)
		treeSort(orig)
		if fmt.Sprintf("%q", c) != fmt.Sprintf("%q", orig) {
			t.Errorf("\n%10s: %q \n%10s: %q", "Sorted", c, "Treesorted", orig)
		}
	}

}
