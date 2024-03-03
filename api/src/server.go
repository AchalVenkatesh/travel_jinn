package main

import(
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
  "os"
  "log"
  "database/sql"
  "github.com/go-sql-driver/mysql"
)

type users struct{
  Full_name string `json:"full_name"`
  Email     string `json:"email"`
  Password  string `json:"password"`
  Username  string `json:"username"`
}

func connectToDb(db *sql.DB)(*sql.DB, error){
  cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "travel_jinn",
    }
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        db.Close()
        return nil, err
    }

    fmt.Println("Connected!")
    return db , nil;
}

func main(){
  var db *sql.DB
  var err error
  db,err = connectToDb(db)
  if err!=nil{
    log.Fatal(err)
  }
  defer db.Close()

  router := gin.Default()
  router.POST("/signin",createUser(db))
  router.POST("/login",login(db))
  router.POST("/travel",travel(db))
  router.POST()
  router.Run("localhost:3050")

}

func createUser(db *sql.DB) gin.HandlerFunc{
  return func(c *gin.Context){
    var err error
    var user users
    if err = c.ShouldBindJSON(&user);err!=nil{
            c.String(http.StatusBadRequest, "Bad request")
            return
        }

    result, err := db.Exec("INSERT INTO users (email, password, username, full_name) VALUES (?, ?, ?, ?)", user.Email, user.Password, user.Username, user.Full_name)
    if err != nil {
      log.Printf("Error inserting user into database: %v", err)
      c.String(http.StatusInternalServerError, "Internal Server Error")
      return
    }

        id, err := result.LastInsertId()
        if err != nil {
            log.Printf("Error retrieving last insert ID: %v", err)
            c.String(http.StatusInternalServerError, "Internal Server Error")
            return
        }

        c.JSON(http.StatusCreated, gin.H{"id": id})
    }
}

func login(db *sql.DB) gin.HandlerFunc{
  return func(c *gin.Context){
    var err error
    var user users
    if err = c.ShouldBindJSON(&user);err!=nil{
      c.String(http.StatusBadRequest, "Bad Request")
    }
    result, err := db.Exec("Select * from users where email=?",user.Email)
    if err != nil{
      c.Status(http.StatusNotFound)
    }
    fmt.Println(result)
    c.JSON(http.StatusOK,"Logged in successfully")
  }
}

func travel(db *sql.DB) gin.HandlerFunc{
  return func(c *gin.Context){
    var err error
    var user users
    
  }
}

