package generator

import (
	"math"
	"math/rand"
	"time"

	"practice/storage"
)

// RandomTimestamp generates random timestamp
func RandomTimestamp() time.Time {
	const unixTimeSecond = 94608000
	randomTime := rand.Int63n(time.Now().Unix()-unixTimeSecond) + unixTimeSecond
	randomNow := time.Unix(randomTime, 0)
	return randomNow
}

// GenerateTelematic generates telematic data and sends it to the channel
func GenerateTelematic(num int, c chan *storage.Car) {
	const (
		minLatitude  = -90
		maxLatitude  = 90
		minLongitude = -180
		maxLongitude = 180
	)
	prevTimestamp := RandomTimestamp()
	prevCoords := storage.Coordinates{minLatitude + rand.Float64()*maxLatitude*2, minLongitude + rand.Float64()*maxLongitude*2} //res[i] = min + rand.Float64() * (max - min)
	for {
		newTimestamp := rand.Intn(60)
		newTime := prevTimestamp.Add(time.Duration(newTimestamp) * time.Second)
		newSpeed := rand.Intn(120)
		//log.Println(".............", t, newSpeed, newTime)
		var s float64
		s = (float64(newSpeed) * float64(newTimestamp)) / 3600 / 111
		x := -float64(s) + rand.Float64()*(float64(s)*2)         // random latitude change
		y := math.Sqrt(math.Pow(float64(s), 2) - math.Pow(x, 2)) // random longitude change
		newCoords := storage.Coordinates{prevCoords.Latitude + x, prevCoords.Longitude + y}

		car := &storage.Car{
			0,
			num,
			newSpeed,
			newCoords,
			newTime}
		c <- car
		prevTimestamp = newTime
		prevCoords = newCoords
	}
}
