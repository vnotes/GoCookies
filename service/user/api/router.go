package api

import "net/http"

// swagger:parameters get-user
type GetUserReq struct {
	// in: query
	ID int `json:"id"`
}

type UserInfo struct {
	Name string `json:"name"`
}

// return user info
// swagger:response userResponse
type UserResponse struct {
	// in: body
	Result []*UserInfo `json:"result"`
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /api/user user get-user
	//
	// Lists users filtered by some parameters.
	//
	//     Responses:
	//       200: userResponse
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

}
