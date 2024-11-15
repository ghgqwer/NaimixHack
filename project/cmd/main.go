package main

//import "backend/internal/server"

import ("fmt";"net/http";
	"html/template"
)
import( "github.com/gorilla/mux"
  //  "database/sql"
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

type User struct{
   Id, Name, Surname, Email, Password string
   Is_that_authorized_user bool
   Is_it_admin bool
}
var store = sessions.NewCookieStore([]byte("super-secret-key"))





func populateUserFromSession(r *http.Request, sessionName string) (User, error) {
    session, err := store.Get(r, sessionName)
    if err != nil{
      panic(err)
    }

    user := User{
        Id:                    getSessionValueAsString(session, "curret_user_id"),
        Name:                  getSessionValueAsString(session, "name"),
        Surname:               getSessionValueAsString(session, "surname"),
        Email:                 getSessionValueAsString(session, "email"),
        Password:              getSessionValueAsString(session, "password"),
        Is_that_authorized_user: getSessionValueAsBool(session, "is_authorized"),
        Is_it_admin:  getSessionValueAsBool(session, "adminAuth"),
    }

    return user, nil
}



func saveUserToSession(r *http.Request, w http.ResponseWriter, sessionName string, user User) error{
    session, err := store.Get(r, sessionName)
    if err != nil {
        return err
    }

    session.Values["curret_user_id"] = user.Id
    session.Values["name"] = user.Name
    session.Values["surname"] = user.Surname
    session.Values["email"] = user.Email
    session.Values["password"] = user.Password
    session.Values["is_authorized"] = user.Is_that_authorized_user
    session.Values["adminAuth"] = user.Is_it_admin


    fmt.Println("Данные о сохранении:  : : :")
    fmt.Println("------------")
    return session.Save(r, w)
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
  http.Handle ("/", r)
  http.ListenAndServe(":8080", nil)


}
