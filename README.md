# Телеграм-бот для управления паролями
* /set - записать новый сервис и пароль
* /get - получить пароль для сервиса
* /del - удалить пароль для сервиса
# Особенности
## Независимость пользователей
Приложение может поддерживать параллельные диалоги со многими пользователями. Каждый пользователь имеет доступ только к своим паролям
## Персистентность данных
Обеспечивается на уровне redis благодаря использованию снепшотов.
## Удалени
