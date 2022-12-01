package objects

import (
	_ "encoding/json"
)

type BalanceHistory struct {
	Date          string `json:"date"`
	BalanceDiff   int    `json:"balanceDiff"`
	TicketUid     string `json:"ticketUid"`
	OperationType string `json:"operationType"`
}

type PrivilegeShortInfo struct {
	Balance int    `json:"balance"`
	Status  string `json:"status"`
}

type PrivilegeInfoResponse struct {
	Balance int              `json:"balance"`
	Status  string           `json:"status"`
	History []BalanceHistory `json:"history"`
}

type AddHistoryRequest struct {
	TicketUID       string `json:"ticketUID"`
	Price           int    `json:"price"`
	PaidFromBalance bool   `json:"paidFromBalance"`
}

type AddHistoryResponce struct {
	PaidByMoney   int                `json:"paidByMoney"`
	PaidByBonuses int                `json:"paidByBonuses"`
	Privilege     PrivilegeShortInfo `json:"privilege"`
}

type UserInfoResponse struct {
	Tickets   []TicketResponse   `json:"tickets"`
	Privilege PrivilegeShortInfo `json:"privilege"`
}
