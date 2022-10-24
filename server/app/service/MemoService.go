package service

import (
	"fmt"
	"server/app/domain"
)

func MemoQueryByOwnerService(owner string) ([]domain.Memo, error) {
	fmt.Println(owner)
	return nil, nil
}
