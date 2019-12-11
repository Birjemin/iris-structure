package routes

import (
	"github.com/birjemin/iris-structure/web/controllers"
	middleware "github.com/birjemin/iris-structure/web/middlware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func InitRouter(app *iris.Application) {
	mvc.New(app.Party("/api/v1/book", middleware.Handler)).Handle(controllers.NewBookController())

	mvc.New(app.Party("/hello")).Handle(controllers.NewClientController())
}
