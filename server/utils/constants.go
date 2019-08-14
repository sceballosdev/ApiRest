package utils

const (
	ApiPort = ":4000"
	EndpointNewsfeed = "/newsfeed"
	DriverName = "postgres"
	UrlConnection = "postgresql://%s@localhost:26257/%s?sslmode=disable"
	ErrorMessageConnection = "Error connecting to the database:"
	ErrorTableCreation = "Error creating the table:"
	UserDb = "steven"
	DatabaseName = "apirest_go"
	NewsfeedGetAllQuery = "SELECT id, title, post FROM items ORDER BY id DESC"
)
