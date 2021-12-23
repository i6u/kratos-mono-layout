package db

import "gorm.io/gorm"

// Paginate 分页
func Paginate(page, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		if size > 100 || size < 0 {
			size = 100
		}
		offset := size * (page - 1)
		return db.Limit(size).Offset(offset)
	}
}
