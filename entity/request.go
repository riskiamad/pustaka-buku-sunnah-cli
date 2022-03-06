package entity

type BookRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Discount    int    `json:"discount"`
	Rating      int    `json:"rating"`
}
