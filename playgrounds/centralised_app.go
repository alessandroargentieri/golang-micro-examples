/*
  One possibility to let each component of an app to reach all other app's component is to hold the pointer of the App itself in each implementation of the component
*/
package main

import (
    "fmt"
)

// App is the whole application containing the various components in terms of interfaces
type App struct {
   DB Database
   Cache Cache
   Server Server
   Queue Queue
}

type AppComponent interface{
   GetParentApp() *App
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

type Database interface{
   AppComponent
   HelloDB() string
}

type MySQLDB struct{
    App *App
    // Connection etc.
}

func (db *MySQLDB) HelloDB() string {
   return "hello from MySQL Database!"
}

func (db *MySQLDB) GetParentApp() *App {
   return db.App
}

func NewDB(app *App) Database{
    return &MySQLDB{App: app}
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~

type Cache interface{
   AppComponent
   HelloCache() string
}

type Redis struct{
   App *App
   // Redis client etc. 
}

func (r *Redis) HelloCache() string {
   return "hello from Redis Cache!"
}

func (r *Redis) GetParentApp() *App {
   return r.App
}

func NewCache(app *App) Cache{
    return &Redis{App: app}
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~

type Server interface{
   AppComponent
   HelloServer() string
}

type HTTPServer struct{
   App *App
   // Endpoints etc.
}

func (s *HTTPServer) HelloServer() string {
   return "hello from HTTP Server!"
}

func (s *HTTPServer) GetParentApp() *App {
   return s.App
}

func NewServer(app *App) Server{
    return &HTTPServer{App: app}
}

// ~~~~~~~~~~~~~~~~~~~~~~~~

type Queue interface{
    AppComponent
    HelloQueue() string
}

// Kafka is the real implementation of the Queue interface
type Kafka struct{
    App *App
    // Kafka Client etc.
}

func (k *Kafka) HelloQueue() string {
   return "hello from kafka queue!"
}

func (k *Kafka) GetParentApp() *App {
   return k.App
}

func NewQueue(app *App) Queue{
    return & Kafka{App: app}
}

// ~~~~~~~~~~~~~~~~~~~~~~


func main() {

   app := initApp()

   server := app.Server

   fmt.Println(server.HelloServer())

   fmt.Println(server.GetParentApp().Queue.HelloQueue())

}

func initApp() App {

   app := App{}
	
   db := NewDB(&app)
   cache := NewCache(&app)
   server := NewServer(&app)
   queue := NewQueue(&app)

   app.DB = db
   app.Cache = cache
   app.Server = server
   app.Queue = queue

   return app

}

