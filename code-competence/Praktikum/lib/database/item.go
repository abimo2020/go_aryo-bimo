package database

import (
	"mini-project/config"
	"mini-project/models"
	"mini-project/models/payload"

	uuid "github.com/satori/go.uuid"
)

func GetItems() (resp []models.Item, err error) {
	if err := config.DB.Preload("Category").Find(&resp).Error; err != nil {
		return resp, err
	}
	return
}

func GetItem(id uuid.UUID) (resp models.Item, err error) {
	if err := config.DB.Preload("Category").Where("id = ?", id).First(&resp).Error; err != nil {
		return resp, err
	}
	return
}

func GetItemByCategoryId(id uint) (resp []models.Item, err error) {
	if err := config.DB.Preload("Category").Where("category_id = ?", id).Find(&resp).Error; err != nil {
		return resp, err
	}
	return resp, nil
}

func GetItemsByName(name string) (resp []models.Item, err error) {
	if err := config.DB.Preload("Category").Where("name like ?", "%"+name+"%").Find(&resp).Error; err != nil {
		return resp, err
	}
	return resp, nil
}

func CreateItem(req *payload.Item) error {
	var item models.Item
	if err := config.DB.Model(&item).Create(&models.Item{
		Name:        req.Name,
		Description: req.Description,
		Stock:       req.Stock,
		Price:       req.Price,
		CategoryID:  req.CategoryID,
	}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateItem(req *payload.Item, id uuid.UUID) error {
	var item models.Item
	if err := config.DB.Model(&item).Where("id = ?", id).Updates(&models.Item{
		Name:        req.Name,
		Description: req.Description,
		Stock:       req.Stock,
		Price:       req.Price,
		CategoryID:  req.CategoryID,
	}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteItem(id uuid.UUID) error {
	var item models.Item
	if err := config.DB.Where("id = ?", id).Delete(&item).Error; err != nil {
		return err
	}
	return nil
}
