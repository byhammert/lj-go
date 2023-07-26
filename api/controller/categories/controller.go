package categories

import (
	"net/http"

	"github.com/byhammert/lj-go/api/controller"
	entities "github.com/byhammert/lj-go/entities/categories"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	CategoryUsecase entities.CategoryUsecaseContract
}

func NewCategoryController(usecase entities.CategoryUsecaseContract) *CategoryController {
	return &CategoryController{
		CategoryUsecase: usecase,
	}
}

func (cc *CategoryController) Create(c *gin.Context) {
	input, err := getInputBody(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	category, err := cc.CategoryUsecase.Create(input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, category)
}

func (cc *CategoryController) Delete(c *gin.Context) {
	id, err := getInputID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	if err = cc.CategoryUsecase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, controller.NewResponseMessage("Estudante exluido com sucesso"))
}

func (cc *CategoryController) Details(c *gin.Context) {
	id, err := getInputID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Problema com seu id"))
		return
	}

	categoryFound, err := cc.CategoryUsecase.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	output, err := getOutputCategory(categoryFound)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, output)
}

func (cc *CategoryController) List(c *gin.Context) {
	categories, err := cc.CategoryUsecase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
	}

	output, err := getOutputListCategories(categories)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
	}

	c.JSON(http.StatusOK, output)
}

func (cc *CategoryController) Update(c *gin.Context) {
	id, err := getInputID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	input, err := getInputBody(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	category, err := cc.CategoryUsecase.Update(id, input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	output, err := getOutputCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, output)
}
