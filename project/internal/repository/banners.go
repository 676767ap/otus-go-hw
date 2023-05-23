package repository

import (
	"github.com/676767ap/project/internal/entity"
	"github.com/676767ap/project/internal/services"
)

func (g *GormRepository) AddBannerToSlot(bannerId int32, slotId int32) error {
	var banner = &entity.Banner{}
	err := g.db.Model(banner).Where("id = ?", bannerId).First(banner).Error
	if err != nil {
		return err
	}

	var slot = &entity.Slot{}
	err = g.db.Model(slot).Where("id = ?", slotId).First(slot).Error
	if err != nil {
		return err
	}

	err = g.db.Exec("INSERT INTO banner_slots VALUES (?, ?)", bannerId, slotId).Error
	return err
}

func (g *GormRepository) CreateStat(stat *entity.Stat) error {
	err := g.db.Create(stat).Error
	return err
}

func (g *GormRepository) RemoveBannerFromSlot(bannerId int32, slotId int32) error {
	var banner = &entity.Banner{}
	err := g.db.Model(banner).Where("id = ?", bannerId).First(banner).Error
	if err != nil {
		return err
	}

	var slot = &entity.Slot{}
	err = g.db.Model(slot).Where("id = ?", slotId).First(slot).Error
	if err != nil {
		return err
	}

	err = g.db.Exec("DELETE FROM banner_slots WHERE banner_id = ? AND slot_id = ?", bannerId, slotId).Error
	return err
}

func (g *GormRepository) ClickOnBanner(bannerId int32, slotId int32, socGroupId int32) (*entity.Stat, error) {
	var banner = &entity.Banner{}
	err := g.db.Model(banner).Where("id = ?", bannerId).First(banner).Error
	if err != nil {
		return nil, err
	}

	var slot = &entity.Slot{}
	err = g.db.Model(slot).Where("id = ?", slotId).First(slot).Error
	if err != nil {
		return nil, err
	}

	var socGroup = &entity.SocGroup{}
	err = g.db.Model(socGroup).Where("id = ?", socGroupId).First(socGroup).Error
	if err != nil {
		return nil, err
	}
	stat := entity.Stat{
		Type:       "click",
		SlotID:     slotId,
		BannerID:   bannerId,
		SocGruopID: socGroupId,
	}
	err = g.CreateStat(&stat)
	if err != nil {
		return nil, err
	}

	return &stat, nil
}

func (g *GormRepository) ChooseBannerForSlot(socGroupId int32, slotId int32) (int, *entity.Stat, error) {
	type Result struct {
		ID int
	}
	var banners []Result
	err := g.db.Raw("SELECT FROM banner_slots WHERE slot_id = ?", slotId).Scan(&banners).Error
	if err != nil {
		return 0, nil, err
	}

	var bannersForCalc []*services.BannerForCalc
	for _, bannerId := range banners {
		var clicks int64
		var shown int64
		txa := g.db.Begin().Model(&entity.Stat{}).Where("banner_id = ?", bannerId.ID).Where("soc_group_id = ?", socGroupId).Where("type = ?", "click")
		txa.Count(&clicks)
		txa.Commit()
		txa = g.db.Begin().Model(&entity.Stat{}).Where("banner_id = ?", bannerId.ID).Where("soc_group_id = ?", socGroupId).Where("type = ?", "shown")
		txa.Count(&shown)
		txa.Commit()
		curBanner := services.BannerForCalc{
			BannerId: bannerId.ID,
			Shown:    int(shown),
			Clicks:   int(clicks),
		}
		bannersForCalc = append(bannersForCalc, &curBanner)
	}
	choosenBannerId := services.MaxUCB1(bannersForCalc)
	var socGroup = &entity.SocGroup{}
	err = g.db.Model(socGroup).Where("id = ?", socGroupId).First(socGroup).Error
	if err != nil {
		return 0, nil, err
	}
	stat := entity.Stat{
		Type:       "shown",
		SlotID:     slotId,
		BannerID:   int32(choosenBannerId),
		SocGruopID: socGroupId,
	}
	err = g.CreateStat(&stat)
	if err != nil {
		return 0, nil, err
	}
	return choosenBannerId, &stat, nil
}
