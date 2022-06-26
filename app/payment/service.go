package payment

import (
	"bwastartup/app/transaction"
	"bwastartup/app/user"
	"strconv"

	midtrans "github.com/veritrans/go-midtrans"
)

type service struct{
	transactionRepository transaction.Repository
}

type Service interface{
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
	ProcessPayment(input transaction.TransactionNotificationInput) error
}

func NewService(transactionRepository transaction.Repository) *service{
	return &service{transactionRepository}
}

func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error){
	midclient := midtrans.NewClient()
	midclient.ServerKey = ""
    midclient.ClientKey = ""
    midclient.APIEnvType = midtrans.Sandbox

    snapGateway := midtrans.SnapGateway{
        Client: midclient,
    }

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},

		TransactionDetails: midtrans.TransactionDetails{
			OrderID: strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "",err
	}

	return snapTokenResp.RedirectURL, err
}

func (s *service) ProcessPayment(input transaction.TransactionNotificationInput) error{
	// get transaction id from input
	transaction_id, _ := strconv.Atoi(input.OrderID)

	// use transaction id for collect transaction
	transaction, err := s.transactionRepository.GetByID(transaction_id)
	if err != nil {
		return err
	}

	// response from midtrans for update status transaction
	if(input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept"){
		transaction.Status = "paid"
	}else if(input.TransactionStatus == "settlement"){
		transaction.Status = "paid"
	}else if(input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel"){
		transaction.Status = "cancelled"
	}

	// update status transaction from response midtrans
	updateTransaction, err := s.transactionRepository.Update(transaction)
	if err != nil {
		return err
	}


}