package controller

import (
	"strconv"

	r "github.com/kiditz/spgku-api/repository"
	t "github.com/kiditz/spgku-api/translate"

	"github.com/labstack/echo/v4"
)

// GetCategories godoc
// @Summary GetCategories used to find all categories
// @Description find all category
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {array} translate.ResultSuccess{data=entity.Category} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /categories [get]
// @Security ApiKeyAuth
func GetCategories(c echo.Context) error {
	categories := r.GetCategories()
	return t.Success(c, categories)
}

// GetSubCategories godoc
// @Summary GetSubCategories used to find all sub category
// @Description find all sub category
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {array} translate.ResultSuccess{data=entity.SubCategory} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /sub-categories [get]
// @Security ApiKeyAuth
func GetSubCategories(c echo.Context) error {
	categories := r.GetSubCategories()
	return t.Success(c, categories)
}

// GetSubCategoriesByCategoryID godoc
// @Summary GetSubCategoriesByCategoryID used to find all sub category by category id
// @Description used to find all sub category by category id
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {array} translate.ResultSuccess{data=entity.SubCategory} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /sub-categories/{id} [get]
func GetSubCategoriesByCategoryID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	categories := r.GetSubCategoriesByCategoryID(id)
	return t.Success(c, categories)
}

// GetExpertises godoc
// @Summary GetExpertises used to find all expertises
// @Description  used to find all expertises
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {array} translate.ResultSuccess{data=entity.Expertise} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /expertises [get]
func GetExpertises(c echo.Context) error {
	expertises := r.GetExpertises()
	return t.Success(c, expertises)
}
