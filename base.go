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

// represent a datetime object from navitia (with its custom formating)
// When unmarshalled from a navitia response the timezone will be lost.
// In most case navitia response are in the local timezone of the coverage,
// the "real" timezone can be obtained from the context object at the root of response
type NavitiaDatetime time.Time

func (t NavitiaDatetime) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return []byte("\"\""), nil
	}
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("20060102T150405"))
	return []byte(stamp), nil
}

func (t *NavitiaDatetime) UnmarshalJSON(data []byte) error {
	//we get the "json" value, with quote, so we remove them
	data = data[1 : len(data)-1]
	if len(data) == 0 {
		return nil
	}
	value, err := time.Parse("20060102T150405", string(data))
	if err != nil {
		return err
	}
	*t = NavitiaDatetime(value)
	return nil
}

type Coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

//represent a coord with lat and lon as string...
type coordString struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

func (c Coord) MarshalJSON() ([]byte, error) {
	return json.Marshal(&coordString{
		Lat: strconv.FormatFloat(c.Lat, 'f', -1, 64),
		Lon: strconv.FormatFloat(c.Lon, 'f', -1, 64),
	})
}

func (c *Coord) UnmarshalJSON(data []byte) error {
	coord := &coordString{}
	err := json.Unmarshal(data, coord)
	if err != nil {
		return err
	}
	if coord == nil {
		return nil
	}
	c.Lat, err = strconv.ParseFloat(coord.Lat, 64)
	if err != nil {
		return err
	}
	c.Lon, err = strconv.ParseFloat(coord.Lon, 64)
	return err
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
