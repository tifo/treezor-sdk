package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	json "github.com/tifo/treezor-sdk/internal/json"
)

func TestDate_MarshalJSON(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		dateUTC0, _ := time.Parse(shortFormat, "2019-10-01")

		ti := &Date{}
		ti.Time = dateUTC0

		jsonDate, err := json.Marshal(ti)
		jsonDateString := string(jsonDate)

		expectedMarshal := `"` + dateUTC0.UTC().Format(shortFormat) + `"`
		assert.Equal(t, expectedMarshal, jsonDateString)
		assert.Nil(t, err)
	})
	t.Run("Success emtpy date", func(t *testing.T) {
		ti := &Date{}

		data, err := json.Marshal(ti)
		assert.Equal(t, []byte(`"0000-00-00"`), data)
		assert.Nil(t, err)
	})
}

func TestDate_UnmarshalJSON(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		dateByte := []byte(`"2019-10-01"`)

		ti := &Date{}
		dateUTC0, _ := time.Parse(`"`+shortFormat+`"`, string(dateByte))
		expectedTi := &Date{
			Time:            dateUTC0,
			OriginalPayload: string(dateByte),
		}

		err := json.Unmarshal(dateByte, ti)
		assert.Equal(t, expectedTi, ti)
		assert.Nil(t, err)
	})
	t.Run("Success empty date", func(t *testing.T) {
		dateByte := []byte(`"0000-00-00"`)

		ti := &Date{}

		expectedTi := &Date{
			OriginalPayload: string(dateByte),
		}

		err := json.Unmarshal(dateByte, ti)
		assert.Equal(t, expectedTi, ti)
		assert.Nil(t, err)
	})
	t.Run("Error wrong format", func(t *testing.T) {
		dateByte := []byte(`"2019-10-01 11:00:00"`)

		ti := &Date{}

		err := json.Unmarshal(dateByte, ti)
		assert.EqualError(t, err, `unmarshalerDecoder: treezor.Date: parsing time ""2019-10-01 11:00:00"" as ""2006-01-02"": cannot parse " 11:00:00"" as """, error found in #10 byte of ...| 11:00:00"|..., bigger context ...|"2019-10-01 11:00:00"|...`)
	})
}
