# conf
## RU
Предоставляет доступ к параметрам конфигурации.
Инициализируется `Source` и `EventSource` с указанием их приоритетности.
Данные полученные из источников мержатся в итоговый конфиг и возвращаются приложению, в виде `map[string]string`

Источники:
* флаги 
* файлы
* env

