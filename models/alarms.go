package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Alarms struct {
	Id             int       `json:"id" gorm:"id"`
	Identification string    `json:"identification" gorm:"identification"`
	Site           string    `json:"site" gorm:"site"`
	Target         string    `json:"target" gorm:"target"`
	Name           string    `json:"name" gorm:"name"`
	Reason         string    `json:"reason" gorm:"reason"`
	Captain        string    `json:"captain" gorm:"captain"`
	Company        string    `json:"company" gorm:"company"`
	Status         int       `json:"status" gorm:"status"`
	CreateAt       time.Time `json:"create_at" gorm:"create_at"`
	AlarmStartAt   time.Time `json:"alarm_start_at" gorm:"alarm_start_at"`
	AlarmEndAt     time.Time `json:"alarm_end_at" gorm:"alarm_end_at"`
	Duration float64 `json:"duration" gorm:"duration"`
	SiteType string  `json:"site_type" gorm:"site_type"`
}

func (a Alarms) TableName() string {
	return "alarms"
}

func GetNewAlarms() ([]*Alarms, error) {
	var alarms []*Alarms
	err := db.Where("status = ?", 1).Find(&alarms).Order("create_at DESC").Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return alarms, nil
}

func GetResolvedAlarms() ([]*Alarms, error) {
	var alarms []*Alarms
	err := db.Model(&Alarms{}).Where("status = ?", -1).Find(&alarms).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return alarms, nil
}
