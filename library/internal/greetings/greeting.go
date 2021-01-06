package greetings

import (
  "fmt"
  "errors"
  "math/rand"
  "time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // no names
    if name == ""{
      return "", errors.New("empty name")
    }
    // Return a greeting that embeds the name in a message.
    // message := fmt.Sprintf("Hi, %v. Welcome.", name)
    message := fmt.Sprintf(randomFormat(), name)
    return message,nil
}

func Hellos(names []string) (map[string]string, error) {
  // 用map存储多个人名信息
  messages := make(map[string]string)
  // 循环来接收人名,对每个人名调用Hello
  for _, name := range names{
    message, err := Hello(name)
    if err != nil{
      return nil, err
    }
    messages[name] = message
  }
  return messages, nil
}


func init(){
  rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
  formats := []string{
    "Hi, %v. Welcome!",
    "Great to see you, %v!",
    "Hail, %v! Well met!",
  }
  return formats[rand.Intn(len(formats))]
}

