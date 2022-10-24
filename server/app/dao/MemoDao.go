package dao

import (
	"server/app/domain"
)

func MemoInsert(memo domain.Memo) error {
	_, err := engine.Insert(&memo)
	return err
}

func MemoFindByOwner(memo domain.Memo) ([]domain.Memo, error) {
	memos := make([]domain.Memo, 0)
	err := engine.Find(&memos)
	return memos, err
}
