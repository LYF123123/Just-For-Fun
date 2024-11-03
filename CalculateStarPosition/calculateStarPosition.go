package calculatestarposition

//! Warning Julian Date Calculate Method in this file not work
import "math"

const degToRad float64 = math.Pi / 180.0
const radToDeg float64 = 180 / math.Pi

func calculateGST(year, month, day int, utcHours float64) float64 {
	// Julian Date Calculate
	a := (14 - month/12)
	y := year + 4800 - a
	m := month + 12*a - 3
	jdn := day + ((153*m + 2) / 5) + 365*y + y/4 - y/100 + y/400 - 32045

	jd := float64(jdn) + (utcHours-12.0)/24.0
	d := jd - 2451545.0 // Calculate the days from 2000/jan/1st 12:00
	t := d / 36525.0
	gst := 6.697374558 + 24.06570982441908*float64(d) + 0.000026*float64(t)*float64(t)
	return math.Mod(gst, 24.0)
}

func calculateAzimuthAndAltitude(HA, Dec, Latitude float64) (Azimuth, Altitude float64) {
	HA_rad := HA * degToRad
	Dec_rad := Dec * degToRad
	Latitude_rad := Latitude * degToRad

	sin_Alt := math.Sin(Dec_rad)*math.Sin(Latitude_rad) + math.Cos(Dec_rad)*math.Cos(Latitude_rad)*math.Cos(HA_rad)
	Altitude = math.Asin(sin_Alt) * radToDeg

	cos_Az := (math.Sin(Dec_rad) - math.Sin(Altitude*degToRad)*math.Sin(Latitude_rad)) / (math.Cos(Altitude*degToRad) * math.Cos(Latitude_rad))
	Azimuth = math.Acos(cos_Az) * radToDeg
	if math.Sin(HA_rad) > 0 {
		Azimuth = 360 - Azimuth
	}
	return Azimuth, Altitude
}

func CalculateStarPosition(year, month, day, hour, minute, second int, longitude, latitude, raHours, decDegrees float64) (azimuth, altitude float64) {
	// UTC 时间
	// https://github.com/nayarsystems/posix_tz_db/blob/master/zones.csv
	// initTime("CST-8"); // Set for Asia/Shanghai

	// float utcHours = 21 + 34.0 / 60 + 56.0 / 3600;
	utcHours := float64(hour) + float64(minute)/60 + float64(second)/3600

	// Calculate GST and LST
	gst := calculateGST(year, month, day, utcHours)
	lst := math.Mod(gst+longitude/15.0, 24.0)
	ha := math.Mod(lst*15-raHours*15, 360)
	if ha < 0 {
		ha += 360
	}
	azimuth, altitude = calculateAzimuthAndAltitude(ha, decDegrees, latitude)

	return azimuth, altitude

}
