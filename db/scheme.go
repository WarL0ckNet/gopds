package db

import (
	"time"

	"gorm.io/gorm"
)

type Genre struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	Genre         string `gorm:"size:32;not null"`
	Section       string `gorm:"size:64;not null"`
	Subsection    string `gorm:"size:100;not null"`
	SectionEng    string `gorm:"size:64;not null"`
	SubsectionEng string `gorm:"size:100;not null"`
}

type Catalog struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Cat_name string `gorm:"size:190;not null"`
	Path     string `gorm:"size:512;not null"`
	Cat_type uint   `gorm:"not null"`
	Cat_size uint
	ParentID *uint
	Parent   *Catalog `gorm:"foreignKey:ParentID"`
}

type Author struct {
	ID               uint   `gorm:"primaryKey;autoIncrement"`
	Full_name        string `gorm:"size:128;not null"`
	Search_full_name string `gorm:"size:128;not null"`
	Lang_code        uint
}

type Series struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	Ser        string `gorm:"size:150;not null"`
	Search_ser string `gorm:"size:150;not null"`
	Lang_code  uint
}

type Book struct {
	ID           uint       `gorm:"primaryKey;autoIncrement"`
	Filename     string     `gorm:"size:512;not null"`
	Path         string     `gorm:"size:512;not null"`
	Filesize     uint       `gorm:"not null"`
	Format       string     `gorm:"size:8;not null"`
	Cat_type     uint       `gorm:"not null"`
	Registredate *time.Time `gorm:"not null"`
	Docdate      string     `gorm:"size:32;not null"`
	Lang         string     `gorm:"size:16;not null"`
	Title        string     `gorm:"size:512;not null"`
	Search_title string     `gorm:"size:512;not null"`
	Annotation   string     `gorm:"size:10000;not null"`
	Lang_code    uint       `gorm:"not null"`
	Avail        uint       `gorm:"not null"`
	CatalogID    *uint      `gorm:"not null"`
	Catalog      *Catalog   `gorm:"foreignKey:CatalogID"`
}

type Bauthor struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	AuthorID *uint
	Author   *Author `gorm:"foreignKey:AuthorID"`
	BookID   *uint
	Book     *Book `gorm:"foreignKey:BookID"`
}

type Bgenre struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	BookID  *uint
	Book    *Book `gorm:"foreignKey:BookID"`
	GenreID *uint
	Genre   *Genre `gorm:"foreignKey:GenreID"`
}

type Bseries struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	Ser_no uint `gorm:"not null"`
	BookID *uint
	Book   *Book `gorm:"foreignKey:BookID"`
	SerID  *uint
	Series *Series `gorm:"foreignKey:SerID"`
}

func MigrateAll(db *gorm.DB) {
	db.AutoMigrate(&Genre{})
	db.AutoMigrate(&Catalog{})
	db.AutoMigrate(&Series{})
	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Book{})
	db.AutoMigrate(&Bauthor{})
	db.AutoMigrate(&Bgenre{})
	db.AutoMigrate(&Bseries{})
}
