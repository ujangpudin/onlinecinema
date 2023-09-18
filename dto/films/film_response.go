package filmsdto

type FilmResponse struct {
	ID            int    `json:"id"`
	Title         string `json:"title" form:"title" validate:"required"`
	Thumbnailfilm string `json:"thumbnailfilm" form:"thumbnailfilm" validate:"required"`
	Price         string `json:"price" form:"price"`
	Linkfilmbuyed string `json:"linkfilmbuyed"`
	Linkfilm      string `json:"linkfilm"`
	CategoryID    string `json:"category_id"`
	Category      string `json:"category"`
	Description   string `json:"description" form:"description" validate:"required"`
}
