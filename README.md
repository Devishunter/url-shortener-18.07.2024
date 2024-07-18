# url-shortener
 Challenge: URL trimming service using Golang and Docker
Задача: сервис обрезки URL-адресов с использованием Golang и Docker

Project Description

The URL Shortener project is a web service that provides the ability to create shortened URLs and redirect to the original long URLs. The service includes two main API endpoints:

 POST /shorten: Takes a JSON request with a long URL and returns a shortened URL.

 GET /{shortURL}: Redirects to the original long URL.

Project structure

url-shortener/

├── main.go

├── main_test.go

├── go.mod

└── go.sum

Main Components

 main.go: main file with application logic.

 main_test.go: file with tests to check the correct operation of the service.

 go.mod and go.sum: files that manage Go dependencies.

Installation and launch

Requirements

 Go 1.18 or higher

 MariaDB/MySQL server

Описание проекта 

Проект URL Shortener - это веб-сервис, который предоставляет возможность создавать сокращенные URL и перенаправлять на исходные длинные URL. Сервис включает два основных API-эндпоинта: 

    POST /shorten: принимает JSON-запрос с длинным URL и возвращает сокращенный URL. 

    GET /{shortURL}: перенаправляет на исходный длинный URL. 

Структура проекта 

url-shortener/ 

├── main.go 

├── main_test.go 

├── go.mod 

└── go.sum 

Основные компоненты 

    main.go: основной файл с логикой приложения. 

    main_test.go: файл с тестами для проверки корректности работы сервиса. 

    go.mod и go.sum: файлы, управляющие зависимостями Go. 

Установка и запуск 

Требования 

    Go 1.18 или выше 

    MariaDB/MySQL сервер 