package controller

import (
	"net/http"

	e "github.com/kiditz/spgku-api/entity"
	r "github.com/kiditz/spgku-api/repository"
	t "github.com/kiditz/spgku-api/translate"
	u "github.com/kiditz/spgku-api/utils"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

// AddDigitalStaff godoc
// @Summary AddDigitalStaff used to categories help digital staff
// @Description Create a new digital staff category
// @Tags digitalStaff
// @Accept json
// @Produce json
// @Param digitallStaff body entity.DigitalStaff true "New DigitalStaff"
// @Success 200 {object} translate.ResultSuccess{data=entity.DigitalStaff} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /digital-staff [post]
func AddDigitalStaff(c echo.Context) error {
	digitalStaff := new(e.DigitalStaff)
	err := c.Bind(&digitalStaff)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var hasErr, tx = t.ValidateTranslator(c, digitalStaff)
	if hasErr != nil {
		return t.Errors(c, http.StatusBadRequest, hasErr)
	}
	digitalStaff.CreatedBy = u.GetUsername(c)

	err = r.AddDigitalStaff(digitalStaff)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			res, _ := tx.T(err.Constraint)
			return t.Errors(c, http.StatusBadRequest, res)
		}
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, digitalStaff)
}