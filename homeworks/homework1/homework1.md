#  Тур исследовательского тестирования сервиса device-api

## Пример запросов

### Получение всех девайсов
```curl
curl -X 'GET' \
'http://localhost:8080/api/v1/devices?page=1&perPage=10' \
-H 'accept: application/json'
```
### Создание дейвайса
```curl
curl -X 'POST' \
  'http://localhost:8080/api/v1/devices' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "platform": "my_test",
  "userId": "1"
}'
```
### Получение выбранного девайса
```curl
curl -X 'GET' \
  'http://localhost:8080/api/v1/devices/1' \
  -H 'accept: application/json'
```

### Удаление выбраного дейвайса
```curl
curl -X 'DELETE' \
  'http://localhost:8080/api/v1/devices/11' \
  -H 'accept: application/json'
```

###  Обновление выбранного девайса
```curl
curl -X 'PUT' \
  'http://localhost:8080/api/v1/devices/11' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "platform": "ав",
  "userId": "1"
}'
```


## Баги
* При запросе получения девайсов отствуют значение по умолчания для параметров page и perPage.
* При запросе получения девайсов без параметров page и perPage возвращается пустой массив.
```curl
curl -X 'GET' \
  'http://localhost:8080/api/v1/devices' \
  -H 'accept: application/json'
```

```json
{
  "items": []
}
```


## Заметки
Сервис позволяет выполнять CRUD операция над девайсами. При попытке удаление или обновления не существующего сервиса
возвращается код 200 с
```
{
  "found": false
}
```
,но при попытке получение девайса по не существующему ID получаем 404
```
{
  "code": 5,
  "message": "device not found",
  "details": []
}
```
необходимо уточнить у аналитика нормальное ли это поведение или необходимо исправить.

## Позитивные кейсы

### Получение всех девайсов
* Запрашиваем [GET] /device.
* Проверяем, что в ответе содержится лист девайсов.

### Получение выбранного девайсов
* Запрашиваем [GET] /device. Выбираем любой {deviceID}.
* Запрашиваем [GET] /device/{deviceID}. Проверяем, что в ответе есть выбранный девайс.

### Создание нового дейвайса
* Запрашиваем [GET] /device. Запоминаем список.
* Создаем новый девайс [POST] /device.
* Запрашиваем [GET] /device. Проверяем, что новый девайс есть в списке и старые девайсы присутствуют.
* Запрашиваем [GET] /device/{deviceID}. Проверяем полученные данные.

### Удаление дейвайса
* Создаем девайс [POST] /device. Запоминаем {deviceID}.
* Запрашиваем [GET] /device. Запоминаем список.
* Делаем [DELETE] /device{deviceID}.
* Запрашиваем [GET] /device. Проверяем, что девайса нет в списке.
* Запрашиваем [GET] /device/{deviceID}. Проверяем, что девайса нет.

### Обновление девайса
* Создаем девайс [POST] /device. Запоминаем {deviceID}.
* Запрашиваем [GET] /device. Запоминаем список.
* Обновляем поле "platform" через [PUT] /device/{deviceID}.
* Запрашиваем [GET] /device. Проверяем, что новый девайс содержит актуальноое поле "platform"
* Запрашиваем [GET] /device/{deviceID}. Проверяем, что девайс содержит актуальноое поле "platform"
