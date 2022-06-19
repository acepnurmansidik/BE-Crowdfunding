package payment

import (
	"bwastartup/app/transaction"
	"bwastartup/app/user"

	"github.com/midtrans/midtrans-go"
)

type service struct{}

type Service interface{
	GetToken(transaction transaction.Transaction, user user.User) string
}

func NewService() *service{
	return &service{}
}

func (s *service) GetToken(transaction transaction.Transaction, user user.User) string{
    midclient := midtrans.NewClient()
    midclient.ServerKey = "YOUR-VT-SERVER-KEY"
    midclient.ClientKey = "YOUR-VT-CLIENT-KEY"
    midclient.APIEnvType = midtrans.Sandbox
}