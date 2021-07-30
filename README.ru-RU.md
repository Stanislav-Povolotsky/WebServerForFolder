﻿# WebServerForFolder

Простой Web-сервер, который "расшаривает" текущую папку через HTTP-протокол
Формат:
	WebServerForFolder [<port> [<folder>]]
где
	<port>   - HTTP-порт, по умолчанию 80
	<folder> - папка, которую надо расшарить (по умолчанию текущий каталог)

Примеры:
	# Расшарить текущий каталог. Web-сервер будет работать по адресу :80
	WebServerForFolder

	# Расшарить каталог c:\Temp, Web-сервер будет работать по адресу :8080
	WebServerForFolder 8080 c:\Temp