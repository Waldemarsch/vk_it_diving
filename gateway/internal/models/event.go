package models

import (
	"encoding/json"
	"fmt"
)

type Event struct {
	Type    string                 `json:"type"`
	EventId int32                  `json:"event_id"`
	V       string                 `json:"v"`
	Object  map[string]interface{} `json:"object"`
	GroupId int32                  `json:"group_id"`
}

func (m *Event) MarshalToJSON() []byte {
	EventJSON, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return EventJSON
}
