package main

type Service struct {
	store Store
}

type Store interface {
	GetTwoMessages() error
	LogSentMessages() error
	GetSentMessages() error
}

func NewService(s Store) *Service {
	return &Service{store: s}
}

func (s *Service) StartSending() error {
	return nil
}

func (s *Service) StopSending() error {
	return nil
}

func (s *Service) GetSentMessages() error {
	return nil
}
