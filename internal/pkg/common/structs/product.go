package structs

type ProductComponent struct {
	ArticleId int `json:"art_id,string"`
	Quantity int `json:"amount_of,string"`
}

type Product struct {
	Name string `json:"name"`
	Price float64                 `json:"price"`
	Components []ProductComponent `json:"contain_articles"`
}

type ProductAvailability struct {
	Product Product `json:"product"`
	Availability int `json:"availability"`
}