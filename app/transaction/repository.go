package transaction

import "gorm.io/gorm"

type Repository interface{
	GetCampaignByID(CampaignID int) ([]Transaction, error)
	GetUserTransaction(userID int) ([]Transaction, error)
}

type repository struct{
	db *gorm.DB
}

func NewRepository (db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetCampaignByID(CampaignID int) ([]Transaction, error){
	var transactions []Transaction
	err := r.db.Preload("User").Where("campaign_id = ?", CampaignID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) GetUserTransaction(userID int) ([]Transaction, error){
	var transactions []Transaction

	err := r.db.Preload("User").Where("user_id = ?", userID).Order("id desc").Find(&transactions).Error

	if err != nil {
		return []Transaction{}, err
	}

	return transactions, nil
}