package payload

import uuid "github.com/satori/go.uuid"

type DashboardUser struct {
	Adoption int `json:"adoption" form:"adoption"`
	Donation int `json:"donation" form:"donation"`
}

type DashboardAdmin struct {
	TotalUser         int `json:"total_user" form:"total_user"`
	TotalAdoption     int `json:"total_adoption" form:"total_adoption"`
	TotalDonation     int `json:"total_donation" form:"total_donation"`
	TotalPetAvailable int `json:"total_pet_available" form:"total_pet_available"`
	TotalPetAdopted   int `json:"total_pet_adopted" form:"total_pet_"`
}

type GetProfil struct {
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
}

type GetItem struct {
	ID          uuid.UUID
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Stock       int    `json:"stock" form:"stock"`
	Price       int    `json:"price" form:"price"`
	Category    string `json:"category" form:"category"`
}
