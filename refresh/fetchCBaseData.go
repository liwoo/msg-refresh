package refresh

import (
	"encoding/json"
	"fmt"
	"log"
	"msgRefresh3/dataset"
	"time"
)

func (r *EntityRefresher) fetchCBData(result chan dataset.DatasetResult) {
	time.Sleep(2 * time.Second)
	//TODO: fetch data from Couchbase
	fmt.Println("Fetching CB Data")
	data := []byte("{\"cb_repayments\": \"1234\"}")
	err, entity := dataset.DatasetFactory(r.Properties.EntityType, "Couchbase")
	if err != nil {
		result <- dataset.DatasetResult{Error: err}
		log.Fatalf(err.Error())
	}
	jsonErr := json.Unmarshal(data, &entity)
	if jsonErr != nil {
		result <- dataset.DatasetResult{Error: err}
		log.Fatalf(err.Error())
	}
	result <- dataset.DatasetResult{Dataset: entity}
}
