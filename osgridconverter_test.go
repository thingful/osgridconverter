package osgridconverter

import (
	"testing"
)

func TestConvertToLatLon(t *testing.T) {
	testcases := []struct {
		easting     float64
		northing    float64
		expectedLat float64
		expectedLon float64
	}{
		{
			easting:     651409,
			northing:    313177,
			expectedLat: 52.65796257381598,
			expectedLon: 1.7160372280352103,
		},
		{
			easting:     438700,
			northing:    114800,
			expectedLat: 50.93135051478858,
			expectedLon: -1.4506775180363458,
		},
	}

	for _, testcase := range testcases {

		c, err := ConvertToLatLon(testcase.easting, testcase.northing)

		if err != nil {
			t.Errorf("Error generating lat/lon coordinates: %v", err)
		}

		if c.Lat != testcase.expectedLat {
			t.Errorf("Unexpected error generating latitude coordinate. Expected %g, got %g", testcase.expectedLat, c.Lat)
		}

		if c.Lon != testcase.expectedLon {
			t.Errorf("Unexpected error generating longitude coordinate. Expected %g, got %g", testcase.expectedLon, c.Lon)
		}
	}
}

func TestConvertToLatLonError(t *testing.T) {
	testcases := []struct {
		easting  float64
		northing float64
	}{
		{
			easting:  -0.1,
			northing: 1,
		},
		{
			easting:  1,
			northing: -0.1,
		},
		{
			easting:  -1,
			northing: -2,
		},
	}

	for _, testcase := range testcases {

		_, err := ConvertToLatLon(testcase.easting, testcase.northing)

		if err == nil {
			t.Errorf("Expecting Error when passing out of bounds arguments")
		}
	}
}

func TestConvertToNorthingEasting(t *testing.T) {
	testcases := []struct {
		lat       float64
		lon       float64
		expectedE float64
		expectedN float64
	}{
		{
			lat:       52.696361,
			lon:       -1.625977,
			expectedE: 425274.082, // SHOULD BE 425374.372
			expectedN: 311070.002, // SHOULD BE 311030.142
		},
		{
			lat:       0.0,
			lon:       0.0,
			expectedE: 622575.703,       // SHOULD BE 622674.837
			expectedN: -5.527062157e+06, // SHOULD BE -5527598.33
		},
	}

	for _, testcase := range testcases {
		o, err := ConvertToNorthingEasting(testcase.lat, testcase.lon)

		if err != nil {
			t.Errorf("Error generating easting/northing coordinates: %v", err)
		}

		if o.Easting != testcase.expectedE {
			t.Errorf("Unexpected error generating easting coordinate. Expected %g, got %g", testcase.expectedE, o.Easting)
		}

		if o.Northing != testcase.expectedN {
			t.Errorf("Unexpected error generating northing coordinate. Expected %g, got %g", testcase.expectedN, o.Northing)
		}
	}
}

func TestConvertToNorthingEastingError(t *testing.T) {
	testcases := []struct {
		lat float64
		lon float64
	}{
		{
			lat: -90.1,
			lon: 1,
		},
		{
			lat: 90.1,
			lon: 1,
		},
		{
			lat: 1,
			lon: 180.1,
		},
		{
			lat: 1,
			lon: -180.1,
		},
	}

	for _, testcase := range testcases {

		_, err := ConvertToNorthingEasting(testcase.lat, testcase.lon)

		if err == nil {
			t.Errorf("Expecting Error when passing out of bounds arguments")
		}
	}
}
