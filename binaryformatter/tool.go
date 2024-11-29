package binaryformatter

import (
	"errors"

	"github.com/fairyming/binary_formatter/common"
)

func LengthPrefixedString(reader *common.DataReader) (string, error) {
	length := 0
	shift := 0
	c := true
	for c {
		byte, err := reader.Read1()
		if err != nil {
			return "", err
		}
		if byte&128 != 0 {
			byte ^= 128
		} else {
			c = false
		}
		length += int(byte) << shift
		shift += 7
	}
	return reader.ReadString(length)
}

func MessageFlagsFromUint32(value uint32) (MessageFlags, error) {
	switch value {
	case 0x00000001:
		return MessageFlags_NoArgs, nil
	case 0x00000002:
		return MessageFlags_ArgsInline, nil
	case 0x00000004:
		return MessageFlags_ArgsIsArray, nil
	case 0x00000008:
		return MessageFlags_ArgsInArray, nil
	case 0x00000010:
		return MessageFlags_NoContext, nil
	case 0x00000020:
		return MessageFlags_ContextInline, nil
	case 0x00000040:
		return MessageFlags_ContextInArray, nil
	case 0x00000080:
		return MessageFlags_MethodSignatureInArray, nil
	case 0x00000100:
		return MessageFlags_PropertiesInArray, nil
	case 0x00000200:
		return MessageFlags_NoReturnValue, nil
	case 0x00000400:
		return MessageFlags_ReturnValueVoid, nil
	case 0x00000800:
		return MessageFlags_ReturnValueInline, nil
	case 0x00001000:
		return MessageFlags_ReturnValueInArray, nil
	case 0x00002000:
		return MessageFlags_ExceptionInArray, nil
	case 0x00008000:
		return MessageFlags_GenericMethod, nil
	default:
		return 0, ErrUnknownMessageFlags
	}
}

func BinaryTypeFromUint8(value uint8) (BinaryType, error) {
	switch value {
	case 0x0:
		return BinaryType_Primitive, nil
	case 0x1:
		return BinaryType_String, nil
	case 0x2:
		return BinaryType_Object, nil
	case 0x3:
		return BinaryType_SystemClass, nil
	case 0x4:
		return BinaryType_Class, nil
	case 0x5:
		return BinaryType_ObjectArray, nil
	case 0x6:
		return BinaryType_StringArray, nil
	case 0x7:
		return BinaryType_PrimitiveArray, nil
	default:
		return 0, ErrUnknownBinaryType
	}
}

func PrimitiveTypeFromUint8(value uint8) (PrimitiveType, error) {
	switch value {
	case 0x1:
		return PrimitiveType_Boolean, nil
	case 0x2:
		return PrimitiveType_Byte, nil
	case 0x3:
		return PrimitiveType_Char, nil
	case 0x5:
		return PrimitiveType_Decimal, nil
	case 0x6:
		return PrimitiveType_Double, nil
	case 0x7:
		return PrimitiveType_Int16, nil
	case 0x8:
		return PrimitiveType_Int32, nil
	case 0x9:
		return PrimitiveType_Int64, nil
	case 0x0a:
		return PrimitiveType_SByte, nil
	case 0x0b:
		return PrimitiveType_Single, nil
	case 0x0c:
		return PrimitiveType_TimeSpan, nil
	case 0x0d:
		return PrimitiveType_DateTime, nil
	case 0x0e:
		return PrimitiveType_Uint16, nil
	case 0x0f:
		return PrimitiveType_Uint32, nil
	case 0x10:
		return PrimitiveType_Uint64, nil
	case 0x11:
		return PrimitiveType_Null, nil
	case 0x12:
		return PrimitiveType_String, nil
	default:
		return 0, errors.New("invalid value for PrimitiveType")
	}
}

func BinaryArrayTypeFromUint8(value uint8) (BinaryArrayType, error) {
	switch value {
	case 0x0:
		return BinaryArrayType_Single, nil
	case 0x1:
		return BinaryArrayType_Jagged, nil
	case 0x2:
		return BinaryArrayType_Rectangular, nil
	case 0x3:
		return BinaryArrayType_SingleOffset, nil
	case 0x4:
		return BinaryArrayType_JaggedOffset, nil
	case 0x5:
		return BinaryArrayType_RectangularOffset, nil
	default:
		return 0, ErrUnknownBinaryArrayType
	}
}

