package controller

import (
	"mini-project/models/payload"
	"mini-project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

func GetItemsController(c echo.Context) error {
	var items []payload.GetItem
	var err error
	name := c.QueryParam("keyword")

	if name != "" {
		it, er := usecase.GetItemsByName(name)
		items = it
		err = er
	} else {
		i, e := usecase.GetItems()
		items = i
		err = e
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all items",
		"data":    items,
	})
}

func GetItemController(c echo.Context) error {
	id := c.Param("id")
	item, err := usecase.GetItem(uuid.FromStringOrNil(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get item",
		"data":    item,
	})
}

func GetItemByCategoryController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("category_id"))
	item, err := usecase.GetItemByCategoryId(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get item",
		"data":    item,
	})
}

func GetItemByNameControllers(c echo.Context) error {
	name := c.QueryParam("item_name")
	item, err := usecase.GetItemsByName(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get item",
		"data":    item,
	})
}

func CreateItemController(c echo.Context) error {
	req := payload.Item{}
	c.Bind(&req)
	if err := c.Validate(req); err != nil {
		return err
	}
	err := usecase.CreateItem(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create item",
	})
}

func UpdateItemController(c echo.Context) error {
	var req payload.Item
	id := c.Param("id")
	c.Bind(&req)
	if err := c.Validate(req); err != nil {
		return err
	}
	err := usecase.UpdateItem(&req, uuid.FromStringOrNil(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update item",
	})
}

func DeleteItemController(c echo.Context) error {
	id := c.Param("id")
	err := usecase.DeleteItem(uuid.FromStringOrNil(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete item",
	})
}
