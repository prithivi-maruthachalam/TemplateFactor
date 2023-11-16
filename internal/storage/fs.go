package storage

import (
	"log"
	"os"
	"path/filepath"

	"github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/errors"
)

const db_file_name = "templatefactory.db" // Name of the db file
const home_dir_permission = 0777          // Permissions for tf home directory
const db_file_permissions = 0644          // Permissions for the db file

var user_home_dir = GetUserHomeDir()                           // Home directory for the user (os specific)
var TF_HOME = filepath.Join(user_home_dir, ".templatefactory") // Home directory for template factory within the user's home directory
var db_path = filepath.Join(TF_HOME, db_file_name)             // Path to the db file within the template factory home directory

// Create Template Factory home directory if it does not exist
func CreateTemplateFactoryHomeIfNotExists() error {
	return os.MkdirAll(TF_HOME, home_dir_permission)
}

// Gets the home directory for the given operating system
func GetUserHomeDir() string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(errors.HomePathNotFoundError{})
	}

	return homedir
}
