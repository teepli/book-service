package books

type BookQuery struct {
	Title       string  `json:"title" binding:"required,notEmptyString"`
	Author      string  `json:"author" binding:"required,notEmptyString"`
	Year        int     `json:"year" binding:"required,numeric"`
	Publisher   *string `json:"publisher,omitempty" binding:"omitempty,notEmptyString"`
	Description string  `json:"description,omitempty" binding:"omitempty"`
}

type BookResponse struct {
	Id int `json:"id,omitempty"`
	BookQuery
}

type BookFilterParams struct {
	Author    *string `form:"author" binding:"omitempty,notEmptyString"`
	Year      *int    `form:"year" binding:"omitempty,numeric"`
	Publisher *string `form:"publisher,omitempty" binding:"omitempty,notEmptyString"`
}

type IdParam struct {
	Id int `uri:"id" binding:"required,numeric"`
}

type CreateBookResponse struct {
	Id int `json:"id" binding:"required"`
}
