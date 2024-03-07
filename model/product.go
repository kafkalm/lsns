package model

type Product struct {
	Name string `json:"name"`
	Price int `json:"price"`
}

type SortableProducts []Product

func (p SortableProducts) Len() int {
	return len(p)
}

func (p SortableProducts) Less(i, j int) bool {
	return p[i].Price > p[j].Price
}

func (p SortableProducts) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}