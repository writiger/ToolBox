package dao

import "server/app/domain"

func ClockInsert(input domain.Clock) error {
	_, err := engine.Insert(input)
	return err
}
