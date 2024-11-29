package binaryformatter

import "github.com/fairyming/binary_formatter/common"

// basic structs
type ClassInfo struct {
	ObjectId    int32
	Name        string
	MemberCount int32
	MemberNames []string
}

func newClassInfo(reader *common.DataReader) (ClassInfo, error) {
	var classInfo ClassInfo
	if objectId, err := reader.ReadU32Le(); err == nil {
		classInfo.ObjectId = int32(objectId)
	} else {
		return classInfo, err
	}
	if name, err := LengthPrefixedString(reader); err == nil {
		classInfo.Name = name
	} else {
		return classInfo, err
	}
	if memberCount, err := reader.ReadU32Le(); err == nil {
		classInfo.MemberCount = int32(memberCount)
	} else {
		return classInfo, err
	}
	classInfo.MemberNames = make([]string, classInfo.MemberCount)
	for i := 0; i < int(classInfo.MemberCount); i++ {
		if name, err := LengthPrefixedString(reader); err == nil {
			classInfo.MemberNames[i] = name
		} else {
			return classInfo, err
		}
	}
	return classInfo, nil
}

type MemberTypeInfo struct {
	BinaryTypeEnums []BinaryType
	AdditionalInfos []interface{}
}

func newMemberTypeInfo(reader *common.DataReader, memberCount int) (MemberTypeInfo, error) {
	var memberTypeInfo MemberTypeInfo
	memberTypeInfo.BinaryTypeEnums = make([]BinaryType, memberCount)
	// isHasPrimitive := false
	// isHasOther := false
	for i := 0; i < memberCount; i++ {
		data, err := reader.Read1()
		if err != nil {
			return memberTypeInfo, err
		}
		if binaryType, err := BinaryTypeFromUint8(data); err == nil {
			// if binaryType == BinaryType_Primitive || binaryType == BinaryType_PrimitiveArray {
			// 	isHasPrimitive = true
			// } else {
			// 	isHasOther = true
			// }
			memberTypeInfo.BinaryTypeEnums[i] = binaryType
		} else {
			return memberTypeInfo, err
		}
	}

	// fmt.Println(memberTypeInfo.BinaryTypeEnums)

	for i := 0; i < memberCount; i++ {
		// The AdditionalInfos sequence MUST NOT contain any item for the BinaryTypeEnum values of String, Object, ObjectArray, or StringArray.
		if memberTypeInfo.BinaryTypeEnums[i] != BinaryType_String && memberTypeInfo.BinaryTypeEnums[i] != BinaryType_Object && memberTypeInfo.BinaryTypeEnums[i] != BinaryType_ObjectArray && memberTypeInfo.BinaryTypeEnums[i] != BinaryType_StringArray {
			if additionalInfo, err := readAdditionalInfo(reader, memberTypeInfo.BinaryTypeEnums[i]); err == nil {
				memberTypeInfo.AdditionalInfos = append(memberTypeInfo.AdditionalInfos, additionalInfo)
			} else {
				return memberTypeInfo, err
			}
		} else {
			memberTypeInfo.AdditionalInfos = append(memberTypeInfo.AdditionalInfos, nil)
		}

	}
	return memberTypeInfo, nil
}

type ArrayInfo struct {
	ObjectId int32
	Length   int32
}

func newArrayInfo(reader *common.DataReader) (ArrayInfo, error) {
	var arrayInfo ArrayInfo
	if objectId, err := reader.ReadU32Le(); err == nil {
		arrayInfo.ObjectId = int32(objectId)
	} else {
		return arrayInfo, err
	}
	if length, err := reader.ReadU32Le(); err == nil {
		arrayInfo.Length = int32(length)
	} else {
		return arrayInfo, err
	}
	return arrayInfo, nil
}

type StringValueWithCode struct {
	PrimitiveTypeEnum PrimitiveType
	StringValue       string
}

func newStringValueWithCode(reader *common.DataReader) (StringValueWithCode, error) {
	var stringValueWithCode StringValueWithCode
	if data, err := reader.Read1(); err == nil {
		if primitiveType, err := PrimitiveTypeFromUint8(data); err == nil {
			stringValueWithCode.PrimitiveTypeEnum = primitiveType
		} else {
			return stringValueWithCode, err
		}
	} else {
		return stringValueWithCode, err
	}
	if stringValue, err := LengthPrefixedString(reader); err == nil {
		stringValueWithCode.StringValue = stringValue
	} else {
		return stringValueWithCode, err
	}
	return stringValueWithCode, nil
}

type ValueWithCode struct {
	PrimitiveTypeEnum PrimitiveType
	Value             interface{}
}

func newValueWithCode(reader *common.DataReader) (ValueWithCode, error) {
	var valueWithCode ValueWithCode
	if data, err := reader.Read1(); err == nil {
		if primitiveType, err := PrimitiveTypeFromUint8(data); err == nil {
			valueWithCode.PrimitiveTypeEnum = primitiveType
		} else {
			return valueWithCode, err
		}
	} else {
		return valueWithCode, err
	}

	if value, err := readPrimitive(reader, valueWithCode.PrimitiveTypeEnum); err == nil {
		valueWithCode.Value = value
	} else {
		return valueWithCode, err
	}
	return valueWithCode, nil
}

type ArrayOfValueWithCode struct {
	Length              int32
	ListOfValueWithCode []ValueWithCode
}

func newArrayOfValueWithCode(reader *common.DataReader) (ArrayOfValueWithCode, error) {
	var arrayOfValueWithCode ArrayOfValueWithCode
	if length, err := reader.ReadU32Le(); err == nil {
		arrayOfValueWithCode.Length = int32(length)
	} else {
		return arrayOfValueWithCode, err
	}
	for i := 0; i < int(arrayOfValueWithCode.Length); i++ {
		if valueWithCode, err := newValueWithCode(reader); err == nil {
			arrayOfValueWithCode.ListOfValueWithCode = append(arrayOfValueWithCode.ListOfValueWithCode, valueWithCode)
		} else {
			return arrayOfValueWithCode, err
		}
	}
	return arrayOfValueWithCode, nil
}

type ClassTypeInfo struct {
	TypeName  string
	LibraryId int32
}

func newClassTypeInfo(reader *common.DataReader) (ClassTypeInfo, error) {
	var classTypeInfo ClassTypeInfo
	if typeName, err := LengthPrefixedString(reader); err == nil {
		classTypeInfo.TypeName = typeName
	} else {
		return classTypeInfo, err
	}
	if libraryId, err := reader.ReadU32Le(); err == nil {
		classTypeInfo.LibraryId = int32(libraryId)
	} else {
		return classTypeInfo, err
	}
	return classTypeInfo, nil
}
