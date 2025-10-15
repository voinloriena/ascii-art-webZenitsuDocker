# ASCII Art Web (ascii-art-webZenitsu)

Простой веб-сервис на Go для генерации ASCII-арт текста с использованием различных баннеров.  
Проект полностью контейнеризован с помощью Docker.

---

## 🔹 Особенности

- Веб-сервер на Go
- Поддержка HTML-форм для ввода текста и выбора баннера
- Статические файлы (CSS, JS, изображения)
- Папка `back` для хранения изображений и ресурсов
- Полностью Dockerized (Dockerfile + контейнер)
- Совместимость с современными версиями Go (≥1.25)

---

## 📂 Структура проекта

ascii-art-webZenitsu/
├── Dockerfile
├── main.go
├── go.mod
├── templates/
│ └── index.html
│ └── error.html
├── static/
│ └── css/
│ └── js/
├── back/
│ └── 1625588496_1040737.webp
├── banners/
├── asciigo/
└── README.md
---

## ⚙️ Требования

- Docker ≥20.10
- Go ≥1.25 (для локальной сборки, опционально)
- Браузер для доступа к веб-серверу

---
## 🛠️ Установка и запуск через Docker

1. Перейти в папку проекта:
   
2.Сборка Docker-образа:

docker build -t ascii-art-webzenitsu:1.0 .

3.Запуск контейнера:

docker run -d --name ascii-art-webzenitsu -p 8080:8080 ascii-art-webzenitsu:1.0
4.Открыть браузер и перейти:

http://localhost:8080

5.Проверка картинки:

http://localhost:8080/back/1625588496_1040737.webp
