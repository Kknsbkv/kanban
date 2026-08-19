package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Task представляет структуру задачи
type Task struct {
	ID     int
	Title  string
	Status string
}

// ValidateTitle проверяет, что название не пустое и не содержит ';'
func ValidateTitle(title string) (string, error) {
	trimmed := strings.TrimSpace(title)
	if trimmed == "" {
		return "", fmt.Errorf("Error: название не может быть пустым.")
	}
	if strings.Contains(trimmed, ";") {
		return "", fmt.Errorf("Error: название не должно содержать символ ';'")
	}
	return trimmed, nil
}

// ParseID переводит строку с ID в число и проверяет на ошибки
func ParseID(input string) (int, error) {
	trimmed := strings.TrimSpace(input)
	id, err := strconv.Atoi(trimmed)
	if err != nil {
		return 0, fmt.Errorf("Error: ID должен быть числом.")
	}
	return id, nil
}

// FindTaskIndex ищет индекс задачи по ID или возвращает ошибку, если задача не найдена
func FindTaskIndex(tasks []Task, id int) (int, error) {
	for i, task := range tasks {
		if task.ID == id {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Error: задача с ID %d не найдена.", id)
}

// GetNextStatus (Бонус 3) переводит статус по цепочке: TODO -> IN_PROGRESS -> DONE
func GetNextStatus(currentStatus string) (string, error) {
	switch currentStatus {
	case "TODO":
		return "IN_PROGRESS", nil
	case "IN_PROGRESS":
		return "DONE", nil
	case "DONE":
		return "", fmt.Errorf("Error: задача уже выполнена.")
	default:
		return "", fmt.Errorf("Error: неизвестный статус задачи.")
	}
}

// ValidateMenuChoice проверяет выбор пункта меню
func ValidateMenuChoice(input string) (int, error) {
	trimmed := strings.TrimSpace(input)
	choice, err := strconv.Atoi(trimmed)
	if err != nil || choice < 1 || choice > 7 {
		return 0, fmt.Errorf("Неверный ввод. Введите число от 1 до 7.")
	}
	return choice, nil
}

// PrintHelp выводит полную справку при флаге --help или -h
func PrintHelp() {
	fmt.Println("Kanban CLI — консольный менеджер задач в интерактивном режиме.")
	fmt.Println()
	fmt.Println("Использование:")
	fmt.Println("  go run .         Запуск интерактивного меню")
	fmt.Println("  go run . --help  Показать эту справку")
	fmt.Println()
	fmt.Println("Доступные команды в меню:")
	fmt.Println("  1. Создать таску          — запрашивает название и добавляет новую задачу (TODO)")
	fmt.Println("  2. Сменить статус         — переводит статус: TODO -> IN_PROGRESS -> DONE")
	fmt.Println("  3. Изменить название      — переименовывает задачу по ID")
	fmt.Println("  4. Удалить таску          — удаляет задачу из списка по ID")
	fmt.Println("  5. Все таски              — выводит текущую доску задач в виде ASCII-таблицы")
	fmt.Println("  6. Фильтр по статусу      — показывает задачи только с выбранным статусом")
	fmt.Println("  7. Завершить работу       — сохраняет состояние и закрывает программу")
}
