package _interface

import "github.com/iacopoGhilardi/amILate/internal/model"

type UserLocationRepositoryInterface interface {
	BaseRepositoryInterface[model.UserLocation]

	Upsert(location *model.UserLocation) error
}
