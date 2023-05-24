package usecase

import (
	"github.com/676767ap/otus-go-hw/project/internal/entity"
	"github.com/676767ap/otus-go-hw/project/util/log"
)

func (rep *Repos) AddBannerToSlot(bannerId int32, slotId int32) error {
	err := rep.BannerRepository.AddBannerToSlot(bannerId, slotId)
	if err != nil {
		log.Error(err)
	}
	return err
}

func (rep *Repos) RemoveBannerFromSlot(bannerId int32, slotId int32) error {
	err := rep.BannerRepository.RemoveBannerFromSlot(bannerId, slotId)
	if err != nil {
		log.Error(err)
	}
	return err
}

func (rep *Repos) ClickOnBanner(bannerId int32, slotId int32, socGroupId int32) (*entity.Stat, error) {
	stat, err := rep.BannerRepository.ClickOnBanner(bannerId, slotId, socGroupId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return stat, err
}

func (rep *Repos) ChooseBannerForSlot(socGroupId int32, slotId int32) (int, *entity.Stat, error) {
	bannerId, stat, err := rep.BannerRepository.ChooseBannerForSlot(socGroupId, slotId)
	if err != nil {
		log.Error(err)
		return 0, nil, err
	}
	return bannerId, stat, nil
}
