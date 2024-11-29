package binaryformatter

import "github.com/fairyming/binary_formatter/common"

// RecordType structs
type SerializedStreamHeader struct {
	RootId       int32
	HeaderId     int32
	MajorVersion int32
	MinorVersion int32
}

func newSerializedStreamHeader(reader *common.DataReader) (SerializedStreamHeader, error) {
	var header SerializedStreamHeader
	if rootId, err := reader.ReadU32Le(); err == nil {
		header.RootId = int32(rootId)
	} else {
		return header, err
	}
	if headerId, err := reader.ReadU32Le(); err == nil {
		header.HeaderId = int32(headerId)
	} else {
		return header, err
	}
	if majorVersion, err := reader.ReadU32Le(); err == nil {
		header.MajorVersion = int32(majorVersion)
	} else {
		return header, err
	}
	if minorVersion, err := reader.ReadU32Le(); err == nil {
		header.MinorVersion = int32(minorVersion)
	} else {
		return header, err
	}
	return header, nil
}

type ClassWithId struct {
	ObjectId   int32
	MetadataId int32
}

func newClassWithId(reader *common.DataReader) (ClassWithId, error) {
	var classWithId ClassWithId
	if objectId, err := reader.ReadU32Be(); err == nil {
		classWithId.ObjectId = int32(objectId)
	} else {
		return classWithId, err
	}
	if metadataId, err := reader.ReadU32Be(); err == nil {
		classWithId.MetadataId = int32(metadataId)
	} else {
		return classWithId, err
	}
	return classWithId, nil
}

type SystemClassWithMembers struct {
	ClassInfo ClassInfo
}

func newSystemClassWithMembers(reader *common.DataReader) (SystemClassWithMembers, error) {
	var systemClassWithMembers SystemClassWithMembers
	if classInfo, err := newClassInfo(reader); err == nil {
		systemClassWithMembers.ClassInfo = classInfo
	} else {
		return systemClassWithMembers, err
	}
	return systemClassWithMembers, nil
}

type ClassWithMembers struct {
	ClassInfo ClassInfo
	LibraryId int32
}

func newClassWithMembers(reader *common.DataReader) (ClassWithMembers, error) {
	var classWithMembers ClassWithMembers
	if classInfo, err := newClassInfo(reader); err == nil {
		classWithMembers.ClassInfo = classInfo
	} else {
		return classWithMembers, err
	}
	if libraryId, err := reader.ReadU32Le(); err == nil {
		classWithMembers.LibraryId = int32(libraryId)
	} else {
		return classWithMembers, err
	}
	return classWithMembers, nil
}

type SystemClassWithMembersAndTypes struct {
	ClassInfo      ClassInfo
	MemberTypeInfo MemberTypeInfo
}

func newSystemClassWithMembersAndTypes(reader *common.DataReader) (SystemClassWithMembersAndTypes, error) {
	var systemClassWithMembersAndTypes SystemClassWithMembersAndTypes
	if classInfo, err := newClassInfo(reader); err == nil {
		systemClassWithMembersAndTypes.ClassInfo = classInfo
	} else {
		return systemClassWithMembersAndTypes, err
	}
	if memberTypeInfo, err := newMemberTypeInfo(reader, int(systemClassWithMembersAndTypes.ClassInfo.MemberCount)); err == nil {
		systemClassWithMembersAndTypes.MemberTypeInfo = memberTypeInfo
	} else {
		return systemClassWithMembersAndTypes, err
	}
	return systemClassWithMembersAndTypes, nil
}

type ClassWithMembersAndTypes struct {
	ClassInfo      ClassInfo
	MemberTypeInfo MemberTypeInfo
	LibraryId      int32
}

func newClassWithMembersAndTypes(reader *common.DataReader) (ClassWithMembersAndTypes, error) {
	var classWithMembersAndTypes ClassWithMembersAndTypes
	if classInfo, err := newClassInfo(reader); err == nil {
		classWithMembersAndTypes.ClassInfo = classInfo
	} else {
		return classWithMembersAndTypes, err
	}
	if memberTypeInfo, err := newMemberTypeInfo(reader, int(classWithMembersAndTypes.ClassInfo.MemberCount)); err == nil {
		classWithMembersAndTypes.MemberTypeInfo = memberTypeInfo
	} else {
		return classWithMembersAndTypes, err
	}
	if libraryId, err := reader.ReadU32Le(); err == nil {
		classWithMembersAndTypes.LibraryId = int32(libraryId)
	} else {
		return classWithMembersAndTypes, err
	}
	return classWithMembersAndTypes, nil
}

