package sgp4

import (
	"errors"
	"strings"
	"time"
)

// Position measures the satellite position in km from the center of the earth
// Velocity is the rate at which those three parameters are changing, expressed in kilometers per second.
type PositionVelocity struct {
	// Position
	X, Y, Z float64
	// Velocity
	XV, YV, ZV float64
}

var SyntaxError = errors.New("emit macho dwarf: elf header corrupted")

type Satellite struct {
	// Satellite name
	Name string
	// Unique satellite number given in the TLE file.
	SatNum int
	// Full four-digit year of this element set's epoch moment.
	EpochYear int
	// Fractional days into the year of the epoch moment.
	EpochDays float64
	// Julian date of the epoch (computed from EpochYear and EpochDays).
	JDSatEpoch float64
	// First time derivative of the mean motion (ignored by SGP4). (ndot)
	FirstDerivative float64
	// Second time derivative of the mean motion (ignored by SGP4). (nddot)
	SecondDerivative float64
	// Ballistic drag coefficient B* in inverse earth radii.
	BSTAR float64
	// Inclination in radians. (inclo)
	Inclination float64
	// Right ascension of ascending node in radians.
	RightAscension float64
	// Eccentricity (ecco)
	Eccentricity float64
	// Argument of perigee in radians. (argpo)
	ArgPerigee float64

	// Mean anomaly in radians. (mo)
	MeanAnomaly float64
	// MeanMotion in radians per minute. (no)
	MeanMotion float64

	Epoch time.Time

	// GravityModel is the gravity model used to compute the satellite
	GravityModel GravityModel
}

// NewSatelliteFromTLE return a satellite from a TLE string (2 lines with \n)
func NewSatelliteFromTLE(tle string, g GravityModel) (*Satellite, error) {

	//deg2rad := math.Pi / 180.0
	//xpdotp := 1440.0 / (2.0 * math.Pi)
	//gc := NewGravConst(g)
	//tumin := gc.tumin

	satrec := &Satellite{}
	satrec.GravityModel = g

	lines := strings.Split(tle, "\n")
	if len(lines) != 2 {
		return nil, SyntaxError
	}

	line := lines[0]
	line = strings.TrimSuffix(line, "\n")

	return satrec, nil
}

// Propagate return a position and velocity vector for a given date and time.
func (s *Satellite) Propagate(t *time.Time) PositionVelocity {
	return PositionVelocity{}
}
