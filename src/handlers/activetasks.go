package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"main/auth"
	"main/call"
	"main/util"
	"net/http"
	"strconv"
	"time"

	"github.com/buger/jsonparser"
)

func (h *BaseHandler) ActiveTasks(w http.ResponseWriter, r *http.Request) {
	// read body to byte buffer
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := ReadToken(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid token"))
		return
	}

	client, err := auth.GetClient(h.Config, token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("unable to create client using provided token"))
		return
	}

	srv, err := auth.GetService(client)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	// get courses for student
	courses, err := call.ListCourses(srv)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("failed to retrieve courses, check credentials"))
		log.Println(err.Error())
	}

	worksubs, err := util.FetchWorksAndSubmissions(srv, courses)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to retrive active tasks"))
		log.Println(err.Error())
	}

	// construct output object
	arr_out := util.MakePlaceholderJSONArray(len(worksubs))

	for ws_idx, ws := range worksubs {
		el := []byte{'{', '}'}

		// process materials for task
		arr_mats := util.MakePlaceholderJSONArray(len(ws.Work.Materials))

		for mat_idx, mat_croom := range ws.Work.Materials {
			mat := util.ParseClassroomMaterial(mat_croom)

			if mat == nil {
				log.Printf(
					"warning: mat fix failed (user: %s, course: %s, mat: %+v)",
					ws.Submission.UserId, ws.Work.CourseId, mat_croom,
				)
				continue
			}

			// construct material object
			mat_obj := []byte{'{', '}'}

			mat_obj, err := jsonparser.Set(mat_obj, []byte(strconv.Quote(mat.Title)), "title")
			if err != nil {
				log.Printf(
					"warning: mat objectification failed (user: %s, course: %s, mat: %+v)",
					ws.Submission.UserId, ws.Work.CourseId, mat_croom,
				)
				continue
			}

			mat_obj, err = jsonparser.Set(mat_obj, []byte(strconv.Quote(mat.Link)), "link")
			if err != nil {
				log.Printf(
					"warning: mat objectification failed (user: %s, course: %s, mat: %+v)",
					ws.Submission.UserId, ws.Work.CourseId, mat_croom,
				)
				continue
			}

			aidx := fmt.Sprintf("[%d]", mat_idx)
			arr_mats, err = jsonparser.Set(arr_mats, mat_obj, aidx)
			if err != nil {
				log.Printf(
					"warning: mat insert failed (user: %s, course: %s, mat: %+v, idx: %d)",
					ws.Submission.UserId, ws.Work.CourseId, mat_croom, mat_idx,
				)
			}
		}

		el, err := jsonparser.Set(el, arr_mats, "materials")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())
		}

		// add attributes
		el, err = jsonparser.Set(el, util.StringToEscBuf(ws.Work.Id), "id")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())
		}

		el, err = jsonparser.Set(el, util.StringToEscBuf(ws.Work.Title), "name")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())
		}

		el, err = jsonparser.Set(el, util.StringToEscBuf(ws.Course.Id), "courseId")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())
		}

		el, err = jsonparser.Set(el, util.StringToEscBuf(ws.Course.Name), "courseName")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())
		}

		el, err = jsonparser.Set(el, util.StringToEscBuf(ws.Work.Description), "description")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())
		}

		due := util.ParseClassroomTime(ws.Work.DueDate, ws.Work.DueTime)
		var due_s []byte
		if due != nil {
			due_s = []byte(fmt.Sprint(due.UTC().Unix())) // literal int, unquoted
		} else {
			due_s = []byte("null")
		}

		el, err = jsonparser.Set(el, due_s, "dueOn")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())
		}

		created, err := time.Parse(time.RFC3339, ws.Work.CreationTime)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())
		}

		el, err = jsonparser.Set(el, []byte(fmt.Sprint(created.UTC().Unix())), "setOn")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())
		}

		el, err = jsonparser.Set(el, util.StringToEscBuf(ws.Submission.Id), "submission", "id")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())
		}

		el, err = jsonparser.Set(el, util.StringToEscBuf(ws.Submission.State), "submission", "state")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())
		}

		el, err = jsonparser.Set(el, util.StringToEscBuf(strconv.FormatBool(ws.Submission.Late)), "submission", "late")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())
		}

		aidx := fmt.Sprintf("[%d]", ws_idx)
		arr_out, err = jsonparser.Set(arr_out, el, aidx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())
		}
	}

	w.Write(arr_out)
}
