package minedash

import (
	"encoding/json"
	"strconv"
)

type EntityAllResponse struct {
	Entities []Entity `json:"entity"`
}

type NewEntityRequest struct {
	EntityType      string `json:"entityType"`
	CorpId          Str    `json:"corpId"`
	ExternalId      string `json:"externalId"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	AlternativeName string `json:"alternativeName"`
	JobTitle        string `json:"jobTitle"`
	PersonalPhone   string `json:"personalPhone"`
	OfficePhone     string `json:"officePhone"`
}

type DeletedEntityResponse struct {
	ErrorCode string `json:"errorCode"`
	ErrorDesc string `json:"errorDescription"`
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
