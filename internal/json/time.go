package json

import (
	"strings"
	"time"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"

	"github.com/tifo/treezor-sdk/internal/xtime"
)

type TimeExtension struct {
	jsoniter.DummyExtension
}

func (e *TimeExtension) UpdateStructDescriptor(sd *jsoniter.StructDescriptor) {

	for _, binding := range sd.Fields {

		var err error
		var isPtr bool

		typeName := binding.Field.Type().String()
		switch typeName {
		case "time.Time":
			isPtr = false
		case "*time.Time":
			isPtr = true
		default:
			continue
		}

		timeFormat := xtime.DefaultFormat
		fmtTag := binding.Field.Tag().Get(xtime.StructTagFormat)
		if timeFmt := xtime.GetFormat(fmtTag); timeFmt != "" {
			timeFormat = timeFmt
		}

		var timeLocation *time.Location
		locTag := binding.Field.Tag().Get(xtime.StructTagLocation)
		if timeLoc, ok := xtime.GetLocation(locTag); ok {
			timeLocation = timeLoc
		} else if locTag != "" {
			err = errors.Errorf("invalid or unloaded location specified: %s", locTag)
		}

		encDec := &timeEncoderDecoder{err: err, isPtr: isPtr, loc: timeLocation, fmt: timeFormat}
		binding.Encoder = encDec
		binding.Decoder = encDec
	}
}

type timeEncoderDecoder struct {
	isPtr bool
	err   error
	loc   *time.Location
	fmt   string
}

func (ed *timeEncoderDecoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	if ed.err != nil {
		stream.Error = ed.err
		return
	}

	var tp *time.Time
	if ed.isPtr {
		tpp := (**time.Time)(ptr)
		tp = *(tpp)
	} else {
		tp = (*time.Time)(ptr)
	}

	if tp != nil {
		lt := *tp
		if ed.loc != nil {
			lt = tp.In(ed.loc)
		}
		str := lt.Format(ed.fmt)
		stream.WriteString(str)
	} else {
		stream.WriteString("null")
	}
}

func (ed *timeEncoderDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	if ed.err != nil {
		iter.Error = ed.err
		return
	}

	str := iter.ReadString()
	var t *time.Time
	switch {
	case str == "" || str == "0000-00-00 00:00:00":
		t = nil
	case strings.HasPrefix(str, "-0001-"):
		// NOTE: failsafe when a negative time is returned
		t = nil
	default:
		var err error
		var tmp time.Time
		if ed.loc != nil {
			tmp, err = time.ParseInLocation(ed.fmt, str, ed.loc)
		} else {
			tmp, err = time.Parse(ed.fmt, str)
		}
		if err != nil {
			iter.Error = err
			return
		}
		t = &tmp
	}

	if ed.isPtr {
		tpp := (**time.Time)(ptr)
		*tpp = t
	} else {
		tp := (*time.Time)(ptr)
		if tp != nil && t != nil {
			*tp = *t
		}
	}
}

func (ed *timeEncoderDecoder) IsEmpty(ptr unsafe.Pointer) bool {
	var tp *time.Time
	if ed.isPtr {
		tpp := (**time.Time)(ptr)
		tp = *(tpp)
	} else {
		tp = (*time.Time)(ptr)
	}
	if tp == nil {
		return true
	}
	return false
}
