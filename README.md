# TODO List REST API

REST API для управления задачами, разработанное на Go с использованием Fiber и PostgreSQL.

## Запуск через Docker
1. Склонируйте репозиторий:
   ```bash
   git clone https://github.com/FolderWithFolders/todo-api.git
   cd todo-api
2. Создайте .env как в репозитории
3. Запустите контейнер
   docker-compose up -d

## API доступно на http://localhost:3000/tasks

## Примеры запросов
### Создание задачи (POST /tasks)
![POST /tasks](images/post-request.png)
### Получение задач (GET /tasks)
![GET /tasks](images/get-request.png)
### Изменение задач (PUT /tasks)
![PUT /tasks](images/put-request.png)
### Удаление задач (DELETE /tasks)
![DELETE /tasks](images/delete-request.png)