package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"todoList/db"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "todo",
		Short: "todoList stores things to do",
		Long:  "todoList stores things to do, it can be used with cli",
	}
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add something to do",
		Run:   add,
	}
	rmCmd = &cobra.Command{
		Use:   "rm",
		Short: "Remove something to do",
		Run:   rm,
	}
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all the things to do",
		Run:   list,
	}
)

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(addCmd, rmCmd, listCmd)
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// add calls the AddTodo function from db package with the string of the todo 
func add(cmd *cobra.Command, args []string) {
	fmt.Println("Add: " + strings.Join(args, " "))
	db.AddTodo(strings.Join(args, " "))
}

// rm calls the RemoveTodo function from db package with the id of the todo to remove
func rm(cmd *cobra.Command, args []string) {
	fmt.Println("Remove: " + strings.Join(args, " "))
	for _, id := range args {
		db.RemoveTodo(id)
	}
}

// list calls the ListTodo function from db package and print the result in a table
func list(cmd *cobra.Command, args []string) {
	var todos []db.Todo
	todos = db.ListTodo()

	todoTable := [][]string{}

	for _, todo := range todos {
		todoTable = append(todoTable, []string{todo.ID, todo.Title, todo.Status})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Description", "Status"})

	for _, v := range todoTable {
		table.Append(v)
	}
	table.Render()
}
