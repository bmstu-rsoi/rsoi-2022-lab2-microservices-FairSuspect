package repository

import (
	"privileges/errors"
	"privileges/objects"
	"time"

	"github.com/jinzhu/gorm"
)

type HistoryRep interface {
	Fetch(privilege_id int) ([]objects.PrivilegeHistory, error)
	Create(entry *objects.PrivilegeHistory) error
	Find(privilege_id int, ticket_uid string) (*objects.PrivilegeHistory, error)
}

type PGHistoryRep struct {
	db *gorm.DB
}

func NewPGHistoryRep(db *gorm.DB) *PGHistoryRep {
	return &PGHistoryRep{db}
}

func (rep *PGHistoryRep) Fetch(privilege_id int) ([]objects.PrivilegeHistory, error) {
	temp := []objects.PrivilegeHistory{}
	err := rep.db.Where(objects.PrivilegeHistory{PrivilegeID: privilege_id}).Find(&temp).Error
	return temp, err
}

func (rep *PGHistoryRep) Create(entry *objects.PrivilegeHistory) error {
	entry.Datetime = time.Now().Format(time.RFC3339)
	return rep.db.Create(entry).Error
}

func (rep *PGHistoryRep) Find(privilege_id int, ticket_uid string) (*objects.PrivilegeHistory, error) {
	temp := new(objects.PrivilegeHistory)
	err := rep.db.
		Where(&objects.PrivilegeHistory{PrivilegeID: privilege_id, TicketUID: ticket_uid}).
		Last(temp).
		Error

	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		temp, err = nil, errors.RecordNotFound
	default:
		temp, err = nil, errors.UnknownError
	}

	return temp, err
}
