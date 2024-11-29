package binaryformatter

import "github.com/fairyming/binary_formatter/common"

func Parse(parseData []byte) ([]interface{}, error) {
	reader, err := common.NewDataReader(parseData)
	if err != nil {
		return nil, err
	}
	parseResult := make([]interface{}, 0)
	for {
		data, err := reader.Read1()
		// fmt.Println(data)
		if err != nil {
			return nil, err
		}
		recordType, err := RecordTypeFromUint8(data)
		if err != nil {
			return nil, err
		}

		if parseValue, err := readRecordType(reader, recordType); err == nil {
			parseResult = append(parseResult, parseValue)
			// r, _ := json.Marshal(parseValue)
			// fmt.Println(string(r))
		} else {
			// r, _ := json.Marshal(parseValue)
			// fmt.Println("err", string(r))
			return nil, err
		}
		// message end
		if recordType == RecordType_MessageEnd {
			break
		}
		if reader.IsEof() {
			break
		}
	}
	return parseResult, nil
}

func readRecordType(reader *common.DataReader, recordType RecordType) (interface{}, error) {
	switch recordType {
	case RecordType_SerializedStreamHeader:
		return newSerializedStreamHeader(reader)
	case RecordType_ClassWithId:
		return newClassWithId(reader)
	case RecordType_SystemClassWithMembers:
		return newSystemClassWithMembers(reader)
	case RecordType_ClassWithMembers:
		return newClassWithMembers(reader)
	case RecordType_SystemClassWithMembersAndTypes:
		return newSystemClassWithMembersAndTypes(reader)
	case RecordType_ClassWithMembersAndTypes:
		return newClassWithMembersAndTypes(reader)
	case RecordType_BinaryObjectString:
		return newBinaryObjectString(reader)
	case RecordType_BinaryArray:
		return newBinaryArray(reader)
	case RecordType_MemberPrimitiveTyped:
		return newMemberPrimitiveTyped(reader)
	case RecordType_MemberReference:
		return newMemberReference(reader)
	case RecordType_ObjectNull:
		return newObjectNull(reader)
	case RecordType_MessageEnd:
		return newMessageEnd(reader)
	case RecordType_BinaryLibrary:
		return newBinaryLibrary(reader)
	case RecordType_ObjectNullMultiple256:
		return newObjectNullMultiple256(reader)
	case RecordType_ObjectNullMultiple:
		return newObjectNullMultiple(reader)
	case RecordType_ArraySinglePrimitive:
		return newArraySinglePrimitive(reader)
	case RecordType_ArraySingleObject:
		return newArraySingleObject(reader)
	case RecordType_ArraySingleString:
		return newArraySingleString(reader)
	case RecordType_MethodCall:
		return newMethodCall(reader)
	case RecordType_MethodReturn:
		return newMethodReturn(reader)
	}
	return nil, ErrUnknownRecordType
}