type BinaryObjectString struct {
	ObjectId int32
	Value    string
}

func newBinaryObjectString(reader *common.DataReader) (BinaryObjectString, error) {
	var binaryObjectString BinaryObjectString
	if objectId, err := reader.ReadU32Le(); err == nil {
		binaryObjectString.ObjectId = int32(objectId)
	} else {
		return binaryObjectString, err
	}
	if value, err := LengthPrefixedString(reader); err == nil {
		binaryObjectString.Value = value
	} else {
		return binaryObjectString, err
	}
	return binaryObjectString, nil
}

type BinaryArray struct {
	ObjectId            int32
	BinaryArrayTypeEnum BinaryArrayType
	Rank                int32
	Lengths             []int32
	LowerBounds         []int32
	TypeEnum            BinaryType
	AdditionalTypeInfo  interface{}
}

func newBinaryArray(reader *common.DataReader) (BinaryArray, error) {
	var binaryArray BinaryArray
	if objectId, err := reader.ReadU32Le(); err == nil {
		binaryArray.ObjectId = int32(objectId)
	} else {
		return binaryArray, err
	}
	if data, err := reader.Read1(); err == nil {
		if binaryArrayType, err := BinaryArrayTypeFromUint8(data); err == nil {
			binaryArray.BinaryArrayTypeEnum = binaryArrayType
		} else {
			return binaryArray, err
		}
	} else {
		return binaryArray, err
	}
	if rank, err := reader.ReadU32Le(); err == nil {
		binaryArray.Rank = int32(rank)
	} else {
		return binaryArray, err
	}

	for i := 0; i < int(binaryArray.Rank); i++ {
		if length, err := reader.ReadU32Le(); err == nil {
			binaryArray.Lengths = append(binaryArray.Lengths, int32(length))
		} else {
			return binaryArray, err
		}
	}
	// 2.4.3.1 BinaryArray
	// if the value of the BinaryArrayTypeEnum field is SingleOffset, JaggedOffset, or RectangularOffset, this field MUST be present in the serialization stream; otherwise, this field MUST NOT be present in the serialization stream.
	if binaryArray.BinaryArrayTypeEnum == BinaryArrayType_SingleOffset || binaryArray.BinaryArrayTypeEnum == BinaryArrayType_JaggedOffset || binaryArray.BinaryArrayTypeEnum == BinaryArrayType_RectangularOffset {
		for i := 0; i < int(binaryArray.Rank); i++ {
			if lowerBound, err := reader.ReadU32Le(); err == nil {
				binaryArray.LowerBounds = append(binaryArray.LowerBounds, int32(lowerBound))
			} else {
				return binaryArray, err
			}
		}
	}

	if data, err := reader.Read1(); err == nil {
		if binaryType, err := BinaryTypeFromUint8(data); err == nil {
			binaryArray.TypeEnum = binaryType
		} else {
			return binaryArray, err
		}
	} else {
		return binaryArray, err
	}

	if additionalTypeInfo, err := readAdditionalInfo(reader, binaryArray.TypeEnum); err == nil {
		binaryArray.AdditionalTypeInfo = additionalTypeInfo
	} else {
		return binaryArray, err
	}

	return binaryArray, nil
}

type MemberPrimitiveTyped struct {
	PrimitiveTypeEnum PrimitiveType
	Value             interface{}
}

