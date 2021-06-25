package types

import (
	"encoding/json"
	"strconv"
)

type Level int32

func (l *Level) UnmarshalJSON(data []byte) error {
	var str json.Number
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	v, err := str.Int64()
	if err != nil {
		return err
	}
	*l = Level(v)
	return nil
}

const (
	LevelNone          Level = 0
	LevelPending       Level = 1
	LevelRegular       Level = 2
	LevelStrong        Level = 3
	LevelRefused       Level = 4
	LevelInvestigating Level = 5
)

var levelNames = map[int32]string{
	0: "LEVEL_NONE",
	1: "LEVEL_PENDING",
	2: "LEVEL_REGULAR",
	3: "LEVEL_STRONG",
	4: "LEVEL_REFUSED",
	5: "LEVEL_INVESTIGATING",
}

func (l Level) String() string {
	s, ok := levelNames[int32(l)]
	if ok {
		return s
	}
	return strconv.Itoa(int(l))
}
