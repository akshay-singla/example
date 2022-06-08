package routes

import (
	"encoding/json"
	"example/src/tree"

	"github.com/gin-gonic/gin"
)

//Elements struct used for the store the user input and return the user output
type Elements struct {
	Deminision []KeyValue `json:"dim"`
	Metrics    []KeyValue `json:"metrics"`
}

//KeyValue struct used for the store the keys and values of the user input and return the user output
type KeyValue struct {
	Key   string      `json:"key"`
	Value interface{} `json:"val"`
}

// api handler to fetch users details
func InsertElement(c *gin.Context) {
	// binding body json with above variable and checking error
	data := Elements{}
	err := c.BindJSON(&data)
	if err != nil {
		// if there is some error passing bad status code
		c.JSON(400, gin.H{"error": true, "message": "Password is required field."})
		return
	}

	//set the values im the map struct so that we can easily manipulate the data
	mapd := make(map[string]interface{})
	for _, dem := range data.Deminision {
		mapd[dem.Key] = dem.Value
	}
	for _, metric := range data.Metrics {
		mapd[metric.Key] = metric.Value
	}

	byteData, err := json.Marshal(mapd)
	if err != nil {
		// if there is some error passing bad status code
		c.JSON(400, gin.H{"error": true, "message": "Internal Server Error"})
		return
	}

	//unmarshal the data into the required structure
	elements := tree.Node{}
	err = json.Unmarshal(byteData, &elements)

	//error handling
	if err != nil {
		// if there is some error passing bad status code
		c.JSON(400, gin.H{"error": true, "message": "Internal Server Error"})
		return
	}

	//Insert the element
	go tree.InsertElement(elements)

	//return the response
	c.JSON(200, gin.H{
		"error":   false,
		"message": "Record Inserted",
	})
}

// api handler to fetch users details
func QueryElement(c *gin.Context) {

	//fetcht the element on the basis of the country
	elements := tree.QueryElement(c.Query("country"))

	//return response
	c.JSON(200, gin.H{
		"error": false,
		"result": Elements{
			Deminision: []KeyValue{
				{
					Key:   "country",
					Value: elements.Country,
				},
			},
			Metrics: []KeyValue{
				{
					Key:   "webreq",
					Value: elements.WebReq,
				},
				{
					Key:   "timespent",
					Value: elements.TimeSpent,
				},
			},
		},
	})
}
