package cars_catalog

import (
	"net/http"

	"cars.import.prices/domain/model"
)

type Service struct {
	Host string
}

func New(host string) *Service {
	return &Service{
		Host: host,
	}
}

func (s *Service) GetUrl(postfix string) string {
	return "http://" + s.Host + postfix
}

func (s *Service) get(path string, logger model.Logger) (*http.Response, error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", s.GetUrl(path), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("X-Request-Id", logger.Trace())
	request.Header.Set("X-Auth-Internal-Request", "cars.catalog")

	return client.Do(request)
}
package cars_catalog
