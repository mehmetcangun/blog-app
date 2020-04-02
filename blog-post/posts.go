package main

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func panicP(e error) {
	if e != nil {
		panic(e)
	}
}

func InitDB(src string) *sql.DB {
	db, err := sql.Open("sqlite3", src)

	panicP(err)

	if db == nil {
		panic("DB not created.")
	}

	return db
}

type UserDetail struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	ImgSrc  string `json:"imgSrc"`
}

type UserLogin struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Post struct {
	ID            int        `json:"id"`
	Title         string     `json:"title"`
	Content       string     `json:"content"`
	ImgSrc        string     `json:"imgSrc"`
	EstimatedTime int        `json:"estimatedTime"`
	CreatedDate   time.Time  `json:"createdDate"`
	UserDetails   UserDetail `json:"userDetails"`
}

func HashPassword(x string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(x), 14)
	panicP(err)
	return string(bytes)
}

func Migrate(db *sql.DB) {
	queryTables := `
			drop table if exists users;
			drop table if exists posts;
			create table if not exists users(
				id integer not null primary key autoincrement,
				username varchar(50) unique,
				password varchar(50),
				name varchar(20),
				surname varchar(20),
				imgSrc varchar(1000)
			);
			
			create table if not exists posts(
				id integer not null primary key autoincrement,
				title varchar(100),
				content varchar,
				imgSrc varchar(1000),
				estimatedTime smallint,
				createdDate timestamp default current_timestamp,
				userId integer references users(id)
			);
			`

	queryUser1 :=
		`
			insert into users(username, password, name, surname, imgSrc) values( 'root', ?, 'Go', 'lang', 'https://www.madebymikal.com/wp-content/uploads/2018/04/golang.png' );
		`
	queryUser2 :=
		`
			insert into users(username, password, name, surname, imgSrc) values( 'root2', ?, 'Docker', 'Compose', 'https://www.pngitem.com/pimgs/m/27-272595_docker-compose-docker-compose-logo-hd-png-download.png' );
		`

	_, errTables := db.Exec(queryTables)
	panicP(errTables)

	hashedPass := HashPassword("toor")

	stmt, errUser1 := db.Prepare(queryUser1)
	panicP(errUser1)
	defer stmt.Close()
	_, err2User1 := stmt.Exec(hashedPass)
	panicP(err2User1)

	stmt2, errUser2 := db.Prepare(queryUser2)
	panicP(errUser2)
	defer stmt2.Close()
	_, err2User2 := stmt2.Exec(hashedPass)
	panicP(err2User2)

	var x UserDetail
	x.ID = 1
	_, errPost := NewPost(db, Post{
		Title:         "How to learn docker",
		Content:       "Docker content",
		ImgSrc:        "https://miro.medium.com/max/3200/1*asSDJQpw1EQPFN-BqQSU0Q.png",
		EstimatedTime: 5,
		UserDetails:   x,
	})

	panicP(errPost)
}

type UserCollection struct {
	UserDetails UserDetail `json:"userDetails"`
}

type PostCollection struct {
	Posts []Post `json:"posts"`
}

func GetPosts(db *sql.DB) PostCollection {
	sql := "select p.id as ID, p.title as Title, p.content as Content, p.imgSrc as ImgSrc, p.estimatedTime as EstimatedTime, p.createdDate as CreatedDate, u.id as UserId, u.name as Name, u.surname as Surname, u.imgSrc as ImgSrc from posts as p inner join users as u on p.Userid = u.id order by p.createdDate desc;"

	rows, err := db.Query(sql)
	panicP(err)
	defer rows.Close()
	result := PostCollection{}
	for rows.Next() {
		post := Post{}
		err2 := rows.Scan(&post.ID, &post.Title, &post.Content, &post.ImgSrc, &post.EstimatedTime, &post.CreatedDate, &post.UserDetails.ID, &post.UserDetails.Name, &post.UserDetails.Surname, &post.UserDetails.ImgSrc)
		panicP(err2)
		result.Posts = append(result.Posts, post)
	}
	return result
}

func HandlePosts(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, GetPosts(db))
	}
}

func NewPost(db *sql.DB, p Post) (int64, error) {

	query := "insert into posts(title, content, imgSrc, estimatedTime, userId) values(?, ?, ?, ?, ?);"

	stmt, err := db.Prepare(query)

	panicP(err)

	defer stmt.Close()
	res, err2 := stmt.Exec(p.Title, p.Content, p.ImgSrc, p.EstimatedTime, p.UserDetails.ID)

	panicP(err2)

	return res.LastInsertId()
}

func HandleCreatePost(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var p Post
		c.Bind(&p)

		id, err := NewPost(db, p)
		if err == nil {
			return c.JSON(http.StatusCreated, map[string]int64{
				"affected": id,
			})
		}
		return err
	}
}

func CheckLogin(db *sql.DB, uname, pass string) UserDetail {

	query := "select id, name, surname, imgsrc, password from users where username=?;"

	stmt, err := db.Query(query, uname, pass)
	panicP(err)

	defer stmt.Close()

	var userDetail UserDetail
	if stmt.Next() {
		var passFromTables string
		err2 := stmt.Scan(&userDetail.ID, &userDetail.Name, &userDetail.Surname, &userDetail.ImgSrc, &pass)
		panicP(err2)

		if err := bcrypt.CompareHashAndPassword([]byte(passFromTables), []byte(pass)); err == nil {
			userDetail.ID = -1
			userDetail.ImgSrc = ""
			userDetail.Name = ""
			userDetail.Surname = ""
		}
	}

	return userDetail
}

func HandleLogin(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var inputUser UserLogin
		c.Bind(&inputUser)

		resUser := CheckLogin(db, inputUser.Username, inputUser.Password)

		if resUser.ID > 0 {
			return c.JSON(http.StatusOK, resUser)
		}
		return c.JSON(http.StatusUnauthorized, resUser)
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	var migrated bool = false

	if _, err := os.Stat("blog.db"); os.IsNotExist(err) {
		migrated = true
	}

	db := InitDB("blog.db")
	if migrated {
		Migrate(db)
	}

	e.GET("/posts", HandlePosts(db))
	e.POST("/posts", HandleCreatePost(db))
	e.POST("/login", HandleLogin(db))

	e.Logger.Fatal(e.Start(":1414"))
}
