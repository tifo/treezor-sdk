package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	json "github.com/tifo/treezor-sdk/internal/json"
)

func TestTimestampLondon_MarshalJSON(t *testing.T) {
	t.Run("Success full format CEST", func(t *testing.T) {
		dateUTC0, _ := time.Parse(fullFormat, "2019-10-01 09:00:00")
		location, _ := time.LoadLocation("Europe/London")

		ti := &TimestampLondon{}
		ti.Time = dateUTC0

		jsonDate, err := json.Marshal(ti)
		jsonDateString := string(jsonDate)

		expectedMarshal := `"` + dateUTC0.In(location).Format(fullFormat) + `"` // +2 Hours
		assert.Equal(t, expectedMarshal, jsonDateString)
		assert.Nil(t, err)
	})
	t.Run("Success full format CET", func(t *testing.T) {
		dateUTC0, _ := time.Parse(fullFormat, "2019-11-01 09:00:00")
		location, _ := time.LoadLocation("Europe/London")

		ti := &TimestampLondon{}
		ti.Time = dateUTC0

		jsonDate, err := json.Marshal(ti)
		jsonDateString := string(jsonDate)

		expectedMarshal := `"` + dateUTC0.In(location).Format(fullFormat) + `"` // +1 Hours
		assert.Equal(t, expectedMarshal, jsonDateString)
		assert.Nil(t, err)
	})
	t.Run("Success emtpy date", func(t *testing.T) {
		ti := &TimestampLondon{}

		data, err := json.Marshal(ti)
		assert.Equal(t, []byte(`"0000-00-00 00:00:00"`), data)
		assert.Nil(t, err)
	})
}

func TestTimestampLondon_UnmarshalJSON(t *testing.T) {
	t.Run("Success full format", func(t *testing.T) {
		dateByte := []byte(`"2019-10-01 11:00:00"`)
		location, _ := time.LoadLocation("Europe/London")

		ti := &TimestampLondon{}
		dateUTC0, _ := time.ParseInLocation(`"`+fullFormat+`"`, string(dateByte), location)
		expectedTi := NewTimestampLondon(dateUTC0)
		expectedTi.OriginalPayload = string(dateByte)

		err := json.Unmarshal(dateByte, ti)
		assert.Equal(t, expectedTi, ti)
		assert.Nil(t, err)
	})
	t.Run("Success empty date", func(t *testing.T) {
		dateByte := []byte(`"0000-00-00 00:00:00"`)

		ti := &TimestampLondon{}

		expectedTi := NewTimestampLondon(time.Time{})
		expectedTi.OriginalPayload = string(dateByte)

		err := json.Unmarshal(dateByte, ti)
		assert.Equal(t, expectedTi, ti)
		assert.Nil(t, err)
	})
}
