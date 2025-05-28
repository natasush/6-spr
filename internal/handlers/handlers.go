package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Читаем содержимое файла index.html
	filePath := "index.html"

	content, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Ошибка при чтении файла", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Отправляем содержимое файла в ответ
	w.Write(content)
}

func HandlerUpload(w http.ResponseWriter, r *http.Request) {
	// 1 парсим html-форму из файла index.html

	if r.Method == http.MethodPost || r.Method == http.MethodGet {
		//Ограничиваем размер загружаемого файла
		err := r.ParseMultipartForm(10 << 20) // 10 MB
		if err != nil {
			fmt.Printf("Ошибка при разборе формы: %v", err)
			http.Error(w, "Ошибка при разборе формы", http.StatusBadRequest)
			return
		}

		//получаем файл из формы
		file, _, err := r.FormFile("myFile")
		if err != nil {
			fmt.Printf("Ошибка при при полученни файла: %v", err)
			http.Error(w, "Ошибка при получении файла", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// 3 Читаем содержимое файла
		data, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "Ошибка при чтении файла", http.StatusInternalServerError)
			return
		}
		// 4 передаем данные из файла в функцию из service
		conv, err := service.Convert(string(data))
		if err != nil {

			log.Printf("error converting:%v", err)
			http.Error(w, "Ошибка при конвертации", http.StatusInternalServerError)
			return
		}
		// 5 6 созд. локальный файл и записываем результат конвертации строки
		fileName := time.Now().UTC().Format("2006-01-02T15-04-05") + filepath.Ext("output.txt")
		err = os.WriteFile(fileName, []byte(conv), 0644)
		if err != nil {
			log.Printf("error writing file: %v", err)
			http.Error(w, "Ошибка при записи в файл", http.StatusInternalServerError)
			return
		}

		//7 возв. результат конвертации строки
		w.Write([]byte(conv))
	} else {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
	}
}
