package dto

// PhotoUpdateDTO is a model that client use when updating a photo
type PhotoUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Title    string `json:"title" form:"title" binding:"required"`
	Caption  string `json:"caption" form:"caption" binding:"required"`
	PhotoURL string ``
	UserID   uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}

// PhotoCreateDTO is is a model that clinet use when create a new photo
type PhotoCreateDTO struct {
	Title    string `json:"title" form:"title" binding:"required"`
	Caption  string `json:"caption" form:"caption" binding:"required"`
	PhotoURL string ``
	UserID   uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}
