package refresh

import (
	"fmt"
	"msgRefresh3/dataset"
)

// EntityRefresher Treat this as a class
type EntityRefresher struct {
	Properties *RefreshEntityDto
}

type RefreshEntityDto struct {
	EntityType       string `json:"entity-type"`
	Channel          string `json:"channel"`
	Headers          []byte `json:"headers"`
	Endpoint         string `json:"endpoint"`
	ComparisonFields []byte `json:"comparison-fields"`
}

func (r *EntityRefresher) Refresh() error {
	rosterChan := make(chan dataset.DatasetResult)
	cbChan := make(chan dataset.DatasetResult)

	go r.fetchRosterData(rosterChan)
	go r.fetchCBData(cbChan)

	rosterResult := <-rosterChan
	cbResult := <-cbChan

	if rosterResult.Error != nil {
		return rosterResult.Error
	}

	if cbResult.Error != nil {
		return cbResult.Error
	}

	rosterData := rosterResult.Dataset
	cbData := cbResult.Dataset

	fmt.Printf("%+v", rosterData)
	fmt.Printf("%+v", cbData)

	//TODO: comparison := r.compare(repayment, couchbaseRepayment)
	//TODO: update := r.update(comparison)

	return nil
}
