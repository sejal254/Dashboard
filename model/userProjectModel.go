package model

import (
	"github.com/beego/beego/orm"
)

type UserProject struct {
	User_id    int `orm:"column(user_id);pk" json:"userId"`
	Project_id int `orm:"column(project_id)"json:"projectId"`
	Project_name string `orm:"column(project_name)" json:"projectName"`
	Status string `orm:"column(status)" json:"status"`
}

func init() {
	orm.RegisterModel(new(UserProject))
}

func (u *UserProject) TableName() string {
	return "userprojects"
}

func AssignProjectByID(up UserProject) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(up)
	if err != nil {
		return -1, err
	}
	return id, nil

}
func GetProjectsByUserID(uid int) (p []UserProject, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(UserProject)).RelatedSel().Filter("Users__User__user_id", uid).All(&p)
	return
}
func IsUserAssignedToProject(u UserProject) bool {
	o := orm.NewOrm()
	return o.QueryTable(new(UserProject)).RelatedSel().Filter("user_id", u.User_id).Filter("Project_id", u.Project_id).Exist()
}
