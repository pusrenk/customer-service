package entitites

import (
	"time"

	"gorm.io/gorm"
)

// Customer model
type Customer struct {
	ID        uint           `gorm:"primaryKey;autoIncrement;column:id"`
	Name      string         `gorm:"type:varchar(100);not null;column:name"`
	Email     string         `gorm:"type:varchar(100);not null;uniqueIndex;column:email"`
	Phone     string         `gorm:"type:varchar(20);not null;column:phone"`
	CreatedAt time.Time      `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;column:created_at"`
	CreatedBy string         `gorm:"type:varchar(100);not null;column:created_by"`
	UpdatedAt time.Time      `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;autoUpdateTime;column:updated_at"`
	UpdatedBy string         `gorm:"type:varchar(100);not null;column:updated_by"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp;index;column:deleted_at"`
}

// TableName set table name
func (Customer) TableName() string {
	return "customers"
}
