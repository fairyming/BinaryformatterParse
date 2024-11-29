# Binaryformatter 解析

Binaryformatter 解析初版实现。

使用
```
go run main.go --path example/MS-NRBF_example.ser
```

解析结果
```
SerializedStreamHeader:
    RootId: 1
    HeaderId: -1
    MajorVersion: 1
    MinorVersion: 0
MethodCall:
    MessageEnum: 20
    MethodName:
        ObjectId: 18
        Value: SendAddress
    TypeName:
        ObjectId: 18
        Value: DOJRemotingMetadata.MyServer, DOJRemotingMetadata, Version=1.0.2622.31326, Culture=neutral, PublicKeyToken=null
    CallContext:
        ObjectId: 0
        Value:
    Args:
        Length: 0
        ListOfValueWithCode:
ArraySingleObject:
    ArrayInfo:
        ObjectId: 1
        Length: 1
MemberReference:
    IdRef: 2
BinaryLibrary:
    LibraryId: 3
    LibraryName: DOJRemotingMetadata, Version=1.0.2622.31326, Culture=neutral, PublicKeyToken=null
ClassWithMembersAndTypes:
    ClassInfo:
        ObjectId: 2
        Name: DOJRemotingMetadata.Address
        MemberCount: 4
        MemberNames:
            - Street
            - City
            - State
            - Zip
BinaryObjectString:
    ObjectId: 4
    Value: One Microsoft Way
BinaryObjectString:
    ObjectId: 5
    Value: Redmond
BinaryObjectString:
    ObjectId: 6
    Value: WA
BinaryObjectString:
    ObjectId: 7
    Value: 98054
MessageEnd
```


还存在很多已知和未知的问题待解决