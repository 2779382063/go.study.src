package main
 
import (
   	 _ "github.com/go-sql-driver/mysql"
//	  "github.com/rs/xid"  
     // "github.com/satori/go.uuid"
	  "github.com/bwmarrin/snowflake"  //雪花id
	  "github.com/pochard/commons/randstr"
	//  "github.com/google/uuid"
      "math/rand"
      "database/sql"
      "time"
      "fmt"
      "sync"
	  "flag"
)
//create table person (id bigint primary key,user_id int ,username varchar(100),sex int,phone varchar(12),email varchar(20));
type Person struct {
	id		 int64  `db:"id"`
    userId   int    `db:"user_id"`
    username string `db:"username"`
    sex      int `db:"sex"`
	phone    string `db:"phone"`
    email    string `db:"email"`
}
var wg sync.WaitGroup
//var node *snowflake.Node
var db *sql.DB //连接池对象
func initDB() (err error) {
	//数据库
	//用户名:密码啊@tcp(ip:端口)/数据库的名字
	dsn := "root:123456@tcp(172.16.13.66:3305)/test"
	//连接数据集
	db, err = sql.Open("mysql", dsn) //open不会检验用户名和密码
	if err != nil {
		return
	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		return
	}
	fmt.Println("连接数据库成功~")
	//设置数据库连接池的最大连接数
	db.SetMaxIdleConns(10)
	return
}

func selectData(dataNum int) {

	var personObject Person
    for i := 0; i < dataNum ;i++ {
			    userId:=rand.Intn(100000000)
			
    /*            err := db.Select(&person, "select user_id, username, sex from person where user_id=?", userId)
                if err != nil {
                    fmt.Println("exec failed, ", err)
                    return
                }

                 fmt.Println("select succ:", person)
    */
        sqlStr := "select user_id, username, sex from person where user_id > ? limit 3"
        rows, err := db.Query(sqlStr, userId)
        if err != nil {
            fmt.Println("数据查询失败. err:", err)
        }
    
        defer rows.Close()
        for rows.Next() {
            if err := rows.Scan(&personObject.userId, &personObject.username,&personObject.sex); err != nil {
                panic(err)
            }
            fmt.Printf("User ID: %d, Name: %s, Sex: %d\n", personObject.userId, personObject.username,personObject.sex)
        }
    }
    wg.Done()
}
func insertData(dataNum int) {
	node, err := snowflake.NewNode(1)
    if err != nil {
        fmt.Println(err)
        return
    }
    conn,err := sql.Open("mysql","root:123456@tcp(172.16.13.66:3305)/test?charset=utf8mb4")
    if err != nil{
        fmt.Printf("open failed, ", err)
        return
    }
  //  var uuid string
  //  var str string
  //  var k int
    //num:=10000
	var personObject Person
    for i := 0; i < dataNum ;i++ {
              //  uuid := uuid.NewV4()
                //str := uuid.String()
               // k := rand.Intn(100)
			   // xId :=xid.New()
			  //	uid := uuid.NewV4()
				//a=uid.BigInt()
			    personObject.id=node.Generate().Int64()
				personObject.userId=rand.Intn(100000000)
				personObject.username=randstr.RandomAlphanumeric(20)
				personObject.sex=rand.Intn(1)
				personObject.phone=randstr.RandomAlphanumeric(12)
				personObject.email=randstr.RandomAlphanumeric(20)
                _,err :=conn.Exec("insert into  person(id,user_id,username,sex,phone,email) values(?,?,?,?,?,?)",personObject.id,personObject.userId,personObject.username,personObject.sex,
				personObject.phone,personObject.email)
                if err != nil{
                   fmt.Printf("exec failed, ", err)
                  // fmt.Println("insert succ:", res)
                   return
                } else {
					//fmt.Println("insert succ:", res)
				}
               // fmt.Println("insert succ:", res)
    }
    defer conn.Close()
    wg.Done()
}
