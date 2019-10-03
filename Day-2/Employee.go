package main

import (
	"fmt"
  
)


type  EmergencyContact struct {
	ename string
	phone string

}
type  contactInfo struct {
	city string
	phone string
	email string
	econtact []EmergencyContact 
}

type payroll struct {

	basicsalary float64
	variablePay float64
	hra float64
	lta float64
	professional_tax float64
	pf float64
}

type employee struct {
	fname string
	lname string
	empid int
	dob string
	contactInfo
	payroll
	leave int	

}

func (e employee) AddEmergencyContact(ename string , phone string)  {

/*	if (e.econtact).size() == 10 {

	}*/
	var tmp EmergencyContact
	tmp.ename = ename
	tmp.phone = phone

	e.econtact = append(e.econtact, tmp)

}


func (e employee) PrintEmergencyContact()  {
	
	for index, tmp := range e.econtact {
		fmt.Println(index ," Name : ",tmp.ename , " Phone",tmp.phone)
	}
	
}
	
func (e employee) CalculateGS() float64 {
/*
	basicsalary float64
	variablePay float64
	hra float64
	lta float64
	professional_tax float64
	pf float64
*/

	var gs float64
	gs = (e.payroll).basicsalary+ (e.payroll).hra + (e.payroll).lta  - (e.payroll).professional_tax - (e.payroll).variablePay - (e.payroll).pf
	return gs
}

func (e employee) CalculateCTC() float64 {


	ctc := e.CalculateGS() *12
	return ctc
}

func (e employee) LeaveBalance() int {


	return 23-e.leave 
}
	
func (e employee) Print (){
	fmt.Println( "Emplyee ID = ",e.empid)
	fmt.Println( "Name  =",e.fname ," " , e.lname)
	fmt.Println( "DOB = ", e.dob)
	fmt.Println( "Contact Info " ,e.contactInfo)
	fmt.Println( "Leave Balance " ,e.LeaveBalance())
	fmt.Println( "Gorss Salary ", e.CalculateGS())
	fmt.Println( "Total CTC " ,e.CalculateCTC())
}
