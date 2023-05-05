package webs

type ProductResponseDTO struct {
	ID                uint   `json:"product_id" form:"product_id"`
	Code              string `json:"code" form:"code"`
	Name              string `json:"name" form:"name"`
	Desc              string `json:"desc" form:"desc"`
	Price             int    `json:"price" form:"price"`
	UnitOfMeasurement string `json:"unit_of_measurement" form:"unit_of_measurement"`
}
