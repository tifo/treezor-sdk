package types

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimestamp_MarshalJSON(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		dateUTC0, _ := time.Parse(fullFormat, "2019-10-01 09:00:00")

		ti := &timestamp{}
		ti.Time = dateUTC0

		jsonDate, err := json.Marshal(ti)
		jsonDateString := string(jsonDate)

		expectedMarshal := `"` + dateUTC0.UTC().Format(fullFormat) + `"`
		assert.Equal(t, expectedMarshal, jsonDateString)
		assert.Nil(t, err)
	})
	t.Run("Success emtpy date", func(t *testing.T) {
		ti := &timestamp{}

		data, err := json.Marshal(ti)
		assert.Equal(t, []byte(`"0000-00-00 00:00:00"`), data)
		assert.Nil(t, err)
	})
}

func TestTimestamp_UnmarshalJSON(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		dateByte := []byte(`"2019-10-01 11:00:00"`)

		ti := &timestamp{}
		dateUTC0, _ := time.Parse(`"`+fullFormat+`"`, string(dateByte))
		expectedTi := &timestamp{
			Time:            dateUTC0,
			OriginalPayload: string(dateByte),
		}

		err := json.Unmarshal(dateByte, ti)
		assert.Equal(t, expectedTi, ti)
		assert.Nil(t, err)
	})
	t.Run("Success empty date", func(t *testing.T) {
		dateByte := []byte(`"0000-00-00 00:00:00"`)

		ti := &timestamp{}

		expectedTi := &timestamp{
			OriginalPayload: string(dateByte),
		}

		err := json.Unmarshal(dateByte, ti)
		assert.Equal(t, expectedTi, ti)
		assert.Nil(t, err)
	})
	t.Run("Success with wrong format", func(t *testing.T) {
		dateByte := []byte(`"2019-10-01"`)

		ti := &timestamp{}

		expectedTi := &timestamp{
			OriginalPayload: string(dateByte),
		}

		err := json.Unmarshal(dateByte, ti)
		assert.Equal(t, expectedTi, ti)
		assert.Nil(t, err)
	})
}
