package sgp4

import "testing"

func TestRLE(t *testing.T) {
	s, err := NewSatelliteFromTLE("1 28654U 05018A   15053.51663152  .00000205  00000-0  13711-3 0  9995\n2 28654  99.1770  43.4120 0013283 282.3899 185.0991 14.12152816502890\n", WGS72)
	if err != nil {
		t.Errorf(err.Error())
	}

	s.Name = "NOAA 18"
}

func TestInvalidRLE(t *testing.T) {
	s, err := NewSatelliteFromTLE("2 28654  99.1770  43.4120 0013283 282.3899 185.0991 14.12152816502890\n", WGS72)
	if err != nil {
		t.Errorf(err.Error())
	}

	s.Name = "NOAA 18"
}
