# Job manager

<!-- vim-markdown-toc GitLab -->

* [Интерфейс взаимодействия](#Интерфейс-взаимодействия)
    * [Внешний API](#Внешний-api)
* [Спецификация](#Спецификация)
    * [Внешний API](#Внешний-api-1)
        * [Запуск задания](#Запуск-задания)
            * [URL](#url)
            * [Method](#method)
            * [Data Params](#data-params)
            * [Success Response](#success-response)
            * [Comments] (#comments)
            * [Error Response](#error-response)
            * [Sample call](#sample-call)
        * [Состояние задания](#Состояние-задания)
            * [URL](#url-3)
            * [Method](#method-3)
            * [Url Params](#url-params-3)
            * [Data Params](#data-params-3)
            * [Success Response](#success-response-3)
            * [Comments] (#comments-3)
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

##### Внешний API
+ `POST /jobs` --- запуск задания.
+ `GET /jobs/JOB_ID` --- получения состояния задания.
+ `GET /jobs` --- получение списка заданий.
+ `POST /jobs/JOBS_ID/cancel` --- отмена задания.
+ `DELETE /jobs/JOB_ID` --- удаление задания.

### Спецификация
#### Внешний API
REST
##### Запуск задания
###### URL
`/jobs`
###### Method
`POST`
###### Data Params
+ Required
    * `tasks = [array of 
       {command=[string] (required), input_files=[array of (uuid, path)], output_files=[array of string], ttl=[uint64]}
+ Optional
    * `type=string`
  `
###### Success Response
+ **Code**: `201`
+ **Content**
```
{
    "job_id": [uuid]
}
```

+ **Comments**: 
```
type = тип задания, сколько тасков внутри, по умолчанию один таск. - simple

time_limit = лимит на время жизни задачи (таска), по истечению которого оно будет принудительно завершено.
command = команда запуска задачи 
input_files = идентификаторы входного файлового ресурса и пути куда его нужно положить.
output_files = идентификаторы выходного файлового ресурса и пути откуда его прочитать, или stdout/stderr, если это стандартный вывод/ошибка.
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
curl https://sky.io/api/jobs -X POST --data '{
    "type": simple
    "tasks": [
        {
			"command": "wc lines.txt",
			"input_files": [["dc4eec6e-4e34-49a7-8fe8-8e19d5bfb8a4", "lines.txt"]].
			"output_files": ["stdout"],
			"time_limit": 500
		},
    ]
}'
```

##### Состояние задания
###### URL
+ `/jobs/JOB_ID`
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
    "results": [list of uuids],
    "spec": [...]
}
```
+ **Comments**: 
```
state = в каком состоянии задание (отправлено - SUBMITTED, запланировано - SCHEDULED, запущено - RUNNING, выполнено - DONE, выполнено с ошибкой - FAILED)
results = uuid выходного файла (или выходных файлов, если их несколько)
spec = параметры задания, достаточные чтобы перезапустить его
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
curl https://sky.io/api/jobs/dc4eec6e-4e34-49a7-8fe8-8e19d5bfb8a4
```

##### Список заданий
###### URL
+ `/jobs`
###### Method
`GET`
###### Data Params
Возможно, разные фильтры, по дате, пользователю и т.д.
###### Success Response
+ **Code**: `200`
+ **Content**
```
{
    array of [
    "job_id": uuid
    "state": [string],
    "results": [list of uuids],
    "spec": [...]]
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
curl https://sky.io/api/jobs
```

##### Отмена задания
###### URL
`/jobs/JOB_ID/cancel`
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
+ **Code**: `400 / 403 / 404`
+ **Content**
```
{
    "error": [string]
}
```

###### Sample call
```
curl https://sky.io/api/jobs/dc4eec6e-4e34-49a7-8fe8-8e19d5bfb8a4/cancel -X POST
```

##### Удаление задания
###### URL
`/jobs/JOB_ID/delete`
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
+ **Code**: `400 / 403 / 404`
+ **Content**
```
{
    "error": [string]
}
```

###### Sample call
```
curl https://sky.io/api/jobs/dc4eec6e-4e34-49a7-8fe8-8e19d5bfb8a4 -X DELETE
```