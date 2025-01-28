package ratement

import (
	"sort"
)

type Rater struct {
	rates  map[int]int
	sorted []int
}

func NewRater(rates map[int]int) *Rater {
	sorted := make([]int, 0, len(rates))

	for k := range rates {
		sorted = append(sorted, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))

	return &Rater{rates, sorted}
}

func (r *Rater) Value(amount int) int {
	total := 0

	for _, v := range r.sorted {
		if amount == 0 {
			break
		}

		usable := amount / v

		if usable > 0 {
			total += usable * r.rates[v]
			amount -= usable * v
		}
	}

	return total
}
