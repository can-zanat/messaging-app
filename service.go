package main

import (
	"errors"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

const timeInterval = 2 * time.Minute

type Service struct {
	store    Store
	mu       sync.Mutex
	stopChan chan struct{}
	running  bool
}

type Store interface {
	GetTwoMessages() (*[]Message, error)
	UpdateSentStatus(*[]Message) error
	GetSentMessages() (*[]Message, error)
}

func NewService(s Store) *Service {
	return &Service{store: s}
}

func (s *Service) StartSending() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.running {
		return errors.New("sending is already running")
	}

	s.running = true
	s.stopChan = make(chan struct{})

	go s.runSendingLoop()

	return nil
}

func (s *Service) StopSending() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.running {
		return errors.New("sending is already not running")
	}

	close(s.stopChan)
	s.running = false
	s.stopChan = nil

	return nil
}

func (s *Service) runSendingLoop() {
	ticker := time.NewTicker(timeInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// This function is called every 2 minutes
			msg, err := s.store.GetTwoMessages()

			if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
				log.Printf("error getting messages: %v", err)
				continue
			}

			if msg == nil || len(*msg) == 0 {
				log.Printf("no messages to send")
				continue
			}

			statusCode, err := SendMessages(msg)
			if err != nil {
				log.Printf("error message sending: %v", err)
				continue
			}

			if statusCode >= 200 && statusCode < 300 {
				err = s.store.UpdateSentStatus(msg)
				if err != nil {
					log.Printf("error update sent messages status: %v", err)
					continue
				}
			}
		case <-s.stopChan:
			// Return here to end the goroutine;
			return
		}
	}
}

func (s *Service) GetSentMessages() (*[]Message, error) {
	return s.store.GetSentMessages()
}
