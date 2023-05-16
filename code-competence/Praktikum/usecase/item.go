package usecase

import (
	"mini-project/lib/database"
	"mini-project/models/payload"

	uuid "github.com/satori/go.uuid"
)

func GetItems() (resp []payload.GetItem, err error) {
	items, err := database.GetItems()
	if err != nil {
		return []payload.GetItem{}, err
	}
	for _, value := range items {
		resp = append(resp, payload.GetItem{
			ID:          value.ID,
			Name:        value.Name,
			Description: value.Description,
			Stock:       value.Stock,
			Price:       value.Price,
			Category:    value.Category.Name,
		})
	}
	return
}

func GetItem(id uuid.UUID) (resp payload.GetItem, err error) {
	item, err := database.GetItem(id)
	if err != nil {
		return payload.GetItem{}, err
	}
	resp = payload.GetItem{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Stock:       item.Stock,
		Price:       item.Price,
		Category:    item.Category.Name,
	}
	return
}

func GetItemByCategoryId(id uint) (resp []payload.GetItem, err error) {
	items, err := database.GetItemByCategoryId(id)
	if err != nil {
		return []payload.GetItem{}, err
	}
	for _, value := range items {
		resp = append(resp, payload.GetItem{
			ID:          value.ID,
			Name:        value.Name,
			Description: value.Description,
			Stock:       value.Stock,
			Price:       value.Price,
			Category:    value.Category.Name,
		})
	}
	return
}

func GetItemsByName(name string) (resp []payload.GetItem, err error) {
	items, err := database.GetItemsByName(name)
	if err != nil {
		return []payload.GetItem{}, err
	}
	for _, value := range items {
		resp = append(resp, payload.GetItem{
			ID:          value.ID,
			Name:        value.Name,
			Description: value.Description,
			Stock:       value.Stock,
			Price:       value.Price,
			Category:    value.Category.Name,
		})
	}
	return
}

func CreateItem(req *payload.Item) error {
	if err := database.CreateItem(req); err != nil {
		return err
	}
	return nil
}

func UpdateItem(req *payload.Item, id uuid.UUID) error {
	if _, err := database.GetItem(id); err != nil {
		return err
	}
	if err := database.UpdateItem(req, id); err != nil {
		return err
	}
	return nil
}

func DeleteItem(id uuid.UUID) error {
	if _, err := database.GetItem(id); err != nil {
		return err
	}
	if err := database.DeleteItem(id); err != nil {
		return err
	}
	return nil
}
