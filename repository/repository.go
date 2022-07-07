package repository

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"

	database "test.com/test/Database"
	model "test.com/test/Model"
	"test.com/test/request"
)

type IRepository interface {
	CreateContact(ctx context.Context, contact model.Contact) error
	DeleteContact(ctx context.Context, name string) error
	GetContact(ctx context.Context) ([]string, []string, error)
	GetContactByName(ctx context.Context, name string) (*[]string, error)
	UpdatedContact(ctx context.Context, newContact request.UpdateContact) error
}
type db struct {
	db *gorm.DB
}

func NewRepository() *db {
	return &db{
		db: database.GetConnection()}
}
func (conn *db) CreateContact(ctx context.Context, contact model.Contact) error {

	err := conn.db.Table(model.TABLE_CONTACT).Create(&contact)
	if err.Error != nil {
		return err.Error

	}

	return nil
}

func (conn *db) GetContact(ctx context.Context) ([]string, []string, error) {

	rowsName, err := conn.db.Table(model.TABLE_CONTACT).Select(model.TABLE_CONTACT_NAME).Rows()
	rowsNUmber, err := conn.db.Table(model.TABLE_CONTACT).Select(model.TABLE_CONTACT_NUMBER).Rows()

	var Contactname []string
	if err != nil {
		return []string{}, []string{}, err
	}
	for rowsName.Next() {
		var s string
		if err := rowsName.Scan(&s); err != nil {
			return []string{}, []string{}, err
		}
		Contactname = append(Contactname, s)
	}

	var ContactNumber []string
	if err != nil {
		return []string{}, []string{}, err
	}
	for rowsNUmber.Next() {
		var s string
		if err := rowsNUmber.Scan(&s); err != nil {
			return []string{}, []string{}, err
		}
		ContactNumber = append(ContactNumber, s)
	}

	return Contactname, ContactNumber, nil
}

func (conn *db) DeleteContact(ctx context.Context, name string) error {

	tx := conn.db.Begin()
	defer tx.Rollback()

	query := fmt.Sprintf("DELETE FROM  %s WHERE NAME= '%s'", model.TABLE_CONTACT, name)

	if err := tx.Exec(query); err.Error != nil {

		return err.Error
	}
	tx.Commit()
	return nil
}
func (conn *db) UpdatedContact(ctx context.Context, newContact request.UpdateContact) error {

	tx := conn.db.Begin()
	defer tx.Rollback()

	if err := tx.Table(model.TABLE_CONTACT).Where(model.TABLE_CONTACT_NUMBER+" = ?", newContact.Previous_number).UpdateColumn(model.TABLE_CONTACT_NUMBER, newContact.Number); err.Error != nil {

		return err.Error
	}
	tx.Commit()
	return nil
}

func (conn *db) GetContactByName(ctx context.Context, contactname string) (*[]string, error) {

	var number []string

	rowsNUmber, err := conn.db.Table(model.TABLE_CONTACT).Select(model.TABLE_CONTACT_NUMBER).Where(model.TABLE_CONTACT_NAME+" = ?", contactname).Rows()

	if err != nil {
		return &[]string{}, err
	}
	for rowsNUmber.Next() {
		var s string
		if err := rowsNUmber.Scan(&s); err != nil {
			return &[]string{}, err
		}
		number = append(number, s)
	}
	return &number, nil

}
