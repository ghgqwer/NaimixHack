package main

//import "backend/internal/server"

import ("fmt";"net/http";
	"html/template"
)
import( "github.com/gorilla/mux"
    "database/sql"
		_ "github.com/go-sql-driver/mysql"
	//	"os"
	//	"strconv"
	//	"path/filepath"

	//	"mime/multipart"
	//	"io"
	//	"time"
	//	"log"
//		"gonum.org/v1/plot"
	//	"gonum.org/v1/plot/plotter"
		//"gonum.org/v1/plot/plotutil"
		//"gonum.org/v1/plot/vg"
		//"gonum.org/v1/plot/vg/draw"

	//	"sort"
		//"log"
	//	"strings"
		"github.com/gorilla/sessions"
		)

type Auth_data struct{
	Is_that_authorized bool

	Is_that_employee bool
	Employee User
	Corporation Corporation_data
}

type User struct{
   Id, Firstname, Surname,Secondname, Email, Password, Date_birth_day, Time_birth, City_of_birth, Work_experience string
   Is_that_authorized_user bool
   Is_it_admin bool
}

type Corporation_data struct{
	Id, Email, Password, Name string
	Employees []User
}
var store = sessions.NewCookieStore([]byte("super-secret-key"))

var adress_web = "http://localhost:8080"
//var authorized_user User
var sessionName = "name_session"
var adress_sql = "root:@tcp(127.127.126.50)/"
var adress_data_base_test = adress_sql + "naimix"





func populateUserFromSession(r *http.Request, sessionName string) (Auth_data, error) {
    session, err := store.Get(r, sessionName)
    if err != nil{
      panic(err)
    }

    user := User{
        Id:                    getSessionValueAsString(session, "Curret_user_id"),
        Firstname:                  getSessionValueAsString(session, "Firstname"),
        Surname:               getSessionValueAsString(session, "Surname"),
        Secondname:                 getSessionValueAsString(session, "Secondname"),
				Email:  getSessionValueAsString(session, "Email"),
        Password:              getSessionValueAsString(session, "Password"),

				Date_birth_day:              getSessionValueAsString(session, "Date_birth_day"),
				Time_birth:              getSessionValueAsString(session, "Time_birth"),
				City_of_birth:              getSessionValueAsString(session, "City_of_birth"),
				Work_experience:              getSessionValueAsString(session, "Work_experience"),


        Is_that_authorized_user: getSessionValueAsBool(session, "Is_authorized"),


    }
		var auth_data Auth_data
		auth_data.Employee = user
    return auth_data, nil
}



func saveUserToSession(r *http.Request, w http.ResponseWriter, sessionName string, user Auth_data) error{
    session, err := store.Get(r, sessionName)
    if err != nil {
        return err
    }


    session.Values["Curret_user_id"] = user.Employee.Id
    session.Values["Firstname"] = user.Employee.Firstname
    session.Values["Secondname"] = user.Employee.Secondname
    session.Values["Surname"] = user.Employee.Surname

		session.Values["Date_birth_day"] = user.Employee.Date_birth_day
		session.Values["Time_birth"] = user.Employee.Time_birth
		session.Values["City_of_birth"] = user.Employee.City_of_birth
		session.Values["Work_experience"] = user.Employee.Work_experience


    session.Values["Email"] = user.Employee.Email
		session.Values["Password"] = user.Employee.Password
    session.Values["Is_authorized"] = user.Employee.Is_that_authorized_user



    fmt.Println("Данные о сохранении:  : : :")
    fmt.Println("------------")
    return session.Save(r, w)
}


func  get_user_employee_by_id(id string) (User){

	db, err := sql.Open("mysql", adress_data_base_test)
	if err != nil{
		panic(err)
	}
	defer db.Close()


	var zapros = fmt.Sprintf("SELECT  firstname, secondname, surname, Work_experience, City_of_birth, Time_birth, Date_birth_day FROM `users` WHERE id = %s", id)
	res,_ := db.Query(zapros)
	fmt.Println(zapros)
	var user User

	for res.Next(){
		err = res.Scan( &user.Firstname,&user.Secondname,&user.Surname, &user.Work_experience,  &user.City_of_birth,  &user.Time_birth, &user.Date_birth_day)
	}
	return user
}



