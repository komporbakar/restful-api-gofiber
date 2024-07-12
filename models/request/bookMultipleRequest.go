package request

type BookMultipleRequest struct {
	CategoryId uint `json:"category_id" form:"category_id" validate:"required"`
	// Name       string `json:"name" validate:"required"`
}
