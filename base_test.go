package gonavitia

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func marshallAndUnmarshallCoord(t *testing.T, c *Coord) (r *Coord) {
	data, err := json.Marshal(c)
	assert.Nil(t, err)
	err = json.Unmarshal(data, &r)
	assert.Nil(t, err)
	return r
}

func marshallAndUnmarshallDate(t *testing.T, d *NavitiaDatetime) (r *NavitiaDatetime) {
	data, err := json.Marshal(d)
	assert.Nil(t, err)
	err = json.Unmarshal(data, &r)
	assert.Nil(t, err)
	return r
}

func TestMarshallingCoord(t *testing.T) {
	c := &Coord{}
	r := marshallAndUnmarshallCoord(t, c)
	assert.Equal(t, c, r)

	r = marshallAndUnmarshallCoord(t, nil)
	assert.Nil(t, r)

	c = &Coord{
		Lat: 2.98,
		Lon: 6.76,
	}
	r = marshallAndUnmarshallCoord(t, c)
	assert.Equal(t, c, r)

	c = &Coord{
		Lat: -2.98,
		Lon: -6.76,
	}
	r = marshallAndUnmarshallCoord(t, c)
	assert.Equal(t, c, r)
}

func TestMarshallingDate(t *testing.T) {
	d := &NavitiaDatetime{}
	r := marshallAndUnmarshallDate(t, d)
	assert.True(t, time.Time(*d).Equal(time.Time(*r)))

	r = marshallAndUnmarshallDate(t, nil)
	assert.Nil(t, r)

	dt := NavitiaDatetime(time.Date(2018, time.June, 11, 13, 19, 22, 0, time.UTC))
	r = marshallAndUnmarshallDate(t, &dt)
	assert.Equal(t, time.Time(dt).String(), time.Time(*r).String())
	assert.True(t, time.Time(dt).Equal(time.Time(*r)))

}
