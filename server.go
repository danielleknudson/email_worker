package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/danielleknudson/email_worker/send"

	"github.com/gorilla/Schema"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexController)
	r.HandleFunc("/email", EmailController)

	http.Handle("/", r)

	fmt.Println("Starting Go server on: 127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func IndexController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		`<!DOCTYPE html>
      <head>
        <title>Test Go Server</title>
        <style>
          body {
            font-size: 15px;
          }
          input, textarea {
            width: 300px;
            display: block;
            border-radius: 3px;
            border: 1px solid #ccc;
            margin-bottom: 20px;
            height: 35px;
          }
          textarea {
            resize: none;
            height: 60px;
          }
          input[type="submit"] {
            background: #1D4C75;
            color: #FFFFFF;
          }
        </style>
      </head>
      <body>
        <h1>Test Go Server</h1>
        <form method="POST" action="/email">
          <input type="text" id="recipient" placeholder="Recipient email address" />
          <input type="text" id="sender" placeholder="Sender email address" />
          <input type="text" id="subject" placeholder="Subject line" />
          <textarea id="body" placeholder="Email body"></textarea>
          <input type="submit" value="Send Email" />
        </form>
        <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.12.0/jquery.min.js"></script>
        <script>
          $(document).ready(function(){
            $('form').on('submit', function(e) {
              e.preventDefault();

              var $recipient = $('#recipient');
              var $sender = $('#sender');
              var $subject = $('#subject');
              var $body = $('#body');

              $.ajax({
                url: '/email',
                type: 'POST',
                dataType: 'json',
                data: {
                  recipient: $recipient.val(),
                  sender: $sender.val(),
                  subject: $subject.val(),
                  body: $body.val()
                },
                success: function(res) {
                  $recipient.val('');
                  $sender.val('');
                  $subject.val('');
                  $body.val('');
                  console.log(res);
                },
                error: function(res) {
                  console.log(res);
                }
              });
            });
          });
        </script>
      </body>
    </html>`)
}

func EmailController(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		panic(err)
	}

	email := new(Email)

	decoder := schema.NewDecoder()

	err = decoder.Decode(email, r.PostForm)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", email)
	js, err := json.Marshal(email)
	if err != nil {
		panic(err)
	}

	SendEmail(email)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(js))
}
