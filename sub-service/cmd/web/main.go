package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
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

	db.Ping()

	// create sessions
  
  session := initSession()

	// create channels

	// create waitgroup

	// setup app config

	// set mail

	// listen for web

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
