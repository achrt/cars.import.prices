package cars_catalog

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"cars.import.prices/domain"
	"cars.import.prices/domain/services"
)

func (s *Service) GetPriceTypeByCode(code []string, logger domain.Logger) ([]*services.PriceType, error) {
	codes := strings.Join(code, ",")
	r, err := s.get(fmt.Sprintf("/api/v1/price_types/by/code?code=%s", codes), logger)

	if err != nil {
		return nil, err
	}

	defer r.Body.Close()
	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return nil, err
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("invalid cars.catalog response status (GetPriceTypeByCode). Body: %s", string(bytes)))
	}

	response := services.PriceTypeResponse{}
	err = json.Unmarshal(bytes, &response)

	if err != nil {
		return nil, err
	}
	return response.Data, nil
}