func newMemberPrimitiveTyped(reader *common.DataReader) (MemberPrimitiveTyped, error) {
	var memberPrimitiveTyped MemberPrimitiveTyped
	if data, err := reader.Read1(); err == nil {
		if primitiveType, err := PrimitiveTypeFromUint8(data); err == nil {
			memberPrimitiveTyped.PrimitiveTypeEnum = primitiveType
		} else {
			return memberPrimitiveTyped, err
		}
	} else {
		return memberPrimitiveTyped, err
	}

	if value, err := readPrimitive(reader, memberPrimitiveTyped.PrimitiveTypeEnum); err == nil {
		memberPrimitiveTyped.Value = value
	} else {
		return memberPrimitiveTyped, err
	}
	return memberPrimitiveTyped, nil
}

type MemberReference struct {
	IdRef int32
}

func newMemberReference(reader *common.DataReader) (MemberReference, error) {
	var memberReference MemberReference
	if idRef, err := reader.ReadU32Le(); err == nil {
		memberReference.IdRef = int32(idRef)
	} else {
		return memberReference, err
	}
	return memberReference, nil
}

type ObjectNull struct {
}

func newObjectNull(reader *common.DataReader) (ObjectNull, error) {
	return ObjectNull{}, nil
}

type MessageEnd struct {
}

func newMessageEnd(reader *common.DataReader) (MessageEnd, error) {
	return MessageEnd{}, nil
}

type BinaryLibrary struct {
	LibraryId   int32
	LibraryName string
}

func newBinaryLibrary(reader *common.DataReader) (BinaryLibrary, error) {
	var binaryLibrary BinaryLibrary
	if libraryId, err := reader.ReadU32Le(); err == nil {
		binaryLibrary.LibraryId = int32(libraryId)
	} else {
		return binaryLibrary, err
	}
	if libraryName, err := LengthPrefixedString(reader); err == nil {
		binaryLibrary.LibraryName = libraryName
	} else {
		return binaryLibrary, err
	}
	return binaryLibrary, nil
}

type ObjectNullMultiple256 struct {
	NullCount int8
}

func newObjectNullMultiple256(reader *common.DataReader) (ObjectNullMultiple256, error) {
	var objectNullMultiple256 ObjectNullMultiple256
	if nullCount, err := reader.Read1(); err == nil {
		objectNullMultiple256.NullCount = int8(nullCount)
	} else {
		return objectNullMultiple256, err
	}
	return objectNullMultiple256, nil
}

type ObjectNullMultiple struct {
	NullCount int32
}

func newObjectNullMultiple(reader *common.DataReader) (ObjectNullMultiple, error) {
	var objectNullMultiple ObjectNullMultiple
	if nullCount, err := reader.ReadU32Le(); err == nil {
		objectNullMultiple.NullCount = int32(nullCount)
	} else {
		return objectNullMultiple, err
	}
	return objectNullMultiple, nil
}

type ArraySinglePrimitive struct {
	ArrayInfo         ArrayInfo
	PrimitiveTypeEnum PrimitiveType
	Values            []interface{}
}

func newArraySinglePrimitive(reader *common.DataReader) (ArraySinglePrimitive, error) {
	var arraySinglePrimitive ArraySinglePrimitive
	if arrayInfo, err := newArrayInfo(reader); err == nil {
		arraySinglePrimitive.ArrayInfo = arrayInfo
	} else {
		return arraySinglePrimitive, err
	}
	if data, err := reader.Read1(); err == nil {
		if primitiveType, err := PrimitiveTypeFromUint8(data); err == nil {
			arraySinglePrimitive.PrimitiveTypeEnum = primitiveType
		} else {
			return arraySinglePrimitive, err
		}
	} else {
		return arraySinglePrimitive, err
	}
	for i := 0; i < int(arraySinglePrimitive.ArrayInfo.Length); i++ {
		if value, err := readPrimitive(reader, arraySinglePrimitive.PrimitiveTypeEnum); err == nil {
			arraySinglePrimitive.Values = append(arraySinglePrimitive.Values, value)
		} else {
			return arraySinglePrimitive, err
		}
	}
	return arraySinglePrimitive, nil
}

type ArraySingleObject struct {
	ArrayInfo ArrayInfo
}

func newArraySingleObject(reader *common.DataReader) (ArraySingleObject, error) {
	var arraySingleObject ArraySingleObject
	if arrayInfo, err := newArrayInfo(reader); err == nil {
		arraySingleObject.ArrayInfo = arrayInfo
	} else {
		return arraySingleObject, err
	}
	return arraySingleObject, nil
}

