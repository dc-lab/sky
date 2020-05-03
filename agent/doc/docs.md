# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [Protocol Documentation](#protocol-documentation)
  - [Table of Contents](#table-of-contents)
  - [agent.proto](#agentproto)
    - [TFromAgentMessage](#tfromagentmessage)
    - [TGreetings](#tgreetings)
    - [THardwareData](#thardwaredata)
    - [THardwareRequest](#thardwarerequest)
    - [TResult](#tresult)
    - [TTask](#ttask)
    - [TTaskRequest](#ttaskrequest)
    - [TTaskResponse](#ttaskresponse)
    - [TToAgentMessage](#ttoagentmessage)
    - [TResult.TErrorCode](#tresultterrorcode)
    - [TResult.TResultCode](#tresulttresultcode)
    - [ResourceManager](#resourcemanager)
  - [stage_in_file.proto](#stageinfileproto)
    - [StageInFilesRequest](#stageinfilesrequest)
    - [StageInFilesResponse](#stageinfilesresponse)
    - [StagedFile](#stagedfile)
  - [Scalar Value Types](#scalar-value-types)

    - [TResult.TErrorCode](#protobuf.TResult.TErrorCode)
    - [TResult.TResultCode](#protobuf.TResult.TResultCode)


    - [ResourceManager](#protobuf.ResourceManager)


- [stage_in_file.proto](#stage_in_file.proto)
    - [StageInFilesRequest](#.StageInFilesRequest)
    - [StageInFilesResponse](#.StageInFilesResponse)
    - [StagedFile](#.StagedFile)





- [Scalar Value Types](#scalar-value-types)



<a name="agent.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## agent.proto



<a name="protobuf.TFromAgentMessage"></a>

### TFromAgentMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hardware_data | [THardwareData](#protobuf.THardwareData) | optional |  |
| greetings | [TGreetings](#protobuf.TGreetings) | optional |  |
| task_response | [TTaskResponse](#protobuf.TTaskResponse) | optional |  |






<a name="protobuf.TGreetings"></a>

### TGreetings



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) | required |  |






<a name="protobuf.THardwareData"></a>

### THardwareData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cores_count | [int32](#int32) | required |  |
| memory_amount | [uint64](#uint64) | required |  |
| disk_amount | [uint64](#uint64) | required |  |






<a name="protobuf.THardwareRequest"></a>

### THardwareRequest







<a name="protobuf.TResult"></a>

### TResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result_code | [TResult.TResultCode](#protobuf.TResult.TResultCode) | required |  |
| error_code | [TResult.TErrorCode](#protobuf.TResult.TErrorCode) | optional |  |






<a name="protobuf.TTask"></a>

### TTask



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) | required |  |
| execution_shell_command | [string](#string) | required |  |
| requirements_shell_command | [string](#string) | optional |  |






<a name="protobuf.TTaskRequest"></a>

### TTaskRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [TTask](#protobuf.TTask) | required |  |






<a name="protobuf.TTaskResponse"></a>

### TTaskResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task_id | [string](#string) | required |  |
| result | [TResult](#protobuf.TResult) | required |  |






<a name="protobuf.TToAgentMessage"></a>

### TToAgentMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hardware_request | [THardwareRequest](#protobuf.THardwareRequest) | optional |  |
| task_request | [TTaskRequest](#protobuf.TTaskRequest) | optional |  |








<a name="protobuf.TResult.TErrorCode"></a>

### TResult.TErrorCode


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 |  |
| INTERNAL | 1 |  |
| INVALID_ARGUMENT | 2 |  |
| UNAUTHENTICATED | 3 |  |
| UNAUTHORIZED | 4 |  |



<a name="protobuf.TResult.TResultCode"></a>

### TResult.TResultCode


| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| WAIT | 1 |  |
| RUN | 2 |  |
| FAILED | 3 |  |
| SUCCESS | 4 |  |







<a name="protobuf.ResourceManager"></a>

### ResourceManager


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Send | [TFromAgentMessage](#protobuf.TFromAgentMessage) stream | [TToAgentMessage](#protobuf.TToAgentMessage) stream |  |





<a name="stage_in_file.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stage_in_file.proto



<a name=".StageInFilesRequest"></a>

### StageInFilesRequest
Возможно достаточно просто файлов, чтобы зарегистрировать у себя. Неявная связь с job&#39;ами возникает только по запросу RunJob.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task_id | [string](#string) | required |  |
| file_id | [string](#string) | repeated |  |






<a name=".StageInFilesResponse"></a>

### StageInFilesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success_files | [StagedFile](#StagedFile) | repeated |  |






<a name=".StagedFile"></a>

### StagedFile
Увидел это в других спеках. Обдумать, почему для этого должен быть отдельный класс (может быть просто класс File).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| file_id | [string](#string) | required |  |















## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |
