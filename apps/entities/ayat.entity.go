package entities

type Ayat struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	IdSurat int    `json:"id_surat"`
	No      int    `json:"no"`
	Name    string `json:"name"`
	NameLtn string `json:"name_ltn"`
	NameId  string `json:"name_id"`
	Type    string `json:"type"`
	Text    string `json:"text"`
	TextLtn string `json:"text_ltn"`
	TextId  string `json:"text_id"`
}

func (Ayat) CollectionName() string {
	return COLLECTION_AYAT
}