func  get_user_corporation_by_id(id string) (Corporation_data){

	db, err := sql.Open("mysql", adress_data_base_test)
	if err != nil{
		panic(err)
	}
	defer db.Close()


	var zapros = fmt.Sprintf("SELECT id, name FROM `corporations` WHERE id = %s", id)
	res,_ := db.Query(zapros)
	fmt.Println(zapros)
	var user Corporation_data

	for res.Next(){
		err = res.Scan(&user.Id,&user.Name)
	}
	return user
}


func returnToLastPage(r *http.Request, w http.ResponseWriter){
  referer := r.Referer()

    // Если Referer пустой, используем URL по умолчанию
    if referer == "" {
        referer = "buzhor13.ru" // Замените на ваш URL по умолчанию
    }

    // Выполняем редирект
    http.Redirect(w, r, referer, http.StatusFound)
}
func getSessionValueAsString(session *sessions.Session, key string) string {
    if val, ok := session.Values[key].(string); ok {
        return val
    }
    return ""
}

func getSessionValueAsBool(session *sessions.Session, key string) bool {
    if val, ok := session.Values[key].(bool); ok {
        return val
    }
    return false
}

func home_page(w http.ResponseWriter, r *http.Request){
      t, err := template.ParseFiles("../frontend/templates/index.html", "../frontend/templates/header.html", "../frontend/templates/footer.html")
      if err != nil{
        fmt.Fprintf(w, err.Error())

      }

			fmt.Println("GO")
      //var data Data_for_personal_page

      //authorized_user, err  := populateUserFromSession(r, sessionName)
      //data.Authorized_user_data  = authorized_user
      t.ExecuteTemplate(w, "index", nil)
}


func employee_profile(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
  w.WriteHeader(http.StatusOK)
//  fmt.Fprintf(w, "Category: %v\n", vars["id_user"])
  var current_user_id = vars["id_employee"]

	fmt.Println("prof")
	fmt.Println(current_user_id)
	user := get_user_employee_by_id(current_user_id)

	//user, _  := populateUserFromSession(r, sessionName)
	fmt.Println(user)
}

func employee_authorization_page(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("../frontend/templates/employee_authorization_page.html", "../frontend/templates/header.html", "../frontend/templates/footer.html")
	if err != nil{
		fmt.Fprintf(w, err.Error())

	}
		db, err := sql.Open("mysql", adress_data_base_test)
		if err != nil{
			panic(err)
		}
		defer db.Close()


		if r.Method == http.MethodPost {
			email := r.FormValue("email")
			password := r.FormValue("password")
			var zapros = fmt.Sprintf("SELECT id, firstname, secondname, surname FROM `users` WHERE email = %s AND password = %s", email, password)
			res,_ := db.Query(zapros)
			fmt.Println(zapros)
			var user Auth_data
			user.Is_that_employee = true
			for res.Next(){
				err = res.Scan(&user.Employee.Id,&user.Employee.Firstname,&user.Employee.Secondname,&user.Employee.Surname)
			}

			user.Employee.Is_that_authorized_user = true
			saveUserToSession(r, w, sessionName, user)
			http.Redirect(w, r, "/employee_profile/" + user.Employee.Id, http.StatusSeeOther)
		}


	t.ExecuteTemplate(w, "employee_authorization_page", nil)
}

func employee_registration_page(w http.ResponseWriter, r *http.Request){
      t, err := template.ParseFiles("../frontend/templates/employee_registration_page.html", "../frontend/templates/header.html", "../frontend/templates/footer.html")
      if err != nil{
        fmt.Fprintf(w, err.Error())

      }


				db, err := sql.Open("mysql", adress_data_base_test)
			  if err != nil{
			    panic(err)
			  }
			  defer db.Close()


			if r.Method == http.MethodPost {
				name := r.FormValue("name")
				surname := r.FormValue("surname")
				//weight := r.FormValue("weight")
				secondname := r.FormValue("secondname")
				email := r.FormValue("email")

				Date_birth_day := r.FormValue("Date_birth_day")
				Time_birth := r.FormValue("Time_birth")
				City_of_birth := r.FormValue("City_of_birth")
				Work_experience := r.FormValue("Work_experience")

				password := r.FormValue("password")
				job_title := r.FormValue("dropdown2")






					result, err := db.Exec("insert into naimix.users ( `firstname`,	`surname`, `secondname`, `job_title`, `email`, `password`, `Date_birth_day`, `Time_birth`, `City_of_birth`, `Work_experience`) values (?, ?,?,?, ?, ?, ?,?,?,?)",name, surname, secondname, job_title, email, password, Date_birth_day, Time_birth, City_of_birth, Work_experience)

					fmt.Println(result)
					if(err != nil){
						fmt.Println(err)
					}



			}

      //var data Data_for_personal_page

      //authorized_user, err  := populateUserFromSession(r, sessionName)
      //data.Authorized_user_data  = authorized_user
      t.ExecuteTemplate(w, "employee_registration_page", nil)
}




