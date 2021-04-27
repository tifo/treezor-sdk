package treezor

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimestampParis_MarshalJSON(t *testing.T) {
	t.Run("Success full format CEST", func(t *testing.T) {
		dateUTC0, _ := time.Parse(fullFormat, "2019-10-01 09:00:00")
		location, _ := time.LoadLocation("Europe/Paris")

		ti := &TimestampParis{}
		ti.Time = dateUTC0

		jsonDate, err := json.Marshal(ti)
		jsonDateString := string(jsonDate)

		expectedMarshal := `"` + dateUTC0.In(location).Format(fullFormat) + `"` // +2 Hours
		assert.Equal(t, expectedMarshal, jsonDateString)
		assert.Nil(t, err)
	})
	t.Run("Success full format CET", func(t *testing.T) {
		dateUTC0, _ := time.Parse(fullFormat, "2019-11-01 09:00:00")
		location, _ := time.LoadLocation("Europe/Paris")

		ti := &TimestampParis{}
		ti.Time = dateUTC0

		jsonDate, err := json.Marshal(ti)
		jsonDateString := string(jsonDate)

		expectedMarshal := `"` + dateUTC0.In(location).Format(fullFormat) + `"` // +1 Hours
		assert.Equal(t, expectedMarshal, jsonDateString)
		assert.Nil(t, err)
	})
	t.Run("Success emtpy date", func(t *testing.T) {
		ti := &TimestampParis{}

		data, err := json.Marshal(ti)
		assert.Equal(t, []byte(`"0000-00-00 00:00:00"`), data)
		assert.Nil(t, err)
	})
}

func TestTimestampParis_UnmarshalJSON(t *testing.T) {
	t.Run("Success full format", func(t *testing.T) {
		dateByte := []byte(`"2019-10-01 11:00:00"`)
		location, _ := time.LoadLocation("Europe/Paris")

		ti := &TimestampParis{}
		dateUTC0, _ := time.ParseInLocation(`"`+fullFormat+`"`, string(dateByte), location)
		expectedTi := NewTimestampParis(dateUTC0)
		expectedTi.OriginalPayload = string(dateByte)

		err := json.Unmarshal(dateByte, ti)
		assert.Equal(t, expectedTi, ti)
		assert.Nil(t, err)
	})
	t.Run("Success empty date", func(t *testing.T) {
		dateByte := []byte(`"0000-00-00 00:00:00"`)

		ti := &TimestampParis{}

		expectedTi := NewTimestampParis(time.Time{})
		expectedTi.OriginalPayload = string(dateByte)

		err := json.Unmarshal(dateByte, ti)
		assert.Equal(t, expectedTi, ti)
		assert.Nil(t, err)
	})
}
