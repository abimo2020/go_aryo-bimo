package usecase

import (
	"mini-project/lib/database"
	"mini-project/models"
	"mini-project/models/payload"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func GetProfil(id uint) (payload.GetProfil, error) {
	user, err := database.GetProfil(id)
	if err != nil {
		return payload.GetProfil{}, err
	}
	profil := payload.GetProfil{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}
	return profil, nil
}

func UpdateProfil(id uint, req *payload.UpdateProfil) error {
	user, err := database.GetUserById(id)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "The password is wrong")
	}
	if req.NewPassword != req.RetypePassword {
		return echo.NewHTTPError(http.StatusBadRequest, "The password is not match")
	}
	password, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	profil := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(password),
	}
	if err := database.UpdateProfil(&profil, user.Username); err != nil {
		return err
	}
	return nil
}

func DeleteUser(id uint, password string) error {
	user, err := database.GetUserById(id)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "The password is wrong")
	}
	if err := database.DeleteUser(id); err != nil {
		return err
	}
	return nil
}
