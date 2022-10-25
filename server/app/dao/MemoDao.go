package dao

import (
	"server/app/domain"
)

func MemoInsert(inputMemo domain.Memo) error {
	_, err := engine.Insert(&inputMemo)
	return err
}

func MemoDelete(inputMemo domain.Memo) error {
	_, err := engine.Delete(inputMemo)
	return err
}

func MemoFindByOwner(inputMemo domain.Memo) ([]domain.Memo, error) {
	memos := make([]domain.Memo, 0)
	err := engine.Where("owner = '?", inputMemo.Owner).Find(&memos)
	return memos, err
}
