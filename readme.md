# Контрольная работа. Алгоритмы и Go Routines

## Описание проекта
Вариант 3
Этот проект включает в себя два алгоритма и задачи, решаемые с использованием Go и горутин. В проекте реализованы следующие функциональные части:

1. Алгоритмы и структуры данных: Дан массив строк. Необходимо найти самую длинную строку, являющуюся префиксом для всех остальных строк в массиве. Опишите ваш алгоритм и приведите пример его работы для массива ["flower","flow","flight"].
2. Go Routines и каналы: Создайте две горутины. Первая генерирует случайные целые числа от 1 до 100 и отправляет их в канал. Вторая горутина получает числа из канала и вычисляет их квадратный корень. Если квадратный корень является целым числом, горутина выводит это число и его корень на экран. Обработайте 20 сгенерированных чисел.

## Структура проекта
```
cr1/
│
├── Makefile                # Makefile для автоматизации сборки, тестирования и запуска
├── README.md               # Описание проекта
├── LICENSE.md              # Лицензия
├── go.mod                  # Модуль Go
├── go.sum                  # Контрольные суммы зависимостей Go
│
├── src/                    # Каталог с исходным кодом
│   ├── task1/              # Задача 1
│   │   ├── task1.go        # Основной код задачи 1
│   │   └── task1_test.go   
│   │
│   ├── task2/              # Задача 2
│   │   └── task2.go        # Основной код задачи 2 (горутин и каналов)
│   │
│   ├── test/              # Тесты
│   │   └── task1_test.go   # Тесты для задачи 1
|   │
│   └── main.go             # Главный исполняемый файл, который импортирует пакеты task1 и task2
│
└── bin/                    # Каталог для скомпилированных файлов
```

## Установка
Для того чтобы запустить проект на своей машине, выполните следующие шаги:

1. Клонируйте репозиторий
```
git clone https://github.com/Arxonait/go-cr1.git
cd go-cr1
```

2. Установите зависимости
Проект использует модуль Go для управления зависимостями. Для установки всех необходимых зависимостей выполните команду:
```
go mod tidy
```

3. Сборка и запуск
Для сборки и запуска проекта используйте команду make.

*Сборка и запуск*
```
make all
```
Эта команда скомпилирует проект и запустит его.

*Запуск тестов*
```
make test
```
Эта команда выполнит все тесты проекта.

## Лицензия
Этот проект лицензируется под лицензией GPLv3. Подробности смотрите в файле LICENSE.md.

## Автор
Мякишев Никита