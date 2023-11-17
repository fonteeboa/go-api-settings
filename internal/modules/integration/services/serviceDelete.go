package services

import (
	types "golang-api-settings/internal/modules/integration/domain"
)

func (s *IntegrationService) Delete(filter types.Integration) error {
	return s.repositories.DeleteByID(filter.ID)
}

func (s *IntegrationService) DeleteAll(filters []types.Integration) error {
	var err error
	for _, filter := range filters {
		err := s.repositories.DeleteByID(filter.ID)
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	return nil
}
