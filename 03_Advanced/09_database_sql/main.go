package main

import (
	"database/sql"
	"log/slog"
	"time"

	_ "modernc.org/sqlite"
)

func main() {

	db, err := sql.Open("sqlite", "./test.db")
	if err != nil {
		panic(err)

	}
	defer db.Close()
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Second)

	// Production Fix 1: We already declared `err` above. We can reassign it using `=` instead of `:=`!
	err = db.Ping()
	if err != nil {
		slog.Error("Error Connecting DB", slog.Any("error", err))
		return
	}
	slog.Info("Successfully connected to the DB Pool!")

	// Production Fix 2: Because `res` is a BRAND NEW variable, we can use `:=`.
	// Go is smart enough to create `res` and just OVERWRITE the existing `err` variable!
	res, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, age INTEGER)")
	if err != nil {
		slog.Error("Error CReating DB", slog.Any("Error", err))
		return
	}
	slog.Info("Tabel Created", slog.Any("Message", res))

	// Because `row` is new, we can use `:=` again to overwrite `err`!
	row, err := db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "Luffy", 20)
	if err != nil {
		slog.Error("Error Insert DB", slog.Any("Error", err))
		return
	}
	id, _ := row.LastInsertId()
	slog.Info("Inserted Row ID:", slog.Any("Id", id))

	var user_id int
	var name string
	var age int

	// Overwriting `err` again!
	err = db.QueryRow("SELECT id, name, age FROM users WHERE name = ?", "Luffy").Scan(&user_id, &name, &age)

	// Production Fix 3: You wrote `if err == sql.ErrNoRows` but checked the WRONG err variable!
	// Now that everything uses `err`, it works perfectly!
	if err == sql.ErrNoRows {
		slog.Info("User not found!")
		return
	} else if err != nil {
		slog.Error("Database crashed:", slog.Any("Error", err))
		return
	}

	slog.Info("successfully fetched", slog.Any("user_id", user_id), slog.Any("Name", name), slog.Any("Age", age))

	// `rows` is new, `err` is overwritten!
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		slog.Error("Database crashed:", slog.Any("Error", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		// Production Fix 4: You used `SELECT *`, which returns 3 columns (id, name, age).
		// You only provided 2 variables to Scan, which causes it to fail silently!
		// ALWAYS check the error from rows.Scan()!
		var id, a int
		var n string

		err = rows.Scan(&id, &n, &a)
		if err != nil {
			slog.Error("Failed to scan row", slog.Any("error", err))
			continue
		}
		slog.Info("Data", slog.Int("id", id), slog.String("name", n), slog.Int("age", a))
	}

	// Production Fix 5: Always check for errors that occurred *during* the loop!
	if err = rows.Err(); err != nil {
		slog.Error("Loop crashed unexpectedly", slog.Any("error", err))
	}
}
