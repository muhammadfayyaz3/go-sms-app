package twilio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

type Config struct {
	ACCOUNT_SID string
	AUTH_TOKEN  string
	API_URL     string
	FROM_NUMBER string
}

type SMS struct {
	From string
	To   string
	Body string
}

func SendSMS(s SMS) string {

	//Decode config JSON
	file, _ := os.Open(getConfigFile())
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Config{}
	err := decoder.Decode(&config)
	if err != nil {
		result := fmt.Sprintf("%v", err)
		return result
	}

	//Setting SMS contents
	v := url.Values{}
	v.Set("To", s.To)
	v.Set("From", config.FROM_NUMBER)
	v.Set("Body", s.Body)
	request_values := *strings.NewReader(v.Encode())

	// Create client
	client := &http.Client{}

	request, _ := http.NewRequest("POST", config.API_URL, &request_values)
	request.SetBasicAuth(config.ACCOUNT_SID, config.AUTH_TOKEN)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	var result = "Something went wrong, unable to send SMS"
	response, _ := client.Do(request)
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		var data map[string]interface{}
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		err := json.Unmarshal(bodyBytes, &data)
		if err == nil {
			result := fmt.Sprintf("%v", data["sid"])
			return "SMS Sent: " + result
		}
		return result
	} else {
		return "Error: " + response.Status
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
