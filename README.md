# Book API

Book API — это RESTful сервис для управления книгами, написанный на Go.

##  Установка и запуск

### 1. Клонирование репозитория
```sh
git clone https://github.com/Oleg-React-Dev/books.git
cd books

2. Настройка окружения
Создайте .env файл в корне проекта и добавьте туда переменные окружения, 
    пример: .env.example

3. Запуск сервера

make up - запуск БД
make run - запуск приложения


 Миграции базы данных
Создание новой миграции

make migrate-create migration_name

Применение миграций

make migrate-up

Откат миграций

make migrate-down

запуск тестов
make test

 Запуск проекта с предварительной сборкой

 make start

 Документация API (Swagger)
Swagger-документация доступна после запуска сервера по адресу:

http://localhost:8080/swagger/index.html

