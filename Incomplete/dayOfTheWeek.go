func dayOfTheWeek(day int, month int, year int) string {
    weekday := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
    local := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
    return weekday[local.Weekday()]
}
