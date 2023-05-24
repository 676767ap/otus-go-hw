package repository

import "github.com/676767ap/otus-go-hw/project/internal/entity"

type BannerRepository interface {
	AddBannerToSlot(bannerId int32, slotId int32) error
	RemoveBannerFromSlot(bannerId int32, slotId int32) error
	ClickOnBanner(bannerId int32, slotId int32, socGroupId int32) (*entity.Stat, error)
	ChooseBannerForSlot(bannerId int32, slotId int32) (int, *entity.Stat, error)
}
