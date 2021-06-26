package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/karlaugust1/code_pix_go/domain/model"
)

type PixKetRepositoryDb struct {
	Db *gorm.DB
}

func (r PixKetRepositoryDb) AddBank(bank *model.Bank) error {
	err := r.Db.Create(bank).Error
	if err != nil {
		return err
	}

	return nil
}

func (r PixKetRepositoryDb) AddAccount(account *model.Account) error {
	err := r.Db.Create(account).Error
	if err != nil {
		return err
	}

	return nil
}

func (r PixKetRepositoryDb) RegisterKey(pixKey *model.PixKey) error {
	err := r.Db.Create(pixKey).Error
	if err != nil {
		return err
	}

	return nil
}

func (r PixKetRepositoryDb) FindKeyByKind(key, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey

	r.Db.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)
	if pixKey.ID == "" {
		return nil, fmt.Errorf("no keys was found")
	}

	return &pixKey, nil
}

func (r PixKetRepositoryDb) FindAccount(id string) (*model.Account, error) {
	var account model.Account

	r.Db.Preload("Bank").First(&account, "id = ?", id)
	if account.ID == "" {
		return nil, fmt.Errorf("no account was found")
	}

	return &account, nil
}

func (r PixKetRepositoryDb) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank

	r.Db.First(&bank, "id = ?", id)
	if bank.ID == "" {
		return nil, fmt.Errorf("no bank was found")
	}

	return &bank, nil
}
