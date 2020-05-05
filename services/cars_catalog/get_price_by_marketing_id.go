package cars_catalog

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"cars.import.prices/domain"
	"cars.import.prices/domain/services"
)

type GetPricesByMarkIdResponse struct {
	Data *services.Price
}

func (s *Service) GetPricesByMarkId(markId string, logger domain.Logger) (*services.Price, error) {
	r, err := s.get(fmt.Sprintf("/api/v1//marketing_complectations/%s/prices", markId), logger)

	if err != nil {
		return nil, err
	}

	defer r.Body.Close()
	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return nil, err
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("invalid cars.catalog response status (GetUpholsteryByCode). Body: %s", string(bytes)))
	}

	response := GetPricesByMarkIdResponse{}
	err = json.Unmarshal(bytes, &response)

	if err != nil {
		return nil, err
	}
	return response.Data, nil
}
