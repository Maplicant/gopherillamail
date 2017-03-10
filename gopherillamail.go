package gopherillamail

import (
    "fmt"
    "net/http"
    "time"
)

type Mail struct {
    guid string
    subject string
    sender string
    time time.Time
    read bool
    excerpt string
    body string
}

func main() {

}
