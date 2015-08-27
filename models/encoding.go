package models

import (
	"encoding/json"
	"reflect"

	"github.com/gogo/protobuf/proto"
	"github.com/pivotal-golang/lager"
)

type SerializationFormat byte

const JSON SerializationFormat = 0
const PROTO SerializationFormat = 1

type Version byte

const V0 Version = 0

type Envelope struct {
	SerializationFormat SerializationFormat
	Version             Version
	Payload             []byte
}

func OpenEnvelope(data []byte) *Envelope {
	e := &Envelope{}

	if !isEncoded(data) {
		e.SerializationFormat = JSON
		e.Version = V0
		e.Payload = data
		return e
	}

	e.SerializationFormat = SerializationFormat(data[0])
	e.Version = Version(V0)
	e.Payload = data[2:]
	return e
}

func Marshal(version Version, v ProtoValidator) ([]byte, *Error) {
	payload, err := toProto(v)
	if err != nil {
		return nil, err
	}
	return append([]byte{byte(PROTO), byte(version)}, payload...), nil
}

func (e *Envelope) Unmarshal(logger lager.Logger, model Validator) *Error {
	switch e.SerializationFormat {
	case JSON:
		err := json.Unmarshal(e.Payload, model)
		if err != nil {
			logger.Error("failed-to-json-unmarshal-payload", err)
			return NewError(InvalidRecord, err.Error())
		}
	case PROTO:
		if pv, ok := model.(ProtoValidator); ok {
			err := proto.Unmarshal(e.Payload, pv)
			if err != nil {
				logger.Error("failed-to-proto-unmarshal-payload", err)
				return NewError(InvalidRecord, err.Error())
			}
		} else {
			logger.Error("cannot-unmarshal-into-non-proto-model", nil)
			return NewError(FailedToOpenEnvelope, "cannot unmarshal protobuffer")
		}
	default:
		logger.Error("cannot-unmarshal-unknown-serialization-format", nil)
		return NewError(FailedToOpenEnvelope, "unknown serialization format")
	}

	if versioner, ok := model.(Versioner); ok {
		versioner.MigrateFromVersion(e.Version)
	}

	err := model.Validate()
	if err != nil {
		logger.Error("invalid-record", err)
		return NewError(InvalidRecord, err.Error())
	}
	return nil
}

func isEncoded(data []byte) bool {
	if len(data) < 2 {
		return false
	}

	switch SerializationFormat(data[0]) {
	case JSON, PROTO:
	default:
		return false
	}

	switch Version(data[1]) {
	case V0:
	default:
		return false
	}

	return true
}

func FromJSON(payload []byte, v Validator) error {
	err := json.Unmarshal(payload, v)
	if err != nil {
		return err
	}
	return v.Validate()
}

func ToJSON(v Validator) ([]byte, *Error) {
	if !isNil(v) {
		if err := v.Validate(); err != nil {
			return nil, NewError(InvalidRecord, err.Error())
		}
	}

	bytes, err := json.Marshal(v)
	if err != nil {
		return nil, NewError(InvalidJSON, err.Error())
	}

	return bytes, nil
}

func toProto(v ProtoValidator) ([]byte, *Error) {
	if !isNil(v) {
		if err := v.Validate(); err != nil {
			return nil, NewError(InvalidRecord, err.Error())
		}
	}

	bytes, err := proto.Marshal(v)
	if err != nil {
		return nil, NewError(InvalidProtobufMessage, err.Error())
	}

	return bytes, nil
}

func isNil(a interface{}) bool {
	if a == nil {
		return true
	}

	switch reflect.TypeOf(a).Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return reflect.ValueOf(a).IsNil()
	}

	return false
}
