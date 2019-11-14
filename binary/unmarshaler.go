package binary

import (
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
)

// Unmarshaler is the interface implemented by types that can unmarshal themselves from binary.
type Unmarshaler interface {
	UnmarshalBinary(r io.Reader) error
}

// Decoder reads and decodes binary data from an input stream.
type Decoder struct {
	r io.Reader
}

// NewDecoder returns a new Decoder that reads  from r.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r: r,
	}
}

// Decode ...
func (d Decoder) Decode(v interface{}) error {
	switch val := v.(type) {
	case *bool:
		d, err := d.decodeBool()
		if err != nil {
			return err
		}

		*val = d

	case *int32:
		d, err := d.decodeInt32()
		if err != nil {
			return err
		}

		*val = d

	case *int64:
		d, err := d.decodeInt64()
		if err != nil {
			return err
		}

		*val = d

	case *uint8:
		d, err := d.decodeUint8()
		if err != nil {
			return err
		}

		*val = d

	case *uint16:
		d, err := d.decodeUint16()
		if err != nil {
			return err
		}

		*val = d

	case *uint32:
		d, err := d.decodeUint32()
		if err != nil {
			return err
		}

		*val = d

	case *uint64:
		d, err := d.decodeUint64()
		if err != nil {
			return err
		}

		*val = d

	case *[magicAndChecksumLength]byte:
		err := d.decodeArray(magicAndChecksumLength, val)
		if err != nil {
			return err
		}

	case *[commandLength]byte:
		err := d.decodeArray(commandLength, val)
		if err != nil {
			return err
		}

	case Unmarshaler:
		err := val.UnmarshalBinary(d.r)
		if err != nil {
			return err
		}

	default:
		if reflect.ValueOf(v).Kind() == reflect.Ptr &&
			reflect.ValueOf(v).Elem().Kind() == reflect.Struct {
			if err := d.decodeStruct(v); err != nil {
				return err
			}
			break
		}

		return fmt.Errorf("unsupported type %s", reflect.TypeOf(v).String())
	}

	return nil
}

func (d Decoder) decodeArray(len int64, out interface{}) (err error) {
	lr := io.LimitReader(d.r, len)

	if err = binary.Read(lr, binary.LittleEndian, out); err != nil {
		return
	}

	return
}

func (d Decoder) decodeStruct(v interface{}) error {
	val := reflect.Indirect(reflect.ValueOf(v))

	for i := 0; i < val.NumField(); i++ {
		if err := d.Decode(val.Field(i).Addr().Interface()); err != nil {
			return err
		}

	}

	return nil
}

func (d Decoder) decodeBool() (out bool, err error) {
	lr := io.LimitReader(d.r, 1)

	if err = binary.Read(lr, binary.LittleEndian, &out); err != nil {
		return
	}

	return
}

func (d Decoder) decodeInt32() (out int32, err error) {
	lr := io.LimitReader(d.r, 4)

	if err = binary.Read(lr, binary.LittleEndian, &out); err != nil {
		return
	}

	return
}

func (d Decoder) decodeInt64() (out int64, err error) {
	lr := io.LimitReader(d.r, 8)

	if err = binary.Read(lr, binary.LittleEndian, &out); err != nil {
		return
	}

	return
}

func (d Decoder) decodeUint8() (out uint8, err error) {
	lr := io.LimitReader(d.r, 1)

	if err = binary.Read(lr, binary.LittleEndian, &out); err != nil {
		return
	}

	return
}

func (d Decoder) decodeUint16() (out uint16, err error) {
	lr := io.LimitReader(d.r, 2)

	if err = binary.Read(lr, binary.BigEndian, &out); err != nil {
		return
	}

	return
}

func (d Decoder) decodeUint32() (out uint32, err error) {
	lr := io.LimitReader(d.r, 4)

	if err = binary.Read(lr, binary.LittleEndian, &out); err != nil {
		return
	}

	return
}

func (d Decoder) decodeUint64() (out uint64, err error) {
	lr := io.LimitReader(d.r, 8)

	if err = binary.Read(lr, binary.LittleEndian, &out); err != nil {
		return
	}

	return
}

// Unmarshal parses the binary-encoded data and stores the result in the value pointed to by v.
func Unmarshal(data []byte, v interface{}) error {
	switch v.(type) {
	case *uint8:
		v = unmarshalUint8(data[0])
	default:
		return fmt.Errorf("unsupported type %s", reflect.TypeOf(v).String())
	}

	return nil
}

func unmarshalUint8(data byte) uint8 {
	return uint8(data)
}
