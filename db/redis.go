package db

import (
	"log"
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"
)

type Todo struct {
	ID    string
	Title string `redis:"title"`
	//TODO	CreationDate string `redis:"creationDate"`
	Status string `redis:"status"`
}

// checkError prints the error and stop the app 
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// AddTodo adds a new member of the todo-id-set and set hash fields with the given string
func AddTodo(todo string) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	checkError(err)
	defer conn.Close()

	id, err := redis.Int(conn.Do(
		"INCR",
		"todoIdCounter",
	))
	checkError(err)
	todoId := "todo:" + strconv.Itoa(int(id))

	_, err = conn.Do(
		"SADD",
		"todo-id-set",
		todoId,
	)
	checkError(err)

	_, err = conn.Do(
		"HMSET",
		todoId,
		"title",
		todo,
		"status",
		"pending",
	)
	checkError(err)
}

// RemoveTodo removes the todo's hash linked with the ginven id
func RemoveTodo(id string) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	checkError(err)
	defer conn.Close()

	n, err := redis.Int(conn.Do("DEL", "todo:"+id))

	if n > 0 {
		_, err = conn.Do(
			"SREM",
			"todo-id-set",
			"todo:"+id)
	}
	checkError(err)
}

// ListTodo returns the todos stored in the redis hashes
func ListTodo() []Todo {
	conn, err := redis.Dial("tcp", "localhost:6379")
	checkError(err)
	defer conn.Close()

	todoHasNames, err := redis.Strings(conn.Do("SMEMBERS", "todo-id-set"))
	if err != nil {
		log.Fatal("failed to get todo IDs", err)
	}

	todos := []Todo{}
	for _, todoHashName := range todoHasNames {
		id := strings.Split(todoHashName, ":")[1]

		todoMap, err := redis.StringMap(conn.Do("HGETALL", todoHashName))
		if err != nil {
			log.Fatalf("failed to get todo from %s - %v\n", todoHashName, err)
		}

		var todo Todo
		todo = Todo{id, todoMap["title"], todoMap["status"]}
		todos = append(todos, todo)
	}

	return todos
}
