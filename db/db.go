package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	Database *sql.DB
}

type Employee struct {
	Id     int    `json:"id"`
	Gender string `json:"gender"`
}

type EmployeeJob struct {
	Id          int    `json:"id"`
	Employee_id int    `json:"employee_id"`
	Department  string `json:"department"`
	Job_title   string `json:"job_title"`
}

func (d *DB) InitDB(databaseName string) error {
	db, err := sql.Open("sqlite3", databaseName)
	if err != nil {
		return err
	}
	d.Database = db
	return nil
}

func (d *DB) AddEmployeeJob(employeeID int, department string, jobTitle string) (int64, error) {
	tx, err := d.Database.Begin()
	if err != nil {
		return 0, err
	}

	stmt, err := tx.Prepare("INSERT INTO EmployeeJobs(employee_id, department, job_title) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(employeeID, department, jobTitle)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

func (d *DB) GetEmployees() ([]Employee, error) {
	query := "SELECT id, gender FROM employees "

	rows, err := d.Database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []Employee

	for rows.Next() {
		var emp Employee
		if err := rows.Scan(&emp.Id, &emp.Gender); err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}

	if len(employees) == 0 {
		return []Employee{}, nil
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func (d *DB) GetEmployeeJobs() ([]EmployeeJob, error) {
	query := "SELECT id, employee_id, department, job_title FROM employeejobs"

	rows, err := d.Database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []EmployeeJob

	for rows.Next() {
		var job EmployeeJob
		if err := rows.Scan(&job.Id, &job.Employee_id, &job.Department, &job.Job_title); err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}

	if len(jobs) == 0 {
		return []EmployeeJob{}, nil
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return jobs, nil
}

func (d *DB) GetEmployeeJobByEmployeeId(employeeID int) ([]EmployeeJob, error) {
	query := "SELECT id, employee_id, department, job_title FROM employeejobs WHERE employee_id = ?"

	rows, err := d.Database.Query(query, employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []EmployeeJob

	for rows.Next() {
		var job EmployeeJob
		if err := rows.Scan(&job.Id, &job.Employee_id, &job.Department, &job.Job_title); err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}

	if len(jobs) == 0 {
		return []EmployeeJob{}, nil
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return jobs, nil
}

func (d *DB) UpdateEmployeeJob(id int, department string, job_title string) (bool, error) {
	tx, err := d.Database.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("update employeejobs set department=?, job_title=? where id = ?")
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(department, job_title, id)
	if err != nil {
		return false, err
	}

	err = tx.Commit()

	return true, err
}
