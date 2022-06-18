package web

// ProductRequest web entity response provides
// a schema for product request input
type ProductRequest struct {
	Name     string
	Unit     string
	Price    int64
	Category string
	isActive bool
}

// ProductResponse web entity 
// response mapper
type ProductResponse struct {
	ID       string
	Name     string
	Unit     string
	Price    int64
	Category string
	isActive bool
}