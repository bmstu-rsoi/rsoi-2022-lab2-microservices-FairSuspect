package models

import (
	"privileges/objects"
	"privileges/repository"
)

type HistoryM struct {
	rep repository.HistoryRep
}

func NewHistoryM(rep repository.HistoryRep) *HistoryM {
	return &HistoryM{rep}
}

func (model *HistoryM) Fetch(privilege_id int) []objects.PrivilegeHistory {
	history, _ := model.rep.Fetch(privilege_id)
	return history
}

func (model *HistoryM) Find(privilege_id int, ticket_uid string) (*objects.PrivilegeHistory, error) {
	return model.rep.Find(privilege_id, ticket_uid)
}

func (model *HistoryM) FillInBalance(privilege_id int, ticket_uid string, balance_diff int) error {
	entry := &objects.PrivilegeHistory{
		PrivilegeID:   privilege_id,
		TicketUID:     ticket_uid,
		BalanceDiff:   balance_diff,
		OperationType: "FILL_IN_BALANCE",
	}

	return model.rep.Create(entry)
}

func (model *HistoryM) FillByMoney(privilege_id int, ticket_uid string, balance_diff int) error {
	entry := &objects.PrivilegeHistory{
		PrivilegeID:   privilege_id,
		TicketUID:     ticket_uid,
		BalanceDiff:   balance_diff,
		OperationType: "FILLED_BY_MONEY",
	}

	return model.rep.Create(entry)
}

func (model *HistoryM) DebitTheAccount(privilege_id int, ticket_uid string, balance_diff int) error {
	entry := &objects.PrivilegeHistory{
		PrivilegeID:   privilege_id,
		TicketUID:     ticket_uid,
		BalanceDiff:   balance_diff,
		OperationType: "DEBIT_THE_ACCOUNT",
	}

	return model.rep.Create(entry)
}
