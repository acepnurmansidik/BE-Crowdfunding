package transaction

import "bwastartup/app/campaign"

type service struct {
	repository Repository
	campaignRepository campaign.Repository

}

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetUserTransactionByUserID(userID int) ([]Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
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
