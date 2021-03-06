# Домашнее задание


# Покрыть тест кейсами device-api

Цель:
Применение на практике техник тест дизайна

## Критерии приемки задания:
1 Как минимум должны быть покрыты все позитивные кейсы
2 Кейсы должны быть правильно приоритизированы
3 Кейсы должны быть пригодны для дальнейшей автоматизации (возможно не все)

******************

## Endpoint для тестирования
* [GET] /api/v1/devices
* [POST] /api/v1/devices
* [GET] /api/v1/devices/{deviceID}
* [DELETE] /api/v1/devices/{deviceID}
* [PUT] /api/v1/devices/{deviceID}

## Скрипты для тестиорвания

* Список активных девайсов 1-5 (0)
```sql
select id, platform, user_id, entered_at from devices where removed is false order by id desc limit 5;
```
* Список активных девайсов 5-10 (1)
```sql
select id, platform, user_id, entered_at from devices where removed is false order by id desc offset 5 limit 5;
```

* Список активных девайсов 1-10 (2)
```sql
select id, platform, user_id, entered_at from devices where removed is false order by id desc limit 10;
```

* Активный(не удаленный) девайс с максимальным id (3)
```sql
select id, platform, user_id from devices where id in (select max(id) from devices where removed is false limit 1);
```

* Активный(не удаленный) девайс с минимальным id (4)
```sql
select id, platform, user_id from devices where id in (select min(id) from devices where removed is false limit 1);
```

* Удаленный девайс в бд (5)
```sql
select id from devices where removed is true limit 1;
```

* Несуществующий девайс (6)
```sql
select max(id)+1000 as id from devices where removed is false limit 1
```


## Тест кейсы


### Тест кейсы [GET] /api/v1/devices

| Test case name                |                   Scenario                    |                                            Expected result |                            Actual result |     Status | Priority |
|-------------------------------|:---------------------------------------------:|-----------------------------------------------------------:|-----------------------------------------:|-----------:|---------:|
| Первая страница девайсов      | Запрос девайсов указанием page=1 и perPage=5  |                              Получение девайсов скрипт (0) |       Вернулись девайсы с id  5,7,8,9,10 | **passed** |        0 |
| Вторая страница девайсов      | Запрос девайсов указанием page=2 и perPage=5  |                              Получение девайсов скрипт (1) |            Вернулись девайсы с id  1,2,3 | **passed** |        0 |
| Третья страница девайсов      | Запрос девайсов указанием page=3 и perPage=5  |                                     Пустой список девайсов |                   Пустой список девайсов | **passed** |        0 |
| Все девайсы на одной странице | Запрос девайсов указанием page=1 и perPage=10 |                              Получение девайсов скрипт (2) | Вернулись девайсы с id  1,2,3,5,7,8,9,10 | **passed** |        0 |
| Девайсы без page              | Запрос девайсов без указания page, perPage=10 |                              Получение девайсов скрипт (2) |                               Ошибка 500 |   **fail** |        1 |
| Девайсы без perPage           | Запрос девайсов без указания perPage, page=1  |        Получение девайсов с id дефолтным значением perPage |                            Пустой список |   **fail** |        1 |
| Девайсы без perPage и page    |  Запрос девайсов без указания perPage и page  | Получение девайсов с id дефолтным значением perPage и Page |                            Пустой список |   **fail** |        1 |


### Тест кейсы [POST] /api/v1/devices

| Test case name            |                      Scenario                      |                   Expected result |                           Actual result |     Status | Priority |
|---------------------------|:--------------------------------------------------:|----------------------------------:|----------------------------------------:|-----------:|---------:|
| Успешное создание девайса | Отправка запроса с {"platform": "1","userId": "2"} | Проверка данных в бд - скрипт (3) | Данные в бд id=12, platform=1, userId=2 | **passed** |        0 |


### Тест кейсы [GET] /api/v1/devices/{deviceID}

| Test case name               |           Scenario            |                            Expected result |                             Actual result |     Status | Priority |
|------------------------------|:-----------------------------:|-------------------------------------------:|------------------------------------------:|-----------:|---------:|
| Получение первого девайса    |  Запрос первого (4) девайса   | Получение девайса с данными из скрипта (4) | Данные в ответе соответствуют данным в бд | **passed** |        0 |
| Получение последнего девайса | Запрос последнего (3) девайса | Получение девайса с данными из скрипта (3) | Данные в ответе соответствуют данным в бд | **passed** |        0 |


### Тест кейсы  [DELETE] /api/v1/devices/{deviceID}

| Test case name                    |                     Scenario                      |                                                                         Expected result |                                                                           Actual result |     Status | Priority |
|-----------------------------------|:-------------------------------------------------:|----------------------------------------------------------------------------------------:|----------------------------------------------------------------------------------------:|-----------:|---------:|
| Удаление первого девайса          |        Удаление первого (скрипт 4) девайса        | В ответе found:true, в бд изменился флаг removed для удаляемого девайся с true на false | В ответе found:true, в бд изменился флаг removed для удаляемого девайся с true на false | **passed** |        0 |
| Удаление удаленного девайса       | Удаление девайса с флагом removed:true (скрипт 5) |                                    В ответе found:false, в бд не изменился флаг removed |                                    В ответе found:false, в бд не изменился флаг removed | **passed** |        1 |
| Удаление не существующего девайса |   Удаление не существующего девайса (скрипт 6)    |                                                                    В ответе found:false |                                                                    В ответе found:false | **passed** |        1 |


### Тест кейсы  [PUT] /api/v1/devices/{deviceID}

| Test case name                                    |                                                      Scenario                                                      |                                        Expected result |                                           Actual result |     Status | Priority |
|---------------------------------------------------|:------------------------------------------------------------------------------------------------------------------:|-------------------------------------------------------:|--------------------------------------------------------:|-----------:|---------:|
| Обновление platform для первого девайса           |      Обновление значения platform (старое значение + _test) у первого девайса (скрипт 4), userID не меняется       |             Значение в бд изменилось для поля platform |              Значение в бд изменилось для поля platform | **passed** |        0 |
| Обновление userId для первого девайса             |        Обновление значения userID (старое значение + 1) у первого девайса (скрипт 4), platform не меняется         |               Значение в бд изменилось для поля userId |              Значение в бд изменилось для поля platform | **passed** |        0 |
| Обновление platform и userId  для первого девайса | Обновление значения platform (старое значение + _test) у первого девайса (скрипт 4) и userID (старое значение + 1) |    Значение в бд изменилось для поля userId и platform |     Значение в бд изменилось для поля userId и platform | **passed** |        0 |
| Обновление platform и userId  старыми значениями  |                               В запросе те же значения(скрипт 4), что находятся в бд                               | Значение в бд не изменилось для поля userId и platform |  Значение в бд не изменилось для поля userId и platform | **passed** |        1 |
