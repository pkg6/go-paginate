package gormp

import (
	"github.com/pkg6/go-paginate"
	"gorm.io/gorm"
)

type GORM struct {
	db *gorm.DB
}

func Adapter(db *gorm.DB) paginate.IAdapter {
	return &GORM{db: db}
}

func (a GORM) Length() (int64, error) {
	var count int64
	if err := a.db.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (a GORM) Slice(offset, length int64, dest any) error {
	return a.db.Session(&gorm.Session{}).Limit(int(length)).Offset(int(offset)).Find(dest).Error
}
