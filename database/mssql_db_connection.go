package database

import (
   "bufio"
   "database/sql"
   "fmt"
   _ "github.com/denisenkom/go-mssqldb"
   "github.com/joho/godotenv"
   "mssql-go-cli/database"
   "os"
)

func MSSQLDBConnectionExample() {
    envErr := godotenv.Load(); if envErr != nil {
       fmt.Printf("Error loading credentials: %v", envErr)
    }
  //TODO NOTE: FILL THE DB VARIABLES TO RUN THE CODE
var (
   password = "YOUR_DB_PASSWORD" //os.Getenv("MSSQL_DB_PASSWORD")
   user = "YOUR_DB_USERNAME" //os.Getenv("MSSQL_DB_USER")
   port = "YOUR_DB_PORT" //os.Getenv("MSSQL_DB_PORT")
   database = "YOUR_DATABASE_NAME" //os.Getenv("MSSQL_DB_DATABASE")
)

connectionString := fmt.Sprintf("user id=%s;password=%s;port=%s;database=%s", user, password, port, database)

   sqlObj, connectionError := sql.Open("mssql", database.ConnectionString); if connectionError != nil {
      fmt.Println(fmt.Errorf("error opening database: %v", connectionError))
   }

   data := database.Database{
      SqlDb: sqlObj,
   }

   fmt.Println("-> Welcome to Reminders Console App, built using Golang and Microsoft SQL Server")
   fmt.Println("-> Select a numeric option; \n [1] Create a new Reminder \n [2] Get a reminder \n [3] Delete a reminder")

   consoleReader := bufio.NewScanner(os.Stdin)
   consoleReader.Scan()
   userChoice := consoleReader.Text()

   switch userChoice {
   case "1":
      var (
         titleInput,
         descriptionInput,
         aliasInput string
      )
      fmt.Println("You are about to create a new reminder. Please provide the following details:")

      fmt.Println("-> What is the title of your reminder?")
      consoleReader.Scan()
      titleInput = consoleReader.Text()

      fmt.Println("-> What is the description of your reminder?")
      consoleReader.Scan()
      descriptionInput = consoleReader.Text()

      fmt.Println("-> What is an alias of your reminder? [ An alias will be used to retrieve your reminder ]")
      consoleReader.Scan()
      aliasInput = consoleReader.Text()

      data.CreateReminder(titleInput, descriptionInput, aliasInput)

   case "2":
      fmt.Println("-> Please provide an alias for your reminder:")
      consoleReader.Scan()
      aliasInput := consoleReader.Text()

      data.RetrieveReminder(aliasInput)

   case "3":
      fmt.Println("-> Please provide the alias for the reminder you want to delete:")
      consoleReader.Scan()
      deleteAlias := consoleReader.Text()

      data.DeleteReminder(deleteAlias)

   default:
      fmt.Printf("-> Option: %v is not a valid numeric option. Try 1 , 2 , 3", userChoice)
   }
}
