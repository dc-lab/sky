# Job manager

<!-- vim-markdown-toc GitLab -->

* [Интерфейс взаимодействия](#Интерфейс-взаимодействия)
    * [Внешнее API](#Внешнее-api)
* [Спецификация](#Спецификация)
    * [Внешнее API](#Внешнее-api-1)
        * [Запуск задания](#Запуск-задания)
            * [URL](#url)
            * [Method](#method)
            * [Data Params](#data-params)
            * [Success Response](#success-response)
            * [Error Response](#error-response)
            * [Sample call](#sample-call)
        * [Состояние задания](#Состояние-задания)
            * [URL](#url-3)
            * [Method](#method-3)
            * [Url Params](#url-params-3)
            * [Data Params](#data-params-3)
            * [Success Response](#success-response-3)
            * [Error Response](#error-response-3)
            * [Sample call](#sample-call-3)
        * [Список заданий](#Список-заданий)
            * [URL](#url-1)
            * [Method](#method-1)
            * [Url Params](#url-params)
            * [Data Params](#data-params-1)
            * [Success Response](#success-response-1)
            * [Error Response](#error-response-1)
            * [Sample call](#sample-call-1)
        * [Отмена задания](#Отмена-задания)
            * [URL](#url-2)
            * [Method](#method-2)
            * [Url Params](#url-params-1)
            * [Data Params](#data-params-2)
            * [Success Response](#success-response-2)
            * [Error Response](#error-response-2)
            * [Sample call](#sample-call-2)
        * [Удаление задания](#Удаление-задания)
            * [URL](#url-4)
            * [Method](#method-4)
            * [Url Params](#url-params-4)
            * [Data Params](#data-params-4)
            * [Success Response](#success-response-4)
            * [Error Response](#error-response-4)
            * [Sample call](#sample-call-4)
    * [Внутреннее API](#Внутреннее-api-1)
        * Пока не предполагается

<!-- vim-markdown-toc -->

### Интерфейс взаимодействия
Задания состоят из бинарного запускаемого файла и входных данных, загружаемых на агента из дата-менеджера.
Все идентификаторы --- UUID.

##### Внешнее API
+ `POST /run` --- запуск задания.
+ `GET /job_info` --- состояние задания.
+ `GET /list_jobs` --- список заданий.
+ `POST /cancel` --- отмена задания.
+ `DELETE /delete` --- удаление задания.

### Спецификация
#### Внешнее API
REST
##### Запуск задания
###### URL
`/run`
###### Method
`POST`
###### Data Params
+ Required
    * `command=[string]`
+ Optional
    * `binary=[uuid]`
    * `input_files=[array of uuid]`
    * `ttl=[uint64]`
###### Success Response
+ **Code**: `201`
+ **Content**
```
{
    "job_id": [uuid]
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
curl https://sky.io/api/job/run -X POST --data '{
    "command": "uname -a | wc -l",
}'
```

##### Состояние задания
###### URL
+ `/job_info/:job_id`
###### Method
`GET`
###### Url Params
* **Required**:
  * `job_id=[uuid]`
###### Data Params
None
###### Success Response
+ **Code**: `200`
+ **Content**
```
{
    "state": [string],
    "answer_id": [uuid],
}
```

###### Error Response
+ **Code**: `400 / 404 / 410`
+ **Content**
```
{
    "error": [string]
}
```

###### Sample call
```
curl https://sky.io/api/job/job_info/dc4eec6e-4e34-49a7-8fe8-8e19d5bfb8a4
```

##### Список заданий
###### URL
+ `/list_jobs`
###### Method
`GET`
###### Data Params
Возможно, разные фильтры, по дате, пользователю и т.д.
###### Success Response
+ **Code**: `200`
+ **Content**
```
{
    "job_ids" : [array of uids]
}
```

###### Error Response
+ **Code**: `400 / 403`
+ **Content**
```
{
    "error": [string]
}
```

###### Sample call
```
curl https://sky.io/api/list_jobs
```

##### Остановка задания
###### URL
`/cancel`
###### Method
`POST`
###### Data Params
+ Required
    * `job_id=[uuid]`
###### Success Response
+ **Code**: `200`
+ **Content**
```
{
    "job_id": [uuid]
}
```

###### Error Response
+ **Code**: `400 / 403`
+ **Content**
```
{
    "error": [string]
}
```

###### Sample call
```
curl https://sky.io/api/job/cancel -X POST --data '{
    "job_id": "dc4eec6e-4e34-49a7-8fe8-8e19d5bfb8a4",
}'
```

##### Удаление задания
###### URL
`/delete`
###### Method
`DELETE`
###### Data Params
+ Required
    * `job_id=[uuid]`
###### Success Response
+ **Code**: `200`
+ **Content**
```
{
    "job_id": [uuid]
}
```

###### Error Response
+ **Code**: `400 / 403`
+ **Content**
```
{
    "error": [string]
}
```

###### Sample call
```
curl https://sky.io/api/job/cancel -X DELETE --data '{
    "job_id": "dc4eec6e-4e34-49a7-8fe8-8e19d5bfb8a4",
}'
```