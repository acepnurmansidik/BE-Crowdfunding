package transaction

import "time"


type CampaignTransactionFormatter struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Amount int `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

// format satuannya
func FormatCampaignTransaction(transactionCampaign Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{}
	formatter.ID = transactionCampaign.ID
	formatter.Name = transactionCampaign.User.Name
	formatter.Amount = transactionCampaign.Amount
	formatter.CreatedAt = transactionCampaign.CreatedAt

	return formatter
}

// format array transactions
func FormatCampaignTransactions(transactionsCampaign []Transaction) []CampaignTransactionFormatter {
	// cek jika tidak ada
	if len(transactionsCampaign) == 0 {
		return []CampaignTransactionFormatter{}
	}

	formatter := []CampaignTransactionFormatter{}

	for _, trx := range transactionsCampaign {
		// panggil objectnya
		campaignTransactionFormatter := FormatCampaignTransaction(trx)
		// masukan ke slice
		formatter = append(formatter, campaignTransactionFormatter)
	}

	return formatter
}

type UserTransactionFormatter struct {
	ID int `json:"id"`
	Amount int `json:"amount"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	Campaign CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter{
	formatter := UserTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = transaction.Campaign.Name
	campaignFormatter.ImageURL = ""

	// cek gambarnya
	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = campaignFormatter

	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter{
	var formatter = []UserTransactionFormatter{}

	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	for _, transaction := range transactions {
		userTransaction := FormatUserTransaction(transaction)

		formatter = append(formatter, userTransaction)
	}


	return formatter
}