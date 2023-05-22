package repository

import (
	"testing"

	"gorm.io/gorm"
)

func TestGormRepository_AddBannerToSlot(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		bannerId int32
		slotId   int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GormRepository{
				db: tt.fields.db,
			}
			if err := g.AddBannerToSlot(tt.args.bannerId, tt.args.slotId); (err != nil) != tt.wantErr {
				t.Errorf("GormRepository.AddBannerToSlot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
