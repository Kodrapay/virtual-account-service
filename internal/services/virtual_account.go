package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/kodra-pay/virtual-account-service/internal/dto"
)

type VirtualAccountService struct{}

func NewVirtualAccountService() *VirtualAccountService { return &VirtualAccountService{} }

func (s *VirtualAccountService) Create(_ context.Context, req dto.VirtualAccountRequest) dto.VirtualAccountResponse {
	return dto.VirtualAccountResponse{
		ID:            "va_" + uuid.NewString(),
		AccountName:   "KodraPay Merchant",
		AccountNumber: "1234567890",
		BankCode:      req.BankCode,
		Status:        "active",
	}
}

func (s *VirtualAccountService) Get(_ context.Context, id string) dto.VirtualAccountResponse {
	return dto.VirtualAccountResponse{
		ID:            id,
		AccountName:   "KodraPay Merchant",
		AccountNumber: "1234567890",
		BankCode:      "999",
		Status:        "active",
	}
}
