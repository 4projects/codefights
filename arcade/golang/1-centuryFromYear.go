import "math"

func centuryFromYear(year int) int {
    
    // string convert, pick first two slice
    // if the year rounded + 1 is === to the amount don't add 1 otherwise do
    y := math.Floor(float64(year)) / 100
    if year % 100 != 0 {
	y ++ 
    }
    return int(y)
}
