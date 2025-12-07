package dto

type VirtualAccountRequest struct {
	MerchantID int `json:"merchant_id"`
	BankCode   string `json:"bank_code"`
	Currency   string `json:"currency"`
}

type VirtualAccountResponse struct {
	ID            int `json:"id"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	BankCode      string `json:"bank_code"`
	Status        string `json:"status"`
}
