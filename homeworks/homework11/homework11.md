# Домашнее задание - написать тесты c использованием grpc клиента

## Задание:

Автоматизировать тесты написанные на ручки: ListDevicesV1, CreateDeviceV1, DescribeDeviceV1

После лекции по тест дизайну нужно было покрыть тестами api методы сервиса, теперь нужно автоматизировать часть этих тестов.
У вас мало времени по этому нужно выбрать то что вы будете автоматизировать в первую очередь
Так же стоит учитывать материал из 9-й лекции по юнит тестам (см раздел "Поддерживаемые тесты")

## Критерии приемки задания:

- Автоматизировано минимум по 3 тест кейса на каждый метод
- Структурированный код теста (Arrange-Act-Assert)
- Понятное для “читателя” имя (Паттерны именования см 9 лекцию)
- Независимый от окружения.
- Достоверный.


## Решение
* [ListDevicesV1](list_devices_test.go)
* [CreateDeviceV1](create_devices_test.go)
* [DescribeDeviceV1](describe_device_test.go)
