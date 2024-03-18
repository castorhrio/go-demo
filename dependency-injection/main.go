package main

import "fmt"

type Transfer struct {
	AccountId   string
	CanTransfer bool
}

type TransferClient interface {
	GetTransfer(accountId string) (Transfer, error)
}

type Account struct {
	UserId    string
	AccountId string
}

type AccountRepository interface {
	GetByUserId(userId string) ([]Account, error)
}

type Service interface {
	GetTransferAccounts(userId string) ([]Account, error)
}

type service struct {
	accountRepository AccountRepository
	transferClient    TransferClient
}

// dependency injection
func NewService(accountRepo AccountRepository, transferClient TransferClient) service {
	return service{
		accountRepo,
		transferClient,
	}
}

func (s *service) GetTransferAccounts(userId string) ([]Account, error) {
	accounts, err := s.accountRepository.GetByUserId(userId)
	if err != nil {
		return nil, err
	}

	output := []Account{}

	for _, account := range accounts {
		transfer, err := s.transferClient.GetTransfer(account.AccountId)
		if err != nil {
			return nil, err
		}

		if transfer.CanTransfer {
			output = append(output, account)
		}
	}

	return output, nil
}

// test
type mockAccountRepository struct{}

func (r *mockAccountRepository) GetByUserId(userId string) ([]Account, error) {
	return []Account{{"1", "0x43sdj40sdfhni4sdf"}}, nil
}

type mockTransferClient struct{}

func (r *mockTransferClient) GetTransfer(accountId string) (Transfer, error) {
	return Transfer{"0x43sdj40sdfhni4sdf", true}, nil
}

func main() {
	accountRepository := mockAccountRepository{}
	transferClient := mockTransferClient{}

	service := NewService(&accountRepository, &transferClient)
	result, _ := service.GetTransferAccounts("1")
	fmt.Println(result)
}
