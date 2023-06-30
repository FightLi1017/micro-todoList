package dao

import "user/model"

// 对象映射数据库表
func migration() {
	_db.Set(`gorm:table_options`, "charset=utf8mb4").
		AutoMigrate(&model.User{})
}
