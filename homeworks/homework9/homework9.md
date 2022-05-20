# Домашнее задание - написать тесты на функцию Greet


## Дано:

В файле `greeter.go` определена функция `Greet`, которая принимает имя и номер часа и печатает приветствие.

Функция должна работать по следующей логике:

1. Имя всегда печатается с заглавной буквы.
2. Если час между 6 и 12, выводится "Good morning <name>".
3. Если час между 18 и 22, выводится "Good evening <name>".
4. Если час между 22 и 6, выводится "Good night <name>".
5. "Hello <name>" выводится во всех остальных случаях.
6. Пробельные символы в начале и конце имени обрезаются.

Чтобы выполнить задание, в проекте `act-device-api` создайте отдельный пакет с названием `greeter`. Перенесите в него файл `greeter.go`. Все unit-тесты на тестируемую функцию следует хранить в этом же пакете.

## Задание:

- Напишите тест, используя паттерн test table
- Покройте основные сценарии
- Сгенерируйте отчет о покрытии кода тестами
- :gem: Найдите и пофиксите как минимум 2 бага в коде, добавьте тесты
- :gem: :gem: Посчитайте необходимое количество тест кейсов для проверки всех возможных состояний функции `Greet`


## Критерии приемки задания:

- Как минимум 6 тестов
- Тесты запускаются с помощью t.Run
- Есть html файл с измеренным покрытием

## Подсказки:

Команды для получения отчета о покрытии:
```bash
go test -coverprofile=cover.txt
go tool cover -html=cover.txt -o cover.html
```

***

# Решение ДЗ
* [Тесты](greeter/greeter_test.go) 
* [Покрытие](cover.html) 

## Баги:
* Диапазон "Hello" пересекаются с "Good evening" (18 включается в обоих диапазонах). Для 18 
ожидаем "Good evening".   
* Диапазон "Good evening" заканчивается в 22, ожидаем в 24 не включая.
* При запросе значения больше 24 получаем надпись "Good night" ожидаем "Hello".

## Необходимое количество кейсов
### Тесты кесы на диапазоны
* 4 диапазона, 2 теста на верхнию и нижнию границу плюс один на середину диапазона. Итого: 4*3=12
* Тесты на диапазон больше 24. Например 99. Итого: 1 кейс
* Тесты на не валидные значения часов. (отрицательные числа). Итого: 1 кейс

### Тесты на имена
* Тесы на имена.  Итого: 6 кейсов
  + Имя с большой буквы.  
  + Имя с маленькой буквы  
  + Имя с пробелом слева  
  + Имя с пробелом справа  
  + Имя с пробелом слева и справа  
  + Пустое имя

Итого мы имеем 14 кейсов на диапазоны и 6 кесов на имена, что бы проверить все возможные значения необходимо 14*6=84 тест кейса.