package handlers

import (
	"net/http"
	"strconv"
	"test-impact/server/models/entities"
	"test-impact/server/models/webs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductHandler interface {
	Products(ctx *gin.Context)
	Create(ctx *gin.Context)
	Product(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type ProductHandlerResource struct {
	db *gorm.DB
}

func NewProductHandler(DB *gorm.DB) ProductHandler {
	return &ProductHandlerResource{db: DB}
}

func ConvertBodyProductResp(product entities.Product) webs.ProductResponseDTO {
	return webs.ProductResponseDTO{
		ID:                product.ID,
		Code:              product.Code,
		Name:              product.Name,
		Desc:              product.Desc,
		Price:             product.Price,
		UnitOfMeasurement: product.UnitOfMeasurement,
	}
}

func (h *ProductHandlerResource) Products(ctx *gin.Context) {
	products := []entities.Product{}
	err := h.db.Model(&products).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, false)
		return
	}

	productsResp := []webs.ProductResponseDTO{}

	for i := 0; i < len(products); i++ {
		productBody := ConvertBodyProductResp(products[i])
		productsResp = append(productsResp, productBody)
	}
	ctx.JSON(http.StatusOK, productsResp)
}

func (h *ProductHandlerResource) Create(ctx *gin.Context) {
	var createProduct entities.Product
	if err := ctx.Bind(&createProduct); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	err := h.db.Create(&createProduct).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	respProduct := ConvertBodyProductResp(createProduct)
	ctx.JSON(http.StatusCreated, respProduct)
}

func (h *ProductHandlerResource) Product(ctx *gin.Context) {
	var product entities.Product
	id := ctx.Param("product_id")
	productID, _ := strconv.ParseUint(id, 10, 32)

	err := h.db.Model(&product).Where("id = ?", productID).First(&product).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	respProduct := ConvertBodyProductResp(product)
	ctx.JSON(http.StatusOK, respProduct)
}

func (h *ProductHandlerResource) Update(ctx *gin.Context) {
	var updateProduct entities.Product

	id := ctx.Param("product_id")
	productID, _ := strconv.ParseUint(id, 10, 32)

	if err := ctx.Bind(&updateProduct); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err := h.db.Model(&updateProduct).Where("id = ?", productID).Updates(&updateProduct).First(&updateProduct).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	respBody := ConvertBodyProductResp(updateProduct)
	ctx.JSON(http.StatusOK, respBody)
}

func (h *ProductHandlerResource) Delete(ctx *gin.Context) {
	var product entities.Product
	id := ctx.Param("product_id")
	productID, _ := strconv.ParseUint(id, 10, 32)

	err := h.db.Delete(&product, productID).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusAccepted, "Delete Has Been Successfull")
}
