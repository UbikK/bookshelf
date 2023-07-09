package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"

	_ "modernc.org/sqlite"
)

// App struct
type App struct {
	ctx context.Context
}

type Book struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}
// Greet returns a greeting for the given name
// func (a *App) Greet(name string) string {
// 	return fmt.Sprintf("Hello %s, It's show time!", name)
// }

// func (a *App) SayHello() string {
// 	return fmt.Sprintf("Hello")
// }

func (a *App) CreateBook(name string) (int64, error){
	db, err := sql.Open("sqlite","./database.db")
	
if err != nil {
	log.Print(err.Error())
    return 0, err
}
	
	stmt, stmtError := db.Prepare("INSERT into book (id, name) values (?, ?)")
	if stmtError != nil {
		log.Print(stmtError.Error())
		return 0, stmtError
	}
	res, resErr := stmt.Exec(uuid.NewString(), name)
	if resErr != nil {
		log.Print(resErr.Error())
		return 0, resErr
	}
	id, idErr := res.LastInsertId()
	if idErr != nil {
		log.Print(idErr.Error())
		return 0, idErr
	}
	return id,nil
}

func (a *App) GetAllBooks() ([]Book, error) {
	db, err := sql.Open("sqlite","./database.db")
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	bookRows, listErr := db.Query("SELECT * from book")

	if listErr != nil {
		log.Print(listErr.Error())
		return nil, listErr
	}

	defer bookRows.Close()

	var bookList []Book

	for bookRows.Next() {
		
		var id string
		var name string
		bookErr := bookRows.Scan(&id, &name)

		if bookErr != nil {
			log.Print(bookErr.Error())
		}
		
		book := Book{Id:id, Name:name}
		bookList = append(bookList, book)
	}
	
	return bookList, nil
	
}