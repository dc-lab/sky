# Data manager


<!-- vim-markdown-toc GitLab -->

- [Интерфейс взаимодействия](#Интерфейс-взаимодействия)
    + [Сервер](#Сервер)
        * [Внешнее API](#Внешнее-api)
        * [Внутреннее API](#Внутреннее-api)
    + [Клиент](#Клиент)
- [Спецификация](#Спецификация)
    + [Внешнее API](#Внешнее-api-1)
        * [Загрузка файла](#Загрузка-файла)
        * [Получение метаинформации о файле](#Получение-метаинформации-о-файле)
        * [Загрузка файла](#Загрузка-файла-1)
    + [Внутреннее API](#Внутреннее-api-1)
        * [Получение списка хостов (агентов) с актуальной версией файла](#Получение-списка-хостов-агентов-с-актуальной-версией-файла)
        * [Распределение файлов по агентам](#Распределение-файлов-по-агентам)
        * [GRPC](#grpc)
        * [Клиент](#Клиент-1)
            - [Отправка файлов на агент](#Отправка-файлов-на-агент)
            - [Выгрузка файлов с агента](#Выгрузка-файлов-с-агента)

<!-- vim-markdown-toc -->

### Интерфейс взаимодействия
Файлы неизменяемые. Для начала храним в отдельном хранилище под data manager, в светлом будущем --- реплицированное распределенное хранилище на самих клиентах.
Все идентификаторы --- UUID.

#### Сервер
##### Внешнее API
+ `POST /upload` --- загрузка файла.
+ `POST /download` --- скачивание файла.
+ `GET /info` --- получение метаинформации.

##### Внутреннее API
+ `GET/GRPC /get_filemap` --- получение списка хостов, хранящих копию файла.
+ `POST/GRPC /stage_in_files` --- распределение файлов по агентам.

#### Клиент
+ `POST/GRPC http://resource_manager/stage_in_files` --- отправка списка файлов на хост.
+ `POST/GRPC http://resource_manager/stage_out_files` --- выгрузка файлов с агентов.

### Спецификация
#### Внешнее API
HTTP
##### Загрузка файла
```protobuf
// POST
// /upload

message UserId {
    required string uuid = 1;
}

enum AccessPerms {
    kNone = 0,
    kRead = 1,
    kReadWrite = 2,
    // kWrite is meaningless
}

message UserWithPerms {
    required UserId user = 1;
    required AccessPerms perms = 2;
}

message FileId {
    required string uuid = 1;
}

message Request {
    required string filename = 1;
    repeated UserWithPerms acl = 2;
    required int32 ttl = 3;
    required int32 replication = 4;
    oneof contents {
        bytes blob = 5;
        string url = 6;
        FileId copy = 7;
    }
}

message Response {
    required FileId id = 1;
}
```

##### Получение метаинформации о файле
```protobuf
// GET
// /info?file_id=FILE_ID

message ResponseSuccess {
    required FileId file_id = 1;
    required string filename = 2;
    repeated UserWithPerms acl = 3;
    required int32 ttl = 4;
    required int32 replication = 5;
    required uint64 byte_size = 6;
    required uint64 occupied_storage = 7;
}

enum ErrorKind {
    AcccessDenied = 1;
    FileDoesNotExist = 2;
    FileDeleted = 3;
}

message ResponseError {
    required ErrorKind kind = 1;
    optional string description = 2;
}

message Repsonse {
    oneof result {
        ResponseSuccess success = 1;
        ResponseError error = 2;
    }
}
```

##### Загрузка файла
```protobuf
// GET
// /download?file_id=FILE_ID

message Response {
    required string filename = 1;
    required bytes contents = 2;
}
```

#### Внутреннее API
Внешнее API + дополнительные ручки:

##### Получение списка хостов (агентов) с актуальной версией файла
```protobuf
// GET or GRPC
// /get_filemap?file_id=FILE_ID

message HostId {
    required string uuid = 1;
}

enum FileMapErrorKind {
    FileDoesNotExist = 1;
}

enum FileMapError {
    required FileMapErrorKind kind = 1;
    optional string description = 2;
}

message FileMapSuccessResponse {
    repeated HostId hosts = 1;
}

message FileMapRequest {
    required FileId file = 1;
}

message FileMapResponse {
    oneof result {
        FileMapSuccessResponse success = 1;
        FileMapError error = 2;
    }
}
```

##### Распределение файлов по агентам
```protobuf
// POST or GRPC
// /stage_in_files

enum StageInFilesErrorKind {
    FileDoesNotExist = 1;
}

enum StageInFilesError {
    required StageInFilesErrorKind kind = 1;
    optional string description = 2;
}

message StageInFilesRequest {
    repeated FileId files = 1;
    repeated HostId hosts = 2;
}

message StageInFilesResponse {
    repeated HostId success_hosts = 1;
}
```

##### GRPC
```protobuf
service DataManger {
    rpc GetFileMap(FileMapRequest) returns (FileMapResponse) {}
    rpc StageInFiles(StageInFilesRequest) returns (StageInFilesResponse) {}
}
```

##### Клиент
###### Отправка файлов на агент
```protobuf
// POST or GRPC
// /stage_in_files

message StageInFilesRequest {
    repeated FileId files = 1;
}

message StageInFilesResponse {
    optional string error = 1;
}
```

###### Выгрузка файлов с агента
```protobuf
// POST or GRPC
// /stage_out_files

message StageOutFilesRequest {
    repeated FileId files = 1;
}

message StageOutFilesResponse {
    optional string error = 1;
}
```

