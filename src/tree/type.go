package tree

//Node struct represents the tree structure and its collect the node information
type Node struct {
	WebReq    int    `json:"webreq"`
	TimeSpent int    `json:"timespent"`
	Devive    string `json:"devive"`
	Country   string `json:"country"`
	Children  []*Node
}

//parentNode is a variable that is used for store the information
var parentNode = Node{}
