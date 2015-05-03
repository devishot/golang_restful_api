package structures


type User struct {
  FullName         string  `json:"full_name"        sql:"full_name"`
  StudyYear         int    `json:"study_year"       sql:"study_year"`
  StudyFaculty      string `json:"study_faculty"    sql:"study_faculty"`
  InstaUserID       int    `json:"insta_user_id"    sql:"insta_user_id"`
  InstaUserName     string `json:"insta_user_name"  sql:"insta_user_name"`
}


// func (signature *Signature) valid() bool {
//     return len(signature.FirstName) > 0 &&
//         len(signature.LastName) > 0 &&
//         len(signature.Email) > 0 &&
//         signature.Age >= 18 && signature.Age <= 180 &&
//         len(signature.Message) < 140
// }
