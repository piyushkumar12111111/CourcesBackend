package models

type Teacher struct {
    ID        string `json:"id"`
    Name      string `json:"name"`
    Specialty string `json:"specialty"`
}

var Teachers = []Teacher{
    {ID: "t1", Name: "Alice Johnson", Specialty: "Mathematics"},
    {ID: "t2", Name: "Bob Smith", Specialty: "Science"},
    {ID: "t3", Name: "Charlie Davis", Specialty: "Literature"},
}

// CourseTeacher represents the association between a course and a teacher.
// Key: CourseID, Value: Slice of TeacherIDs
var CourseTeachers = map[string][]string{
    "c1": {"t1", "t2"}, // Assuming "c1" is an existing course ID
    "c2": {"t3"},       // Assuming "c2" is an existing course ID
}
