# Data manager

<!-- vim-markdown-toc GitLab -->

* [Интерфейс взаимодействия](#Интерфейс-взаимодействия)
    * [Как сервер](#Как-сервер)
        * [Внешнее API](#Внешнее-api)
        * [Внутреннее API](#Внутреннее-api)
    * [Как клиент](#Как-клиент)
* [Спецификация](#Спецификация)
    * [Внешнее API](#Внешнее-api-1)
        * [Загрузка файла в облако](#Загрузка-файла-в-облако)
            * [URL](#url)
            * [Method](#method)
            * [Data Params](#data-params)
            * [Success Response](#success-response)
            * [Error Response](#error-response)
            * [Sample call](#sample-call)
        * [Выгрузка файла](#Выгрузка-файла)
            * [URL](#url-1)
            * [Method](#method-1)
            * [Url Params](#url-params)
            * [Data Params](#data-params-1)
            * [Success Response](#success-response-1)
            * [Error Response](#error-response-1)
            * [Sample call](#sample-call-1)
        * [Получение метаинформации файла](#Получение-метаинформации-файла)
            * [URL](#url-2)
            * [Method](#method-2)
            * [Url Params](#url-params-1)
            * [Data Params](#data-params-2)
            * [Success Response](#success-response-2)
            * [Error Response](#error-response-2)
            * [Sample call](#sample-call-2)
    * [Внутреннее API](#Внутреннее-api-1)
        * [GRPC](#grpc)
        * [Распределение файлов по агентам](#Распределение-файлов-по-агентам)
        * [Клиент](#Клиент)
            * [Отправка файлов на агент](#Отправка-файлов-на-агент)
            * [Выгрузка файлов с агента](#Выгрузка-файлов-с-агента)

<!-- vim-markdown-toc -->

### Интерфейс взаимодействия
Файлы неизменяемые. Для начала храним в отдельном хранилище под data manager, в светлом будущем --- реплицированное распределенное хранилище на самих клиентах.
Все идентификаторы --- UUID.

#### Как сервер
##### Внешнее API
+ `POST /upload` --- загрузка файла.
+ `GET /download` --- скачивание файла.
+ `GET /info` --- получение метаинформации.

##### Внутреннее API
+ `gRPC /get_filemap` --- получение списка хостов, хранящих копию файла.
+ `gRPC /stage_in_files` --- распределение файлов по агентам.

#### Как клиент
+ `gRPC http://resource_manager/stage_in_files` --- отправка списка файлов на хост.
+ `gRPC http://resource_manager/stage_out_files` --- выгрузка файлов с агентов.

### Спецификация
#### Внешнее API
REST
##### Загрузка файла в облако
###### URL
`/upload`
###### Method
`POST`
###### Data Params
+ Required
    * `filename=[string]`
+ Optional
    * `ttl=[uint64]`
    * `acl=[array of { "user": [uuid], "perms": [oneof Read, ReadWrite]}]`
    * `replication=[uint32]`
    * `blob=[bytes]`
    * `url=[string]`
    * `copy_from=[uuid]`
###### Success Response
+ **Code**: `200`
+ **Content**
```
{
    "file_id": [uuid]
}
```

###### Error Response
+ **Code**: `400`
+ **Content**
```
{
    "error": [string]
}
```

###### Sample call
```
curl https://sky.io/api/file/upload -X POST --data '{
    "filename": "file.txt",
    "ttl": 14,
    "blob": "Hello, world!"
}'
```

##### Выгрузка файла
###### URL
+ `/download/:file_id`
+ `/download/:job_id/:filename`
###### Method
`GET`
###### Url Params
* **Required**:
  * `file_id=[uuid]`
  * `job_id=[uuid]`
  * `filename=[string]`
###### Data Params
None
###### Success Response
+ **Code**: `200`
+ **Content**
```
{
    "filename": [string],
    "contents": [bytes],
}
```

###### Error Response
+ **Code**: `400 / 404`
+ **Content**
```
{
    "error": [string]
}
```

###### Sample call
```
curl https://sky.io/api/files/dc4eec6e-4e34-49a7-8fe8-8e19d5bfb8a4/download
```

##### Получение метаинформации файла
###### URL
+ `/download/:file_id`
+ `/download/:job_id/:filename`
###### Method
`GET`
###### Url Params
* **Required**:
  * `file_id=[uuid]`
  * `job_id=[uuid]`
  * `filename=[string]`
###### Data Params
None
###### Success Response
+ **Code**: `200`
+ **Content**
```
{
    "file_id": [uuid],
    "filename": [string],
    "acl": [array of { "user": [uuid], "perms": [oneof Read, ReadWrite]}],
    "ttl": [uint64],
    "replication": [uint32],
    "byte_size": [uint64],
    "disk_size": [uint64]
}
```

###### Error Response
+ **Code**: `400 / 403 / 404`
+ **Content**
```
{
    "error": [string]
}
```

###### Sample call
```
curl https://sky.io/api/dc4eec6e-4e34-49a7-8fe8-8e19d5bfb8a4/info
```

#### Внутреннее API
Внешнее API + дополнительные ручки:

##### GRPC
```protobuf
service DataManger {
    rpc GetFileMap(FileMapRequest) returns (FileMapResponse) {}
    rpc StageInFiles(StageInFilesRequest) returns (StageInFilesResponse) {}
}
```

##### Распределение файлов по агентам
```protobuf
enum StageInFilesErrorKind {
    FileDoesNotExist = 1;
}

enum StageInFilesError {
    required StageInFilesErrorKind kind = 1;
    optional string description = 2;
}

message JobId {
    required string uuid = 1;
}

message StageInFilesRequest {
    required JobId job = 1;
    repeated FileId files = 2;
    repeated ResourceId resources = 3;
}

message StagedFile {
    required file_id file = 1;
    repeated ResourceId resources = 2;
};

message StageInFilesResponse {
    repeated StagedFile success_files = 1;
}
```

##### Клиент
###### Отправка файлов на агент
```protobuf
message StageInFilesRequest {
    repeated FileId files = 1;
}

message StageInFilesResponse {
    optional string error = 1;
}
```

###### Выгрузка файлов с агента
```protobuf
message StageOutFilesRequest {
    repeated string files = 1;
}

message StageOutFilesResponse {
    optional string error = 1;
}
```

