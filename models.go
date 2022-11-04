package minedash

import "encoding/json"

type EntityAllResponse struct {
	Entities []Entity `json:"entity"`
}

type Entity struct {
	CorpId      json.Number `json:"corpId"`
	EntityLabel string      `json:"entityLabel"`
	EntityType  string      `json:"entityType"`
	ExternalId  string      `json:"externalId"`
	OnSite      bool        `json:"onSite"`
	SysId       int         `json:"sysId"`
}
