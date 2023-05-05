package entities

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	Code              string `gorm:"not null;unique" valid:"required~Your code product is required,minstringlength(3)~Code Product has to have a minimum length of 3 characters" json:"code" form:"code"`
	Name              string `gorm:"not null;unique" valid:"required~Your Product Name is required,minstringlength(5)~Product Name has to have a minimum length of 6 characters" json:"name" form:"name"`
	Desc              string `gorm:"type:text" valid:"required~Your Desc Product is required" json:"desc" form:"desc"`
	Price             int    `gorm:"not null" valid:"required~Your Price is required" json:"price" form:"price"`
	UnitOfMeasurement string `gorm:"type:text" valid:"required~Your Unit Of Measurement is required" json:"unit_of_measurement" form:"unit_of_measurement"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
