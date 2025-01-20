package models

type Product struct {
	ID          string  `json:"id" bson:"_id"`
	SellerID    string  `json:"seller_id" bson:"seller_id"`
	Name        string  `json:"name" bson:"name"`
	SKU         string  `json:"sku" bson:"sku"`
	Description string  `json:"description" bson:"description"`
	Price       float64 `json:"price" bson:"price"`
	Stock       int     `json:"stock" bson:"stock"`
	//UserID
}
