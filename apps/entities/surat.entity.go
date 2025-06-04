package entities

type Surat struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	No      int    `json:"no"`
	Name    string `json:"name"`
	NameLtn string `json:"name_ltn"`
	NameId  string `json:"name_id"`
	Type    string `json:"type"`
	Total   int    `json:"total"`
}

type SuratDetail struct {
	Surat

	Description string `json:"description"`
	Audio       string `json:"audio"`
}

func (Surat) CollectionName() string {
	return COLLECTION_SURAT
}
