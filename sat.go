package sgp4

import (
	"errors"
	"math"
	"strconv"
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

var SyntaxError = errors.New("TLE file syntax error")

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
	// Right ascension of ascending node in radians. (nodeo)
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

	a, alta, altp float64
}

// NewSatelliteFromTLE return a satellite from a TLE string (2 lines with \n)
func NewSatelliteFromTLE(tle string, g GravityModel) (*Satellite, error) {

	//deg2rad := math.Pi / 180.0
	//xpdotp := 1440.0 / (2.0 * math.Pi)
	gc := newGravConst(g)

	var err error
	parseInt := func(s string) (int, error) {
		if err != nil {
			return 0, err
		}
		return strconv.Atoi(strings.TrimSpace(s))
	}

	parseFloat := func(s string) (float64, error) {
		if err != nil {
			return 0, err
		}
		return strconv.ParseFloat(strings.TrimSpace(s), 64)
	}

	satrec := &Satellite{}
	satrec.GravityModel = g

	lines := strings.Split(tle, "\n")
	if len(lines) != 3 {
		return nil, SyntaxError
	}

	// 1st line
	line := lines[0]
	line = strings.TrimSuffix(line, "\n")
	if !strings.HasPrefix(line, "1 ") || len(line) != 69 {
		return nil, SyntaxError
	}

	satrec.SatNum, err = parseInt(line[2:7])
	twoDigitYear, err := parseInt(line[18:20])
	satrec.EpochDays, err = parseFloat(line[20:32])
	satrec.FirstDerivative, err = parseFloat(line[33:43])
	satrec.SecondDerivative, err = parseFloat(string(line[44]) + "." + line[45:50])
	nexp, err := parseFloat(line[50:52])
	satrec.BSTAR, err = parseFloat(string(line[53]) + "." + line[54:59])
	ibexp, err := parseFloat(line[59:61])

	// second line
	line = lines[1]
	if !strings.HasPrefix(line, "2 ") || len(line) != 69 {
		return nil, SyntaxError
	}
	satnum, err := parseInt(line[2:7])
	if satnum != satrec.SatNum {
		return nil, SyntaxError
	}

	satrec.Inclination, err = parseFloat(line[8:16])
	satrec.RightAscension, err = parseFloat(line[17:25])
	satrec.Eccentricity, err = parseFloat("0." + strings.Replace(line[26:33], " ", "0", -1))
	satrec.ArgPerigee, err = parseFloat(line[34:42])
	satrec.MeanMotion, err = parseFloat(line[52:63])

	// this test is checking for the whole parseFloat and parseInt calls
	if err != nil {
		return nil, err
	}

	xpdotp := 1440.0 / (2.0 * math.Pi)
	deg2rad := math.Pi / 180.0

	satrec.MeanMotion = satrec.MeanMotion / xpdotp
	satrec.SecondDerivative = satrec.SecondDerivative * math.Pow(10.0, nexp)
	satrec.BSTAR = satrec.BSTAR * math.Pow(10.0, ibexp)

	// convert to sgp4 units
	satrec.a = math.Pow(satrec.MeanMotion*gc.tumin, (-2.0 / 3.0))
	satrec.FirstDerivative = satrec.FirstDerivative / (xpdotp * 1440.0)
	satrec.SecondDerivative = satrec.SecondDerivative / (xpdotp * 1440.0 * 1440)

	// find standard orbital elements ----
	satrec.Inclination = satrec.Inclination * deg2rad
	satrec.RightAscension = satrec.RightAscension * deg2rad
	satrec.ArgPerigee = satrec.ArgPerigee * deg2rad
	satrec.MeanAnomaly = satrec.MeanAnomaly * deg2rad

	satrec.alta = satrec.a*(1.0+satrec.Eccentricity) - 1.0
	satrec.altp = satrec.a*(1.0-satrec.Eccentricity) - 1.0

	var year int
	if twoDigitYear < 57 {
		year = twoDigitYear + 2000
	} else {
		year = twoDigitYear + 1900
	}

	// to remove
	year = year + 1
	return satrec, nil
}

// Propagate return a position and velocity vector for a given date and time.
func (s *Satellite) Propagate(t *time.Time) PositionVelocity {
	return PositionVelocity{}
}
