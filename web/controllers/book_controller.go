package controllers

import (
	"github.com/birjemin/iris-structure/models"
	"github.com/birjemin/iris-structure/service"
	"github.com/kataras/iris/v12"
	"github.com/spf13/cast"
	"time"
)

type BookController struct {
	Ctx     iris.Context
	Service service.IBookService
}

func NewBookController() *BookController {
	return &BookController{Service: service.NewBookService()}
}

//GET http://localhost:8081/api/v1/book?page=1&size=10
func (g *BookController) Get() (result models.Result) {
	r := g.Ctx.Request()
	m := make(map[string]interface{})
	page := r.FormValue("page")
	size := r.FormValue("size")
	if page == "" {
		result.Msg = "page不能为空"
		result.Code = -1
		return
	}
	if size == "" {
		result.Code = -1
		result.Msg = "size不能为空"
		return
	}
	m["page"] = page
	m["size"] = size
	return g.Service.List(m)
}

//GET http://localhost:8081/api/v1/book/1
func (g *BookController) GetBy(id uint) (result models.Result) {
	//uid := cast.ToUint(id)
	return g.Service.Get(id)
}

//POST http://localhost:8081/api/v1/book
func (g *BookController) Post() (result models.Result) {
	r := g.Ctx.Request()
	book := models.Book{}
	//book.ID = cast.ToUint(r.PostFormValue("id"))
	book.Name = r.PostFormValue("name")
	book.Count = r.PostFormValue("count")
	book.Author = r.PostFormValue("author")
	book.Type = r.PostFormValue("type")
	book.CreatedAt = cast.ToUint64(time.Now().Unix())
	return g.Service.Save(book)
}

//PUT http://localhost:8081/api/v1/book/1
func (g *BookController) PutBy(id uint) (result models.Result) {
	r := g.Ctx.Request()
	book := models.Book{}
	book.ID = id
	book.Name = r.PostFormValue("name")
	book.Count = r.PostFormValue("count")
	book.Author = r.PostFormValue("author")
	book.Type = r.PostFormValue("type")
	return g.Service.Save(book)
}

//DELETE http://localhost:8081/api/v1/book/2
func (g *BookController) DeleteBy(id uint) (result models.Result) {
	book := models.Book{}
	book.ID = id
	return g.Service.Del(book)
}
