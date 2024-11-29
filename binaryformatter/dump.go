package binaryformatter

import (
	"bytes"
	"fmt"
	"strings"
)

func Dump(datas []interface{}) string {
	buff := bytes.NewBufferString("")

	for _, data := range datas {
		switch ser := data.(type) {
		case SerializedStreamHeader:
			buff.WriteString("SerializedStreamHeader:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20RootId: %d\n", ser.RootId))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20HeaderId: %d\n", ser.HeaderId))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20MajorVersion: %d\n", ser.MajorVersion))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20MinorVersion: %d\n", ser.MinorVersion))
		case ClassWithId:
			buff.WriteString("ClassWithId:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20ObjectId: %d\n", ser.ObjectId))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20MetadataId: %d\n", ser.MetadataId))
		case SystemClassWithMembers:
			buff.WriteString("SystemClassWithMembers:\n")
			buff.WriteString("\x20\x20\x20\x20ClassInfo:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20ObjectId: %d\n", ser.ClassInfo.ObjectId))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20Name: %s\n", ser.ClassInfo.Name))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20MemberCount: %d\n", ser.ClassInfo.MemberCount))
			buff.WriteString("\x20\x20\x20\x20\x20\x20\x20\x20MemberNames: \n")
			for _, name := range ser.ClassInfo.MemberNames {
				buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20- %s\n", name))
			}
		case ClassWithMembers:
			buff.WriteString("ClassWithMembers:\n")
			buff.WriteString("\x20\x20\x20\x20ClassInfo:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20ObjectId: %d\n", ser.ClassInfo.ObjectId))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20Name: %s\n", ser.ClassInfo.Name))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20MemberCount: %d\n", ser.ClassInfo.MemberCount))
			buff.WriteString("\x20\x20\x20\x20\x20\x20\x20\x20MemberNames: \n")
			for _, name := range ser.ClassInfo.MemberNames {
				buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20- %s\n", name))
			}
		case SystemClassWithMembersAndTypes:
			buff.WriteString("SystemClassWithMembersAndTypes:\n")
			buff.WriteString("\x20\x20\x20\x20ClassInfo:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20ObjectId: %d\n", ser.ClassInfo.ObjectId))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20Name: %s\n", ser.ClassInfo.Name))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20MemberCount: %d\n", ser.ClassInfo.MemberCount))
			buff.WriteString("\x20\x20\x20\x20\x20\x20\x20\x20MemberNames: \n")
			for _, name := range ser.ClassInfo.MemberNames {
				buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20- %s\n", name))
			}
		case ClassWithMembersAndTypes:
			buff.WriteString("ClassWithMembersAndTypes:\n")
			buff.WriteString("\x20\x20\x20\x20ClassInfo:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20ObjectId: %d\n", ser.ClassInfo.ObjectId))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20Name: %s\n", ser.ClassInfo.Name))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20MemberCount: %d\n", ser.ClassInfo.MemberCount))
			buff.WriteString("\x20\x20\x20\x20\x20\x20\x20\x20MemberNames: \n")
			for _, name := range ser.ClassInfo.MemberNames {
				buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20- %s\n", name))
			}
		case BinaryObjectString:
			buff.WriteString("BinaryObjectString:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20ObjectId: %d\n", ser.ObjectId))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20Value: %s\n", strings.ReplaceAll(ser.Value, "\n", "\n\x20\x20\x20\x20       ")))
		case BinaryArray:
			buff.WriteString("BinaryArray:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20ObjectId: %d\n", ser.ObjectId))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20BinaryArrayTypeEnum: %d\n", ser.BinaryArrayTypeEnum))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20Rank: %d\n", ser.Rank))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20Lengths: %v\n", ser.Lengths))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20LowerBounds: %v\n", ser.LowerBounds))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20TypeEnum: %d\n", ser.TypeEnum))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20AdditionalTypeInfo: %v\n", ser.AdditionalTypeInfo))
		case MemberPrimitiveTyped:
			buff.WriteString("MemberPrimitiveTyped:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20PrimitiveTypeEnum: %d\n", ser.PrimitiveTypeEnum))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20Value: %v\n", ser.Value))
		case MemberReference:
			buff.WriteString("MemberReference:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20IdRef: %d\n", ser.IdRef))
		case ObjectNull:
			buff.WriteString("ObjectNull\n")
		case BinaryLibrary:
			buff.WriteString("BinaryLibrary:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20LibraryId: %d\n", ser.LibraryId))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20LibraryName: %s\n", ser.LibraryName))
		case ObjectNullMultiple256:
			buff.WriteString("ObjectNullMultiple256:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20NullCount: %d\n", ser.NullCount))
		case ObjectNullMultiple:
			buff.WriteString("ObjectNullMultiple:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20NullCount: %d\n", ser.NullCount))
		case ArraySinglePrimitive:
			buff.WriteString("ArraySinglePrimitive:\n")
			buff.WriteString("\x20\x20\x20\x20ArrayInfo:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20ObjectId: %d\n", ser.ArrayInfo.ObjectId))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20Length: %d\n", ser.ArrayInfo.Length))
			buff.WriteString("\x20\x20\x20\x20\x20\x20\x20\x20Values: \n")
			for _, value := range ser.Values {
				buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20- %v\n", value))
			}
		case ArraySingleObject:
			buff.WriteString("ArraySingleObject:\n")
			buff.WriteString("\x20\x20\x20\x20ArrayInfo:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20ObjectId: %d\n", ser.ArrayInfo.ObjectId))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20Length: %d\n", ser.ArrayInfo.Length))
		case ArraySingleString:
			buff.WriteString("ArraySingleString:\n")
			buff.WriteString("\x20\x20\x20\x20ArrayInfo:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20ObjectId: %d\n", ser.ArrayInfo.ObjectId))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20Length: %d\n", ser.ArrayInfo.Length))
		case MethodCall:
			buff.WriteString("MethodCall:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20MessageEnum: %d\n", ser.MessageEnum))
			buff.WriteString("\x20\x20\x20\x20MethodName:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20ObjectId: %d\n", ser.MethodName.PrimitiveTypeEnum))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20Value: %s\n", ser.MethodName.StringValue))
			buff.WriteString("\x20\x20\x20\x20TypeName:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20ObjectId: %d\n", ser.TypeName.PrimitiveTypeEnum))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20Value: %s\n", ser.TypeName.StringValue))
			buff.WriteString("\x20\x20\x20\x20CallContext:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20ObjectId: %d\n", ser.CallContext.PrimitiveTypeEnum))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20Value: %s\n", ser.CallContext.StringValue))
			buff.WriteString("\x20\x20\x20\x20Args:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20Length: %d\n", ser.Args.Length))
			buff.WriteString("\x20\x20\x20\x20\x20\x20\x20\x20ListOfValueWithCode: \n")
			for _, value := range ser.Args.ListOfValueWithCode {
				buff.WriteString("\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20- \n")
				buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20ObjectId: %d\n", value.PrimitiveTypeEnum))
				buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20Value: %s\n", value.Value))
			}
		case MethodReturn:
			buff.WriteString("MethodReturn:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20MessageEnum: %d\n", ser.MessageEnum))
			buff.WriteString("\x20\x20\x20\x20ReturnValue:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20ObjectId: %d\n", ser.ReturnValue.PrimitiveTypeEnum))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20Value: %s\n", ser.ReturnValue.Value))
			buff.WriteString("\x20\x20\x20\x20CallContext:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20ObjectId: %d\n", ser.CallContext.PrimitiveTypeEnum))
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20Value: %s\n", ser.CallContext.StringValue))
			buff.WriteString("\x20\x20\x20\x20Args:\n")
			buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20Length: %d\n", ser.Args.Length))
			buff.WriteString("\x20\x20\x20\x20\x20\x20\x20\x20ListOfValueWithCode: \n")
			for _, value := range ser.Args.ListOfValueWithCode {
				buff.WriteString("\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20- \n")
				buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20ObjectId: %d\n", value.PrimitiveTypeEnum))
				buff.WriteString(fmt.Sprintf("\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20Value: %s\n", value.Value))
			}
		case MessageEnd:
			buff.WriteString("MessageEnd")

		}
	}

	return buff.String()
}
