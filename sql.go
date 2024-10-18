package minutil

import (
	"errors"

	"gorm.io/gorm"
)

// 定义一个泛型函数来获取单个记录
func GetOne[T any](db *gorm.DB, dest *T, query interface{}, args ...interface{}) error {
	result := db.Where(query, args...).First(dest)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

// 定义一个泛型函数来获取所有记录
func GetAll[T any](db *gorm.DB, dest *[]T, query interface{}, args ...interface{}) error {
	result := db.Where(query, args...).Find(dest)
	return result.Error
}

// 定义一个泛型函数来创建记录
func Create[T any](db *gorm.DB, dest *T) error {
	result := db.Create(dest)
	return result.Error
}

// 定义一个泛型函数来更新记录
func Update[T any](db *gorm.DB, dest *T, query interface{}, args ...interface{}) error {
	result := db.Model(dest).Where(query, args...).Updates(dest)
	return result.Error
}

// 定义一个泛型函数来删除记录
func Delete[T any](db *gorm.DB, dest *T, query interface{}, args ...interface{}) error {
	result := db.Where(query, args...).Delete(dest)
	return result.Error
}

// 定义一个泛型函数来执行模糊查询
func Like[T any](db *gorm.DB, dest *[]T, field string, value string) error {
	result := db.Where(field+" LIKE ?", "%"+value+"%").Find(dest)
	return result.Error
}

// 定义一个泛型函数来执行多字段搜索
func Search[T any](db *gorm.DB, dest *[]T, fields []string, value string) error {
	query := db
	for _, field := range fields {
		query = query.Or(field+" LIKE ?", "%"+value+"%")
	}
	result := query.Find(dest)
	return result.Error
}