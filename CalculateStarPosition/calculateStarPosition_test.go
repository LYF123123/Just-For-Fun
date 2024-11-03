package calculatestarposition

import "testing"

func TestCalculateStarPosition(t *testing.T) {

	year := 2024
	month := 11
	day := 3
	hour := 20
	minute := 43
	second := 30
	// Shanghai
	longitude := 121.93
	latitude := 30.9
	raHours := 14.261
	decDegrees := 19.1822
	azimuth, altitude := CalculateStarPosition(year, month, day, hour, minute, second, longitude, latitude, raHours, decDegrees)
	t.Logf("Azimuth: %f", azimuth)
	t.Logf("Altitude: %f", altitude)
	
}
