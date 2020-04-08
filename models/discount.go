package models

import (
	"context"
	"github.com/hyecheonlee/echosample/factory"
	"time"
)

type Discount struct {
	Id             int64     `json:"id"`
	Name           string    `json:"name"`
	Desc           string    `json:"desc"`
	StartAt        time.Time `json:"startAt"`
	EndAt          time.Time `json:"endAt"`
	ActionType     string    `json:"actionType"`
	DiscountAmount float64   `json:"discountAmount"`
	Enable         bool      `json:"enable"`
	CreatedAt      time.Time `json:"createdAt" xorm:"created"`
	UpdatedAt      time.Time `json:"updatedAt" xorm:"updated"`
}

func (d *Discount) Create(ctx context.Context) (int64, error) {
	affected, err := factory.DB(ctx).Insert(d)
	if err != nil {
		return 0, factory.ErrorDB.New(err)
	}
	return affected, nil
}
