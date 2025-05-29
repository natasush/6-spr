package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Читаем содержимое файла index.html и отправляем в ответ
	http.ServeFile(w, r, "./index.html")

}

func HandlerUpload(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// 1 парсим html-форму из файла index.html

	if r.Method != http.MethodPost {

		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}
	//Ограничиваем размер загружаемого файла
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		log.Printf("Ошибка при разборе формы: %v\n", err)
		http.Error(w, "Ошибка при разборе формы", http.StatusBadRequest)
		return
	}

	//2 получаем файл из формы
	file, _, err := r.FormFile("myFile")
	if err != nil {
		log.Printf("Ошибка при при полученни файла: %v\n", err)
		http.Error(w, "Ошибка при получении файла", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 3 Читаем содержимое файла
	data, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Ошибка при чтении файла: %v\n", err)
		http.Error(w, "Ошибка при чтении файла", http.StatusInternalServerError)
		return
	}
	// 4 передаем данные из файла в функцию из service
	conv, err := service.Convert(string(data))
	if err != nil {
		log.Printf("error converting:%v\n", err)
		http.Error(w, "Ошибка при конвертации", http.StatusInternalServerError)
		return
	}
	// 5 6 созд. локальный файл и записываем результат конвертации строки
	fileName := time.Now().UTC().Format("2006-01-02T15-04-05") + filepath.Ext("output.txt")
	err = os.WriteFile(fileName, []byte(conv), 0644)
	if err != nil {
		log.Printf("Ошибка при создании и записи файла:%v\n", err)
		http.Error(w, "Ошибка при создании и записи файла", http.StatusInternalServerError)
		return
	}

	//7 возв. результат конвертации строки
	_, err = w.Write([]byte(conv))
	if err != nil {
		log.Printf("Ошибка при выводе ответа:%v\n", err)
		http.Error(w, "Ошибка при выводе ответа", http.StatusInternalServerError)
		return

	}
}
