package companymethods

import (
	"errors"
	"fmt"
	"template-app/models/company"
	"template-app/pkg/store"

	"github.com/TheLazarusNetwork/go-helpers/logo"
	"gorm.io/gorm"
)

type sortBy string

const (
	SORT_BY_PRICE               sortBy = "price"
	SORT_BY_USERS               sortBy = "active_users"
	SORT_BY_CURRENT_YEAR_PROFIT sortBy = "current_year_profit"
	SORT_BY_NONE                sortBy = "none"
)

// Add adds user with given Email id to database
func Add(emailId string) error {
	db := store.DB
	newUser := company.Company{
		EmailId: emailId,
	}
	err := db.Model(&newUser).Create(&newUser).Error
	if err != nil {
		return nil
	}
	return err
}

func Update(emailId string, updatedCompany company.Company) error {
	db := store.DB
	return db.Model(&company.Company{}).
		Where("email_id = ?", emailId).
		Updates(updatedCompany).
		Error
}

func AddIfNotExist(emailId string) error {
	_, err := Get(emailId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = Add(emailId)
	}

	return err
}

// Get returns company with given Email id from database
func Get(emailId string) (*company.Company, error) {
	db := store.DB
	var _company company.Company
	res := db.First(&_company, company.Company{
		EmailId: emailId,
	})
	if err := res.Error; err != nil {
		err = fmt.Errorf("failed to get user from database: %w", err)
		return nil, err
	}

	return &_company, nil
}

func GetCompaniesOpenToAquire(offSet int, take int, sortBy sortBy) (companies []company.Company, err error) {
	db := store.DB
	companyModel := db.Limit(take).Offset(offSet).Model(&company.Company{}).Not("name = ?", "").
		Not("price = 0").
		Not("current_year_profit = 0").
		Not("product = ?", "").
		Not("domain = ?", "").
		Where("open_to_aquire = ?", true)
	if sortBy != SORT_BY_NONE {
		orderString := fmt.Sprintf("%s desc", sortBy)
		logo.Info("max ", orderString)
		companyModel = companyModel.Order(orderString)
	}
	err = companyModel.Find(&companies).Error
	return
}
