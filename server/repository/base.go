package repository

import "gorm.io/gorm"

type Repository interface {
    Create(interface{}) error
    Update(interface{}) error
    Delete(interface{}) error
    FindByID(id uint, v interface{}) error
    List(v interface{}) error
}

type BaseRepository struct {
    db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) *BaseRepository {
    return &BaseRepository{db: db}
}

func (r *BaseRepository) Create(v interface{}) error {
    return r.db.Create(v).Error
}

func (r *BaseRepository) Update(v interface{}) error {
    return r.db.Save(v).Error
}

func (r *BaseRepository) Delete(v interface{}) error {
    return r.db.Delete(v).Error
} 

func (r *BaseRepository) FindByID(id uint, v interface{}) error {
    return r.db.First(v, id).Error
}

func (r *BaseRepository) List(v interface{}) error {
    return r.db.Find(v).Error
}