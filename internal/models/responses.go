package models

import "github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/types"

type GetBalanceSuccessResponse struct {
	Success bool                 `json:"success"`
	Data    types.BalanceDataDTO `json:"data"`
	Message string               `json:"message,omitempty"`
}

type GetBalancesSuccessResponse struct {
	Success bool                   `json:"success"`
	Data    []types.BalanceDataDTO `json:"data"`
	Message string                 `json:"message,omitempty"`
}

type GetConfigSuccessResponse struct {
	Success bool               `json:"success"`
	Data    types.SystemConfig `json:"data"`
	Message string             `json:"message,omitempty"`
}

type GenericFailureResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"error description"`
}

type HistoryDataWrapper struct {
	History []types.TransactionHistoryDTO `json:"history"`
}

type GetHistorySuccessResponse struct {
	Success bool               `json:"success" example:"true"`
	Data    HistoryDataWrapper `json:"data"`
}

type GetUtxosSuccessResponse struct {
	Success bool            `json:"success" example:"true"`
	Data    []types.MneeTxo `json:"data"`
}

type GetTicketSuccessResponse struct {
	Success bool         `json:"success" example:"true"`
	Data    types.Ticket `json:"data"`
}

type TransferSyncSuccessResponse struct {
	Success bool                      `json:"success" example:"true"`
	Data    types.TransferResponseDTO `json:"data"`
}

type TicketIdWrapper struct {
	TicketId string `json:"ticketId" example:"KKJS-..."`
}

type TransferAsyncSuccessResponse struct {
	Success bool            `json:"success" example:"true"`
	Data    TicketIdWrapper `json:"data"`
}

type RawTxWrapper struct {
	RawTxHex string `json:"rawTxHex" example:"02000000..."`
}

type PartialSignSuccessResponse struct {
	Success bool         `json:"success" example:"true"`
	Data    RawTxWrapper `json:"data"`
}
