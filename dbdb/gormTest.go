package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func ConnectToDB() *gorm.DB{
	db, err := gorm.Open("mysql","root:123456@/test_db")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return nil
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	return db

}
type Student struct {
	Sid int
	Sname string
}
type User struct {
Id string
Name string
Agg int
}

//findall
func (student Student) FindAll(db *gorm.DB,tableName string) (students []Student) {
	db.Table(tableName).Find(&students)
	return students
}

//query
func (student *Student) QueryById(db *gorm.DB,tableName string) (result int64) {

		student.Sid=6
	 db.Table(tableName).Where("sId = ?", student.Sid).Count(&result)
	db.Table(tableName).Where("sId = ?", student.Sid).First(&student)
	return result
}

func (student *Student) InsertNamespace2(db *gorm.DB,tableName string) (result int64) {
		affected := db.Table(tableName).Create(&student).RowsAffected
		return affected
}
//update return 1
func (student *Student) UpdateNamespace(db *gorm.DB,tableName string) (result int64) {

	id := student.DeleteNamespace(db, tableName)
	if id==1{
		println("yijing")
		student.Sname="12455"
		namespace := student.InsertNamespace2(db, tableName)

		return namespace
	}else {
		println("没有此人")
		return 0}
}
//delete return 1
func (student *Student) DeleteNamespace(db *gorm.DB,tableName string) (result int64) {

	id := student.QueryById(db, tableName)
	if id==1{
		println("yijing")

		rowsAffected := db.Table(tableName).Where("sId= ?", student.Sid).Delete(&student).RowsAffected

		return rowsAffected
	}else {
		println("没有此人")

		return 0}
}
//insert return 1
func (student *Student) InsertNamespace(db *gorm.DB,tableName string) (result int64) {

	id := student.QueryById(db, tableName)
	if id==0{
		println("没有此人")
		student.Sname="12345"

		affected := db.Table(tableName).Create(&student).RowsAffected
		return affected
	}else {
		println("yijing")
		return 0}
}
func main() {
	db, err := gorm.Open("mysql", "mxd:@#MA847547125**@tcp(127.0.0.1:3306)/mxd")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	defer db.Close()
	student := new(Student)
//namespace := student.UpdateNamespace(db, "student")
	//namespace := student.DeleteNamespace(db, "student")
//	println("nam+",namespace)
	all := student.FindAll(db, "student")
	s := all[0]
	fmt.Println(s)
	//println(args)
	student.Sname="123"
	student.Sid=5
	//insert
	//affected := db.Table("student").Create(&student).RowsAffected
//	println(affected)
	fmt.Println(student.Sid)
	//先查询一条记录, 保存在模型变量food
	//等价于: SELECT * FROM `foods`  WHERE (id = '2') LIMIT 1
	var counts int64 = 0
	//db.Table("student").Where("sid = ?", 2).Count(&counts)
	db.Table("student").Where("sid = ?", 20).Count(&counts)
	println("counts+",counts)
	println("----")
	//修改food模型的值
	student.Sname="ss"
	fmt.Println(student.Sid,student.Sname)
	//db.Save(&student)
	db.Table("student").UpdateColumn(&student)

}

