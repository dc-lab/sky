# Resource manager
Для внутренних пользователей Resource manager является некой умной proxy, которая доставляет запросы от компонент системы до Агента и события от Агента до компонент системы.
В каждом запросе к Resource manger должен находиться идентификатор ресурса, к которому хочет обратиться компонента. Дальше этот id достается из данных запроса и по нему определяется текущий адрес машины, все остальные данные прокидываются дальше. Так что в этой схеме я бы отделял данные реального запроса в специальный proto-объект.
События описываются заранее оговоренным идентификатором и неким специфичным для события набором данных.


### Управление ресурсами
#### Создание ресурса
`POST /resources`
###### Data params:
* `hostname=[string]`
* `user=[string]`

###### Success Response
+ **Code**: `200`
+ **Content**
```
{
    "id": [uuid],
    "token": [uuid]
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
curl https://sky.io/api/resources -X POST --data '{
    "hostname": "my first one",
    "user": "steve"
}'
```

#### Удаление ресурса
`DELETE /resources/<id>`

###### Data params:
None

###### Success Response
+ **Code**: `200`

###### Error Response
+ **Code**: `404`

###### Sample call
```
curl https://sky.io/api/resources/57bd00cb-e657-4ca0-b462-ff0e6878eb68 -X DELETE
```

#### Информация о ресурсе
`GET /resources/<id>`

###### Data params:
None

###### Success Response
+ **Code**: `200`
+ **Content**
```
{
    "id": [uuid],
    "hostname": [string],
    ...
}
```

###### Error Response
+ **Code**: `404`

###### Sample call
```
curl https://sky.io/api/resources/57bd00cb-e657-4ca0-b462-ff0e6878eb68 -X GET
```

#### Найти ресурс
`GET /resources/?hostname=<hostname>`

###### Data params:
None

###### Success Response
+ **Code**: `200`
+ **Content**
```
{
    "id": [uuid]
}
```

###### Error Response
+ **Code**: `404`

###### Sample call
```
curl https://sky.io/api/resources?hostname=my%20first%20one -X GET
```

### Обработка джобы
#### Stage in files
gRPC-метод, проксирующий запрос на загрузку файлов от DataManager к Агенту на конкретной машине

#### Files staged in
gRPC-метод, проксирующий событие загрузки файлов от Агента к DataManager

#### Submit task
gRPC-метод, проксирующий запрос на запуск задачи от JobManager к Агенту на конкретной машине

#### Task accepted
gRPC-метод, проксирующий событие принятия джобы от Агента к JobManager

#### Task running
gRPC-метод, проксирующий событие запуска джобы от Агента к JobManager

#### Task completed
gRPC-метод, проксирующий событие завершения джобы от Агента к JobManager

#### Stage out file
gRPC-метод, проксирующий запрос на загрузку результатов работы джобы от DataManager к Агенту на конкретной машине

#### File staged out
gRPC-мтеод, проксирующий событие завершения результатов работы джобы от Агента к DataManager


