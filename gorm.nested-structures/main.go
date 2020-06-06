package main

import (
	"github.com/modelware-ru/examples.go/gorm.nested-structures/mysql"
	"log"
)

type Country struct {
	Id   int64  `gorm:id`
	Name string `gorm:name`
}

type Region struct {
	Id   int64  `gorm:id`
	Name string `gorm:name`
	CountryId int64 `gorm:country_id`
}

type City struct {
	Id   int64  `gorm:id`
	Name string `gorm:name`
	RegionId int64 `gorm:region_id`
	CountryId int64 `gorm:country_id`
}

type Line struct {

}

func main() {

	if err := mysql.DB.DB().Ping(); err != nil {
		panic(err)
	}

	// get all countries
	var allCountries []Country
	if err := mysql.DB.Table("country").Find(&allCountries).Error; err != nil {
		panic(err)
	}
	log.Printf("all countries = %#v", allCountries)

	// get country by id
	var countryById Country
	if err := mysql.DB.Table("country").Where("id = ?", 1).Last(&countryById).Error; err != nil {
		panic(err)
	}
	log.Printf("country by id 1 = %#v", countryById)

	// get all regions
	var allRegions []Region
	if err := mysql.DB.Table("region").Find(&allRegions).Error; err != nil {
		panic(err)
	}
	log.Printf("all regions = %#v", allRegions)

	// get all regions by country id
	var allRegionsByCountryId []Region
	if err := mysql.DB.Table("region").Where("country_id = ?", 2).Find(&allRegionsByCountryId).Error; err != nil {
		panic(err)
	}
	log.Printf("all regions by country id 2 = %#v", allRegionsByCountryId)

	// get region by id
	var regionById Country
	if err := mysql.DB.Table("region").Where("id = ?", 3).Last(&regionById).Error; err != nil {
		panic(err)
	}
	log.Printf("region by id 3 = %#v", regionById)

	// get all cities

	// get all cities by region id

	// get all cities by country id

	// get city by id

	log.Printf("the end\n")
}
