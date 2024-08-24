package pkg_sqlite3

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestListTables(t *testing.T) {
	// 创建内存中的 SQLite 数据库
	//db, err := sql.Open("sqlite3", ":memory:")
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建示例表
	createTableSQL := `CREATE TABLE example (id INTEGER NOT NULL PRIMARY KEY, name TEXT);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}

	// 查询数据库中的所有表
	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table';")
	if err != nil {
		t.Fatalf("Failed to query tables: %s", err)
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			t.Fatalf("Failed to scan table name: %s", err)
		}
		tables = append(tables, tableName)
	}

	// 输出所有表名
	fmt.Println("Tables in the database:")
	for _, table := range tables {
		fmt.Println(table)
	}

	// 测试期望
	expectedTable := "example"
	found := false
	for _, table := range tables {
		if table == expectedTable {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected table %s not found in database", expectedTable)
	}
}
