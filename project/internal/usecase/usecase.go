package usecase

import (
	"github.com/676767ap/otus-go-hw/project/internal/repository"

	"gorm.io/gorm"
)

type Repos struct {
	BannerRepository repository.BannerRepository
}

func NewRepos(db *gorm.DB) Repos {
	return Repos{
		BannerRepository: repository.NewGormRepository(db),
	}
}
