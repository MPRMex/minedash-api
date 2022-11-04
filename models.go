package minedash

import (
	"encoding/json"
	"strconv"
)

type EntityAllResponse struct {
	Entities []Entity `json:"entity"`
}

type Entity struct {
	CorpId      Str         `json:"corpId"`
	EntityLabel string      `json:"entityLabel"`
	EntityType  string      `json:"entityType"`
	ExternalId  string      `json:"externalId"`
	OnSite      bool        `json:"onSite"`
	SysId       json.Number `json:"sysId"`
}

type Str string

func (w *Str) UnmarshalJSON(data []byte) (err error) {
	var s string
	var i int
	result := ""
	if err := json.Unmarshal(data, &s); err == nil {
		result = s
	} else if err := json.Unmarshal(data, &i); err == nil {
		result = strconv.Itoa(i)
	}
	*w = Str(result)
	return nil
}

func (w Str) String() string {
	return string(w)
}
