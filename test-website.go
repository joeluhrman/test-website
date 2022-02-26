package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type MuscleGroup string

const (
	Legs MuscleGroup = "Legs"
	Push MuscleGroup = "Push"
	Pull MuscleGroup = "Pull"
)

type Page struct {
	Title string
	Body  []byte
}

type ExerciseType struct {
	Name   string
	BdyPrt MuscleGroup
}

type Exercise struct {
	Type  ExerciseType
	Sets  int
	Reps  int
	Notes string
}

type Workout struct {
	Date      string
	Exercises []Exercise
	Notes     string
}

func newExercise(name string, sets int, reps int, notes string) *Exercise {
	exercise := Exercise{Name: name, Sets: sets, Reps: reps, Notes: notes}
	return &exercise
}

func newWorkout(date string, exercises []Exercise, notes string) *Workout {
	workout := Workout{Date: date, Exercises: exercises, Notes: notes}
	return &workout
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, _ := os.ReadFile(filename)
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	p, _ := loadPage(title)
	t, _ := template.ParseFiles("homepage.html")
	t.Execute(w, p)
}

func main() {
	//http.HandleFunc("/", viewHandler)
	//log.Fatal(http.ListenAndServe(":8080", nil))

	exercise1 := newExercise("Squat", 3, 5, "good")
	exercise2 := newExercise("Bench", 3, 5, "yes")
	exercise3 := newExercise("Deadlift", 1, 5, "oh")

	exercises := make([]Exercise, 3, 3)
	exercises[0] = *exercise1
	exercises[1] = *exercise2
	exercises[2] = *exercise3

	workout := Workout{Date: "2/25/2022", Exercises: exercises, Notes: "yuh"}

	//fmt.Println(exercise.Name, exercise.Sets, exercise.Reps)
	fmt.Println(workout.Date, workout.Exercises, workout.Notes)
}
