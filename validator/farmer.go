package validator

import "database/sql"

type CreateFarmerRequest struct {
	Entity          string          `json:"entity" binding:"required"`
	Brand           string          `json:"brand" binding:"required"`
	PondArea        float64         `json:"pond_area" binding:"required"`
	InsuredArea     sql.NullFloat64 `json:"insured_area"`
	PondNum         uint            `json:"pond_num" binding:"required"`
	Exp             float64         `json:"exp" binding:"required"`
	ContactMan      string          `json:"contact_man" binding:"required"`
	Phone           string          `json:"phone" binding:"required,len=11"`
	EntityType      uint            `json:"entity_type" binding:"required,oneof=0 1 2 3"`
	AddressProvince string          `json:"address_province" binding:"required"`
	AddressCity     string          `json:"address_city" binding:"required"`
	AddressDistrict string          `json:"address_district" binding:"required"`
	AddressDetail   string          `json:"address_detail"`
	Pic             string          `json:"pic"`
}

type UpdateFarmerRequest struct {
	Entity          string          `json:"entity"`
	Brand           string          `json:"brand"`
	PondArea        float64         `json:"pond_area"`
	InsuredArea     sql.NullFloat64 `json:"insured_area"`
	PondNum         uint            `json:"pond_num"`
	Exp             float64         `json:"exp"`
	ContactMan      string          `json:"contact_man"`
	Phone           string          `json:"phone"`
	EntityType      uint            `json:"entity_type" binding:"oneof=0 1 2 3"`
	AddressProvince string          `json:"address_province"`
	AddressCity     string          `json:"address_city"`
	AddressDistrict string          `json:"address_district"`
	AddressDetail   string          `json:"address_detail"`
	Pic             string          `json:"pic"`
}
