package sgp4

import "math"

type GravityModel int

const (
	// Legacy support for old SGP4 behavior
	WGS72old GravityModel = iota
	// Standard WGS 72 model
	WGS72
	// More recent WGS 84 model
	WGS84
)

type GravityConst struct {
	tumin, mu, radiusearthkm, xke, j2, j3, j4, j3oj2 float64
}

// NewGravConst is returning all the gravity constants based on gravity type
func NewGravConst(g GravityModel) *GravityConst {
	switch g {
	case WGS72old:
		xke := 0.0743669161
		j2 := 0.001082616
		j3 := -0.00000253881

		return &GravityConst{
			mu:            398600.79964,
			radiusearthkm: 6378.135,
			xke:           xke,
			tumin:         1.0 / xke,
			j2:            j2,
			j3:            j3,
			j4:            -0.00000165597,
			j3oj2:         j3 / j2,
		}

	case WGS72:
		mu := 398600.8
		radiusearthkm := 6378.135
		xke := 60.0 / math.Sqrt(radiusearthkm*radiusearthkm*radiusearthkm/mu)
		j2 := 0.001082616
		j3 := -0.00000253881
		return &GravityConst{
			mu:            mu,
			radiusearthkm: radiusearthkm,
			xke:           xke,
			tumin:         1.0 / xke,
			j2:            j2,
			j3:            j3,
			j4:            -0.00000165597,
			j3oj2:         j3 / j2,
		}

	case WGS84:
		mu := 398600.5
		radiusearthkm := 6378.137
		xke := 60.0 / math.Sqrt(radiusearthkm*radiusearthkm*radiusearthkm/mu)
		j2 := 0.00108262998905
		j3 := -0.00000253215306
		return &GravityConst{
			mu:            398600.5,
			radiusearthkm: radiusearthkm,
			xke:           60.0 / math.Sqrt(radiusearthkm*radiusearthkm*radiusearthkm/mu),
			tumin:         1.0 / xke,
			j2:            j2,
			j3:            j3,
			j4:            -0.00000161098761,
			j3oj2:         j3 / j2,
		}
	}
	return nil
}
