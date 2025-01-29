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

func (r *Rater) Has(id int) bool {
	_, ok := r.rates[id]

	return ok
}

func (r *Rater) Get(id int) int {
	return r.rates[id]
}

func (r *Rater) Set(id int, value int) {
	r.rates[id] = value
}

func (r *Rater) Delete(id int) {
	delete(r.rates, id)
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
