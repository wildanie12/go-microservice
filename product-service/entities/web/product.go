package web

// ProductRequest web entity response provides
// a schema for product request input
type ProductRequest struct {
	Name     string 	`form:"name" json:"name"`
	Unit     string 	`form:"unit" json:"unit"`
	Price    int64 		`form:"price" json:"price"`
	Category string 	`form:"category" json:"category"`
	IsActive bool 		`form:"is_active" json:"is_active"`
}

// ProductResponse web entity 
// response mapper
type ProductResponse struct {
	ID       string		`json:"id"`
	Name     string		`json:"name"`
	Unit     string		`json:"unit"`
	Price    int64		`json:"price"`
	Category string		`json:"category"`
	IsActive bool		`json:"is_active"`
}