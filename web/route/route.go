package route

import (
	"github.com/birjemin/iris-structure/web/controllers"
	middleware "github.com/birjemin/iris-structure/web/middlweare"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func InitRouter(app *iris.Application) {
	mvc.New(app.Party("/api/v1/book", middleware.Handler)).Handle(controllers.NewBookController())
}
