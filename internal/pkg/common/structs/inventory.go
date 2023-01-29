package structs

type Article struct {
	Id int `json:"art_id,string"`
	Name string `json:"name"`
	Stock int `json:"stock,string"`
}