type ArraySingleString struct {
	ArrayInfo ArrayInfo
}

func newArraySingleString(reader *common.DataReader) (ArraySingleObject, error) {
	var arraySingleObject ArraySingleObject
	if arrayInfo, err := newArrayInfo(reader); err == nil {
		arraySingleObject.ArrayInfo = arrayInfo
	} else {
		return arraySingleObject, err
	}
	return arraySingleObject, nil
}

type MethodCall struct {
	MessageEnum int32
	MethodName  StringValueWithCode
	TypeName    StringValueWithCode
	CallContext StringValueWithCode
	Args        ArrayOfValueWithCode
}

func newMethodCall(reader *common.DataReader) (MethodCall, error) {
	var methodCall MethodCall
	if messageEnum, err := reader.ReadU32Le(); err == nil {
		methodCall.MessageEnum = int32(messageEnum)
	} else {
		return methodCall, err
	}

	if methodName, err := newStringValueWithCode(reader); err == nil {
		methodCall.MethodName = methodName
	} else {
		return methodCall, err
	}
	if typeName, err := newStringValueWithCode(reader); err == nil {
		methodCall.TypeName = typeName
	} else {
		return methodCall, err
	}
	// A StringValueWithCode that represents the Logical Call ID. This field is
	// conditional. If the MessageEnum field has the ContextInline bit set, the field MUST be present;
	// otherwise, the field MUST NOT be present
	if int(methodCall.MessageEnum)&int(MessageFlags_ContextInline) != 0 {
		if callContext, err := newStringValueWithCode(reader); err == nil {
			methodCall.CallContext = callContext
		} else {
			return methodCall, err
		}
	}
	// An ArrayOfValueWithCode that contains the Output Arguments of the method.
	// This field is conditional. If the MessageEnum field has the ArgsInline bit set, the field MUST be
	// present; otherwise, the field MUST NOT be present.
	if int(methodCall.MessageEnum)&int(MessageFlags_ArgsInline) != 0 {
		if args, err := newArrayOfValueWithCode(reader); err == nil {
			methodCall.Args = args
		} else {
			return methodCall, err
		}
	}
	return methodCall, nil
}

type MethodReturn struct {
	MessageEnum int32
	ReturnValue ValueWithCode
	CallContext StringValueWithCode
	Args        ArrayOfValueWithCode
}

func newMethodReturn(reader *common.DataReader) (MethodReturn, error) {
	var methodReturn MethodReturn
	if messageEnum, err := reader.ReadU32Le(); err == nil {
		methodReturn.MessageEnum = int32(messageEnum)
	} else {
		return methodReturn, err
	}

	// A ValueWithCode that contains the Return Value of a Remote Method. If
	// the MessageEnum field has the ReturnValueInline bit set, this field MUST be present; otherwise,
	// this field MUST NOT be present.
	if int(methodReturn.MessageEnum)&int(MessageFlags_ReturnValueInline) != 0 {
		if returnValue, err := newValueWithCode(reader); err == nil {
			methodReturn.ReturnValue = returnValue
		} else {
			return methodReturn, err
		}
	}

	// A StringValueWithCode that represents the Logical Call ID. This field is
	// conditional. If the MessageEnum field has the ContextInline bit set, the field MUST be present;
	// otherwise, the field MUST NOT be present.
	if int(methodReturn.MessageEnum)&int(MessageFlags_ContextInline) != 0 {
		if callContext, err := newStringValueWithCode(reader); err == nil {
			methodReturn.CallContext = callContext
		} else {
			return methodReturn, err
		}
	}

	// An ArrayOfValueWithCode that contains the Output Arguments of the method.
	// This field is conditional. If the MessageEnum field has the ArgsInline bit set, the field MUST be
	// present; otherwise, the field MUST NOT be present
	if int(methodReturn.MessageEnum)&int(MessageFlags_ArgsInline) != 0 {
		if args, err := newArrayOfValueWithCode(reader); err == nil {
			methodReturn.Args = args
		} else {
			return methodReturn, err
		}
	}

	return methodReturn, nil
}
