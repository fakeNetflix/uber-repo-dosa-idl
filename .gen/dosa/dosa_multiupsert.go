// Code generated by thriftrw v1.13.0. DO NOT EDIT.
// @generated

package dosa

import (
	"errors"
	"fmt"
	"go.uber.org/multierr"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// Dosa_MultiUpsert_Args represents the arguments for the Dosa.multiUpsert function.
//
// The arguments for multiUpsert are sent and received over the wire as this struct.
type Dosa_MultiUpsert_Args struct {
	Request *MultiUpsertRequest `json:"request,omitempty"`
}

// ToWire translates a Dosa_MultiUpsert_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Dosa_MultiUpsert_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Request != nil {
		w, err = v.Request.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _MultiUpsertRequest_Read(w wire.Value) (*MultiUpsertRequest, error) {
	var v MultiUpsertRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Dosa_MultiUpsert_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Dosa_MultiUpsert_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Dosa_MultiUpsert_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Dosa_MultiUpsert_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _MultiUpsertRequest_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a Dosa_MultiUpsert_Args
// struct.
func (v *Dosa_MultiUpsert_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Request != nil {
		fields[i] = fmt.Sprintf("Request: %v", v.Request)
		i++
	}

	return fmt.Sprintf("Dosa_MultiUpsert_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Dosa_MultiUpsert_Args match the
// provided Dosa_MultiUpsert_Args.
//
// This function performs a deep comparison.
func (v *Dosa_MultiUpsert_Args) Equals(rhs *Dosa_MultiUpsert_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Request == nil && rhs.Request == nil) || (v.Request != nil && rhs.Request != nil && v.Request.Equals(rhs.Request))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Dosa_MultiUpsert_Args.
func (v *Dosa_MultiUpsert_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v.Request != nil {
		err = multierr.Append(err, enc.AddObject("request", v.Request))
	}
	return err
}

// GetRequest returns the value of Request if it is set or its
// zero value if it is unset.
func (v *Dosa_MultiUpsert_Args) GetRequest() (o *MultiUpsertRequest) {
	if v.Request != nil {
		return v.Request
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "multiUpsert" for this struct.
func (v *Dosa_MultiUpsert_Args) MethodName() string {
	return "multiUpsert"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Dosa_MultiUpsert_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Dosa_MultiUpsert_Helper provides functions that aid in handling the
// parameters and return values of the Dosa.multiUpsert
// function.
var Dosa_MultiUpsert_Helper = struct {
	// Args accepts the parameters of multiUpsert in-order and returns
	// the arguments struct for the function.
	Args func(
		request *MultiUpsertRequest,
	) *Dosa_MultiUpsert_Args

	// IsException returns true if the given error can be thrown
	// by multiUpsert.
	//
	// An error can be thrown by multiUpsert only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for multiUpsert
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// multiUpsert into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by multiUpsert
	//
	//   value, err := multiUpsert(args)
	//   result, err := Dosa_MultiUpsert_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from multiUpsert: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*MultiUpsertResponse, error) (*Dosa_MultiUpsert_Result, error)

	// UnwrapResponse takes the result struct for multiUpsert
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if multiUpsert threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Dosa_MultiUpsert_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Dosa_MultiUpsert_Result) (*MultiUpsertResponse, error)
}{}

func init() {
	Dosa_MultiUpsert_Helper.Args = func(
		request *MultiUpsertRequest,
	) *Dosa_MultiUpsert_Args {
		return &Dosa_MultiUpsert_Args{
			Request: request,
		}
	}

	Dosa_MultiUpsert_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *BadRequestError:
			return true
		case *InternalServerError:
			return true
		default:
			return false
		}
	}

	Dosa_MultiUpsert_Helper.WrapResponse = func(success *MultiUpsertResponse, err error) (*Dosa_MultiUpsert_Result, error) {
		if err == nil {
			return &Dosa_MultiUpsert_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_MultiUpsert_Result.ClientError")
			}
			return &Dosa_MultiUpsert_Result{ClientError: e}, nil
		case *InternalServerError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_MultiUpsert_Result.ServerError")
			}
			return &Dosa_MultiUpsert_Result{ServerError: e}, nil
		}

		return nil, err
	}
	Dosa_MultiUpsert_Helper.UnwrapResponse = func(result *Dosa_MultiUpsert_Result) (success *MultiUpsertResponse, err error) {
		if result.ClientError != nil {
			err = result.ClientError
			return
		}
		if result.ServerError != nil {
			err = result.ServerError
			return
		}

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// Dosa_MultiUpsert_Result represents the result of a Dosa.multiUpsert function call.
//
// The result of a multiUpsert execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Dosa_MultiUpsert_Result struct {
	// Value returned by multiUpsert after a successful execution.
	Success     *MultiUpsertResponse `json:"success,omitempty"`
	ClientError *BadRequestError     `json:"clientError,omitempty"`
	ServerError *InternalServerError `json:"serverError,omitempty"`
}

// ToWire translates a Dosa_MultiUpsert_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Dosa_MultiUpsert_Result) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.ClientError != nil {
		w, err = v.ClientError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.ServerError != nil {
		w, err = v.ServerError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Dosa_MultiUpsert_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _MultiUpsertResponse_Read(w wire.Value) (*MultiUpsertResponse, error) {
	var v MultiUpsertResponse
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Dosa_MultiUpsert_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Dosa_MultiUpsert_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Dosa_MultiUpsert_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Dosa_MultiUpsert_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _MultiUpsertResponse_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.ClientError, err = _BadRequestError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.ServerError, err = _InternalServerError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.ClientError != nil {
		count++
	}
	if v.ServerError != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Dosa_MultiUpsert_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Dosa_MultiUpsert_Result
// struct.
func (v *Dosa_MultiUpsert_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [3]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.ClientError != nil {
		fields[i] = fmt.Sprintf("ClientError: %v", v.ClientError)
		i++
	}
	if v.ServerError != nil {
		fields[i] = fmt.Sprintf("ServerError: %v", v.ServerError)
		i++
	}

	return fmt.Sprintf("Dosa_MultiUpsert_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Dosa_MultiUpsert_Result match the
// provided Dosa_MultiUpsert_Result.
//
// This function performs a deep comparison.
func (v *Dosa_MultiUpsert_Result) Equals(rhs *Dosa_MultiUpsert_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.ClientError == nil && rhs.ClientError == nil) || (v.ClientError != nil && rhs.ClientError != nil && v.ClientError.Equals(rhs.ClientError))) {
		return false
	}
	if !((v.ServerError == nil && rhs.ServerError == nil) || (v.ServerError != nil && rhs.ServerError != nil && v.ServerError.Equals(rhs.ServerError))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Dosa_MultiUpsert_Result.
func (v *Dosa_MultiUpsert_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v.Success != nil {
		err = multierr.Append(err, enc.AddObject("success", v.Success))
	}
	if v.ClientError != nil {
		err = multierr.Append(err, enc.AddObject("clientError", v.ClientError))
	}
	if v.ServerError != nil {
		err = multierr.Append(err, enc.AddObject("serverError", v.ServerError))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *Dosa_MultiUpsert_Result) GetSuccess() (o *MultiUpsertResponse) {
	if v.Success != nil {
		return v.Success
	}

	return
}

// GetClientError returns the value of ClientError if it is set or its
// zero value if it is unset.
func (v *Dosa_MultiUpsert_Result) GetClientError() (o *BadRequestError) {
	if v.ClientError != nil {
		return v.ClientError
	}

	return
}

// GetServerError returns the value of ServerError if it is set or its
// zero value if it is unset.
func (v *Dosa_MultiUpsert_Result) GetServerError() (o *InternalServerError) {
	if v.ServerError != nil {
		return v.ServerError
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "multiUpsert" for this struct.
func (v *Dosa_MultiUpsert_Result) MethodName() string {
	return "multiUpsert"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Dosa_MultiUpsert_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
