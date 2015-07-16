package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/sessions"
)

type index struct {
	Photos []string
}

//////LOGIN PAGE/////////
/////////////////////////
var store = sessions.NewCookieStore([]byte("something-very-secret"))

func pleaseLogIn(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	username := req.FormValue("username")
	password := req.FormValue("password")
	if req.Method == "POST" {
		if username == "123" && password == "123" {
			session.Values["logged_in"] = "YES"
			// save session
			session.Save(req, res)
			// redirect to main page
			http.Redirect(res, req, "/admin", 302)
			return
		} else {
			http.Error(res, "Wrong username/password", 401)
			return
		}

	}
	tpl, err := template.ParseFiles("admin.tpl")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(res, map[string]interface{}{
		"username": username,
		"password": password,
	})
}

//////LOGOUT FUNC/////////
/////////////////////////
func logout(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	delete(session.Values, "logged_in")
	store.Save(req, res, session)
	http.Redirect(res, req, "/", http.StatusSeeOther)
}

//////UPLOAD PAGE////////
/////////////////////////
func uploadPage(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		src, hdr, err := req.FormFile("fileUploads")
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer src.Close()

		//create a new file
		dst, err := os.Create("img/" + hdr.Filename)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer dst.Close()
		io.Copy(dst, src)
	}

	tpl, err := template.ParseFiles("loggedIn.tpl")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(res, nil)
}

//////UPLOAD FILE////////
/////////////////////////
func myFiles() []string {
	files := make([]string, 0)
	filepath.Walk("img", func(p string, info os.FileInfo, _ error) error {
		if strings.HasSuffix(p, ".jpeg") || strings.HasSuffix(p, ".jpg") {
			files = append(files, p)
		}
		return nil
	})
	return files
}

//////MAIN FUNC//////////
/////////////////////////
func main() {
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		tpl := template.New("index.tpl")
		tpl = tpl.Funcs(template.FuncMap{
		//function go here {},
		})
		tpl, err := tpl.ParseFiles("index.tpl")
		if err != nil {
			log.Fatalln(err)
		}
		err = tpl.Execute(res, index{
			Photos: myFiles(),
		})
	})

	http.HandleFunc("/admin", func(res http.ResponseWriter, req *http.Request) {
		session, _ := store.Get(req, "session")
		str, _ := session.Values["logged_in"].(string)
		// TODO: get from session
		loggedIn := str == "YES"

		if !loggedIn {
			pleaseLogIn(res, req)
		} else {
			uploadPage(res, req)
		}
	})
	http.HandleFunc("/logout", func(res http.ResponseWriter, req *http.Request) {
		logout(res, req)
	})

	http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
}
