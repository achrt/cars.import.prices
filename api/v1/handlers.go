package v1

import (
	"net/http"

	"cars.import.prices/api"
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
