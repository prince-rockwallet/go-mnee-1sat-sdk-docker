package models

import "github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/types"

type GetBalanceSuccessResponse struct {
	Success bool                 `json:"success"`
	Data    types.BalanceDataDTO `json:"data"`
	Message string               `json:"message,omitempty"`
}

type GetBalanceFailureResponse struct {
	Success bool   `json:"success" example:"false" default:"false"`
	Message string `json:"message,omitempty" example:"failed to fetch balances"`
}

type GetBalancesSuccessResponse struct {
	Success bool                   `json:"success"`
	Data    []types.BalanceDataDTO `json:"data"`
	Message string                 `json:"message,omitempty"`
}

type GetBalancesFailureResponse struct {
	Success bool   `json:"success" example:"false" default:"false"`
	Message string `json:"message,omitempty" example:"failed to fetch balances"`
}

type GetConfigSuccessResponse struct {
	Success bool               `json:"success"`
	Data    types.SystemConfig `json:"data"`
	Message string             `json:"message,omitempty"`
}

type GetConfigFailureResponse struct {
	Success bool   `json:"success" example:"false" default:"false"`
	Message string `json:"message,omitempty" example:"failed to fetch config"`
}
