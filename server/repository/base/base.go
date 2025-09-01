package base

import "gorm.io/gorm"

type Repository interface {
    Create(interface{}) error
    Update(interface{}) error
    Delete(interface{}) error
    FindByID(id uint, v interface{}) error
    List(v interface{}) error
}

type BaseRepository struct {
    DB *gorm.DB
}

func NewBaseRepository(DB *gorm.DB) *BaseRepository {
    return &BaseRepository{DB: DB}
}

func (r *BaseRepository) Create(v interface{}) error {
    return r.DB.Create(v).Error
}

func (r *BaseRepository) Update(v interface{}) error {
    return r.DB.Save(v).Error
}

func (r *BaseRepository) Delete(v interface{}) error {
    return r.DB.Delete(v).Error
} 

func (r *BaseRepository) FindByID(id uint, v interface{}) error {
    return r.DB.First(v, id).Error
}

func (r *BaseRepository) List(v interface{}) error {
    return r.DB.Find(v).Error
}