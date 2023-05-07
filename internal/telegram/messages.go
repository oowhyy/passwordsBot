package telegram

const (
	commandStart  = "start"
	commandSet    = "set"
	commandGet    = "get"
	commandDelete = "del"

	msgStart = "TODO"

	msgNewSet = "Введи название сервиса и пароль через пробел"
	msgNewGet = "Введи название сервиса"
	msgNewDel = "Введи название сервиса"

	msgBadSet = "Неверный формат аргументов. Отменяю запись"
	msgBadGet = "Пароль для сервиса %s не найден"
	msgBadDel = "Пароль для сервиса %s не найден или уже был удален"

	msgOKSet = "Установлен пароль %s\nдля сервиса %s"
	msgOKDel = "Пароль для сервиса %s удален"

	msgUnknownCommand = "Я не знаю такую команду"
)
