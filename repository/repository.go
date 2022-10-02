package repository

import (
	"rafit/models"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	CreateIp(req models.IPModel) (*models.IPModel, error)
	GetAllIp() ([]*models.IPModel, error)
	FindByIp(ip string) (models.IPModel, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateIp(ip models.IPModel) (*models.IPModel, error) {
	err := r.db.Create(&ip).Error
	if err != nil {
		return nil, err
	}
	return &ip, nil
}
func (r *repository) FindByIp(ip string) (models.IPModel, error) {
	var resIp models.IPModel
	err := r.db.Where("ip = ?", ip).Find(&resIp).Error
	if err != nil {
		return resIp, err
	}
	return resIp, nil
}

func (r *repository) GetAllIp() ([]*models.IPModel, error) {
	today := time.Now()

	ip := []*models.IPModel{}
	now := today.Format("2006-01-02")
	err := r.db.Where("created_at BETWEEN ? AND ?", now+" 00:00", now+" 23:59").Find(&ip).Error
	if err != nil {
		return nil, err
	}
	return ip, nil
}
