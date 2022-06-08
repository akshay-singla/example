package tree

import "testing"

func TestInsertElement(t *testing.T) {

	//element using IN country and mobile Devive
	element := Node{
		WebReq:    120,
		TimeSpent: 120,
		Country:   "IN",
		Devive:    "Mobile",
	}
	//InsertElement the element into the tree
	InsertElement(element)

	//element using IN country and web Devive
	element = Node{
		WebReq:    120,
		TimeSpent: 120,
		Country:   "IN",
		Devive:    "Web",
	}
	//InsertElement the element into the tree
	InsertElement(element)

	//element using IN country and web Devive
	element = Node{
		WebReq:    60,
		TimeSpent: 160,
		Country:   "IN",
		Devive:    "Web",
	}
	//InsertElement the element into the tree
	InsertElement(element)

	//element using UK country and web Devive
	element = Node{
		WebReq:    120,
		TimeSpent: 120,
		Country:   "UK",
		Devive:    "Web",
	}
	//InsertElement the element into the tree
	InsertElement(element)

	//element using IN country and web Devive
	element = Node{
		WebReq:    120,
		TimeSpent: 120,
		Country:   "UK",
		Devive:    "Web",
	}
	//InsertElement the element into the tree
	InsertElement(element)

	//element using IN country and web Devive
	element = Node{
		WebReq:    120,
		TimeSpent: 120,
		Country:   "CA",
		Devive:    "Web",
	}
	//InsertElement the element into the tree
	InsertElement(element)

}

func TestQueryElement(t *testing.T) {

	//test case for existing country
	element := QueryElement("UK")

	if element.TimeSpent != 240 {
		t.Error("Test case fail")
	}

	//test case for Not existing country
	element = QueryElement("AUS")

	if element.TimeSpent > 0 {
		t.Error("Test case fail")
	}

}
