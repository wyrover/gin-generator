package models

import (
	"time"

	"XXXXX/config"
)

type @@@@@ struct {
	ID             uint   
	
}

func FindAll@@@@@() ([]@@@@@, error) {
	var X@@@@s []@@@@@
	err := config.GetDB().Find(&X@@@@s).Error
	return X@@@@s, err
}

func Find@@@@@ById(id string) (@@@@@, error) {
	var X@@@@ @@@@@
	err := config.GetDB().Where("id = ?", id).First(&X@@@@).Error
	return X@@@@, err
}

func Create@@@@@(X@@@@ *@@@@@) (*@@@@@, error) {
	err := config.GetDB().Create(&X@@@@).Error
	return X@@@@, err

}

func Update@@@@@(X@@@@ *@@@@@) (*@@@@@, error) {
	err := config.GetDB().Save(&X@@@@).Error
	return X@@@@, err

}

func Delete@@@@@(id string) error {
	return config.GetDB().Where("id = ?", id).Delete(&@@@@@{}).Error
}