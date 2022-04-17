package database
 
import (
    "fmt"
    "log"
 
    "github.com/gocql/gocql"
)


func syllaDbConnectionExample() {
    cluster := gocql.NewCluster("127.0.0.1", "127.0.0.2", "127.0.0.3")
    cluster.Keyspace = "example" //DB NAME
    cluster.Consistency = gocql.Quorum
    session, _ := cluster.CreateSession()
    defer session.Close()
 
    //CREATE A (tweet) TABLE FIRST THEN EXECUTE BELLOW QUERY
    if err := session.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
        "me", gocql.TimeUUID(), "hello world").Exec(); err != nil {
        log.Fatal(err)
    }
 
    var id gocql.UUID
    var text string
 
    if err := session.Query(`SELECT id, text FROM tweet WHERE timeline = ? LIMIT 1`,
        "me").Consistency(gocql.One).Scan(&id, &text); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Tweet:", id, text)
 
    iter := session.Query(`SELECT id, text FROM tweet WHERE timeline = ?`, "me").Iter()
    for iter.Scan(&id, &text) {
        fmt.Println("Tweet:", id, text)
    }
    if err := iter.Close(); err != nil {
        log.Fatal(err)
    }
}
