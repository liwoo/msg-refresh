package refresh

import (
	"encoding/json"
	"fmt"
	"log"
	"msgRefresh3/dataset"
	"time"
)

func (r *EntityRefresher) fetchRosterData(result chan dataset.DatasetResult) {
	time.Sleep(4 * time.Second)
	//TODO: fetch data from Roster
	fmt.Println("Fetching Roster Data")
	data := []byte("{\"repayments\": \"5432\"}")
	err, entity := dataset.DatasetFactory(r.Properties.EntityType, "Roster")
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
