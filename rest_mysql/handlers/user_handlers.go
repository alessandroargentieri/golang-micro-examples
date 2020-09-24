package handlers

  import (
	  "log"
	  "net/http"
	  "encoding/json"
	  "goql/dto"
	  "goql/dao"
	  "goql/model"
	  "math/rand"
      "time"
  )




func AddUser(w http.ResponseWriter, r *http.Request) {
	  jsonDecoder := json.NewDecoder(r.Body)

	  userDto := dto.UserDto{}
	  err := jsonDecoder.Decode(&userDto)
	  if err != nil {
	    log.Fatal(err)
	  }

	  userId := randomId()
	  user := model.User{ Id: userId, 	Name: userDto.Name,	  	Email: userDto.Email	  }

	  dao.Add(user)

	  w.Write([]byte("User added."))
}


func randomId() int {
  	rand.Seed(time.Now().UnixNano())
    return rand.Intn(999)
}


func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	
	  var users []model.User = dao.GetAll()

	  jsonBytes, _ := json.Marshal(users)
	  log.Println(string(jsonBytes))
	  w.Write([]byte(jsonBytes))
}

