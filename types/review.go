package types

import (
	"encoding/json"
	"strconv"
)

type Review int32

func (r *Review) UnmarshalJSON(data []byte) error {
	var str json.Number
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	v, err := str.Int64()
	if err != nil {
		return err
	}
	*r = Review(v)
	return nil
}

const (
	ReviewNone      Review = 0
	ReviewPending   Review = 1
	ReviewValidated Review = 2
	ReviewRefused   Review = 3
)

var reviewNames = map[int32]string{
	0: "REVIEW_NONE",
	1: "REVIEW_PENDING",
	2: "REVIEW_VALIDATED",
	3: "REVIEW_REFUSED",
}

func (r Review) String() string {
	s, ok := reviewNames[int32(r)]
	if ok {
		return s
	}
	return strconv.Itoa(int(r))
}
