package v1

import (
	"net/http"
	"strconv"

	"cars.import.prices/api"
	"cars.import.prices/domain/cases/filter_by_year"
	"cars.import.prices/domain/cases/replace_price"
	"cars.import.prices/presenters/jsonapi"
)

func ReplacePrice(c *api.Context, a *api.Application, w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	reqData, err := jsonapi.UnmarshalCarHook(c.Body)
	ucRequest := &replace_price.Request{
		CarHook: reqData.Model(),
	}
	result, err := replace_price.Run(c, ucRequest)
	if err != nil {
		return InternalServerError(err)
	}
	return OK(jsonapi.MarshalCarHook(result.CarHook))
}

func FilterByYear(c *api.Context, a *api.Application, w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	reqData, err := jsonapi.UnmarshalCarHook(c.Body)
	yearFrom, err := strconv.Atoi(c.Vars["year_from"])
	if err != nil {
		return InternalServerError(err)
	}
	ucRequest := &filter_by_year.Request{
		CarHook:  reqData.Model(),
		YearFrom: yearFrom,
	}
	result, err := filter_by_year.Run(c, ucRequest)
	if err != nil {
		return InternalServerError(err)
	}
	return OK(jsonapi.MarshalCarHook(result.CarHook))
}
