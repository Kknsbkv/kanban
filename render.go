package main

import (
	"fmt"
	"strings"
)

// ANSI-коды для цвета
const (
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
)

// PrintBoard выводит список всех задач в виде ASCII-таблицы
func PrintBoard(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println(Yellow + "\nСписок задач пуст." + Reset)
		return
	}

	fmt.Println("\n" + Bold + Cyan + "========================================================" + Reset)
	fmt.Printf("%-5s | %-30s | %-15s\n", "ID", "НАЗВАНИЕ", "СТАТУС")
	fmt.Println("--------------------------------------------------------")

	for _, task := range tasks {
		statusColor := Yellow
		if task.Status == "IN_PROGRESS" {
			statusColor = Blue
		} else if task.Status == "DONE" {
			statusColor = Green
		}

		fmt.Printf("%-5d | %-30s | %s%-15s%s\n", task.ID, task.Title, statusColor, task.Status, Reset)
	}
	fmt.Println(Bold + Cyan + "========================================================" + Reset)
}

// PrintFilteredTasks выводит задачи только с выбранным статусом
func PrintFilteredTasks(tasks []Task, targetStatus string) {
	var filtered []Task
	for _, task := range tasks {
		if strings.EqualFold(task.Status, targetStatus) {
			filtered = append(filtered, task)
		}
	}

	if len(filtered) == 0 {
		fmt.Printf(Yellow+"\nЗадач со статусом %s не найдено.\n"+Reset, targetStatus)
		return
	}

	PrintBoard(filtered)
}
