package persistence

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type BaseRepository[T any, ID comparable] struct {
	db *gorm.DB
}

func NewBaseRepository[T any, ID comparable](db *gorm.DB) *BaseRepository[T, ID] {
	return &BaseRepository[T, ID]{
		db: db,
	}
}

func (r *BaseRepository[T, ID]) GetByID(ctx context.Context, id ID) (*T, error) {
	var entity T
	result := r.db.WithContext(ctx).First(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entity, nil
}

func (r *BaseRepository[T, ID]) GetAll(ctx context.Context) ([]T, error) {
	var entities []T
	result := r.db.WithContext(ctx).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return entities, nil
}

func (r *BaseRepository[T, ID]) GetPaged(ctx context.Context, pageNumber, pageSize int) ([]T, int64, error) {
	var entities []T
	var totalCount int64

	if err := r.db.WithContext(ctx).Model(new(T)).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	offset := (pageNumber - 1) * pageSize
	result := r.db.WithContext(ctx).
		Offset(offset).
		Limit(pageSize).
		Find(&entities)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return entities, totalCount, nil
}

func (r *BaseRepository[T, ID]) Find(ctx context.Context, condition interface{}, args ...interface{}) ([]T, error) {
	var entities []T
	result := r.db.WithContext(ctx).Where(condition, args...).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return entities, nil
}

func (r *BaseRepository[T, ID]) FirstOrDefault(ctx context.Context, condition interface{}, args ...interface{}) (*T, error) {
	var entity T
	result := r.db.WithContext(ctx).Where(condition, args...).First(&entity)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &entity, nil
}

func (r *BaseRepository[T, ID]) Count(ctx context.Context) (int64, error) {
	var count int64
	result := r.db.WithContext(ctx).Model(new(T)).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (r *BaseRepository[T, ID]) Insert(ctx context.Context, entity *T) error {
	result := r.db.WithContext(ctx).Create(entity)
	return result.Error
}

func (r *BaseRepository[T, ID]) InsertMany(ctx context.Context, entities []T) error {
	result := r.db.WithContext(ctx).Create(&entities)
	return result.Error
}

func (r *BaseRepository[T, ID]) Update(ctx context.Context, entity *T) error {
	result := r.db.WithContext(ctx).Save(entity)
	return result.Error
}

func (r *BaseRepository[T, ID]) Delete(ctx context.Context, id ID) error {
	result := r.db.WithContext(ctx).Delete(new(T), id)
	return result.Error
}

func (r *BaseRepository[T, ID]) SoftDelete(ctx context.Context, id ID, userID int) error {

	return fmt.Errorf("soft delete not implemented for this entity type")
}

func (r *BaseRepository[T, ID]) GetDB() *gorm.DB {
	return r.db
}
