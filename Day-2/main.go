package main


func main() {

	var e employee
	e.empid = 101
	e.fname ="Dharmesh"
	e.lname = "Patil"
	e.dob = "2006-01-02"
	e.payroll.basicsalary = 25000
	e.payroll.variablePay  =2000
	e.payroll.hra =3000
	e.payroll.lta =500
	e.payroll.professional_tax =200 
	e.payroll.pf =3000
	e.contactInfo.city = "Pune" 
	e.contactInfo.phone ="123456789"
	e.contactInfo.email = "abc@gmail.com"

	e.AddEmergencyContact("DDd","1234")
	e.AddEmergencyContact("www","1234")
	e.AddEmergencyContact("wwww","1234")

	e.Print()

}
