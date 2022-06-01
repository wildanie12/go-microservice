package web

type ProductRequest struct {
	Name     string
	Unit     string
	Price    int64
	Category string
	isActive bool
}

type ProductResponse struct {
	ID       string
	Name     string
	Unit     string
	Price    int64
	Category string
	isActive bool
}