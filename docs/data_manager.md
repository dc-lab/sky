# Data manager

<!-- vim-markdown-toc GitLab -->

* [Интерфейс взаимодействия](#Интерфейс-взаимодействия)
    * [Как сервер](#Как-сервер)
        * [Внешний API](#Внешний-api)
        * [Внутренний API](#Внутренний-api)
    * [Как клиент](#Как-клиент)
* [Спецификация](#Спецификация)
    * [Внешний API](#Внешний-api-1)
    * [Внутренний API](#Внутренний-api-1)
        * [gRPC](#grpc)
        * [Распределение файлов по агентам](#Распределение-файлов-по-агентам)
        * [Клиент](#Клиент)
            * [Отправка файлов на агент](#Отправка-файлов-на-агент)
            * [Выгрузка файлов с агента](#Выгрузка-файлов-с-агента)

<!-- vim-markdown-toc -->

### Интерфейс взаимодействия
Файлы неизменяемые. Для начала храним в отдельном хранилище под data manager, в светлом будущем --- реплицированное распределенное хранилище на самих клиентах.
Идентификатор файла --- хеш (пока SHA-1 для определенности), остальные идентификаторы --- UUID по модулю удобства взаимодействия с остальными микросервисами.

#### Как сервер
##### Внешний API
+ `POST /files` --- загрузка файла.
+ `GET /files/{fileId}/data` --- скачивание файла.
+ `GET /files/{fileId}/info` --- получение метаинформации о файле.
+ `PUT /files/{fileId}` --- обновление метаинформации о файле (само содержимое файла неизменяемое).

##### Внутренний API
+ `gRPC /get_filemap` --- получение списка агентов, хранящих копию файла.
+ `gRPC /stage_in_files` --- распределение файлов по агентам.

#### Как клиент
+ `gRPC http://resource_manager/stage_in_files` --- загрузка списка файлов для последующего скачивания на агента.
+ `gRPC http://resource_manager/stage_out_files` --- выгрузка файлов с агента.

### Спецификация
#### Внешний API
[Swagger 3.0 описание](../api/generated/data_manager/index.html), [исходный data_manager.yaml](../api/openapi/data_manager.yaml)

#### Внутренний API
Внешний API + gRPC:

##### gRPC
```protobuf
service DataManager {
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

