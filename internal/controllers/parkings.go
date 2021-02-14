package controller

import (
	"github.com/go-pg/pg/v10"
	"github.com/kcobos/SIGA-Cloud/common"
	model "github.com/kcobos/SIGA-Cloud/internal/models"
)

// Parkings is the controller for Parking model.
// It uses a database
type Parkings struct {
	db *pg.DB
}

// NewParkings initialize Parkings struct
func NewParkings(db *pg.DB, conf *common.Conf) *Parkings {
	p := new(Parkings)
	p.db = db
	common.CreateTable(db, conf, (*model.Parking)(nil))
	common.CreateTable(db, conf, (*model.ParkingStatusHistory)(nil))
	return p
}

// Len returns length of parking list
func (p *Parkings) Len() (int, error) {
	count, err := p.db.Model(&model.Parking{}).Count()
	return count, err
}

// NewParking appends a new parking to the list of parking lots
// Return the new id
func (p *Parkings) NewParking() (int64, error) {
	parking := model.NewParking()
	_, err := p.db.Model(parking).Insert()
	return parking.ID, err
}

// GetParking retrieve a Parking from DB using the ID
func (p *Parkings) GetParking(id int64) (model.Parking, error) {
	parking := model.Parking{
		ID: id,
	}
	err := p.db.Model(&parking).WherePK().Select()
	return parking, err
}

// CountHistory counts how many times the parking status has changed
func (p *Parkings) CountHistory(id int64) (int, error) {
	return p.db.Model(&model.ParkingStatusHistory{
		ParkingID: id,
	}).Count()
}

// ChangeStatus sets new status to Parking if it's valid
func (p *Parkings) ChangeStatus(id int64, newStatus string) error {
	parking, err := p.GetParking(id)
	if err != nil {
		return err
	}
	err = parking.ChangeStatus(newStatus)
	if err != nil {
		return err
	}
	_, err = p.db.Model(&parking).WherePK().Update()
	if err != nil {
		return err
	}
	for i := 0; i < len(parking.History); i++ {
		_, err = p.db.Model(&parking.History[i]).Insert()
		if err != nil {
			return err
		}
	}
	return nil
}

// AttachToPlace attach a parking to a place
func (p *Parkings) AttachToPlace(parkingID, placeID int64) error {
	parking, err := p.GetParking(parkingID)
	if err != nil {
		return err
	}
	parking.PlaceID = placeID
	_, err = p.db.Model(&parking).WherePK().Update()
	return err
}
