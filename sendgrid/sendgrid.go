package sendgrid

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Config struct {
	SENDGRID_API_KEY string
}

type EMAIL struct {
	FromName string
	From     string
	ToName   string
	To       string
	Subject  string
	Body     string
}

func SendEmail(e EMAIL) {

	//Decode config JSON
	file, _ := os.Open(getConfigFile())
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Config{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", config)
	os.Exit(0)

	from := mail.NewEmail(e.FromName, e.From)
	subject := e.Subject
	to := mail.NewEmail(e.ToName, e.To)
	plainTextContent := e.Body
	htmlContent := ""
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(config.SENDGRID_API_KEY)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

//Getting Environment config file
func getConfigFile() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "development"
	}
	filename := []string{"../config/", "config.", env, ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	return filePath
}
