# conf
## RU
Предоставляет доступ к параметрам конфигурации.
Инициализируется `Source` и `EventSource` с указанием их приоритетности.
Данные полученные из источников мержатся в итоговый конфиг и возвращаются приложению, в виде `map[string]string`

Источники:
* флаги 
* файлы
* env


ROADMAD:
v0.1 - jsonfile source, yaml source, env source, flags source
v0.2 - подписка на изменение sources
v0.3 - merge services configs 
v0.4 - options