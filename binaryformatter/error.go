package binaryformatter

import "errors"

var ErrUnknownBinaryType = errors.New("unknown binary type")

var ErrUnknownPrimitiveType = errors.New("unknown primitive type")

var ErrUnknownBinaryArrayType = errors.New("unknown binary array type")

var ErrUnknownRecordType = errors.New("unknown record type")

var ErrUnknownMessageFlags = errors.New("unknown message flags")
