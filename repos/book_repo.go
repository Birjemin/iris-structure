package repos

import (
	"github.com/birjemin/iris-structure/datasource"
	"github.com/birjemin/iris-structure/models"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
)

type IBookRepository interface {
	List(m map[string]interface{}) (total int, books []models.Book)
	Save(book models.Book) (err error)
	Get(id uint) (book models.Book, err error)
	Del(book models.Book) (err error)
}

func NewBookRepository() IBookRepository {
	return &bookRepository{db: datasource.GetDB()}
}

type bookRepository struct {
	db *gorm.DB
}

func (b bookRepository) List(m map[string]interface{}) (total int, books []models.Book) {
	b.db.Table("book").Count(&total)
	err := b.db.Limit(cast.ToInt(m["size"])).Offset((cast.ToInt(m["page"]) - 1) * cast.ToInt(m["size"])).Find(&books).Error
	if err != nil {
		panic("select Error")
	}
	return
}

func (b bookRepository) Save(book models.Book) (err error) {
	if book.ID != 0 {
		err := b.db.Save(&book).Error
		return err
	} else {
		err := b.db.Create(&book).Error
		return err
	}
}

func (b bookRepository) Get(id uint) (book models.Book, err error) {
	err = b.db.First(&book, id).Error
	return
}

func (b bookRepository) Del(book models.Book) (err error) {
	err = b.db.Unscoped().Delete(&book).Error
	return
}
