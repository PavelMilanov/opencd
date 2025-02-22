package main

const MENU_TEXT = `
opencd deploy -s=[merge,docker] -e=[opencd.environments.name]	- произвести обновление репозитория. Подробнее - opencd deploy -h;
opencd version							- показать версию;

`

var STEPS_WITH_CACHE = []string{"[1/4] Анализ изменений проекта", "[2/4] Сборка проекта", "[3/4] Обновление проекта", "[4/4] Очистка кеша docker"}
var STEPS_WITHOUT_CACHE = []string{"[1/3] Анализ изменений проекта", "[2/3] Сборка проекта", "[3/3] Обновление проекта"}
