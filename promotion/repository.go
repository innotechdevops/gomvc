package promotion

import "gorm.io/gorm"

type Repository interface {
	GetAll() []Promotion
	GetById(id int) Promotion
	Add(data Promotion) error
	Update(id int, data Promotion) error
	Delete(id int) error
}

type repository struct {
	DB *gorm.DB
}

func (r *repository) GetAll() []Promotion {
	var promotions []Promotion

	tx := r.DB.Find(&promotions)
	if tx.Error != nil {
		return []Promotion{}
	}
	return promotions
}

func (r *repository) GetById(id int) Promotion {
	var promo Promotion
	r.DB.First(&promo, "id = ?", id)
	return promo
}

func (r *repository) Add(data Promotion) error {
	tx := r.DB.Create(&data)
	return tx.Error
}

func (r *repository) Update(id int, data Promotion) error {
	promo := Promotion{Id: id}
	tx := r.DB.Model(&promo).Updates(data)
	return tx.Error
}

func (r *repository) Delete(id int) error {
	// Will update field "deleted_at"
	tx := r.DB.Where("id = ?", id).Delete(&Promotion{})
	return tx.Error
}

func NewRepository(db *gorm.DB) Repository {

	// Migrate the schema
	_ = db.AutoMigrate(&Promotion{})

	return &repository{db}
}
