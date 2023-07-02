package model

import (
	"fmt"

	"github.com/beego/beego/orm"
)

type Project struct {
	Project_id int    `orm:"column(project_id);auto;pk" json:"projectId"`
	Name       string `orm:"column(name)" json:"name"`
	Status     string `orm:"column(status)" json:"status"`
}

func init() {
	orm.RegisterModel(new(Project))
}
func (p *Project) TableName() string {
	return "project"
}

func AddProject(p *Project) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(p)
	return
}

func GetAllProjects() (p []Project, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(Project)).RelatedSel().All(&p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func RemoveProject(id int) error {
	o := orm.NewOrm()
	_, err := o.QueryTable(new(Project)).Filter("project_id", id).Update(orm.Params{"status": "Inactive"})
	if err != nil {
		return err
	}
	return nil
}

func ChangeStatusForAssign(id int) error {
	o := orm.NewOrm()
	_, err1 := o.QueryTable(new(Project)).Filter("project_id", id).Update(orm.Params{"status": "Active"})
	if err1 != nil {
		fmt.Println(err1)
	}
	return err1
}
