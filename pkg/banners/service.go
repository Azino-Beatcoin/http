package banners

import (
	"context"
	"errors"
	"sync"
)

// Service struct
type Service struct {
	mu    sync.RWMutex
	items []*Banner
	id    int64
}

// NewService func
func NewService() *Service {
	return &Service{items: make([]*Banner, 0)}
}

// Banner struct
type Banner struct {
	ID      int64
	Title   string
	Content string
	Button  string
	Link    string
}

// ByID method
func (s *Service) ByID(ctx context.Context, id int64) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, banner := range s.items {
		if banner.ID == id {
			return banner, nil
		}
	}

	return nil, errors.New("item not found")
}

// All method
func (s *Service) All(ctx context.Context) ([]*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.items, nil
}

// RemoveByID method
func (s *Service) RemoveByID(ctx context.Context, id int64) (*Banner, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for index, banner := range s.items {
		if banner.ID == id {
			tmp := append([]*Banner{}, s.items[:index]...)
			tmp = append(tmp, s.items[index+1:]...)
			removed := s.items[index]
			s.items = tmp
			return removed, nil
		}
	}

	return nil, errors.New("post not found")
}

// Save method
func (s *Service) Save(ctx context.Context, item *Banner) (*Banner, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if item.ID == 0 {
		s.id++
		item.ID = s.id
		s.items = append(s.items, item)
		return item, nil
	}

	for _, banner := range s.items {
		if banner.ID == item.ID {
			banner.Button = item.Button
			banner.Content = item.Content
			banner.Link = item.Link
			banner.Title = item.Title
			return banner, nil
		}
	}

	return nil, errors.New("post not found")
}
