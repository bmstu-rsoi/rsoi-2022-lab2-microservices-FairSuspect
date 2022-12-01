package objects

type PrivilegeHistory struct {
	Id            int       `json:"id" gorm:"primary_key;index"`
	Privilege     Privilege `json:"privilege" gorm:"foreignKey:PrivilegeID"`
	PrivilegeID   int       `gorm:"index"`
	TicketUID     string    `json:"ticketUID" gorm:"not null"`
	Datetime      string    `json:"datetime" gorm:"not null"`
	BalanceDiff   int       `json:"balanceDiff" gorm:"not null"`
	OperationType string    `json:"operationType" gorm:"not null"`
}

func (PrivilegeHistory) TableName() string {
	return "privilege_history"
}

type BalanceHistory struct {
	Date          string `json:"date"`
	BalanceDiff   int    `json:"balanceDiff"`
	TicketUid     string `json:"ticketUid"`
	OperationType string `json:"operationType"`
}

func (history *PrivilegeHistory) ToBalanceHistory() *BalanceHistory {
	return &BalanceHistory{
		history.Datetime,
		history.BalanceDiff,
		history.TicketUID,
		history.OperationType,
	}
}

type AddTicketRequest struct {
	TicketUID       string `json:"ticketUID"`
	Price           int    `json:"price"`
	PaidFromBalance bool   `json:"paidFromBalance"`
}

type AddTicketResponce struct {
	PaidByMoney   int                `json:"paidByMoney"`
	PaidByBonuses int                `json:"paidByBonuses"`
	Privilege     PrivilegeShortInfo `json:"privilege"`
}
