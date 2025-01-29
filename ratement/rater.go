package ratement

import (
	"sort"
)

type Rater struct {
	rates  map[int]int
	sorted []int
}

func NewRater(rates map[int]int) *Rater {
	rater := Rater{rates: rates}

	rater.sortRates()

	return &rater
}

func (r *Rater) sortRates() {
	r.sorted = make([]int, 0, len(r.rates))

	for k := range r.rates {
		r.sorted = append(r.sorted, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(r.sorted)))
}

func (r *Rater) Has(id int) bool {
	_, ok := r.rates[id]

	return ok
}

func (r *Rater) Get(id int) int {
	return r.rates[id]
}

func (r *Rater) Set(id int, value int) {
	replace := r.Has(id)

	r.rates[id] = value

	if !replace {
		r.sortRates()
	}
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
			total += usable * r.Get(v)
			amount -= usable * v
		}
	}

	return total
}
