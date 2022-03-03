package dataset

import "fmt"

type Dataset interface {
	toString() string
}

type Repayment struct {
	Repayments string `json:"repayments"`
}

func (r *Repayment) toString() string {
	return "Todo"
}

func (r *CouchbaseRepayment) toString() string {
	return "Todo"
}

type CouchbaseRepayment struct {
	Repayments string `json:"cb_repayments"`
}

type DatasetResult struct {
	Dataset Dataset
	Error   error
}

func DatasetFactory(entity string, source string) (error, Dataset) {
	switch {
	case entity == "Repayment" && source == "Roster":
		var r Repayment
		return nil, &r
	case entity == "Repayment" && source == "Couchbase":
		var r CouchbaseRepayment
		return nil, &r
	}

	return fmt.Errorf("Unsupported entity/source combination"), nil
}
