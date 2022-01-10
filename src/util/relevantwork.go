package util

import (
	"main/call"
	"sync"

	"google.golang.org/api/classroom/v1"
)

type WorkAndSubmission struct {
	Work       *classroom.CourseWork
	Submission *classroom.StudentSubmission
	Course     *classroom.Course
}

func FetchWorksAndSubmissions(srv *classroom.Service, courses []*classroom.Course) ([]*WorkAndSubmission, error) {
	ch_ci := make(chan *WorkAndSubmission)
	var wg_ci sync.WaitGroup

	// for each course
	for _, course := range courses {
		// add a new goroutine to the wait group
		wg_ci.Add(1)

		// call the function that will fetch the relevant for this course in parallel
		go func(course *classroom.Course) {
			defer wg_ci.Done() // release one from wait group on function end

			subs, err := call.ListAllActiveSubmissions(srv, course)
			if err != nil {
				return
			}

			// waitgroup for all coursework fetches
			var wg_cw sync.WaitGroup

			// iterate all submissions
			for _, sub := range subs {
				// add one goroutine to track in wait group
				wg_cw.Add(1)
				// for a submission, spawn a goroutine
				go func(sub *classroom.StudentSubmission, course *classroom.Course) {
					defer wg_cw.Done() // release one from wait group on function end

					// call the api (this takes a long time)
					work, err := call.GetCourseWorkByID(srv, course, sub.CourseWorkId)
					if err != nil {
						return
					}

					// send back a complete info struct on complete info channel
					ch_ci <- &WorkAndSubmission{work, sub, course}
				}(sub, course)
			}

			// at this point a lot of goroutines are spawned
			// wait for them to finish
			wg_cw.Wait()
		}(course)
	}

	go func() {
		// this goroutine makes it so the channel is closed when all of the goroutines are done
		wg_ci.Wait()
		close(ch_ci)
	}()

	// collect all the complete info structs
	var out []*WorkAndSubmission

	for ci := range ch_ci {
		// add something that came in on the channel to the output
		out = append(out, ci)
	}

	return out, nil
}
