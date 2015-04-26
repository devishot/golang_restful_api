package structures


type User struct {
  FirstName         string `json:"first_name"`
  LastName          string `json:"last_name"`
  StudyYear         int    `json:"study_year"`
  StudyFaculty      int    `json:"study_faculty"`
  Age               int    `json:"age"`
  InstaUserID       int `json:"insta_user_id"`
  InstaUserName     string `json:"insta_user_name"`
}


// func (signature *Signature) valid() bool {
//     return len(signature.FirstName) > 0 &&
//         len(signature.LastName) > 0 &&
//         len(signature.Email) > 0 &&
//         signature.Age >= 18 && signature.Age <= 180 &&
//         len(signature.Message) < 140
// }
