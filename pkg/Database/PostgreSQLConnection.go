package database

import (
	"BlogApplication/pkg/Code"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "bhumi"
	dbname   = "postgres"
)

type PostgresStorage struct {
	db *sql.DB
}

func SetUpPostgreSQL() (*PostgresStorage, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to Postgres!")
	return &PostgresStorage{db: db}, nil

}

func (ps *PostgresStorage) AddNewBlog(b Code.Blog) error {

	sqlStatement := `INSERT INTO blog (blog_title, blog_author, blog_category, blog_content) VALUES ($1, $2, $3, $4) RETURNING id`
	id := 0
	err := ps.db.QueryRow(sqlStatement, b.BlogTitle, b.BlogAuthor.AuthorName, b.BlogCategory.CategoryName, b.BlogContent).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
	return nil
}
func (ps *PostgresStorage) DeleteBlog(title string) error {
	sqlStatement := `DELETE FROM blog WHERE blog_title = $1;`
	_, err := ps.db.Exec(sqlStatement, title)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PostgresStorage) GetAllBlog() (Code.BlogList, error) {

	var blogList = Code.BlogList{}
	rows, err := ps.db.Query("SELECT blog_title, blog_author, blog_category, blog_content FROM blog")
	if err != nil {
		return blogList, err
	}
	defer rows.Close()
	for rows.Next() {
		var blog Code.Blog
		err := rows.Scan(&blog.BlogTitle, &blog.BlogAuthor.AuthorName, &blog.BlogCategory.CategoryName, &blog.BlogContent)
		if err != nil {

			return blogList, err
		}

		blogList.Blogs = append(blogList.Blogs, blog)

	}
	err = rows.Err()
	if err != nil {
		return blogList, err
	}
	return blogList, nil
}
