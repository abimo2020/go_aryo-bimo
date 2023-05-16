package payload

type LoginForm struct {
	ID       string `json:"id" form:"id" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type Register struct {
	Name           string `json:"name" form:"name" validate:"required"`
	Username       string `json:"username" form:"username" validate:"required,min=4,max=12"`
	Email          string `json:"email" form:"email" validate:"required,email"`
	Password       string `json:"password" form:"password" validate:"required,min=6,max=12"`
	RetypePassword string `json:"retype_password" form:"retype_password" validate:"required,min=6,max=12"`
}

type UpdateProfil struct {
	Password       string `json:"password" form:"password" validate:"required"`
	Name           string `json:"name" form:"name"`
	Email          string `json:"email" form:"email"`
	NewPassword    string `json:"new_password" form:"new_password" validate:"min=6,max=12"`
	RetypePassword string `json:"retype_password" form:"retype_password" validate:"min=6,max=12"`
}

type CreateCategory struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type Item struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Stock       int    `json:"stock" form:"stock" validate:"required"`
	Price       int    `json:"price" form:"price" validate:"required"`
	CategoryID  uint   `json:"category_id" form:"category_id" validate:"required"`
}
