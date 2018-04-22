package data

import "sort"

const (
	WATSONS_SHOP = "屈臣氏"
	POYA_SHOP    = "寶雅"
)

// ProductInfo defines the product information
type ProductInfo struct {
	ID    string
	Name  string
	Price int
	Shop  string
}

// ProductInfoList defines the lists product information
type ProductInfoList []ProductInfo

func (p ProductInfoList) Len() int           { return len(p) }
func (p ProductInfoList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ProductInfoList) Less(i, j int) bool { return p[i].Price < p[j].Price }

// SortAscendByPrice sorts product with price with ascending order.
func SortAscendByPrice(p []ProductInfo) {
	sort.Sort(ProductInfoList(p))
}

// SortDescendByPrice sorts product with price with descending order.
func SortDescendByPrice(p []ProductInfo) {
	sort.Sort(sort.Reverse(ProductInfoList(p)))
}
