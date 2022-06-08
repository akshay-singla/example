package tree

func InsertElement(element Node) {

	count := 0
	for _, node := range parentNode.Children {

		//Loop for the Parent nodes
		if node.Country == element.Country {
			//Loop for the child node
			for _, child := range node.Children {
				if child.Devive == element.Devive {

					oldWebReq := child.WebReq
					oldTimepent := child.TimeSpent
					child.TimeSpent = (oldTimepent + element.TimeSpent) / 2
					child.WebReq = (oldWebReq + element.WebReq) / 2
					child.Devive = element.Devive

					//update parent node

					node.TimeSpent += child.TimeSpent - oldTimepent
					node.WebReq += child.WebReq - oldWebReq

					parentNode.TimeSpent += child.TimeSpent - oldTimepent
					parentNode.WebReq += child.WebReq - oldWebReq

					//Increment the count
					count++
				}
			}

			//check the count value
			if count == 0 {
				newChild := Node{
					TimeSpent: element.TimeSpent,
					WebReq:    element.WebReq,
					Devive:    element.Devive,
				}
				node.Children = append(node.Children, &newChild)
				node.TimeSpent += element.TimeSpent
				node.WebReq += element.WebReq

				parentNode.TimeSpent += element.TimeSpent
				parentNode.WebReq += element.WebReq
				//decrement the count value
				count--
			}

		}
	}

	//check the count value
	if count == 0 {
		newChild := Node{
			TimeSpent: element.TimeSpent,
			WebReq:    element.WebReq,
			Country:   element.Country,
			Children: []*Node{
				{
					TimeSpent: element.TimeSpent,
					WebReq:    element.WebReq,
					Devive:    element.Devive,
				}},
		}
		parentNode.Children = append(parentNode.Children, &newChild)
		parentNode.TimeSpent += element.TimeSpent
		parentNode.WebReq += element.WebReq

		//decrement the count value
		count--
	}

}

func QueryElement(country string) Node {

	//fetcht the data from tree structure
	for _, node := range parentNode.Children {

		//Loop for the Parent nodes
		if node.Country == country {
			return *node
		}
	}

	return Node{}
}
