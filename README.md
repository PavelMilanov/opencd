# **opencd** - cli-утилита для CD (continuous delivery) кода из вашего удаленного репозитория

___

## требования и зависимости

- ОС на базе дистрибутивов Linux (Debian, Ubuntu);
- установленные git, docker;
- наличие файла docker-compose.* в вашем проекте;
- доступ к удаленному репозиторию. (если репозиторий приватный - предварительно установить ssh-ключи для авторизации в репозитории).

## для чего нужен

- осуществляет непрерывную достаку кода из вашего удаленного репозитория с помощью файла конфигурации [opencd.yaml](opencd.yaml.template);
- собирает и перезапускает docker-контейнеры;
- следит за состоянием образов и контейнеров, удаляя ненеужные.

## как работает

- анализирует изменения удаленного и локального репозитория с помощью git;
- анализирует файл конфигурации docker-compose;
- автоматизирует сборку и обновление ваших компонентов docker-compose на основе изменений git;
- удаляет неиспользуемые контейнеры и образы docker.

## установка

- выполнить последовательно следующие команды:

```bash
wget https://github.com/PavelMilanov/opencd/releases/download/v0.2.4/install.sh
sudo bash install ./install.sh
```

- проверка установленной версии opencd:

```bash
opencd version

opencd version 0.2.4
git version 2.45.0
Docker Compose version v2.27.0-desktop.2
```

- добавить в корень проекта файл [opencd.yaml](opencd.yaml.template)

## использование

- обновить проект:

```bash
opencd deploy -e [environment] -s [stage]
```

- помощь при обновлении проекта:

```bash
opencd deploy -h
```

- посмотреть версию:

```bash
opencd version
```

- посмотреть структуру команд:

```bash
opencd help
```

## справочная информация

- краткое описание структуры [opencd.yaml](docs/opencd.yaml.manual)
