package models

import "time"

type Auditory struct {
	Id                   string    `json:"id,omitempty" bson:"_id,omitempty"`
	EntityId             string    `json:"entityId" bson:"entity_id"`
	Type                 string    `json:"type" bson:"type"`
	Entity               string    `json:"entity" bson:"entity"`
	DatedIn              time.Time `json:"datedIn" bson:"dated_in"`
	ClientId             string    `json:"clientId" bson:"client_id"`
	UserId               string    `json:"userId" bson:"user_id"`
	InternetProtocolAddr string    `json:"internetProtocolAddr" bson:"internet_protocol_addr"`
	Data                 any       `json:"data" bson:"data"`
}
