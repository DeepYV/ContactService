package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
	"test.com/test/request"

	"github.com/stretchr/testify/require"
	"test.com/test/Model"
)

func TestDB(t *testing.T) {

	repo := NewRepository()
	if (repo.db == &gorm.DB{}) {

		t.Error("Erro while creating Db connection ")
	}
}

func TestCreatecontact(t *testing.T) {

	err := NewRepository().CreateContact(context.TODO(), Model.Contact{
		Number: "1234567809",
		Name:   "Test4"})

	Delerr := NewRepository().DeleteContact(context.TODO(), "Test4")

	require.NoError(t, err)
	require.NoError(t, Delerr)

}

func TestGetContact(t *testing.T) {

	_, _, err := NewRepository().GetContact(context.TODO())
	require.NoError(t, err)
	v := reflect.ValueOf(*t)
	name := v.FieldByName("name")
	t.Log("pass", name)

}

func TestUpdat(t *testing.T) {
	previousnumber := "1234567809"
	number := "123456799"
	err := NewRepository().CreateContact(context.TODO(), Model.Contact{
		Number: "1234567809",
		Name:   "Test4"})

	Uerr := NewRepository().UpdatedContact(context.TODO(), request.UpdateContact{
		Previous_number: &previousnumber,
		Number:          &number,
	})
	Delerr := NewRepository().DeleteContact(context.TODO(), "Test4")
	require.NoError(t, err)
	require.NoError(t, Uerr)
	require.NoError(t, Delerr)
	v := reflect.ValueOf(*t)
	name := v.FieldByName("name")
	t.Log("pass", name)
}
