package main

import (
	"fmt"
	"reflect"
)

type Employee struct {
	Id         int    `db:"emp_id"`
	Department string `db:"emp_department"`
}

func PrintDataBaseColumns(s any) {
	// 1. Get the TYPE (The blueprint: what fields exist, what are their tags?)
	t := reflect.TypeOf(s)

	// 2. Get the VALUE (The actual data inside this specific struct instance)
	v := reflect.ValueOf(s)

	if t.Kind() != reflect.Struct {
		fmt.Println("Error: Not a Struct!")
		return
	}

	for i := 0; i < t.NumField(); i++ {
		// Field() on TYPE gives us the Blueprint (Name, DB Tag, etc)
		typeField := t.Field(i)

		// Field() on VALUE gives us the actual data (1, "AI", etc)
		valueField := v.Field(i)

		fmt.Printf("Column: %s | Value: %v\n", typeField.Tag.Get("db"), valueField.Interface())
	}
	fmt.Println("-------------------")
}
func main() {
	PrintDataBaseColumns(Employee{Id: 1, Department: "AI"})
	PrintDataBaseColumns(11)
}
