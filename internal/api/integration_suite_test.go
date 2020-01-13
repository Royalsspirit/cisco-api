package api

import (
	"database/sql"
	"os"

	"github.com/Royalsspirit/cisco-api/internal/logger"
	"github.com/Royalsspirit/cisco-api/internal/validator"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

// APISuite represents the environment used for the tests
type APISuite struct {
	suite.Suite
	Environment string
	Port        string
	Logger      *log.Entry
	DB          *sql.DB
	Validator   *validator.Val
}

// SetupSuite setup the tests suite
func (p *APISuite) SetupSuite() {

	db, err := sql.Open("sqlite3", "../../"+os.Getenv("DB"))
	if err != nil {
		log.Fatal(err)
	}

	// Init logger
	logger := logger.NewLogger(&logger.Config{
		Name:        "cisco-test-api",
		Environment: "develop",
	})

	// Init validator
	v := validator.NewValidator()

	p.DB = db
	p.Logger = logger
	p.Validator = v
}
