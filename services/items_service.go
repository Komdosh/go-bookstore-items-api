package services

import (
	"github.com/Komdosh/go-bookstore-items-api/domain/items"
	"github.com/Komdosh/go-bookstore-items-api/domain/queries"
	"github.com/Komdosh/go-bookstore-utils/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
	Search(queries.EsQuery) ([]items.Item, rest_errors.RestErr)
	Delete(id string) rest_errors.RestErr
	Update(partial bool, request items.Item) (*items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	item := items.Item{Id: id}

	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Search(query queries.EsQuery) ([]items.Item, rest_errors.RestErr) {
	dao := items.Item{}
	return dao.Search(query)
}

func (s *itemsService) Delete(id string) rest_errors.RestErr {
	item := items.Item{Id: id}

	if err := item.Delete(); err != nil {
		return err
	}
	return nil
}

func (s *itemsService) Update(partial bool, item items.Item) (*items.Item, rest_errors.RestErr) {
	current, err := s.Get(item.Id)

	if err != nil {
		return nil, err
	}

	if partial {
		if item.Title != "" {
			current.Title = item.Title
		}
		if item.Description.Html != "" {
			current.Description.Html = item.Description.Html
		}
		if item.Description.PlainText != "" {
			current.Description.PlainText = item.Description.PlainText
		}
		if len(item.Pictures) != len(current.Pictures) && len(item.Pictures) > 0 {
			current.Pictures = item.Pictures
		}
		if item.Video != "" {
			current.Video = item.Video
		}
		if item.Price != 0 {
			current.Price = item.Price
		}
		if item.AvailableQuantity != 0 {
			current.AvailableQuantity = item.AvailableQuantity
		}
		if item.SoldQuantity != 0 {
			current.SoldQuantity = item.SoldQuantity
		}
		if item.Status != "" {
			current.Status = item.Status
		}

	} else {
		current.Title = item.Title
		current.Description.Html = item.Description.Html
		current.Description.PlainText = item.Description.PlainText
		current.Pictures = item.Pictures
		current.Video = item.Video
		current.Price = item.Price
		current.AvailableQuantity = item.AvailableQuantity
		current.SoldQuantity = item.SoldQuantity
		current.Status = item.Status
	}
	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}
