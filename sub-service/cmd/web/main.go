package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	webPort = "80"
)

// main function
func main() {

	// connect to database
	db := initDB()

	// create sessions
  
  session := initSession()
  
  // create loggers

  infoLog := log.New(os.Stdout, "INFO\t",log.Ldate|log.Ltime)
  errorLog := log.New(os.Stdout, "ERROR\t",log.Ldate|log.Ltime|log.Lshortfile)

	// create channels

	// create waitgroup
  wg := sync.WaitGroup{}

	// setup app config
  app := Config{
    Session: session,
    DB: db,
    Wait: &wg,
    InfoLog: infoLog,
    ErrorLog: errorLog,
  }

	// set mail

	// listen for web


  app.serve()

}


func (app *Config) serve() {
  srv := &http.Server{
    Addr: fmt.Sprintf(":%s", webPort),
    Handler: app.routes(),
  }


  app.InfoLog.Println("starting web server")


  if err := srv.ListenAndServe(); err != nil{
    app.ErrorLog.Fatal("failed to start server")
  }
}

func initDB() *sql.DB {
	conn := connectToDB()

	if conn == nil {
		log.Panic("cannot connect to db")
	}

	return conn

}

func connectToDB() *sql.DB {
	counts := 0

	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)

		if err != nil {
			log.Println("postgres not yet ready")
		} else {
			log.Println("connected to db")
			return connection
		}

		if counts > 10 {
			return nil
		}

		log.Println("waiting for 1 second")
		time.Sleep(1 * time.Second)
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}


func initSession() *scs.SessionManager{

  session := scs.New()
  session.Store = redisstore.New(initRedis())
  session.Lifetime = 24 * time.Hour
  session.Cookie.Persist = true
  session.Cookie.SameSite = http.SameSiteLaxMode
  session.Cookie.Secure = true
  
  return session
}


func initRedis() *redis.Pool {
  redisPool := &redis.Pool{
    MaxIdle: 10,
    Dial: func() (redis.Conn, error){
      return redis.Dial("tcp", os.Getenv("REDIS"))
    },
  }


  return redisPool
  
}
