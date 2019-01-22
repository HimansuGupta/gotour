package dbmock

import (
	"fmt"

	"github.com/thylong/mongo_mock_go_example/model"
)

func main() {
	session := model.NewMockSession()
	documents, _ := session.DB("employeedb").C("employeedetails").GetMyDocuments()

	fmt.Println(documents)
}
