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
		if err := d.decodeBool(val); err != nil {
			return err
		}

	case *int32:
		if err := d.decodeInt32(val); err != nil {
			return err
		}

	case *int64:
		if err := d.decodeInt64(val); err != nil {
			return err
		}

	case *uint8:
		if err := d.decodeUint8(val); err != nil {
			return err
		}

	case *uint16:
		if err := d.decodeUint16(val); err != nil {
			return err
		}

	case *uint32:
		if err := d.decodeUint32(val); err != nil {
			return err
		}

	case *uint64:
		if err := d.decodeUint64(val); err != nil {
			return err
		}

	case *[magicAndChecksumLength]byte:
		if err := d.decodeArray(magicAndChecksumLength, val[:]); err != nil {
			return err
		}

	case *[commandLength]byte:
		if err := d.decodeArray(commandLength, val[:]); err != nil {
			return err
		}

	case *[hashLength]byte:
		if err := d.decodeArray(hashLength, val[:]); err != nil {
			return err
		}

	case Unmarshaler:
		if err := val.UnmarshalBinary(d.r); err != nil {
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

func (d Decoder) decodeArray(len int64, out []byte) error {
	if _, err := io.LimitReader(d.r, len).Read(out); err != nil {
		return err
	}

	return nil
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

func (d Decoder) decodeBool(out *bool) error {
	lr := io.LimitReader(d.r, 1)

	if err := binary.Read(lr, binary.LittleEndian, out); err != nil {
		return err
	}

	return nil
}

func (d Decoder) decodeInt32(out *int32) error {
	lr := io.LimitReader(d.r, 4)

	if err := binary.Read(lr, binary.LittleEndian, out); err != nil {
		return err
	}

	return nil
}

func (d Decoder) decodeInt64(out *int64) error {
	lr := io.LimitReader(d.r, 8)

	if err := binary.Read(lr, binary.LittleEndian, out); err != nil {
		return err
	}

	return nil
}

func (d Decoder) decodeUint8(out *uint8) error {
	lr := io.LimitReader(d.r, 1)

	if err := binary.Read(lr, binary.LittleEndian, out); err != nil {
		return err
	}

	return nil
}

func (d Decoder) decodeUint16(out *uint16) error {
	lr := io.LimitReader(d.r, 2)

	if err := binary.Read(lr, binary.BigEndian, out); err != nil {
		return err
	}

	return nil
}

func (d Decoder) decodeUint32(out *uint32) error {
	lr := io.LimitReader(d.r, 4)

	if err := binary.Read(lr, binary.LittleEndian, out); err != nil {
		return err
	}

	return nil
}

func (d Decoder) decodeUint64(out *uint64) error {
	lr := io.LimitReader(d.r, 8)

	if err := binary.Read(lr, binary.LittleEndian, out); err != nil {
		return err
	}

	return nil
}
