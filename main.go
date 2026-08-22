package main
import (
 "bufio"
 "fmt"
 "os"
 "strings"
)
func main() {
 if len(os.Args) > 1 {
  flag := os.Args[1]
  if flag == "--help" || flag == "-h" {
   PrintHelp()
   return
  }
 }
scanner := bufio.NewScanner(os.Stdin)
tasks := []Task{
 {ID: 1, Title: "Изучить Go", Status: "DONE"},
 {ID: 2, Title: "Создать валидатор", Status: "IN_PROGRESS"},
 {ID: 3, Title: "Написать меню", Status: "TODO"},
}
nextID := 4

for {
 fmt.Println("\n--- KANBAN MENU ---")
 fmt.Println("1. Создать таску")
 fmt.Println("2. Сменить статус")
 fmt.Println("3. Изменить название")
 fmt.Println("4. Удалить таску")
 fmt.Println("5. Все таски")
 fmt.Println("6. Фильтр по статусу")
 fmt.Println("7. Завершить работу")
 fmt.Print("Выберите действие (1-7): ")

 if !scanner.Scan() {
  break
 }
 input := scanner.Text()

 choice, err := ValidateMenuChoice(input)
 if err != nil {
  fmt.Println(Red + err.Error() + Reset)
  continue
 }

 switch choice {
 case 1:
  fmt.Print("Введите название задачи: ")
  scanner.Scan()
  titleInput := scanner.Text()
  title, err := ValidateTitle(titleInput)
  if err != nil {
   fmt.Println(Red + err.Error() + Reset)
   continue
  }
  tasks = append(tasks, Task{ID: nextID, Title: title, Status: "TODO"})
  fmt.Printf(Green+"Задача '%s' успешно создана (ID: %d)!\n"+Reset, title, nextID)
  nextID++

 case 2:
  fmt.Print("Введите ID задачи для смены статуса: ")
  scanner.Scan()
  id, err := ParseID(scanner.Text())
  if err != nil {
   fmt.Println(Red + err.Error() + Reset)
   continue
  }
  idx, err := FindTaskIndex(tasks, id)
  if err != nil {
   fmt.Println(Red + err.Error() + Reset)
   continue
  }
  newStatus, err := GetNextStatus(tasks[idx].Status)
  if err != nil {
   fmt.Println(Red + err.Error() + Reset)
   continue
  }
  tasks[idx].Status = newStatus
  fmt.Printf(Green+"Статус задачи #%d изменен на %s\n"+Reset, id, newStatus)

 case 3:
  fmt.Print("Введите ID задачи: ")
  scanner.Scan()
  id, err := ParseID(scanner.Text())
  if err != nil {
   fmt.Println(Red + err.Error() + Reset)
   continue
  }
  idx, err := FindTaskIndex(tasks, id)
  if err != nil {
   fmt.Println(Red + err.Error() + Reset)
   continue
  }
  fmt.Print("Введите новое название: ")
  scanner.Scan()
  newTitle, err := ValidateTitle(scanner.Text())
  if err != nil {
   fmt.Println(Red + err.Error() + Reset)
   continue
  }
  tasks[idx].Title = newTitle
  fmt.Println(Green + "Название успешно изменено!" + Reset)

 case 4:
  fmt.Print("Введите ID задачи для удаления: ")
  scanner.Scan()
  id, err := ParseID(scanner.Text())
  if err != nil {
   fmt.Println(Red + err.Error() + Reset)
   continue
  }
  idx, err := FindTaskIndex(tasks, id)
  if err != nil {
   fmt.Println(Red + err.Error() + Reset)
   continue
  }
  tasks = append(tasks[:idx], tasks[idx+1:]...)
  fmt.Println(Green + "Задача успешно удалена!" + Reset)

 case 5:
  PrintBoard(tasks)

 case 6:
  fmt.Print("Введите статус для фильтрации (TODO / IN_PROGRESS / DONE): ")
  scanner.Scan()
  status := strings.TrimSpace(scanner.Text())
  PrintFilteredTasks(tasks, status)

 case 7:
  fmt.Println(Green + "Работа завершена. До свидания!" + Reset)
  return
 }
}
}