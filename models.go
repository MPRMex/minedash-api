package minedash

type EntityAllResponse struct {
	Entities []Entity `json:"entity"`
}

type Entity struct {
	CorpId      Str    `json:"corpId"`
	EntityLabel string `json:"entityLabel"`
	EntityType  string `json:"entityType"`
	ExternalId  string `json:"externalId"`
	OnSite      bool   `json:"onSite"`
	SysId       int    `json:"sysId"`
}

type Str string

func (w *Str) UnmarshalJSON(data []byte) (err error) {
	v := Str(string(data))
	w = &v
	return nil
}
