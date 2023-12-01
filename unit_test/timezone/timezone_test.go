package timezone_test

import (
	"backend/unit_test/timezone"
	"testing"
)

func TestLoadFromTrustLocationOrNil(t *testing.T) {
	t.Run("returns location on valid timezone", func(t *testing.T) {
		location := timezone.LoadFromTrustLocationOrNil("Asia/Tehran")
		if location.String() == "" {
			t.Fatal("the location is nil for Asia/Tehran")
		}
	})

	t.Run("returns nill on invalid timezone", func(t *testing.T) {
		location := timezone.LoadFromTrustLocationOrNil("Tehran")
		if location.String() != "" {
			t.Fatal("the location is not nil for Tehran")
		}
	})

}
