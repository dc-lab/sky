# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [AgentService.proto](#AgentService.proto)
  
  
  
    - [Agent](#.Agent)
  

- [Common.proto](#Common.proto)
    - [FileId](#.FileId)
    - [JobId](#.JobId)
  
  
  
  

- [RunJob.proto](#RunJob.proto)
    - [JobRequirement](#.JobRequirement)
    - [RunJobRequest](#.RunJobRequest)
    - [RunJobResponse](#.RunJobResponse)
  
    - [ResponseCode](#.ResponseCode)
  
  
  

- [StageInFile.proto](#StageInFile.proto)
    - [StageInFilesRequest](#.StageInFilesRequest)
    - [StageInFilesResponse](#.StageInFilesResponse)
    - [StagedFile](#.StagedFile)
  
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="AgentService.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## AgentService.proto


 

 

 


<a name=".Agent"></a>

### Agent


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| StageInFiles | [.StageInFilesRequest](#StageInFilesRequest) | [.StageInFilesResponse](#StageInFilesResponse) |  |
| RunJob | [.RunJobRequest](#RunJobRequest) | [.RunJobResponse](#RunJobResponse) |  |

 



<a name="Common.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## Common.proto



<a name=".FileId"></a>

### FileId



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) | required |  |






<a name=".JobId"></a>

### JobId



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) | required |  |





 

 

 

 



<a name="RunJob.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## RunJob.proto



<a name=".JobRequirement"></a>

### JobRequirement



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| python_version | [string](#string) | optional |  |






<a name=".RunJobRequest"></a>

### RunJobRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| job | [JobId](#JobId) | required |  |
| requirements | [JobRequirement](#JobRequirement) | repeated |  |
| files | [FileId](#FileId) | repeated | execution line ??? |






<a name=".RunJobResponse"></a>

### RunJobResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [ResponseCode](#ResponseCode) | required |  |





 


<a name=".ResponseCode"></a>

### ResponseCode
Подумать success and fail

| Name | Number | Description |
| ---- | ------ | ----------- |
| WAIT | 0 |  |
| RUN | 1 |  |
| DONE | 2 |  |


 

 

 



<a name="StageInFile.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## StageInFile.proto



<a name=".StageInFilesRequest"></a>

### StageInFilesRequest
Возможно достаточно просто файлов, чтобы зарегистрировать у себя. Неявная связь с job&#39;ами возникает только по запросу RunJob.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| job | [JobId](#JobId) | required |  |
| files | [FileId](#FileId) | repeated |  |






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
| file | [FileId](#FileId) | required |  |





 

 

 

 



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

