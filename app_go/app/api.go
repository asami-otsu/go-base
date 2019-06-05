package main

import (
    "database/sql"
    "fmt"

    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)

type User struct {
    ID   int
    Name string
}

func main() {
    // http://g-hyoga.hatenablog.com/entry/2018/01/29/003033
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/local_insta")
        defer db.Close()
        if err != nil {
            fmt.Println(err.Error())
        }

        rows, err := db.Query("SELECT * FROM users")
        defer rows.Close()
        if err != nil {
            fmt.Println(err)
        }

        user := User{}
        for rows.Next() {
            err = rows.Scan(&user.ID, &user.Name)
            if err != nil {
                fmt.Println(err)
            }
            fmt.Println(user)
        }

        c.JSON(200, gin.H{
            "hello": user.Name,
        })
    })
    r.Run()

  // https://qiita.com/taizo/items/54f5f49c6102f86194b8
    // db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/local_insta")
    // if err != nil {
    //     panic(err.Error())
    // }
    // defer db.Close() // 関数がリターンする直前に呼び出される

    // rows, err := db.Query("SELECT * FROM users") // 
    // if err != nil {
    //     panic(err.Error())
    // }

    // columns, err := rows.Columns() // カラム名を取得
    // if err != nil {
    //     panic(err.Error())
    // }

    // values := make([]sql.RawBytes, len(columns))

    // //  rows.Scan は引数に `[]interface{}`が必要.

    // scanArgs := make([]interface{}, len(values))
    // for i := range values {
    //     scanArgs[i] = &values[i]
    // }

    // for rows.Next() {
    //     err = rows.Scan(scanArgs...)
    //     if err != nil {
    //         panic(err.Error())
    //     }

    //     var value string
    //     for i, col := range values {
    //         // Here we can check if the value is nil (NULL value)
    //         if col == nil {
    //             value = "NULL"
    //         } else {
    //             value = string(col)
    //         }
    //         fmt.Println(columns[i], ": ", value)
    //     }
    //     fmt.Println("-----------------------------------")
    // }
}

