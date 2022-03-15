# crud-mongo-go

### This is a training project that allows you to work with golang in conjunction with mongoDB

This project used the mgo library to work with NO-SQL **mongoDB**. You can read [documentation](https://pkg.go.dev/gopkg.in/mgo.v2?utm_source=godoc) to learn how it works.
Also it used package [gorilla/mux](https://github.com/gorilla/mux) for **easier routing**.

Before you run the project, you need to make adjustments.
To get started, change the **connection string and any other configuration parameters** in the file [db.go](./api/db.go):

```
func init() {
  	session, err := mgo.Dial("Your URL")
  	...
  	db = session.DB("Name Your DB")
 }
 
 func collection() *mgo.Collection {
	return db.C("Name Your Collection")
}
```

Next, we can change the address of the server that we are listening on in [main.go](./main.go):

```
func main() {
	...
	http.ListenAndServe("Your address", route)
}
```

Then you can run the project with command:  `go run main.go`
If you have Postman or other alternatives, you can send requests to  `http://YourAddress/api/...` Instead of three dots, you can write:
-  `create`  (to create new data)
-  `get`  (to get all data)
-  `get/SomeID`  (to get data by ID)
-  `delete/SomeID`  (to delete data by ID)

Happy coding!
