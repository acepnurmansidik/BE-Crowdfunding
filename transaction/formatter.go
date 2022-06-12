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