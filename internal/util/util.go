package util

import "time"

var (
	DefaultLayout = time.RFC3339
	USLocation, _ = time.LoadLocation("America/New_York")
)

func Loop(fns ...func() error) error {
	for _, fn := range fns {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}

func InRegularMarketingTime(tm time.Time) bool {
	start := time.Date(tm.Year(), tm.Month(), tm.Day(), 9, 29, 59, 59, USLocation)
	end := time.Date(tm.Year(), tm.Month(), tm.Day(), 16, 0, 0, 1, USLocation)
	return tm.After(start) && tm.Before(end)
}
