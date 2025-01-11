package db

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Restaurant struct {
	gorm.Model
	Name      string    `gorm:"unique;not null"`
	City      string    `gorm:"not null"`
	District  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"index;autoCreateTime"`
	Address   string    `gorm:""`
	Link      string    `gorm:""`
	Memo      string    `gorm:""`
	Rank      float32   `gorm:""`
	Tags      string    `gorm:""`
}

func GetAllRestaurants() []Restaurant {
	var restaurants []Restaurant
	DB.Order("name").Find(&restaurants)
	return restaurants
}

func DeleteRestaurant(id uint) error {
	result := DB.Delete(&Restaurant{}, id)
	if result.Error != nil {
		return fmt.Errorf("刪除餐廳失敗: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("找不到 ID 為 %d 的餐廳", id)
	}
	return nil
}
