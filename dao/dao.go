package dao

import (

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/naveenbalaji2001/Oauth-with-go/model"
)

type DatabaseAccessor struct {
	DB *gorm.DB
}

func NewDatabaseAccessor(db *gorm.DB) *DatabaseAccessor {
	return &DatabaseAccessor{DB: db}
}

func (d *DatabaseAccessor) SaveTrack(track *model.Track) error {
	return d.DB.Create(track).Error
}

// GetTrackByISRC retrieves a track from the database by ISRC code
func (d *DatabaseAccessor) GetTrackByISRC(isrc string) (*model.Track, error) {
	var track model.Track
	err := d.DB.Where("isrc = ?", isrc).First(&track).Error
	return &track, err
}


