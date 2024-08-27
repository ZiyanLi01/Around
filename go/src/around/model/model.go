package model

type Post struct {
    Id      string `json:"id"`
    User    string `json:"user"`
    Message string `json:"message"`
    Url    string `json:"url"`
    Type    string `json:"type"`
    Caption string `json:"caption"` 
}

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Age      int64  `json:"age"`
    Gender   string `json:"gender"`
}

type Image struct {
    Id          string `json:"id"`
    Url         string `json:"url"`
    Description string `json:"description"`
    UserId      string `json:"user_id"`
}

