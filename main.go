package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	IsClocked bool
}

type AttendanceLog struct {
	EmployeeID int
	ClockTime  time.Time
	IsClockIn  bool
}

var employees = map[int]Employee{
	1: {ID: 1, Name: "John", IsClocked: false},
	2: {ID: 2, Name: "Alice", IsClocked: false},
	// Add more employees as needed
}

var attendanceLogs []AttendanceLog

func findEmployeeByName(name string) (Employee, bool) {
	for _, emp := range employees {
		if emp.Name == name {
			return emp, true
		}
	}
	return Employee{}, false
}

func clockIn(employee Employee) {
	if !employee.IsClocked {
		employee.IsClocked = true
		employees[employee.ID] = employee
		clockTime := time.Now()
		attendanceLogs = append(attendanceLogs, AttendanceLog{EmployeeID: employee.ID, ClockTime: clockTime, IsClockIn: true})
		fmt.Printf("%s clocked in at %s\n", employee.Name, clockTime.Format(time.RFC3339))
	} else {
		fmt.Printf("%s is already clocked in\n", employee.Name)
	}
}

func clockOut(employee Employee) {
	if employee.IsClocked {
		employee.IsClocked = false
		employees[employee.ID] = employee
		clockTime := time.Now()
		attendanceLogs = append(attendanceLogs, AttendanceLog{EmployeeID: employee.ID, ClockTime: clockTime, IsClockIn: false})
		fmt.Printf("%s clocked out at %s\n", employee.Name, clockTime.Format(time.RFC3339))
	} else {
		clockTime := time.Now()
		fmt.Printf("%s clocked out at %s\n", employee.Name, clockTime.Format(time.RFC3339))
	}
}

func main() {
	var employeeName string
	fmt.Print("Enter employee name: ")
	fmt.Scanln(&employeeName)

	employee, found := findEmployeeByName(employeeName)
	if !found {
		fmt.Println("Employee not found")
		return
	}

	var action string
	fmt.Print("Enter action (clockin/clockout): ")
	fmt.Scanln(&action)

	switch action {
	case "clockin":
		clockIn(employee)
	case "clockout":
		clockOut(employee)
	default:
		fmt.Println("Invalid action")
	}
}
