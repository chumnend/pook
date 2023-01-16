package service

import (
	"errors"

	"github.com/chumnend/pook/internal/entity"
)

type pageService struct {
	repo entity.PageRepository
}

// NewPageService returns a PagePageService utilizing provided PageRepository
func NewPageService(repo entity.PageRepository) entity.PageService {
	return &pageService{repo: repo}
}

func (srv *pageService) FindAllByBookID(id uint) ([]entity.Page, error) {
	pages, err := srv.repo.FindAllByBookID(id)
	if err != nil {
		return pages, err
	}
	return pages, nil
}

func (srv *pageService) FindByID(id uint) (*entity.Page, error) {
	page, err := srv.repo.FindByID(id)
	if err != nil {
		return page, err
	}
	return page, nil
}

func (srv *pageService) Create(page *entity.Page) error {
	return srv.repo.Create(page)
}

func (srv *pageService) Update(page *entity.Page) error {
	return srv.repo.Update(page)
}

func (srv *pageService) Delete(page *entity.Page) error {
	return srv.repo.Delete(page)
}

func (srv *pageService) Validate(page *entity.Page) error {
	if page == nil {
		return errors.New("page is empty")
	}

	if page.Content == "" || page.BookID == 0 {
		return errors.New("invalid page")
	}
	return nil
}
