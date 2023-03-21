package account

import (
	"app/commons/helpers"
	"app/db"
)

type AccountRepository struct{}

func (m AccountRepository) Create(data *AccountCreateInput) (account Account, err error) {

	account = Account{Name: data.Name}

	account.Name = data.Name

	if err = db.GetDB().Create(&account).Error; err != nil {
		return account, err
	}

	return account, err

}

func (m AccountRepository) FindOneById(id uint) (account Account, err error) {

	account = Account{ID: id}

	if err = db.GetDB().First(&account).Error; err != nil {
		return account, err
	}

	return account, err

}

func (m AccountRepository) Find(queries map[string][]string) (pagination helpers.Pagination, err error) {

	var accounts []*Account

	cond := helpers.SetConds(queries, &pagination)

	db.GetDB().Scopes(helpers.Paginate(accounts, &pagination, db.GetDB())).Where(cond).Find(&accounts)

	pagination.Data = accounts

	return pagination, nil

}

func (m AccountRepository) Update(id uint, data *AccountUpdateInput) (account Account, err error) {

	account = Account{ID: id}

	if err = db.GetDB().First(&account).Error; err != nil {
		return account, err
	}

	account.Name = data.Name

	if err = db.GetDB().Save(&account).Error; err != nil {
		return account, err
	}

	return account, err

}

func (m AccountRepository) Delete(id uint) (account Account, err error) {

	account = Account{ID: id}

	if err = db.GetDB().Delete(&account).Error; err != nil {
		return account, err
	}

	return account, err

}
