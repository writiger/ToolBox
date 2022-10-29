package service

import (
	"server/app/dao"
	"server/app/domain"
)

func MemoQueryByOwnerService(owner int64) ([]domain.Memo, error) {
	memos, err2 := dao.MemoFindByOwner(domain.Memo{
		Owner: owner,
	})
	return memos, err2
}

func MemoAdd(inputMemo domain.Memo) error {
	err := dao.MemoInsert(inputMemo)
	return err
}

func MemoDelete(inputMemo domain.Memo) error {
	err := dao.MemoDelete(inputMemo)
	return err
}

func MemoPut(inputMemo domain.Memo) error {
	err := dao.MemoUpdate(inputMemo)
	return err
}
