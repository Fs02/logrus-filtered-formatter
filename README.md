# Logrus Filtered Formatter
Use this formatter to log json with filtered fields

## Usage
```
package main

import (
	"github.com/Fs02/logrus-filtered-formatter"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	fields := []string{
		"password",
		"email",
	}
	log.Formatter = filtered.New(fields, &logrus.JSONFormatter{})

	log.WithFields(logrus.Fields{
		"password": "asdfasdf",
		"email":    "johndoe@gmail.com",
		"name":     "john doe",
	}).Info("new user created")
}
```
Output:
```
{"email":[FILTERED],"level":"info","msg":"new user created","name":"john doe","password":[FILTERED],"time":"2017-10-05T16:05:29+07:00"}
```