func RecordTypeFromUint8(value uint8) (RecordType, error) {
	switch value {
	case 0x0:
		return RecordType_SerializedStreamHeader, nil
	case 0x1:
		return RecordType_ClassWithId, nil
	case 0x2:
		return RecordType_SystemClassWithMembers, nil
	case 0x3:
		return RecordType_ClassWithMembers, nil
	case 0x4:
		return RecordType_SystemClassWithMembersAndTypes, nil
	case 0x5:
		return RecordType_ClassWithMembersAndTypes, nil
	case 0x6:
		return RecordType_BinaryObjectString, nil
	case 0x7:
		return RecordType_BinaryArray, nil
	case 0x8:
		return RecordType_MemberPrimitiveTyped, nil
	case 0x9:
		return RecordType_MemberReference, nil
	case 0x0a:
		return RecordType_ObjectNull, nil
	case 0x0b:
		return RecordType_MessageEnd, nil
	case 0x0c:
		return RecordType_BinaryLibrary, nil
	case 0x0d:
		return RecordType_ObjectNullMultiple256, nil
	case 0x0e:
		return RecordType_ObjectNullMultiple, nil
	case 0x0f:
		return RecordType_ArraySinglePrimitive, nil
	case 0x10:
		return RecordType_ArraySingleObject, nil
	case 0x11:
		return RecordType_ArraySingleString, nil
	case 0x15:
		return RecordType_MethodCall, nil
	case 0x16:
		return RecordType_MethodReturn, nil
	default:
		return 0, ErrUnknownRecordType
	}
}

func readAdditionalInfo(reader *common.DataReader, binaryType BinaryType) (interface{}, error) {
	switch binaryType {
	case BinaryType_Primitive, BinaryType_PrimitiveArray:
		data, err := reader.Read1()
		if err != nil {
			return nil, err
		}
		if primitiveType, err := PrimitiveTypeFromUint8(data); err == nil {
			// if data, err := readPrimitive(reader, primitiveType); err == nil {
			// 	return data, nil
			// } else {
			// 	return nil, err
			// }
			return primitiveType, nil
		} else {
			return nil, err
		}
	case BinaryType_String, BinaryType_Object, BinaryType_ObjectArray, BinaryType_StringArray:
		return nil, nil
	case BinaryType_SystemClass:
		if systemClass, err := LengthPrefixedString(reader); err == nil {
			return systemClass, nil
		} else {
			return nil, err
		}
	case BinaryType_Class:
		if classInfo, err := newClassTypeInfo(reader); err == nil {
			return classInfo, nil
		} else {
			return nil, err
		}
	default:
		return nil, ErrUnknownBinaryType
	}
}

func readPrimitive(reader *common.DataReader, primitiveType PrimitiveType) (interface{}, error) {
	switch primitiveType {
	case PrimitiveType_Boolean, PrimitiveType_Byte, PrimitiveType_Char, PrimitiveType_SByte, PrimitiveType_Null:
		if data, err := reader.Read1(); err == nil {
			return data, nil
		} else {
			return nil, err
		}
	case PrimitiveType_Decimal:
		return LengthPrefixedString(reader)
	case PrimitiveType_Double, PrimitiveType_Int64, PrimitiveType_Uint64:
		if data, err := reader.Read(8); err == nil {
			return data, nil
		} else {
			return nil, err
		}
	case PrimitiveType_Int16, PrimitiveType_Uint16:
		if data, err := reader.ReadU16Be(); err == nil {
			return data, nil
		} else {
			return nil, err
		}
	case PrimitiveType_Int32, PrimitiveType_Uint32:
		if data, err := reader.ReadU32Be(); err == nil {
			return data, nil
		} else {
			return nil, err
		}
	case PrimitiveType_Single:
		if data, err := reader.Read(4); err == nil {
			return data, nil
		} else {
			return nil, err
		}
	case PrimitiveType_TimeSpan, PrimitiveType_DateTime:
		if data, err := reader.Read(8); err == nil {
			return data, nil
		} else {
			return nil, err
		}
	case PrimitiveType_String:
		return LengthPrefixedString(reader)
	default:
		return nil, ErrUnknownPrimitiveType
	}
}
