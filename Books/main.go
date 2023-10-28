package main

import (
	"log"
	"net/http"
	"os"
	"fmt"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/lowry/book-crud/book"
	"github.com/lowry/book-crud/storage"

	"gorm.io/gorm"
)

type Book struct {
	Title string
	Author string
	Rating int
}

type Repository struct{
	DB *gorm.DB
}


func (r *Repository) CreateBook(c *fiber.Ctx)error {
	book := book.Books{}
	err := c.BodyParser(&book)
	if err != nil{
		c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message":"request failed"})
	}
	err = r.DB.Create(&book).Error
	if err != nil{
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message":"Book not created"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message":"Book added"})
	return nil
}


func (r *Repository) GetBook(c *fiber.Ctx) error{
	bookModel := &book.Books{} //returns a single book based on ID
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message":"Book don't exist"})
		return nil
	}
	err := r.DB.Where("id= ?", id).First(bookModel).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message":"could not return book"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "book is fetched successfully", "data": bookModel})
	return nil
}


func (r *Repository) GetAllBooks(c *fiber.Ctx) error{
	bookModel := &[]book.Books{}  //initializes a pointer to slice of books
	err := r.DB.Find(bookModel).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message":"Could not get books"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message":"books fetched successfully", "data":bookModel}) //returns all books in json format
	fmt.Println(bookModel)

	return nil
}

// func (r *Repository) DeleteBook(c *fiber.Ctx){
// 	db := database.DBConn
// 	id := c.Params("id")
// 	var book Book
// 	db.Delete(&book, id)
// 	c.JSON(book)
// }

// func UpdateBook(c *fiber.Ctx){
// 	db := database.DBConn
// 	id := c.Params("id")
// 	var book Book
// 	c.JSON(book)
// }


// Request and response endpoints for a CRUD API
func (r *Repository) setupResources(app *fiber.App){
	app.Group("/api/v1")
	app.Get("/books", r.GetAllBooks)
	app.Get("/book/:id", r.GetBook)
	app.Post("/create_book", r.CreateBook)
	// app.Delete("/book/delete/:id", r.DeleteBook)
	// app.Put("/api/v1/book/:id", r.UpdateBook)
}



func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User: os.Getenv("DB_USER"),
		SSLMode: os.Getenv("DB_SSLMODE"),
		DBName: os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)
	if err != nil{
		log.Fatal("could not load the database")
	}

	err = book.Migration(db)
	if err != nil {
		log.Fatal("Could not migrate DB")
	}

	r := Repository{
		DB: db,
	}

	app :=fiber.New()
	r.setupResources(app)
	app.Listen(":3000")
}
