package dao

import (
	"server/app/domain"
)

func MemoInsert(memo domain.Memo) error {
	_, err := engine.Insert(&memo)
	if err != nil {
		return err
	}
	return nil
}

func MemoFindByOwner(memo domain.Memo) ([]domain.Memo, error) {

}
