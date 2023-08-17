/*
 * @Description: 后端接口代码示例
 在这个示例中，我们创建了一个简单的HTTP服务器，它提供了一个名为/getdata的端点来从数据库中检索用户数据。
 请注意，您需要根据实际情况替换数据库连接字符串中的username、password、host、port和database
 * @Author: HCQ
 * @Company(School): UCAS
 * @Email: 1756260160@qq.com
 * @Date: 2023-08-17 23:38:12
 * @LastEditTime: 2023-08-17 23:53:17
 * @FilePath: /go-learning/backend-go/demo/demo.go
 */

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// 初始化数据库连接
	db, err := sql.Open("mysql", "username:password@tcp(host:port)/database")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建HTTP处理函数
	http.HandleFunc("/getdata", func(w http.ResponseWriter, r *http.Request) {
		// 查询数据库
		rows, err := db.Query("SELECT id, name, age FROM users")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// 解析查询结果
		var users []map[string]interface{}
		for rows.Next() {
			var id int
			var name string
			var age int
			err := rows.Scan(&id, &name, &age)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			user := map[string]interface{}{
				"id":   id,
				"name": name,
				"age":  age,
			}
			users = append(users, user)
		}

		// 将结果转为JSON并发送响应
		response, err := json.Marshal(users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	})

	// 启动HTTP服务器
	fmt.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
