package config

import "os"

// DBURL URL of the database.
var DBURL string = os.Getenv("DB_URL")
