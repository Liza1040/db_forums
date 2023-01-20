# API базы данных для форума
Данный репозиторий содержит реализацию API для курса СУБД из этой спецификации:
https://github.com/mailcourses/technopark-dbms-forum/blob/master/swagger.yml

## Состав Docker-контейнеров
Docker-контейнер организован по следующему приципу:

* Приложение:
    * использует и объявляет порт 5000 (http);
* PostgreSQL:
    * использует и объявляет порт 5342 (http);

## Как собрать и запустить контейнер
Для сборки контейнера можно выполнить команду вида:
```bash
docker build -t liza https://github.com/Liza1040/DB_forums.git
```
Или команды:
```bash
git clone https://github.com/Liza1040/DB_forums.git DB_forums
cd DB_forums/
docker build -t liza .
```

После этого будет создан Docker-образ с именем `liza` (опция `-t`).

Запустить ранее собранный контейнер можно командой вида:
```bash
docker run -p 5000:5000 --name liza -t liza
```
После этого можно получить доступ к запущенному в контейнере приложению по URL: ```http://localhost:5000/``` (базовая часть URL: ```/api```)

Получить список запущенных контейнеров можно командой:
```bash
docker ps
```

Остановить контейнер можно командой:
```bash
docker kill liza
```
