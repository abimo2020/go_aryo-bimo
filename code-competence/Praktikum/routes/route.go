package route

import (
	"mini-project/controller"
	m "mini-project/middlewares"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	m.LoggerMiddleware(e)

	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/register", controller.RegisterController)
	e.POST("/login", controller.LoginController)
	e.POST("/logout", controller.LogoutController)

	p := e.Group("/profil")
	p.Use(m.JWTMiddlewareConfig, m.IsUser)
	p.GET("", controller.GetProfilController)
	p.PUT("", controller.UpdateProfilController)
	p.DELETE("", controller.DeleteUserController)

	i := e.Group("/items")
	i.Use(m.JWTMiddlewareConfig)
	i.GET("", controller.GetItemsController)
	i.GET("/:id", controller.GetItemController)
	i.POST("", controller.CreateItemController)
	i.PUT("/:id", controller.UpdateItemController)
	i.DELETE("/:id", controller.DeleteItemController)
	i.GET("/category/:category_id", controller.GetItemByCategoryController)

	a := e.Group("/admin")
	a.Use(m.JWTMiddlewareConfig, m.IsAdmin)
	a.GET("/category", controller.GetCategoryController)
	a.POST("/category", controller.CreateCategoryController)

	return e
}
