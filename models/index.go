package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// BaseModel definition
type BaseModel struct {
	ID         uint       `gorm:"id" json:"id"`
	CreatedOn  *time.Time `gorm:"created_on" json:"created_on"`
	ModifiedOn *time.Time `gorm:"modified_on" json:"modified_on"`
	Deleted    uint       `gorm:"deleted" json:"deleted"`
}

//BeforeCreate : Hooks
func (m *BaseModel) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("CreatedOn", timeLocalNow())
	return nil
}

//BeforeUpdate : Hooks
func (m *BaseModel) BeforeUpdate(scope *gorm.Scope) (err error) {
	scope.SetColumn("ModifiedOn", timeLocalNow())
	return nil
}

// TimeLocalNow ...
func timeLocalNow() *time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	return &now
}
