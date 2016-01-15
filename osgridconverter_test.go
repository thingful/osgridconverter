package osgridconverter

import (
	"testing"
)

func TestConvertToLatLon(t *testing.T) {
	testcases := []struct {
		northing    float64
		easting     float64
		expectedLat float64
		expectedLon float64
	}{
		{
			easting:     651409,
			northing:    313177,
			expectedLat: 52.65755427802226,  // SHOULD BE 52.657977
			expectedLon: 1.7179068604835606, // SHOULD BE 1.716038
		},
		{
			easting:     438700,
			northing:    114800,
			expectedLat: 50.930793682009025, // SHOULD BE 50.931358
			expectedLon: -1.449243402413577, // SHOULD BE -1.450677
		},
	}

	for _, testcase := range testcases {
		lat, lon := ConvertToLatLon(testcase.easting, testcase.northing)

		if lat != testcase.expectedLat {
			t.Errorf("Unexpected error generating latitude coordinate. Expected %g, got %g", testcase.expectedLat, lat)
		}

		if lon != testcase.expectedLon {
			t.Errorf("Unexpected error generating longitude coordinate. Expected %g, got %g", testcase.expectedLon, lon)
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
		easting, northing := ConvertToNorthingEasting(testcase.lat, testcase.lon)

		if easting != testcase.expectedE {
			t.Errorf("Unexpected error generating easting coordinate. Expected %g, got %g", testcase.expectedE, easting)
		}

		if northing != testcase.expectedN {
			t.Errorf("Unexpected error generating northing coordinate. Expected %g, got %g", testcase.expectedN, northing)
		}
	}
}
