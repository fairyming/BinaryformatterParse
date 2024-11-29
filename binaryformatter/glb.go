package binaryformatter

type MessageFlags uint32

const (
	MessageFlags_NoArgs                 MessageFlags = 0x00000001
	MessageFlags_ArgsInline             MessageFlags = 0x00000002
	MessageFlags_ArgsIsArray            MessageFlags = 0x00000004
	MessageFlags_ArgsInArray            MessageFlags = 0x00000008
	MessageFlags_NoContext              MessageFlags = 0x00000010
	MessageFlags_ContextInline          MessageFlags = 0x00000020
	MessageFlags_ContextInArray         MessageFlags = 0x00000040
	MessageFlags_MethodSignatureInArray MessageFlags = 0x00000080
	MessageFlags_PropertiesInArray      MessageFlags = 0x00000100
	MessageFlags_NoReturnValue          MessageFlags = 0x00000200
	MessageFlags_ReturnValueVoid        MessageFlags = 0x00000400
	MessageFlags_ReturnValueInline      MessageFlags = 0x00000800
	MessageFlags_ReturnValueInArray     MessageFlags = 0x00001000
	MessageFlags_ExceptionInArray       MessageFlags = 0x00002000
	MessageFlags_GenericMethod          MessageFlags = 0x00008000
)

type RecordType uint8

const (
	RecordType_SerializedStreamHeader         RecordType = 0x0
	RecordType_ClassWithId                    RecordType = 0x1
	RecordType_SystemClassWithMembers         RecordType = 0x2
	RecordType_ClassWithMembers               RecordType = 0x3
	RecordType_SystemClassWithMembersAndTypes RecordType = 0x4
	RecordType_ClassWithMembersAndTypes       RecordType = 0x5
	RecordType_BinaryObjectString             RecordType = 0x6
	RecordType_BinaryArray                    RecordType = 0x7
	RecordType_MemberPrimitiveTyped           RecordType = 0x8
	RecordType_MemberReference                RecordType = 0x9
	RecordType_ObjectNull                     RecordType = 0x0a
	RecordType_MessageEnd                     RecordType = 0x0b
	RecordType_BinaryLibrary                  RecordType = 0x0c
	RecordType_ObjectNullMultiple256          RecordType = 0x0d
	RecordType_ObjectNullMultiple             RecordType = 0x0e
	RecordType_ArraySinglePrimitive           RecordType = 0x0f
	RecordType_ArraySingleObject              RecordType = 0x10
	RecordType_ArraySingleString              RecordType = 0x11
	RecordType_MethodCall                     RecordType = 0x15
	RecordType_MethodReturn                   RecordType = 0x16
)

type BinaryType uint8

const (
	BinaryType_Primitive      = 0x0
	BinaryType_String         = 0x1
	BinaryType_Object         = 0x2
	BinaryType_SystemClass    = 0x3
	BinaryType_Class          = 0x4
	BinaryType_ObjectArray    = 0x5
	BinaryType_StringArray    = 0x6
	BinaryType_PrimitiveArray = 0x7
)

type PrimitiveType uint8

const (
	PrimitiveType_Boolean  PrimitiveType = 0x1
	PrimitiveType_Byte     PrimitiveType = 0x2
	PrimitiveType_Char     PrimitiveType = 0x3
	PrimitiveType_Decimal  PrimitiveType = 0x5
	PrimitiveType_Double   PrimitiveType = 0x6
	PrimitiveType_Int16    PrimitiveType = 0x7
	PrimitiveType_Int32    PrimitiveType = 0x8
	PrimitiveType_Int64    PrimitiveType = 0x9
	PrimitiveType_SByte    PrimitiveType = 0x0a
	PrimitiveType_Single   PrimitiveType = 0x0b
	PrimitiveType_TimeSpan PrimitiveType = 0x0c
	PrimitiveType_DateTime PrimitiveType = 0x0d
	PrimitiveType_Uint16   PrimitiveType = 0x0e
	PrimitiveType_Uint32   PrimitiveType = 0x0f
	PrimitiveType_Uint64   PrimitiveType = 0x10
	PrimitiveType_Null     PrimitiveType = 0x11
	PrimitiveType_String   PrimitiveType = 0x12
)

type BinaryArrayType uint8

const (
	BinaryArrayType_Single            BinaryArrayType = 0x0
	BinaryArrayType_Jagged            BinaryArrayType = 0x1
	BinaryArrayType_Rectangular       BinaryArrayType = 0x2
	BinaryArrayType_SingleOffset      BinaryArrayType = 0x3
	BinaryArrayType_JaggedOffset      BinaryArrayType = 0x4
	BinaryArrayType_RectangularOffset BinaryArrayType = 0x5
)
