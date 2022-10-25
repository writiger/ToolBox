package service

import (
	"server/app/dao"
	"server/app/domain"
	"server/error_code"
	"strconv"
)

func MemoQueryByOwnerService(owner string) ([]domain.Memo, error) {
	ownerId, err := strconv.ParseInt(owner, 10, 64)
	if err != nil {
		return nil, errorcode.GetErr(errorcode.ErrParam)
	}
	memos, err2 := dao.MemoFindByOwner(domain.Memo{
		Owner: ownerId,
	})
	return memos, err2
}

func MemoAdd(memoInput domain.Memo) error {
	err := dao.MemoInsert(memoInput)
	return err
}

func MemoDelete(memoInput domain.Memo) error {
	err := dao.MemoDelete(memoInput)
	return err
}
