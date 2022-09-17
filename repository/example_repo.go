package repository

import (
	"github.com/akmamun/gorm-pagination"
	"go-fication-examples/models"
)

type ExampleRepo interface {
	GetExamples(limit, offset int64) (res interface{}, err error)
	CreateExample(exp *models.Example) error
	GetExample(id int) (res *models.Example, err error)
}

func (r *GormRepository) GetExamples(limit, offset int64) (res interface{}, err error) {
	var example []*models.Example
	res = pagination.Paginate(&pagination.Param{
		DB:      r.db.Database,
		Limit:   limit,
		Offset:  offset,
		OrderBy: "id ASC",
	}, &example)
	return
}
func (r *GormRepository) GetExamplesList() (exp []*models.Example, err error) {
	err = r.db.Database.Find(&exp).Error
	return
}

func (r *GormRepository) CreateExample(exp *models.Example) (err error) {
	err = r.db.Database.Create(exp).Error
	return
}

func (r *GormRepository) GetExample(id int) (res *models.Example, err error) {
	query := models.Example{Id: id}
	err = r.db.Database.Where(&query).Last(&res).Error
	return
}
