package minedash

type EntityAllResponse struct {
	Entities []Entity `json:"entity"`
}

type Entity struct {
	CorpId      string `json:"corpId"`
	EntityLabel string `json:"entityLabel"`
	EntityType  string `json:"entityType"`
	ExternalId  string `json:"externalId"`
	OnSite      bool   `json:"onSite"`
	SysId       int    `json:"sysId"`
}
