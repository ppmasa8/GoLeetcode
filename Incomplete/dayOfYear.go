func dayOfYear(date string) int {
    year, _ := strconv.Atoi(date[0:4])
    month, _  := strconv.Atoi(date[5:7])
    day, _  := strconv.Atoi(date[8:10])
    start := time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)
    end := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)

    return int(end.Sub(start).Hours() / 24 + 1)
}