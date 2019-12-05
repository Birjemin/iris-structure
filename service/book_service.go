package service

import (
	"encoding/json"
	"github.com/birjemin/iris-structure/cache"
	"github.com/birjemin/iris-structure/models"
	"github.com/birjemin/iris-structure/repo"
)

type IBookService interface {
	List(m map[string]interface{}) (result models.Result)
	Save(book models.Book) (result models.Result)
	Get(id uint) (result models.Result)
	Del(book models.Book) (result models.Result)
}

type bookService struct {
	repoB repo.IBookRepository
}

func NewBookService() IBookService {
	return &bookService{repoB: repo.NewBookRepository()}
}

func (b bookService) List(m map[string]interface{}) (result models.Result) {
	total, books := b.repoB.List(m)
	maps := make(map[string]interface{}, 2)
	maps["Total"] = total
	maps["List"] = books
	result.Data = maps
	result.Code = 0
	result.Msg = "SUCCESS"
	return
}
func (b bookService) Save(book models.Book) (result models.Result) {
	err := b.repoB.Save(book)
	if err != nil {
		result.Data = nil
		result.Code = -1
		result.Msg = "保存失败"
	} else {
		result.Data = nil
		result.Code = 1
		result.Msg = "保存成功"
	}
	return
}
func (b bookService) Get(id uint) (result models.Result) {
	res, _ := cache.Get("tian")
	if len(res) != 0 {
		result.Data = res
		result.Code = 0
		result.Msg = "SUCCESS"
		return
	}

	book, err := b.repoB.Get(id)
	str, _ := json.Marshal(book)
	cache.Set("tian", string(str), 600)
	if err != nil {
		result.Data = nil
		result.Code = -1
		result.Msg = err.Error()
	} else {
		result.Data = book
		result.Code = 0
		result.Msg = "SUCCESS"
	}
	return
}

func (b bookService) Del(book models.Book) (result models.Result) {
	err := b.repoB.Del(book)
	if err != nil {
		result.Data = nil
		result.Code = -1
		result.Msg = err.Error()
	} else {
		result.Data = nil
		result.Code = 0
		result.Msg = "SUCCESS"
	}
	return
}
