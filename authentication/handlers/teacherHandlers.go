package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/piyushkumar/authentication/authentication/models"
)

func AddTeacherToCourseHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    courseID := vars["courseID"]

    var teacherID string
    if err := json.NewDecoder(r.Body).Decode(&teacherID); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Assuming teacherID is valid and exists in Teachers
    models.CourseTeachers[courseID] = append(models.CourseTeachers[courseID], teacherID)
    json.NewEncoder(w).Encode("Teacher added to course successfully")
}

func GetTeachersOfCourseHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    courseID := vars["courseID"]

    teacherIDs, exists := models.CourseTeachers[courseID]
    if !exists {
        http.Error(w, "No teachers found for this course", http.StatusNotFound)
        return
    }

    var teachers []models.Teacher
    for _, id := range teacherIDs {
        for _, teacher := range models.Teachers {
            if teacher.ID == id {
                teachers = append(teachers, teacher)
                break
            }
        }
    }

    json.NewEncoder(w).Encode(teachers)
}
