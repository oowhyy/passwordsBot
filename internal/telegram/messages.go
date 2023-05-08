package telegram

const (
	commandStart  = "start"
	commandSet    = "set"
	commandGet    = "get"
	commandDelete = "del"

	msgStart = `Привет!
	Я умею хранить твои пароли для различных сервисов
	/set - записать новый сервис и пароль
	/get - получить пароль для сервиса
	/del - удалить пароль для сервиса
	Время жизни пароля: %s
	`

	msgNewSet = "Введи название сервиса и пароль через пробел"
	msgNewGet = "Введи название сервиса"
	msgNewDel = "Введи название сервиса"

	msgBadSet = "Неверный формат аргументов. Отменяю запись"
	msgBadGet = "Пароль для сервиса %s не найден"
	msgBadDel = "Пароль для сервиса %s не найден или уже был удален"

	msgOKSet = "Установлен пароль %s\nдля сервиса %s"
	msgOKDel = "Пароль для сервиса %s удален"

	msgErrDel     = "Ошибка базы данных"
	msgErrGet     = "Ошибка базы данных"
	msgErrSet     = "Ошибка базы данных"
	msgErrUnknown = "Что-то пошло не так..."

	msgUnknownCommand = "Я не знаю такую команду"
)
