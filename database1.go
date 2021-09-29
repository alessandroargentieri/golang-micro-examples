package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"  //uses the blank identifier because it serves only to recall its init() function
)

// table object
type Tag struct {
    Id   string    `json:"id"`
    Name string    `json:"name"`
}

///////////////////////////////////////////////////////////////////////////
type Connection interface {
    Open(configs map[string]string) Connection
    Close()
}
//-------------------------------------------------------------------------
// IMPLEMENTS Connection interface
type Conn struct {
    Db *sql.DB
}
func (conn Conn) Open(configs map[string]string) Connection {
    db, err := sql.Open(configs["dbms"], fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs["username"], configs["password"], configs["host"], configs["port"], configs["dbname"]))
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    // set the struct variable
    conn.Db = &db 
    return conn
}
func (conn Conn) Close() {
    conn.Db.Close()
}
///////////////////////////////////////////////////////////////////////////

type TagRepository interface {
        //Initialize(conn Connection)
        FindByID(id string) (Tag, error)
        Find() ([]Tag, error)
        Create(tag Tag) error
        Update(tag Tag) error
        Delete(id string) error
}

/////////////////////////////////
type TagRepositoryImpl struct {
    Conn Connection
}
//func (tagRepo TagRepositoryImpl) Initialize(conn Connection) {
//    tagRepo.Conn = conn
//}
func (tagRepo TagRepositoryImpl) Create(tag Tag) error {
    insert, err := tagRepo.Conn.Db.Query("INSERT INTO tags VALUES ( '" + tag.Id + "', '" + tag.Name + "' )")
    if err != nil {
        return err
    }
    defer insert.Close()
    return nil
}
func (tagRepo TagRepositoryImpl) Update(tag Tag) error {
    update, err := tagRepo.Conn.Db.Query("UPDATE tags SET name='" + tag.Name + "' WHERE id= '" + tag.Id + "'")
    if err != nil {
        return err
    }
    defer update.Close()
    return nil
}
func (tagRepo TagRepositoryImpl) Find() ([]Tag, error) {
    results, err := tagRepo.Conn.Db.Query("SELECT id, name FROM tags")
    if err != nil {
        return nil, err
    }
    var tags []Tag
    for results.Next() {
        var tag Tag
        err = results.Scan(&tag.Id, &tag.Name)
        if err != nil {
            return tags, err
        }
        tags = append(tags, tag)
    }
    return tags, nil
}
func (tagRepo TagRepositoryImpl) FindById(id string) (Tag, error) {
   var tag Tag
   err = tagRepo.Conn.Db.QueryRow("SELECT id, name FROM tags WHERE id = ?", id).Scan(&tag.Id, &tag.Name)
   if err != nil {
       return nil, err
   }
   return tag, nil
}
func (tagRepo TagRepositoryImpl) Delete(id string) error {
   _, err := tagRepo.Conn.Db.Query("DELETE FROM tags WHERE id='" + id + "'")
   if err != nil {
       return err
   }
   return nil     
}



/////////////////////////////////////////

func main() {

    configs := map[string]string {
        "dbms": "mysql",
        "username":"root",
        "password":"password",
        "host":"172.17.0.2",
        "port":"3306",
        "dbname":"db_example",
    }

    var connection Conn = Conn{}
    connection = connection.Open(configs)

    var repo TagRepositoryImpl = TagRepositoryImpl{ connection }

    repo.Create( Tag{"33f92adc-ff49-11ea-adc1-0242ac120002", "Alessandro"} )
    repo.Create( Tag{"55b74bcc-ff49-11ea-adc1-0242ac120002", "Francesca"} )
    repo.Create( Tag{"5f8e2a9e-ff49-11ea-adc1-0242ac120002", "Marco"} )
    repo.Create( Tag{"6974f470-ff49-11ea-adc1-0242ac120002", "Monica"} )

    tags := repo.Find()
    fmt.Println(tags)


}



//////////

type Database interface {
    query (q string) (result interface{}, error Error)
}

//////////

type DbImpl struct {
    Db *sql.Db
}
func 


///////////////////////////////

package database

////////////////////////////////////////
type DB interface {
    Connect(...)
    Query(...)
    Close()
}
/////////////////////////////////////////
type DBImpl struct {
    Db *sql.Db
}
func(db DBImpl) Connect() {}
func(db DBImpl) Query() {}
func(db DBImpl) Close() {}
/////////////////////////////////////////

func init() {
    Db = new(DBImpl)
}

var Db DB = new(DBImpl)

func Query (q string) (result interface{}, error Error) {
   Db.Query(....)
   return result, error
}
////////////////
///////////////////////

func TestQuery(t *testing.T) {

    Db = new(DBMock)

    result := Query("INSERT INTO...")
    expected := ...
    if(result != expected) {
        t.Errorf("Failed!")
    }
}



