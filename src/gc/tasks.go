package gc

import (
	"context"
	"encoding/base64"
	"log"
	"sync"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/classroom/v1"
	"google.golang.org/api/option"
)

func getTask(s *classroom.StudentSubmission, srv *classroom.Service, taskstr *string) {
	var overdue string

	if s.Late {
		overdue = "true"
	} else {
		overdue = "false"
	}

	gctask, err := srv.Courses.CourseWork.Get(s.CourseId, s.CourseWorkId).Do()

	if err != nil {
		log.Fatalf("500: Unable to retrieve task. %v", err)
	}

	title := base64.StdEncoding.EncodeToString([]byte(gctask.Title))
	desc := base64.StdEncoding.EncodeToString([]byte(gctask.Description))

	task := "{\"task\":\"" + title + "\","
	task += "\"class\":\"" + "CLASS" + "\","
	task += "\"desc\":\"" + desc + "\","
	task += "\"link\":\"" + gctask.AlternateLink + "\","
	task += "\"res\":\"" + "RES" + "\","
	task += "\"duedate\":\"" + "TIMESTAMP" + "\","
	task += "\"overdue\":\"" + overdue + "\"},"

	*taskstr += task
}

func getSubmissions(c *classroom.Course, srv *classroom.Service, taskstr *string, wg *sync.WaitGroup) {
	defer wg.Done()
	submissions, err := srv.Courses.CourseWork.StudentSubmissions.List(c.Id, "-").Do()

	if err != nil {
		log.Fatalf("500: Unable to retrieve submissions. %v", err)
	}

	for _, s := range submissions.StudentSubmissions {
		go getTask(s, srv, taskstr)
	}
}

func GetTasks(secret []byte) []byte {
	var wg sync.WaitGroup
	var tasks []byte

	ctx := context.Background()

	config, err := google.ConfigFromJSON(secret, classroom.ClassroomCoursesReadonlyScope)

	if err != nil {
		log.Fatalf("500: Unable to parse client secret file to config: %v", err)
	}

	client := getClient(config, secret)
	srv, err := classroom.NewService(ctx, option.WithHTTPClient(client))

	if err != nil {
		log.Fatalf("500: Unable to create classroom Client %v", err)
	}

	r, err := srv.Courses.List().CourseStates("ACTIVE").Do()
	if err != nil {
		log.Fatalf("500: Unable to retrieve courses. %v", err)
	}

	if len(r.Courses) > 0 {
		taskstr := "["

		for _, c := range r.Courses {
			wg.Add(1)
			go getSubmissions(c, srv, &taskstr, &wg)
		}

		wg.Wait()
		tasks = []byte(taskstr)
		tasks[len(tasks)-1] = byte(']')
	}

	return tasks
}
