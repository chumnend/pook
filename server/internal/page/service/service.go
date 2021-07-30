package service

import (
	"errors"

	"github.com/chumnend/pook/server/internal/domain"
)

type pageSrv struct {
	repo domain.PageRepository
}

// NewService returns a PageService utilizing provided PageRepository
func NewService(repo domain.PageRepository) domain.PageService {
	return &pageSrv{repo: repo}
}

func (srv *pageSrv) FindAllByBookID(id uint) ([]domain.Page, error) {
	pages, err := srv.repo.FindAllByBookID(id)
	if err != nil {
		return pages, err
	}
	return pages, nil
}

func (srv *pageSrv) FindByID(id uint) (*domain.Page, error) {
	page, err := srv.repo.FindByID(id)
	if err != nil {
		return page, err
	}
	return page, nil
}

func (srv *pageSrv) Create(page *domain.Page) error {
	return srv.repo.Create(page)
}

func (srv *pageSrv) Update(page *domain.Page) error {
	return srv.repo.Update(page)
}

func (srv *pageSrv) Delete(page *domain.Page) error {
	return srv.repo.Delete(page)
}

func (srv *pageSrv) Validate(page *domain.Page) error {
	if page == nil {
		return errors.New("page is empty")
	}

	if page.Content == "" || page.BookID == 0 {
		return errors.New("invalid page")
	}
	return nil
}
