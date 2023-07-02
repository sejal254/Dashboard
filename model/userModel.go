package model

import (
	"fmt"

	"github.com/beego/beego/orm"
)

type User struct {
	User_id  int    `orm:"column(user_id);auto;pk" json:"userId"`
	Username string `orm:"column(username)" json:"userName"`
	Email    string `orm:"column(email);unique" json:"email"`
	Role     string `orm:"column(role)" json:"role"`
	Status   string `orm:"column(status)" json:"status"`
	Password string `orm:"column(password)" json:"password"`
}

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) TableName() string {
	return "user"
}
func AddUser(u *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(u)
	return id, err
}

func GetAllUsers() (u []User, err error) {
	o := orm.NewOrm()
	//_, err = o.QueryTable(new(User)).Filter("status","Active").All(&u)
	_, err = o.Raw("SELECT * FROM user where status=?", "Active").QueryRows(&u)

	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetUserByCred(email string) (u User, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(new(User)).Filter("email", email).One(&u)
	return u, err

}
func DeleteUser(id int) error {
	o := orm.NewOrm()
	p1, err1 := GetProjectsByUserID(id)
	fmt.Println(err1)
	for _, val := range p1 {
		//_, err2 := o.QueryTable(new(Project)).Filter("project_id", val.Project_id).Update(orm.Params{"status": "Inactive"})
		sql := "Update userprojects SET stauts=? WHERE project_id=?"
		_, err2 := o.Raw(sql, "Inactive", val.Project_id).Exec()

		if err2 != nil {
			fmt.Println(err2)
		}

	}
	sql := "Update user SET status=? WHERE user_id=?"
	_, err := o.Raw(sql, "Inactive", id).Exec()
	//_, err := o.QueryTable(new(User)).Filter("user_id", id).Update(orm.Params{"status": "Inactive"})
	return err
}

func UpdateRole(id int) error {

	o := orm.NewOrm()

	_, err := o.QueryTable(new(User)).Filter("user_id", id).Update(orm.Params{"role": "Admin"})

	return err
}

func GetDetailsOfUserID(id int) (u User, err error) {
	o := orm.NewOrm()

	err1 := o.QueryTable(new(User)).Filter("user_id", id).One(&u)
	fmt.Println(err1)
	return u, err1

}

func CheckIfUserExists(email string) bool {
	o := orm.NewOrm()
	exists := o.QueryTable(new(User)).Filter("email", email).Exist()
	return exists
}
