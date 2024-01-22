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
var db *sql.DB //连接池对象
var node *snowflake.Node
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
	/*node, err := snowflake.NewNode(1)
    if err != nil {
        fmt.Println(err)
        return
    }*/
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
                   fmt.Printf("exec failed:%s \n", err)
                    //return
                } 
               // fmt.Println("insert succ:", res)
    }
    defer conn.Close()
    wg.Done()
}

func useTransaction(dataNum int) {
	var personObject Person
    for i := 0; i < dataNum ;i++ {
		personObject.id=node.Generate().Int64()
		personObject.userId=rand.Intn(100000000)
		personObject.username=randstr.RandomAlphanumeric(20)
		personObject.sex=rand.Intn(1)
		personObject.phone=randstr.RandomAlphanumeric(12)
		personObject.email=randstr.RandomAlphanumeric(20)

        updateUserId:=rand.Intn(100000000)
        //deleteUserId:=rand.Intn(100000000)
        //事务
        tx, err := db.Begin()
        if err != nil {
            if tx != nil {
                _ = tx.Rollback()
            }
            fmt.Printf("begin trans action failed, err:%v\n", err)
            wg.Done()
            return
        }
        //语句1：insert
        sqlStr1 := "insert into  person(id,user_id,username,sex,phone,email) values(?,?,?,?,?,?)"
	    result1, err := tx.Exec(sqlStr1,personObject.id,personObject.userId,personObject.username,personObject.sex,personObject.phone,personObject.email)
	    if err != nil {
		    _ = tx.Rollback()
		    fmt.Printf("insert failed, err:%v\n", err)
            wg.Done()
		    return
	    }
	    id, err := result1.LastInsertId()
	    if err != nil {
		    _ = tx.Rollback()
		    fmt.Printf("insert result1.RowsAffected() failed, err:%v\n", err)
            wg.Done()
		    return
	    }
        //fmt.Println("insert succ:", id)
        //语句2：update
	    sqlStr2 := "UPDATE person SET username = ? WHERE user_id = ?"
	    result2, err := tx.Exec(sqlStr2,"sz4450", updateUserId)
	    if err != nil {
		    _ = tx.Rollback()
		    fmt.Printf("update failed, err:%v\n", err)
            wg.Done()
		    return
	    }
	    row2, err := result2.RowsAffected()
	    if err != nil {
		    _ = tx.Rollback()
		    fmt.Printf("update result2.RowsAffected() failed, err:%v\n", err)
            wg.Done()
		    return
	    }
        //语句3：delete
        /*//delete打开，并发会死锁
        sqlStr3 := "delete from person WHERE user_id = ?"
	    result3, err := tx.Exec(sqlStr3,deleteUserId)
	    if err != nil {
		    _ = tx.Rollback()
		    fmt.Printf("delete failed, err:%v\n", err)
            wg.Done()
		    return
	    }
	    row3, err := result3.RowsAffected()
	    if err != nil {
		    _ = tx.Rollback()
		    fmt.Printf("delete result3.RowsAffected() failed, err:%v\n", err)
            wg.Done()
		    return
	    }
        fmt.Println("transaction commit succ:",id,row2,row3)
        */
        fmt.Println("transaction commit succ:",id,row2)
        tx.Commit()

    }
    wg.Done()
}

func main() {
	//node= snowflake.NewNode(1)
    initDB()
    var err error
    node,err = snowflake.NewNode(1)
    if err != nil {
        fmt.Println(err)
        return
    }
    var threadNum,dataNum,opType int
	flag.IntVar(&threadNum, "threadNum", 1, "并发线程数")
	flag.IntVar(&dataNum, "dataNum", 100000, "每个线程数据量")
    flag.IntVar(&opType, "opType", 1, "操作类型")
	flag.Parse()
	fmt.Println(threadNum,dataNum,opType)
    //threads=100
    //node= snowflake.NewNode(1)
    st := time.Now().UnixNano()
    if opType==1 {
        fmt.Printf("insert操作开始\n")
        for i := 0; i < threadNum; i++ {
            wg.Add(1)
            go insertData(dataNum)
        }
        wg.Wait()
    } else if opType==2 {
        fmt.Printf("select操作开始\n")
        for i := 0; i < threadNum; i++ {
            wg.Add(1)
            go selectData(dataNum)
        }
        wg.Wait()  
    } else if opType==3 {
        fmt.Printf("事务操作开始\n")
        for i := 0; i < threadNum; i++ {
            wg.Add(1)
            go useTransaction(dataNum)
        }
        wg.Wait() 
    } else {
        fmt.Printf("不支持该操作类型\n")
    }
    et := time.Now().UnixNano()
    fmt.Printf("finish time: %6.6fs\n", (float32(et-st))/1e9)
}