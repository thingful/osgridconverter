// osgridConverter package contains contains utility functions to convert
// Ordnance Survey grid references to latitude/longitude coordinates.
// Original javascript implementaion by Chris Veness.
// For more information see http://www.movable-type.co.uk/scripts/latlong-gridref.html
package osgridConverter

import (
	"math"
)

const (
	a  = 6377563.396        // Airy 1830 major & minor semi-axes
	b  = 6356256.909        // Airy 1830 major & minor semi-axes
	F0 = 0.9996012717       // NatGrid scale factor on central meridian
	φ0 = 49 * math.Pi / 180 // NatGrid true origin
	λ0 = -2 * math.Pi / 180 // NatGrid true origin
	N0 = -100000            // northing of true origin, metres
	E0 = 400000             // easting of true origin, metres
	e2 = 1 - (b*b)/(a*a)    // eccentricity squared
	n  = (a - b) / (a + b)  // n
	n2 = n * n              // n²
	n3 = n * n * n          // n³
)

// OsGridToLatLon converts Ordnance Survey grid reference easting/northing coordinate to latitude/longitude
func OsGridToLatLon(easting, northing float64) (lat, lon float64) {
	φ := φ0
	M := float64(0)

	for northing-N0-M >= 0.00001 {
		φ = (northing-N0-M)/(a*F0) + φ
		Ma := (1 + n + (5/4)*n2 + (5/4)*n3) * (φ - φ0)
		Mb := (3*n + 3*n*n + (21/8)*n3) * math.Sin(φ-φ0) * math.Cos(φ+φ0)
		Mc := ((15/8)*n2 + (15/8)*n3) * math.Sin(2*(φ-φ0)) * math.Cos(2*(φ+φ0))
		Md := (35 / 24) * n3 * math.Sin(3*(φ-φ0)) * math.Cos(3*(φ+φ0))
		M = b * F0 * (Ma - Mb + Mc - Md) // meridional arc
	}

	cosφ := math.Cos(φ)
	sinφ := math.Sin(φ)
	ν := a * F0 / math.Sqrt(1-e2*sinφ*sinφ)                // nu = transverse radius of curvature
	ρ := a * F0 * (1 - e2) / math.Pow(1-e2*sinφ*sinφ, 1.5) // rho = meridional radius of curvature
	η2 := ν/ρ - 1

	tanφ := math.Tan(φ)
	tan2φ := tanφ * tanφ
	tan4φ := tan2φ * tan2φ
	tan6φ := tan4φ * tan2φ
	secφ := 1 / cosφ
	ν3 := ν * ν * ν
	ν5 := ν3 * ν * ν
	ν7 := ν5 * ν * ν
	VII := tanφ / (2 * ρ * ν)
	VIII := tanφ / (24 * ρ * ν3) * (5 + 3*tan2φ + η2 - 9*tan2φ*η2)
	IX := tanφ / (720 * ρ * ν5) * (61 + 90*tan2φ + 45*tan4φ)
	X := secφ / ν
	XI := secφ / (6 * ν3) * (ν/ρ + 2*tan2φ)
	XII := secφ / (120 * ν5) * (5 + 28*tan2φ + 24*tan4φ)
	XIIA := secφ / (5040 * ν7) * (61 + 662*tan2φ + 1320*tan4φ + 720*tan6φ)

	dE := (easting - E0)
	dE2 := dE * dE
	dE3 := dE2 * dE
	dE4 := dE2 * dE2
	dE5 := dE3 * dE2
	dE6 := dE4 * dE2
	dE7 := dE5 * dE2
	φ = φ - VII*dE2 + VIII*dE4 - IX*dE6
	λ := λ0 + X*dE - XI*dE3 + XII*dE5 - XIIA*dE7

	return toDegrees(φ), toDegrees(λ)
}

// toDegrees converts radians to numeric degrees
func toDegrees(input float64) float64 {
	return input * 180 / math.Pi
}
