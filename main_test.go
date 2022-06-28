package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

// Convert any struct to JSON String with Pretty Print
func ToJson(val interface{}) (string, error) {
	b, err := json.Marshal(val)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func expectedResult() string {
	result := `[{"id":"b","name":"test"},{"id":"a","name":""}]`

	return result
}

func TestGORM(t *testing.T) {
	car1 := Car{ID: "a"}
	DB.Create(&car1)
	DB.Model(&car1).Select("Name").Updates(map[string]interface{}{"name": gorm.Expr("NULL")})

	car2 := Car{ID: "b", Name: "test"}
	DB.Create(&car2)

	var records []Car
	DB.Model(&Car{}).Order("id desc").Find(&records)

	jsonstr, _ := ToJson(records)
	fmt.Println(jsonstr)

	if expectedResult() != jsonstr {
		t.Errorf("Failed, got error: %v", "Name should be Null but it copied")
	}
}
