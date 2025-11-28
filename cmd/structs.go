package cmd

type Activity struct {
	Id         string  `json:"id"`
	Type       string  `json:"type"`
	Public     bool    `json:"public"`
	Created_at string  `json:"created_at"`
	Actor      Actor   `json:"actor"`
	Repo       Repo    `json:"repo"`
	Payload    Payload `json:"payload"`
	Org        Org     `json:"org"`
}

type Actor struct {
	Id            int    `json:"id"`
	Login         string `json:"login"`
	Display_login string `json:"display_login"`
	Gravatar_id   string `json:"gravatar_id"`
	Url           string `json:"url"`
	Avatar_url    string `json:"avatar_url"`
}

type Repo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Payload struct {
	Action        string   `json:"action"`
	Push_id       int      `json:"push_id"`
	Size          int      `json:"size"`
	Distinct_size int      `json:"distinct_size"`
	Ref           string   `json:"ref"`
	Head          string   `json:"head"`
	Before        string   `json:"before"`
	Commits       []Commit `json:"commits"`
	Issue         Issue    `json:"issue"`
	Comment       Comment  `json:"comment"`
}

type Commit struct {
	Sha    string `json:"sha"`
	Author struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"author"`
	Message  string `json:"message"`
	Distinct bool   `json:"distinct"`
	Url      string `json:"url"`
}

type Org struct {
	Id          int    `json:"id"`
	Login       string `json:"login"`
	Gravatar_id string `json:"gravatar_id"`
	Url         string `json:"url"`
	Avatar_url  string `json:"avatar_url"`
}

type Issue struct {
	Id     int    `json:"id"`
	Number int    `json:"number"`
	Title  string `json:"title"`
	User   struct {
		Id    int    `json:"id"`
		Login string `json:"login"`
		Url   string `json:"url"`
	} `json:"user"`
	State      string `json:"state"`
	Comments   int    `json:"comments"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Url        string `json:"url"`
}

type Comment struct {
	Id   int `json:"id"`
	User struct {
		Id    int    `json:"id"`
		Login string `json:"login"`
		Url   string `json:"url"`
	} `json:"user"`
	Body       string `json:"body"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Url        string `json:"url"`
}
