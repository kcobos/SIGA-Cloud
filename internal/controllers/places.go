package controller

import (
	"github.com/go-pg/pg/v10"
	"github.com/kcobos/SIGA-Cloud/common"
	"github.com/kcobos/SIGA-Cloud/internal/errors"
	model "github.com/kcobos/SIGA-Cloud/internal/models"
)

// Places is the controller for Place model
// It uses a database
type Places struct {
	db *pg.DB
}

const minimun_coordinate_change float64 = 0.000001

// NewPlaces initialize Places struct
// It sets the map and the last ID
func NewPlaces(db *pg.DB, conf *common.Conf) *Places {
	p := new(Places)
	p.db = db
	common.CreateTable(db, conf, (*model.Place)(nil))
	return p
}

// Len returns length of place list if it is initialized
func (p *Places) Len() (int, error) {
	count, err := p.db.Model(&model.Place{}).Count()
	return count, err
}

// NewPlace appends a new place to the list of places
// Return the new id
func (p *Places) NewPlace(latitude, longitude float64, address string, parkingIDs []int64, parkings *Parkings) (int64, error) {
	count, err := p.db.Model(&model.Place{}).
		Where("address = ?", address).
		Count()
	if err != nil {
		return -1, err
	}
	if count != 0 {
		return -1, &errors.PlaceAlreadyExists{}
	}
	count, err = p.db.Model(&model.Place{}).
		Where("longitude BETWEEN ? and ?", longitude-minimun_coordinate_change, longitude+minimun_coordinate_change).
		Where("latitude BETWEEN ? and ?", latitude-minimun_coordinate_change, latitude+minimun_coordinate_change).
		Count()

	if err != nil {
		return -1, err
	}
	if count != 0 {
		return -1, &errors.PlaceAlreadyExists{}
	}

	var parkingsForPlace []*model.Parking
	for i := 0; i < len(parkingIDs); i++ {
		parking, _ := parkings.GetParking(parkingIDs[i])
		parkingsForPlace = append(parkingsForPlace, &parking)
	}
	place, err := model.NewPlace(latitude, longitude, address, parkingsForPlace)
	if err != nil {
		return -1, err
	}
	_, err = p.db.Model(place).Insert()
	if err != nil {
		return -1, err
	}
	for i := 0; i < len(parkingIDs); i++ {
		err = parkings.AttachToPlace(parkingIDs[i], place.ID)
		if err != nil {
			return -1, err
		}
	}
	return place.ID, err
}
