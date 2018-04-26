package gonavitia

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// nothing to do with https://golang.org/pkg/context/
type Context struct {
	Timezone        string          `json:"timezone,omitempty"`
	CurrentDatetime NavitiaDatetime `json:"current_datetime,omitempty"`
}

type Place struct {
	Id           *string    `json:"id"`
	Name         *string    `json:"name"`
	EmbeddedType *string    `json:"embedded_type"`
	Quality      *int32     `json:"quality,omitempty"`
	StopPoint    *StopPoint `json:"stop_point,omitempty"`
	StopArea     *StopArea  `json:"stop_area,omitempty"`
	Admin        *Admin     `json:"administrative_region,omitempty"`
	Address      *Address   `json:"address,omitempty"`
}

type Pagination struct {
	StartPage    int32 `json:"start_page"`
	ItemsOnPage  int32 `json:"items_on_page"`
	ItemsPerPage int32 `json:"items_per_page"`
	TotalResult  int32 `json:"total_result"`
}

type NavitiaDatetime time.Time

func (t NavitiaDatetime) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return []byte("\"\""), nil
	}
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("20060102T150405"))
	return []byte(stamp), nil
}

type Coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func (c Coord) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	}{
		Lat: strconv.FormatFloat(c.Lat, 'f', -1, 64),
		Lon: strconv.FormatFloat(c.Lon, 'f', -1, 64),
	})
}

type Code struct {
	Type  *string `json:"type"`
	Value *string `json:"value"`
}

type StopArea struct {
	Id              *string          `json:"id"`
	Name            *string          `json:"name"`
	Label           *string          `json:"label"`
	Timezone        *string          `json:"timezone,omitempty"`
	Coord           *Coord           `json:"coord"`
	Admins          []*Admin         `json:"administrative_regions"`
	Codes           []*Code          `json:"codes,omitempty"`
	StopPoints      []*StopPoint     `json:"stop_points,omitempty"`
	Links           []*Link          `json:"links"`
	CommercialModes []CommercialMode `json:"commercial_modes,omitempty"`
	PhysicalModes   []PhysicalMode   `json:"physical_modes,omitempty"`
}

type StopPoint struct {
	Id              *string          `json:"id"`
	Name            *string          `json:"name"`
	Label           *string          `json:"label"`
	Coord           *Coord           `json:"coord"`
	Admins          []*Admin         `json:"administrative_regions"`
	Codes           []*Code          `json:"codes,omitempty"`
	StopArea        *StopArea        `json:"stop_area,omitempty"`
	Equipments      []string         `json:"equipments"`
	Links           []*Link          `json:"links"`
	Address         *Address         `json:"address,omitempty"`
	CommercialModes []CommercialMode `json:"commercial_modes,omitempty"`
	PhysicalModes   []PhysicalMode   `json:"physical_modes,omitempty"`
}

type Admin struct {
	Id      *string `json:"id"`
	Name    *string `json:"name"`
	Label   *string `json:"label"`
	Coord   *Coord  `json:"coord"`
	Insee   *string `json:"insee,omitempty"`
	ZipCode *string `json:"zip_code,omitempty"`
	Level   int32   `json:"level"`
}

type Address struct {
	Id          *string  `json:"id"`
	Name        *string  `json:"name"`
	Label       *string  `json:"label"`
	Coord       *Coord   `json:"coord"`
	HouseNumber *int32   `json:"house_number,omitempty"`
	Admins      []*Admin `json:"administrative_regions"`
}

type FeedPublisher struct {
	Id      *string `json:"id"`
	Name    *string `json:"name,omitempty"`
	Url     *string `json:"url,omitempty"`
	License *string `json:"license,omitempty"`
}
