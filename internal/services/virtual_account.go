package services

import (
	"context"

	"github.com/kodra-pay/virtual-account-service/internal/dto"
)

type VirtualAccountService struct{}

func NewVirtualAccountService() *VirtualAccountService { return &VirtualAccountService{} }

func (s *VirtualAccountService) Create(_ context.Context, req dto.VirtualAccountRequest) dto.VirtualAccountResponse {
	// In a real scenario, this would generate a unique int ID.
	// For this mock implementation, we return a placeholder.
	// req.MerchantID (int) is now available in the request.
	return dto.VirtualAccountResponse{
		ID:            1, // Placeholder for an auto-generated int ID
		AccountName:   "KodraPay Merchant",
		AccountNumber: "1234567890",
		BankCode:      req.BankCode,
		Status:        "active",
	}
}

func (s *VirtualAccountService) Get(_ context.Context, id int) dto.VirtualAccountResponse {
	return dto.VirtualAccountResponse{
		ID:            id, // id is int
		AccountName:   "KodraPay Merchant",
		AccountNumber: "1234567890",
		BankCode:      "999",
		Status:        "active",
	}
}
