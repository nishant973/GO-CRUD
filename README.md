# GO-CRUD
Demo Project for webservice implementing for a CRUD operation in GO with  handler , routing and establishing and listening web server.


Run :
go build .
./webservice.exe ----> server started on port 3000

Get Users : localhost:3000/users ---- HTTP VERB  ---- GET


Add Users : localhost:3000/users ---- HTTP VERB  ---- POST ---- Pass Body ----- {"FirstName" : "Nishant" , "LastName" : "J"}


Get UserByID : localhost:3000/users/{ID} ---- HTTP VERB  ---- GET 


Update User : localhost:3000/users/{ID} ----- HTTP VERB  ---- PUT ------ Pass Body ----- {"ID" : 1 , "FirstName" : "Nishi" , "LastName" : "J"}


Delete USer : localhost:3000/users/{ID} ----- HTTP VERB  ---- DELETE

Model --------- user.go ------- Contains the logic for CRUD operation and object User


Controllers --- user.go ------- Contains the logic of mapping to call the appropriate handler method , parsing the requets from json to User in user.go in Model package , sending JSON response



Controller  --- front.go -----  Contains the logic for Routing and cerating the json encoder which is used in user.go present in controllers package.



main.go  ------ listen the request .
