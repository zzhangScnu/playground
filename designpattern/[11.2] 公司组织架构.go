package designpattern

import (
	"fmt"
	"strings"
)

/**
题目描述
小明所在的公司内部有多个部门，每个部门下可能有不同的子部门或者员工。
请你设计一个组合模式来管理这些部门和员工，实现对公司组织结构的统一操作。部门和员工都具有一个通用的接口，可以获取他们的名称以及展示公司组织结构。

输入描述
第一行是一个整数 N（1 <= N <= 100），表示后面有 N 行输入。
接下来的 N 行，每行描述一个部门或员工的信息。部门的信息格式为 D 部门名称，员工的信息格式为 E 员工名称，其中 D 或 E 表示部门或员工。

输出描述
输出公司的组织结构，展示每个部门下的子部门和员工

输入示例
MyCompany
8
D HR
E HRManager
D Finance
E AccountantA
E AccountantB
D IT
E DeveloperA
E DeveloperB

输出示例
Company Structure:
MyCompany
  HR
    HRManager
  Finance
    AccountantA
    AccountantB
  IT
    DeveloperA
    DeveloperB
*/

type Component interface {
	Display(depth int)
}

type Employee struct {
	name string
}

func NewEmployee(name string) *Employee {
	return &Employee{
		name: name,
	}
}

func (e *Employee) Display(depth int) {
	fmt.Println(strings.Repeat(" ", depth) + e.name)
}

type Department struct {
	name     string
	children []Component
}

func NewDepartment(name string) *Department {
	return &Department{
		name:     name,
		children: make([]Component, 0),
	}
}

func (d *Department) Display(depth int) {
	fmt.Println(strings.Repeat(" ", depth) + d.name)
	for _, child := range d.children {
		child.Display(depth + 1)
	}
}

func (d *Department) Add(child Component) {
	d.children = append(d.children, child)
}

type Company struct {
	name       string
	department *Department
}

func NewCompany(name string) *Company {
	return &Company{
		name:       name,
		department: NewDepartment(""),
	}
}

func (d *Company) Display(depth int) {
	fmt.Println("Company Structure:")
	fmt.Println(strings.Repeat(" ", depth) + d.name)
	for _, child := range d.department.children {
		child.Display(depth + 1)
	}
}

func (d *Company) Add(child Component) {
	d.department.Add(child)
}

func main() {
	var companyName, childType, childName string
	var count int
	fmt.Scan(&companyName, count)
	company := NewCompany(companyName)
	var department *Department
	for i := 0; i < count; i++ {
		fmt.Scan(childType, childName)
		switch childType {
		case "D":
			department = NewDepartment(childName)
			company.Add(department)
		case "E":
			department.Add(NewEmployee(childName))
		}
	}
	company.Display(0)
}
