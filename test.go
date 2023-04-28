package main

// import (
// 	"database/sql"
// 	"fmt"
// 	// "challange-8/controllers"

// 	_ "github.com/lib/pq"
// )

// const (
// 	host		= "localhost"
// 	port		= 5432
// 	user		= "postgres"
// 	password 	= "123"
// 	dbname		= "book-api"
// )

// var (
// 	db *sql.DB
// 	err error
// )

// func main () {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// 	db, err = sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Berhasil konek ke database")

// 	// AddBook()
// 	// GetBooks()
// 	// GetBookId()
// 	// UpdateBook()
// 	// DeleteBook()
// }


// type Book struct {
// 	ID int `json: "book_id"`
// 	Title string `json: "title"`
// 	Author string `json: "author"`
// 	Description string `json: "description"`
// }


// func AddBook() {
// 	var book = Book{}

// 	sqlStatement := `
// 	INSERT INTO books (title, author, description)
// 	VALUES ($1, $2, $3)
// 	Returning *
// 	`

// 	err = db.QueryRow(sqlStatement, "Buku Golang", "Ardiansyah", "Buku tentang bahasa pemrograman").
// 	Scan(&book.ID, &book.Title, &book.Author, &book.Description)

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("Data Buku terbaru : %+v\n", book)
// }

// func GetBooks() {
// 	var results = []Book{}

// 	sqlStatement := `SELECT * from books`

// 	rows, err := db.Query(sqlStatement)

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var book = Book{}

// 		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description)

// 		if err != nil {
// 			panic(err)
// 		}

// 		results = append(results, book)
// 	}

// 	fmt.Println("Data Buku: ", results)
// }

// // func GetBookId() {

// // }

// func UpdateBook() {

// 	sqlStatement := `UPDATE books set title = $2, author = $3, description = $4 where id = $1;`

// 	res, err := db.Exec(sqlStatement, 1, "Buku PHP", "Fadil Roni", "Buku Tentang PHP")

// 	if err != nil {
// 		panic(err)
// 	}

// 	count, err := res.RowsAffected()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Buku berhasil di Update: ", count)
// }

// func DeleteBook() {
// 	sqlStatement := `DELETE from books where id = $1;`

// 	res, err := db.Exec(sqlStatement, 1)
// 	if err != nil{
// 		panic(err)
// 	}

// 	count, err := res.RowsAffected()
// 	if err != nil{
// 		panic(err)
// 	}

// 	fmt.Println("Buku yang dihapus: ", count)
// }