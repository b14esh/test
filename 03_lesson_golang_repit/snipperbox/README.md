//Окружеине GOLANG
go env


// Запускать 
// go run ./cmd/web
// go run ./cmd/web -addr="127.0.0.1:9999"
// go run ./cmd/web -help
// go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log


// error fix
// go: go.mod file not found in current directory or any parent directory; see 'go help modules'  
// go mod init main
// Где main название директории проэкта
go mod init snipperbox
go mod tidy




// страницы
// <название>.<роль>.tmpl
// ui/html/home.page.tmpl - главная страница
// ui/html/base.layout.tmpl - отображаем на каждой странице


// файлы
cmd - все что можно собрать в GO
cmd/web/routes.go - маршруты
cmd/web/main.go - главное тело, запуск сервера, работа с ошибками
cmd/web/handlers.go - что запустить что-бы отрисовать сайт, указывает этот файл
ui - файлы для отрисовки сайта
ui/html/base.layout.tmpl - базовый шаблон
ui/html/footer.partial.tmpl - подвал
ui/html/home.page.tmpl - главная страница
ui/static/css/main.css - стиль css
ui/static/img/logo.png - лого на главной странице
ui/static/img/favicon.ico - иконка в браузере:
ui/static/js/main.js - скрипты
ui/html/show.page.tmpl - вывод из MYSQL 
pkg - файлы пакетов 
pkg/models/models.go - модель, уровень обслуживания или уровень доступа к данным MYSQL 
pkg/models/mysql/snippets.go - файл для работы с заметками


// методы
Метод	URL      	     Обработчик  	    Действие
ANY  	/	             home      	        Отображение домашней страницы
ANY	    /                snippet?id=1	    showSnippet	Отображение определенной заметки
POST	/snippet/create	 createSnippet	    Создание новой заметки
ANY	    /static/	     http.FileServer	Обслуживание определенного статического файла






// curl test
проверка скачивания
curl -i -H "Range: bytes=100-199" --output - http://localhost:4000/static/img/logo.png
проверка запроса POST
curl -i -X POST http://127.0.0.1:4000/snippet/create
проверка запроса PUT
curl -i -X PUT http://127.0.0.1:4000/snippet/create
проверка запроса GET можно просто открыть страницу в браузере или 
curl -i -X GET http://127.0.0.1:4000/snippet/create






// собрать бинарник
// под x86
go build ./cmd/web
GOOS=windows GOARCH=386 go build ./cmd/web
GOOS=windows GOARCH=amd64 go build ./cmd/web
// под arm
GOOS=linux GOARCH=arm go build ./cmd/web
// под arm work macos
GOOS=darwin GOARCH=arm64 go build ./cmd/web



//Установка драйвера для работы с MYSQL
go get github.com/go-sql-driver/mysql
//Показать установленные модули
go list -m all
//Удаление модуля(вся магия в @none)
go get github.com/foo/bar@none
//Установка модуля нужной версии
go get -u github.com/foo/bar@v2.0.0
//Или, если вы удалили все упоминания о пакете в своем коде, вы можете запустить команду "go mod tidy", которая автоматически удалит все неиспользуемые пакеты из файлов
go mod tidy -v

//Использование прокси для установки модуля
export {http,https,ftp}_proxy="http://10.72.0.1:3128"
go get github.com/go-sql-driver/mysql









Запросы:
http://127.0.0.1:4000
http://127.0.0.1:4000/snippet
http://127.0.0.1:4000/snippet?id=1
http://127.0.0.1:4000/snippet/create
http://127.0.0.1:4000/static
















Конвертирование типов из MySQL в Go
Под капотом метода rows.Scan(), драйвер базы данных автоматически преобразует MySQL типы в типы языка программирования Go:

CHAR, VARCHAR и TEXT соответствуют типу string;
BOOLEAN соответствует bool;
INT соответствует int;
BIGINT соответствует int64;
DECIMAL и NUMERIC соответствуют float;
TIME, DATE и TIMESTAMP соответствуют time.Time.


Пока вы используете пакет database/sql, ваш код будет работать с любым типом SQL базы данных — будь то MySQL, PostgreSQL, SQLite или что-то еще. 
Это означает, что ваше приложение не зависит от определённой базы данных, которую вы используете в настоящее время. 
Следовательно, вы можете менять базу данных в будущем, не переписывая весь код (за исключением специфических особенностей драйвера и SQL запросов).




//offtop 
//install mariadb archlinux
pacman -Sy mariadb 
mariadb-install-db --user=mysql --basedir=/usr --datadir=/var/lib/mysql
systemctl start   mariadb
systemctl status mariadb
systemctl enable  mariadb
mariadb-secure-installation

Создание базы:
mariadb

-- Создание новой UTF-8 базы данных `snippetbox`
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
-- Переключение на использование базы данных `snippetbox`
USE snippetbox;

-- Создание таблицы `snippets`
CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);
 
-- Добавление индекса для созданного столбца
CREATE INDEX idx_snippets_created ON snippets(created);

-- Добавляем несколько тестовых записей
INSERT INTO snippets (title, content, created, expires) VALUES (
    'Не имей сто рублей',
    'Не имей сто рублей,\nа имей сто друзей.',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);
 
INSERT INTO snippets (title, content, created, expires) VALUES (
    'Лучше один раз увидеть',
    'Лучше один раз увидеть,\nчем сто раз услышать.',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);
 
INSERT INTO snippets (title, content, created, expires) VALUES (
    'Не откладывай на завтра',
    'Не откладывай на завтра,\nчто можешь сделать сегодня.',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);

Создание нового пользователя
CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT, UPDATE ON snippetbox.* TO 'web'@'localhost';
 
-- Важно: Не забудьте заменить 'pass' на свой пароль, иначе это и будет паролем.
ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';


Тестирование нового пользователя базы данных
mysql -D snippetbox -u web -p
SELECT id, title, expires FROM snippets;


