import (
    "time"
    "fmt"
    "math"
)
const dateFormat = "2006-01-02"

func daysBetweenDates(date1 string, date2 string) int {
    t1, _ := time.Parse(dateFormat, date1)
    t2, _ := time.Parse(dateFormat, date2)

    return int(math.Abs(float64(t2.Sub(t1).Hours() / 24)))
}
