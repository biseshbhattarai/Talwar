package models

type SubDomain struct {
	ID       int64 `gorm:"primaryKey"`
	Name     string
	DomainID int64 `gorm:"foreignkey:ID"`
	Domain   Domain
}
