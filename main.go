package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Tasks struct {
	ID         string `json:"id"`
	TaskName   string `json:"task_name"`
	TaskDetail string `json:"task_detail"`
	Date       string `json:"data"`
}

var tasks []Tasks

func allTasks() {
	task := Tasks{
		ID:         "1",
		TaskName:   "New Project",
		TaskDetail: "You must lead the project",
		Date:       "02/02/22"}

	tasks = append(tasks, task)

	task1 := Tasks{
		ID:         "2",
		TaskName:   "New Project2",
		TaskDetail: "You must lead the project2",
		Date:       "03/02/22"}

	tasks = append(tasks, task1)

	task2 := Tasks{
		ID:         "3",
		TaskName:   "New Project2",
		TaskDetail: "You must lead the project2",
		Date:       "03/02/22"}

	tasks = append(tasks, task2)

	task3 := Tasks{
		ID:         "4",
		TaskName:   "New Project2",
		TaskDetail: "You must lead the project2",
		Date:       "03/02/22"}

	tasks = append(tasks, task3)

	fmt.Println(tasks)

}

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Endpoint is hit")

}
func gettask(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hallow There")

	//fmt.Fprintf(w, "gettask is hit")

	// for _, v := range mux.Vars(r) {
	// 	//fmt.Fprintf(w, v)
	// 	i, _ := strconv.Atoi(v)
	// 	json.NewEncoder(w).Encode(tasks[i])
	// }

	//taskId := mux.Vars(r)
	//fmt.Fprintf(w,)
	//flag := false
	// for i := 0; i < len(tasks); i++ {
	// 	if taskId["Id"] == tasks[i].ID {
	// 		json.NewEncoder(w).Encode(tasks[i])
	// 		flag = true
	// 		break
	// 	}

	// }

	// if flag == false {
	// 	json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	// }

}
func gettasks(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(w, "gettasks is hit")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)

}
func createTask(w http.ResponseWriter, r *http.Request) {
	
	err:=r.ParseForm()

	fmt.Fprintf(w, "create is hit")
	w.Header().Set("Content-Type", "application/json")
	var task Tasks
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = strconv.Itoa(rand.Intn(10000))
	current_time := time.Now().Format("01-02-2006")
	task.Date = current_time
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(tasks)

}
func deleteTask(w http.ResponseWriter, r *http.Request) {

	// for index, item := range tasks {

	// 	fmt.Fprint(w, "the id is", item.ID)
	// 	tasks = append(tasks[:index], tasks[index+1:]...)
	// 	var task Tasks
	// 	_ = json.NewDecoder(r.Body).Decode(&task)
	// 	task.Date = time.Now().Format("01-02-2006")
	// 	task.ID = params["id"]
	// 	tasks = append(tasks, task)
	// 	flag = true
	// 	json.NewEncoder(w).Encode(task)
	// 	return
	// }

	fmt.Fprintf(w, "delete is hit")

}
func updateTask(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "update is hit")

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	flag := false
	for index, item := range tasks {

		fmt.Fprint(w, "the id is", item.ID)
		tasks = append(tasks[:index], tasks[index+1:]...)
		var task Tasks
		_ = json.NewDecoder(r.Body).Decode(&task)
		task.Date = time.Now().Format("01-02-2006")
		task.ID = params["id"]
		tasks = append(tasks, task)
		flag = true
		json.NewEncoder(w).Encode(task)
		return
	}

	if flag == false {
		json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	}

}

func handleRoute() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/gettasks", gettasks).Methods("GET")
	router.HandleFunc("/gettask/{id}", gettask).Methods("GET")
	router.HandleFunc("/create", createTask).Methods("POST")
	router.HandleFunc("/delete{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/update", updateTask).Queries("id", "{id}").Methods("PUT")
	http.ListenAndServe(":8080", router)
}

func main() {

	allTasks()
	fmt.Println("Hello World")
	handleRoute()
}
