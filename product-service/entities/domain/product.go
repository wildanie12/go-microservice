package domain

// Product main entity model
// that used for data modeling using mongodb
type Product struct {
	ID       string `bson:"_id,omitempty"`
	Name     string
	Unit     string
	Price    int64
	Category string
	isActive bool `bson:"is_active"`
}