package models

type Todos struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	Tanggal  string     `json:"tanggal"`
	Kegiatan []Kegiatan `json:"kegiatan" gorm:"foreignKey:todos_id"`
}

type Kegiatan struct {
	NamaKegiatan string `json:"nama_kegiatan"`
	JamKegiatan  string `json:"jam_kegiatan"`
}
