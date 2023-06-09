# Задание
Приложение представляет собой API по работе с сущностью User, где хранилищем выступает файл json.

Ограничения:
- Хранилищем должен оставаться файл в json формате.
- Структура пользователя не должна быть уменьшена.
- Приложение не должно потерять существующую функциональность. 

Мы понимаем, что пределу совершенства нет и ожидаем, что объем рефакторинга вы определяете на свое усмотрение.  

После того как вы выполните задание, вы так же можете написать, как бы улучшили проект в перспективе текстом.

Что следует знать:
- В будущем это приложение ожидает увеличение количества функций и сущностей. 
- Вопрос авторизации умышленно опущен, о нем не стоит беспокоиться.
- API еще не выпущено, вы в праве скорректировать интерфейс / форматы ответов.

# Предложение по улучшиению 

1. Добавление Docker контейнеров для проекта, что позволит упростить процесс установки и запуска проекта на разных платформах и в разных средах.(Инструмент, который помогает генерировать Dockerfile с меньшим объемом образа - https://github.com/zeromicro/go-zero)
2. Добавление реляционной БД. (Реализация защиты от SQL-инъекций).
3. Добавление регулярок на валидацию входящих данных.
4. В случае реализации авторизации необходимо добавить jwt. Также следует реализовать с помощью jwt токена middleware для аутентификации.
5. Реализовать изменение email.
6. Покрыть код тестами.
7. Добавление документации, которая описывает архитектуру, API, используемые технологии, инструкции по развертыванию и использованию проекта.
8. Добавление комментариев к коду.
9. Доработка обработки ошибок и распределение логирования по всему коду.