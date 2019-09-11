package app

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetConnectionString(t *testing.T){
	c := Config{
		dbUser: "postgres",
		dbPass: "postgres",
		dbHost: "localhost",
		dbPort: "5432",
		dbName: "app",
	}

	connString := c.getConnectionString()

	expectedString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
			c.dbUser, c.dbPass, c.dbHost, c.dbPort, c.dbName)

	assert.Equal(t,expectedString,connString)
}

func TestGetEnvVar(t *testing.T){
	err := os.Setenv("KEY","value")
	if err != nil{
		t.Fatalf("Could not set env var: %s",err)
	}
	envVar := getEnvVar("KEY")
	assert.Equal(t,"value",envVar)

	envVar = getEnvVar("NOT_EXISTING_KEY")
	assert.Equal(t,"",envVar)

}

func TestNewConfig(t *testing.T) {
	setEnv(t,"DB_USER","postgres")
	setEnv(t,"DB_PASS","postgres")
	setEnv(t,"DB_HOST","localhost")
	setEnv(t,"DB_PORT","5432")
	setEnv(t,"DB_NAME","app")

	c := NewConfig()

	assert.Equal(t,"postgres",c.dbUser)
	assert.Equal(t,"postgres",c.dbPass)
	assert.Equal(t,"localhost",c.dbHost)
	assert.Equal(t,"5432",c.dbPort)
	assert.Equal(t,"app",c.dbName)

}

func setEnv(t *testing.T, key , value string){
	err := os.Setenv(key,value)
	if err != nil{
		t.Fatalf("could not set env var: %s",err)
	}
}