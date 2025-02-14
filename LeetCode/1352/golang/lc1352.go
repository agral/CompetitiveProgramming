package lc1352

type ProductOfNumbers struct {
	Prefix []int
}

func Constructor() ProductOfNumbers {
	p := ProductOfNumbers{}
	p.Prefix = []int{1}
	return p
}

func (p *ProductOfNumbers) Add(num int) {
	if num == 0 {
		p.Prefix = []int{1}
	} else {
		lastVal := p.Prefix[len(p.Prefix)-1]
		p.Prefix = append(p.Prefix, lastVal*num)
	}
}

func (p *ProductOfNumbers) GetProduct(k int) int {
	length := len(p.Prefix)
	if length <= k { // has to be <=, as I've decided to initialize the prefix slice with "1".
		// This approach simplifies corner cases, at (a negligible) cost of one extra stored value.
		// But this one extra stored value must be taken into account, hence <= instead of <.
		return 0 // there's not enough new values yet after the last 0. The product is 0.
	} else {
		// A prefix sum (or rather, a prefix product) of last k numbers is:
		return p.Prefix[length-1] / p.Prefix[length-k-1]
	}
}
