package current_next

import (
	"github.com/jubnzv/go-taskwarrior"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// Errors
var NoTaskFoundError = errors.New("A task with that id could not be found")
var NoHomeDirError = errors.New("No home directory found for this user")
var NoDataFileFoundError = errors.New("No data file found")
var ReadingDataFileError = errors.New("Error reading the data file") 

 func GetTaskWithId(t *taskwarrior.TaskWarrior, id int32) (*taskwarrior.Task, error) {
	 for _, task := range t.Tasks {
		 if task.Id == id {
			 return &task, nil
		 }
	 }
	 return nil, NoTaskFoundError
 }

func ReadCurrentData() (string, error) {
	homedir, err := os.UserHomeDir()
        if err != nil {
		return "", NoHomeDirError
        }

        data_path := filepath.Join(homedir, ".current-next", "current.data")
        if _, err := os.Stat(data_path); os.IsNotExist(err) {
		return "", NoDataFileFoundError 
        }

        data, err := os.ReadFile(data_path)
        if err != nil {
		return "", ReadingDataFileError
        }

        temp := strings.Replace(string(data), "\n", "", -1)
	return temp, nil
}


