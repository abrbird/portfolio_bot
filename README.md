# Бот-монитор портфеля

> Бот, который позволяет мониторить баланс портфеля из разных финансовых инструментов.
Необходимые возможности:
> 1. Бот хранит информацию об объёме разных финансовых инструментов (фалюты, акции, фонды)
> 2. Бот выводит информацию о текущем объёме портфеля, изменениях за час, день, неделю
> 3. Можно придумать как сделать так, чтобы хранить информацию о разных пользователях без раскрытия данных


## Компоненты
1. БД postgresql
2. Go сервис для сбора данных о котировках, курсах, ценах
   1. различные финансовые инструменты: (крипто-)валюты, драг.металлы, акции и пр.
   2. адаптор под каждый внешний сервис
   3. запрашивать данные периодически и хранить историю (шедулинг тасков каждые n минут)
3. Go сервис для обработки запросов пользователя
   1. создание пользователя
   2. информация об имеющихся возможностях
   3. информация о портфеле, добавление, редактирование, удаление элемента портфеля
   4. представление информации о портфеле с изменениями по каждому элементу портфеля + итог в виде таблицы, графика.
   5. шифрование данных  - TODO
4. Telegram-бот на фронте

## Запуск
```
docker-compose -f docker-compose.yml up
```
