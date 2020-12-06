# CLI
Интерфейс командной строки собирается командой `go build` из папки sky/cli.
В результате сборки появляется исполняемый файл `cli`.

Настройки приложения передаются через конфигурационный файл, путь до которого указывается с флагом `--config`.
По умолчанию приложение пытается найти файл `$HOME/.sky_cli.yaml`.

Конфигурация представляет собой yaml-файл со следующими ключами:
* token - токен пользователя
* url - адрес системы. По умолчанию: "http://localhost:4000"

Всю основную информацию о командах и параметрах можно получить, указав флаг `--help`.