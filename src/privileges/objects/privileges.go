package objects

type Privilege struct {
	Id       int    `json:"id" gorm:"primary_key;index"`
	Username string `json:"username" gorm:"not null;unique"`
	Status   string `json:"status" gorm:"not null" sql:"DEFAULT:'BRONZE'"`
	Balance  int    `json:"balance"`
}

func (Privilege) TableName() string {
	return "privilege"
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

func ToPrivilegeInfoResponse(privilege *Privilege, history []PrivilegeHistory) *PrivilegeInfoResponse {
	balance_history := make([]BalanceHistory, len(history))
	for k, v := range history {
		balance_history[k] = *v.ToBalanceHistory()
	}

	return &PrivilegeInfoResponse{
		privilege.Balance,
		privilege.Status,
		balance_history,
	}
}
