package service

import "fmt"

type PartnerFactory interface {
	CreatePartner(patnerID int) (Partner, error)
}

type DefaultPartnerFactory struct {
	partnerBaseURLs map[int]string
}

func NewPartnerFactory(partnerBaseURLs map[int]string) PartnerFactory {
	return &DefaultPartnerFactory{
		partnerBaseURLs: partnerBaseURLs,
	}
}

func (f *DefaultPartnerFactory) CreatePartner(patnerID int) (Partner, error) {
	baseURL, ok := f.partnerBaseURLs[patnerID]

	if !ok {
		return nil, fmt.Errorf("partner with id %d not found", patnerID)
	}

	switch patnerID {
	case 1:
		return &Partner1{
			BaseURL: baseURL,
		}, nil
	case 2:
		return &Partner2{
			BaseURL: baseURL,
		}, nil
	default:
		return nil, fmt.Errorf("partner with id %d not found", patnerID)
	}
}
