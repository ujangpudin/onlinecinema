package categorydto

type CreateCategoryRequest struct {
	Name string `json:"name"  form:"name" validate:"required"`
	// Subscribe bool   `json:"subscribe"  form:"form"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name"  form:"name"`
}
