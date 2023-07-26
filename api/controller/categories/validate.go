package categories

import (
	entities "github.com/byhammert/lj-go/entities/categories"
	"github.com/byhammert/lj-go/entities/shared"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gookit/validate"
)

func getInputBody(c *gin.Context) (input InputBody, err error) {
	err = c.Bind(&input)
	if err != nil {
		return input, err
	}

	validation := validate.Struct(input)
	if !validation.Validate() {
		return input, validation.Errors
	}

	return input, err
}

func getInputID(c *gin.Context) (id uuid.UUID, err error) {
	inputID := c.Params.ByName("id")

	id, err = shared.GetUuidByString(inputID)
	if err != nil {
		return id, err
	}

	return id, nil
}

func getOutputListCategories(categories []entities.Category) (output OutputCategories, err error) {
	for _, s := range categories {
		outputCategory, err := getOutputCategory(s)
		if err != nil {
			return output, err
		}

		output.Categories = append(output.Categories, outputCategory)
	}

	return output, err
}

func getOutputCategory(category entities.Category) (output OutputCategory, err error) {
	return OutputCategory{
		ID:   category.ID,
		Name: category.Name,
	}, err
}
