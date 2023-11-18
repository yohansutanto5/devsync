package model

import "time"

type Student struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	FirstName   string
	LastName    string
	Enrollments []Enrollment // One-to-Many: One student can be enrolled in multiple courses
}

type AddStudentIn struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
}

type Course struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	Title      string
	TeacherID  uint         // Many-to-One: Many courses are taught by one teacher
	Enrollment []Enrollment `gorm:"many2many:enrollments"` // Many-to-Many: Many students can enroll in many courses
}

type Teacher struct {
	ID           uint `gorm:"primaryKey;autoIncrement"`
	Name         string
	Courses      []Course // One-to-Many: One teacher can teach multiple courses
	DepartmentID uint
}

type Enrollment struct {
	ID             uint `gorm:"primaryKey;autoIncrement"`
	StudentID      uint
	CourseID       uint
	EnrollmentDate time.Time
	Student        Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // One-to-One: Each enrollment is associated with one student
	Course         Course  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // One-to-One: Each enrollment is associated with one course
}

type Department struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Name     string
	Teachers []Teacher // One-to-Many: One department can have multiple teachers
}

type CustomLog struct {
	TransactionID int    `json:"transactionID"`
	Code          string `json:"code"`
	Status        int    `json:"status"`
	Message       string `json:"message"`
	Method        string `json:"method"`
	Path          string
	Duration      time.Duration
	ClientIp      string
	Agent         string

	Data interface{} `json:"data"`
}
