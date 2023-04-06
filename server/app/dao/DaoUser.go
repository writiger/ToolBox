package dao

import (
	"server/app/domain"
	"server/error_code"
	"time"
)

func UserInsert(inputUser domain.User) error {
	has, _ := UserIsExist(inputUser)
	if has {
		return errorcode.GetErr(errorcode.ErrUserAlreadyExist)
	}
	_, err := engine.Insert(&inputUser)
	return err
}

func UserEmailRedisSet(inputUser domain.User, code int) error {
	err := rc.Set(inputUser.Account, code, time.Duration(c.RegisterCodeFailureTime)*time.Second).Err()
	return err
}

func UserEmailCodeGet(inputUser domain.User) string {
	res, _ := rc.Get(inputUser.Account).Result()
	return res
}

func UserIsExist(inputUser domain.User) (bool, error) {
	has, err := engine.Exist(&inputUser)
	return has, err
}

func UserLogin(inputUser *domain.User) (domain.User, error) {
	has, err := engine.Get(inputUser)
	if err != nil {
		return domain.User{}, err
	}

	if has {
		return *inputUser, nil
	} else {
		return domain.User{}, errorcode.GetErr(errorcode.ErrUserWrongPassword)
	}
}

func UserChangePayInfo(inputUser *domain.User) error {
	_, err := engine.ID(inputUser.Id).Update(inputUser)
	return err
}

func UserGetAll() ([]domain.User, error) {
	var res []domain.User
	err := engine.Cols("base_disk", "pay_day").Find(&res)
	return res, err
}
