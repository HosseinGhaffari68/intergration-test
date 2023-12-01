package timezone

import "time"

func LoadFromTrustLocationOrNil(location string) *time.Location {
	tz, err := time.LoadLocation(location)
	if err != nil {
		//todo: log error
		return &time.Location{}
	}
	return tz
}
