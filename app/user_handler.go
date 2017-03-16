package app

import (
	"net/http"
	"fmt"
	"gopkg.in/mgo.v2"
	"errors"
	"time"
)

type UserHandler struct {
	Db mgo.Database
	Users UserRepository
}

type email string
type sessionid string

var authUsers map[sessionid]email

var (
	notLoggedIn = errors.New("Not logged in")
)

func init() {
	authUsers = make(map[sessionid]email)

}

func setSessionId(w http.ResponseWriter, sessionId string) {
	http.SetCookie(w, &http.Cookie{Name: "sessionId", Value: sessionId})
}

func addUserSession(w http.ResponseWriter, gUser GoogleUser) {
	token := "rto" //todo: fix this shit. Generate some random hash
	setSessionId(w, token)
	authUsers[sessionid(token)] = email(gUser.Email)
}

func removeUserSession(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "sessionId", Value: "", Expires: time.Time{} })
	delete(authUsers, getSessionId(r))
}

func getSessionId(r *http.Request) sessionid {
	cookie, err := r.Cookie("sessionId")
	if err != nil {
		return ""
	}

	return sessionid(cookie.Value)
}

func (u UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	user, err := u.getUser(r)

	if err == nil {
		fmt.Fprintf(w,"<html><body>Welcome %v. <a href='/logout'>logout</a> </body></html>", user.Email)
		return
	}

	if err == notLoggedIn {
		fmt.Fprint(w,"<html><body>You are not logged in:  <a href='/login'>login here</a></body></html>")
	}


}

func LogOutHandler(w http.ResponseWriter, r *http.Request) {
	removeUserSession(w, r)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func (u UserHandler) getUser(r *http.Request) (User, error){

	userEmail, ok := authUsers[getSessionId(r)]
	if !ok {
		return User{}, notLoggedIn
	}

	user, err := u.Users.GetUser(string(userEmail))

	if err == nil {
		return user, nil
	}


	if err == mgo.ErrNotFound {
		user, err := u.Users.New(string(userEmail))
		if err != nil {
			return User{}, err
		}
		return user, nil
	}

	return User{}, err
}

