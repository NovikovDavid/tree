# TREE

tree - это инструмент для вывода дерева директорий


## Run locally
Для работы потребуется версия Go >= **1.21.1**
### Сборка 
Собрать можно командой ```make build```

### Запуск линтера
Осуществляется командой ```make lint```

### Тестирование
Запустить тестирование можно командой ```make it```



### Запуск в CLI
```sh
go run cmd/cli/main.go -f <path>
```
Ключ ```-f``` выводит более подробную информацию о директориях и их файлах, без ключа выводит просто директории

### С ключом ```-f```
```
├───project
│   ├───file.txt (19)
│   └───gopher.png (70372)
├───static
│   ├───a_lorem
│   │   ├───dolor.txt (empty)
│   │   ├───gopher.png (70372)
│   │   └───ipsum
│   │       └───gopher.png (70372)
│   ├───css
│   │   └───body.css (28)
│   ├───empty.txt (empty)
│   ├───html
│   │   └───index.html (57)
│   ├───js
│   │   └───site.js (10)
│   └───z_lorem
│       ├───dolor.txt (empty)
│       ├───gopher.png (70372)
│       └───ipsum
│           └───gopher.png (70372)
├───zline
│   ├───empty.txt (empty)
│   └───lorem
│       ├───dolor.txt (empty)
│       ├───gopher.png (70372)
│       └───ipsum
│           └───gopher.png (70372)
└───zzfile.txt (empty)
```
### Без ключа
```
├───project
├───static
│   ├───a_lorem
│   │   └───ipsum
│   ├───css
│   ├───html
│   ├───js
│   └───z_lorem
│       └───ipsum
└───zline
    └───lorem
        └───ipsum
```