func corporation_registration_page(w http.ResponseWriter, r *http.Request){
      t, err := template.ParseFiles("../frontend/templates/corporation_registration_page.html", "../frontend/templates/header.html", "../frontend/templates/footer.html")
      if err != nil{
        fmt.Fprintf(w, err.Error())

      }


				db, err := sql.Open("mysql", adress_data_base_test)
			  if err != nil{
			    panic(err)
			  }
			  defer db.Close()


			if r.Method == http.MethodPost {
				name := r.FormValue("name")
				email := r.FormValue("email")
				//weight := r.FormValue("weight")
				password := r.FormValue("password")

					result, err := db.Exec("insert into naimix.corporations ( `name`,	`email`, `password`) values (?, ?,?)",name, email, password)

					fmt.Println(result)
					if(err != nil){
						fmt.Println(err)
					}



			}

      //var data Data_for_personal_page

      //authorized_user, err  := populateUserFromSession(r, sessionName)
      //data.Authorized_user_data  = authorized_user
      t.ExecuteTemplate(w, "corporation_registration_page", nil)
}

func corporation_profile(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
  w.WriteHeader(http.StatusOK)
//  fmt.Fprintf(w, "Category: %v\n", vars["id_user"])
  var current_user_id = vars["id_corporation"]

	fmt.Println("prof")
	fmt.Println(current_user_id)
	user := get_user_corporation_by_id(current_user_id)

	//user, _  := populateUserFromSession(r, sessionName)
	fmt.Println(user)
}

func corporation_authorization_page(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("../frontend/templates/corporation_authorization_page.html", "../frontend/templates/header.html", "../frontend/templates/footer.html")
	if err != nil{
		fmt.Fprintf(w, err.Error())

	}




		db, err := sql.Open("mysql", adress_data_base_test)
		if err != nil{
			panic(err)
		}
		defer db.Close()


		if r.Method == http.MethodPost {
			email := r.FormValue("email")
			password := r.FormValue("password")
			var zapros = fmt.Sprintf("SELECT id, name FROM `corporations` WHERE email = %s AND password = %s", email, password)
			res,_ := db.Query(zapros)
			fmt.Println(zapros)
			var user Auth_data
			user.Is_that_employee = true
			for res.Next(){
				err = res.Scan(&user.Corporation.Id,&user.Corporation.Name)
			}


			saveUserToSession(r, w, sessionName, user)
			http.Redirect(w, r, "/corporation_profile/" + user.Corporation.Id, http.StatusSeeOther)
		}


	t.ExecuteTemplate(w, "corporation_authorization_page", nil)
}


func main() {

	store.Options = &sessions.Options{
        Path:     "/",       // Путь для куков
        MaxAge:   86400 * 7, // Время жизни куков - 7 дней
        HttpOnly: true,      // Запрет доступа к кукам через JavaScript
        Secure:   false,     // Для HTTP должен быть false
    }


 	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r := mux.NewRouter()
	fmt.Println("Start")



	r.HandleFunc("/", home_page)
	r.HandleFunc("/employee_registration_page", employee_registration_page)
	r.HandleFunc("/employee_authorization_page", employee_authorization_page)
	r.HandleFunc("/employee_profile/{id_employee}", employee_profile)

	r.HandleFunc("/corporation_registration_page", corporation_registration_page)
	r.HandleFunc("/corporation_authorization_page", corporation_authorization_page)
	r.HandleFunc("/corporation_profile/{id_corporation}", corporation_profile)



  http.Handle ("/", r)
  http.ListenAndServe(":8080", nil)


}
