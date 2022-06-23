package transaction

import (
	"bwastartup/app/campaign"
	"bwastartup/app/payment"
	"crypto/rand"
	"fmt"
	"math/big"
)

type service struct {
	repository Repository
	campaignRepository campaign.Repository
	paymentService payment.Service

}

type Service interface {
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetUserTransactionByUserID(userID int) ([]Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{repository, campaignRepository, paymentService}
}

func (s *service) GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error){
	// get campaign
	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	// check campaign user.id dengan user yang login
	if campaign.User.ID != input.User.ID{
		return []Transaction{}, err
	}
	
	transactions, err := s.repository.GetCampaignByID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetUserTransactionByUserID(userID int) ([]Transaction, error){
	transactions, err := s.repository.GetByUserID(userID)

	if err != nil{
		return transactions, err
	}

	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput)(Transaction, error){
	// create code random
	randomCrypto1, _ := rand.Int(rand.Reader, big.NewInt(9999999999))
	randomCrypto2, _ := rand.Int(rand.Reader, big.NewInt(99999))
	// mapping data transaction from input user to db
	trx := Transaction{}
	trx.Amount = input.Amount
	trx.CampaignID = input.CampaignID
	trx.User.ID = input.User.ID
	trx.Status = "pending"
	trx.Code = fmt.Sprintf("%v%s", randomCrypto1, randomCrypto2)

	transaction, err := s.repository.Save(trx)

	if err != nil{
		return transaction, err
	}

	return transaction, nil
}