package cars_catalog

import (
	"net/http"

	"cars.import.prices/domain"
)

type Service struct {
	Host string
	ctx  domain.Context
}

func New(host string) *Service {
	return &Service{
		Host: host,
	}
}

func (s *Service) WithContext(ctx domain.Context) *Service {
	s.ctx = ctx
	return s
}

func (s *Service) GetUrl(postfix string) string {
	return "http://" + s.Host + postfix
}

func (s *Service) get(path string, logger domain.Logger) (*http.Response, error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", s.GetUrl(path), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("X-Request-Id", logger.Trace())
	request.Header.Set("X-Auth-Internal-Request", "cars.catalog")

	return client.Do(request)
}
