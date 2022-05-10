package service_impl

import "gitlab.ozon.dev/zBlur/homework-2/internal/service"

type Service struct {
	userService        *UserService
	currencyService    *CurrencyService
	marketItemService  *MarketItemService
	marketPriceService *MarketPriceService
}

func New() *Service {
	return &Service{}
}

func (s *Service) User() service.UserService {
	if s.userService != nil {
		return s.userService
	}

	s.userService = &UserService{}

	return s.userService
}

func (s *Service) Currency() service.CurrencyService {
	if s.currencyService != nil {
		return s.currencyService
	}

	s.currencyService = &CurrencyService{}

	return s.currencyService
}

func (s *Service) MarketItem() service.MarketItemService {
	if s.marketItemService != nil {
		return s.marketItemService
	}

	s.marketItemService = &MarketItemService{}

	return s.marketItemService
}

func (s *Service) MarketPrice() service.MarketPriceService {
	if s.marketPriceService != nil {
		return s.marketPriceService
	}

	s.marketPriceService = &MarketPriceService{}

	return s.marketPriceService
}
