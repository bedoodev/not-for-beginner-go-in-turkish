package main

import (
	"fmt"
	"time"
)

func main() {
	// Current local time
	fmt.Println(time.Now()) // 2025-06-30 12:36:50.667074 +0300 +03 m=+0.000160084

	// Spesific time
	specificTime := time.Date(2025, time.June, 30, 12, 36, 50, 0, time.UTC)
	fmt.Println("Spesific Time:", specificTime) // Spesific Time: 2025-06-30 12:36:50 +0000 UTC

	// Parse time => Mon Jan 2 15:04:05 MST 2006
	parsedTime1, _ := time.Parse("2006-01-02", "2025-06-30")
	parsedTime2, _ := time.Parse("06-01-02", "25-06-30")
	parsedTime3, _ := time.Parse("06-1-2", "25-6-30")
	parsedTime4, _ := time.Parse("06-1-2 15-04", "25-6-30 12-36")

	fmt.Println("Parsed Time1:", parsedTime1) // Parsed Time1: 2025-06-30 00:00:00 +0000 UTC
	fmt.Println("Parsed Time2:", parsedTime2) // Parsed Time2: 2025-06-30 00:00:00 +0000 UTC
	fmt.Println("Parsed Time3:", parsedTime3) // Parsed Time3: 2025-06-30 00:00:00 +0000 UTC
	fmt.Println("Parsed Time4:", parsedTime4) // Parsed Time4: 2025-06-30 12:36:00 +0000 UTC

	// Formatting time
	t := time.Now()
	fmt.Println("Formatted time:", t.Format("01-02-2006"))                  // Formatted time: 06-30-2025
	fmt.Println("Formatted time2:", t.Format("2006-01-02"))                 // Formatted time2: 2025-06-30
	fmt.Println("Formatted time3:", t.Format("2006-01-02 15-04-05"))        // Formatted time3: 2025-06-30 13-26-03
	fmt.Println("Formatted time4:", t.Format("Monday 2006-01-02 15-04-05")) // Formatted time4: Monday 2025-06-30 13-29-48
	fmt.Println("Formatted time5:", t.Format("Mon 2006-01-02 15-04-05"))    // Formatted time5: Mon 2025-06-30 13-30-13

	// Calculation
	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println("One day later:", oneDayLater)             // One day later: 2025-07-01 13:53:07.387853 +0300 +03 m=+86400.000772418
	fmt.Println("Day of the week:", oneDayLater.Weekday()) // Day of the week: Tuesday

	// Round time
	fmt.Println("Rounded Time:", t.Round(time.Hour))    // Rounded Time: 2025-06-30 14:00:00 +0300 +03
	fmt.Println("Rounded Time2:", t.Round(time.Minute)) // Rounded Time2: 2025-06-30 13:56:00 +0300 +03

	// Convert to spesific time zone
	loc, _ := time.LoadLocation("Europe/Istanbul")
	loc2, _ := time.LoadLocation("America/New_York")
	t = time.Date(2025, time.June, 30, 13, 56, 0, 0, time.UTC)
	tLocal := t.In(loc)
	tLocal2 := t.In(loc2)
	fmt.Println("Europe/Istanbul Time", tLocal)   // Europe/Istanbul Time 2025-06-30 16:56:00 +0300 +03
	fmt.Println("America/New_York Time", tLocal2) // America/New_York Time 2025-06-30 09:56:00 -0400 EDT

	// Truncate time
	truncatedTime := t.Truncate(time.Hour)
	fmt.Println("Truncated Time:", truncatedTime) // Truncated Time: 2025-06-30 13:00:00 +0000 UTC

	// Compare time
	t1 := time.Date(2025, time.June, 30, 17, 10, 0, 0, time.UTC)
	t2 := time.Date(2025, time.June, 30, 12, 25, 0, 0, time.UTC)

	fmt.Println("Time Difference:", t1.Sub(t2)) // Time Difference: 4h45m0s
	fmt.Println("Time Difference:", t2.Sub(t1)) // Time Difference: -4h45m0s

	fmt.Println("t2 is after t1?", t2.After(t1))   // t2 is after t1? false
	fmt.Println("t2 is before t1?", t2.Before(t1)) // t2 is before t1? true

	// Unix Time (Epoch)
	now := time.Now()
	unixTime := now.Unix()
	fmt.Println("Unix Time:", unixTime) // Unix Time: 1751288398
	ut := time.Unix(unixTime, 0)
	fmt.Println("Time:", ut) // Time: 2025-06-30 16:01:23 +0300 +03
	fmt.Println("Formatted Time:", ut.Format("2006-01-02")) // Formatted Time: 2025-06-30

}
