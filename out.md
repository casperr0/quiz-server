

# VerifyPostProvokes


## New Provoke

- [ method and url ]
```
$ POST http://0.0.0.0:8080/v1/provokes
```

- [ example payload ]
```
{
	"correct":true,
	"message":"test provoke message"
}
```

- [ example response ]
```
{
  "_links": {
    "self": {
      "href": "/v1/provokes"
    }
  },
  "status": {
    "status_code": 201,
    "message": "provoke created successfully."
  },
  "data": {
    "ID": 0,
    "Correct": true,
    "Message": "test provoke message"
  },
  "_embedded": null
}
```


## Duplicate Provoke

- [ method and url ]
```
$ POST http://0.0.0.0:8080/v1/provokes
```

- [ example payload ]
```
{
	"correct":true,
	"message":"test provoke message"
}
```

- [ example response ]
```
{
  "_links": {
    "self": {
      "href": "/v1/provokes"
    }
  },
  "status": {
    "status_code": 201,
    "message": "provoke created successfully."
  },
  "data": {
    "ID": 0,
    "Correct": true,
    "Message": "test provoke message"
  },
  "_embedded": null
}
```


# VerifyGetProvokes


## All Provokes

- [ method and url ]
```
$ GET http://0.0.0.0:8080/v1/provokes
```

- [ example response ]
```
{
  "_links": {
    "self": {
      "href": "/v1/provokes"
    }
  },
  "status": {
    "status_code": 200,
    "message": "provokes listed successfully."
  },
  "data": [
    {
      "ID": 1,
      "Correct": true,
      "Message": "啊不就好棒棒?"
    },
    {
      "ID": 2,
      "Correct": true,
      "Message": "這題就當作送分啦!"
    },
    {
      "ID": 3,
      "Correct": true,
      "Message": "你是不是偷看過我們的 database?"
    },
    {
      "ID": 4,
      "Correct": false,
      "Message": "你踩到陷阱，損失了 1 點生命值"
    },
    {
      "ID": 5,
      "Correct": false,
      "Message": "系統管理員使用了影分身，你的命中率降低了"
    },
    {
      "ID": 6,
      "Correct": true,
      "Message": "test provoke message"
    },
    {
      "ID": 7,
      "Correct": true,
      "Message": "test provoke message"
    },
    {
      "ID": 8,
      "Correct": true,
      "Message": "test provoke message"
    },
    {
      "ID": 9,
      "Correct": true,
      "Message": "test provoke message"
    },
    {
      "ID": 10,
      "Correct": true,
      "Message": "test provoke message"
    },
    {
      "ID": 11,
      "Correct": true,
      "Message": "test provoke message"
    }
  ],
  "_embedded": null
}
```


## Query Provokes by Correctness

- [ method and url ]
```
$ GET http://0.0.0.0:8080/v1/provokes?correct=true
```

- [ example response ]
```
{
  "_links": {
    "self": {
      "href": "/v1/provokes"
    }
  },
  "status": {
    "status_code": 200,
    "message": "provokes listed successfully."
  },
  "data": [
    {
      "ID": 1,
      "Correct": true,
      "Message": "啊不就好棒棒?"
    },
    {
      "ID": 2,
      "Correct": true,
      "Message": "這題就當作送分啦!"
    },
    {
      "ID": 3,
      "Correct": true,
      "Message": "你是不是偷看過我們的 database?"
    },
    {
      "ID": 6,
      "Correct": true,
      "Message": "test provoke message"
    },
    {
      "ID": 7,
      "Correct": true,
      "Message": "test provoke message"
    },
    {
      "ID": 8,
      "Correct": true,
      "Message": "test provoke message"
    },
    {
      "ID": 9,
      "Correct": true,
      "Message": "test provoke message"
    },
    {
      "ID": 10,
      "Correct": true,
      "Message": "test provoke message"
    },
    {
      "ID": 11,
      "Correct": true,
      "Message": "test provoke message"
    }
  ],
  "_embedded": null
}
```


# VerifyPostTags


## New Tag

- [ method and url ]
```
$ POST http://0.0.0.0:8080/v1/tags
```

- [ example payload ]
```
{
	"name":"TestTag"
}
```

- [ example response ]
```
{
  "_links": {
    "self": {
      "href": "/v1/tags"
    }
  },
  "status": {
    "status_code": 409,
    "message": "tag TestTag already existed"
  },
  "data": null,
  "_embedded": null
}
```


## Duplicate Tag

- [ method and url ]
```
$ POST http://0.0.0.0:8080/v1/tags
```

- [ example payload ]
```
{
	"name":"TestTag"
}
```

- [ example response ]
```
{
  "_links": {
    "self": {
      "href": "/v1/tags"
    }
  },
  "status": {
    "status_code": 409,
    "message": "tag TestTag already existed"
  },
  "data": null,
  "_embedded": null
}
```


# VerifyGetTags


## All Tags

- [ method and url ]
```
$ GET http://0.0.0.0:8080/v1/tags
```

- [ example response ]
```
{
  "_links": {
    "self": {
      "href": "/v1/tags"
    }
  },
  "status": {
    "status_code": 200,
    "message": "tags listed successfully."
  },
  "data": [
    {
      "id": 1,
      "name": "Network"
    },
    {
      "id": 2,
      "name": "Language"
    },
    {
      "id": 3,
      "name": "Security"
    },
    {
      "id": 4,
      "name": "Hardware"
    },
    {
      "id": 5,
      "name": "Animation"
    },
    {
      "id": 6,
      "name": "Game"
    },
    {
      "id": 7,
      "name": "SysAdmin"
    },
    {
      "id": 8,
      "name": "School"
    },
    {
      "id": 9,
      "name": "CCNS"
    },
    {
      "id": 10,
      "name": "Engineering"
    },
    {
      "id": 11,
      "name": "Math"
    },
    {
      "id": 12,
      "name": "Others"
    },
    {
      "id": 13,
      "name": "TestTag"
    }
  ],
  "_embedded": null
}
```


# VerifyPostPlayers


## New Player

- [ method and url ]
```
$ POST http://0.0.0.0:8080/v1/players
```

- [ example payload ]
```
{
	"name":"testplayer"
}
```

- [ example response ]
```
{
  "_links": {
    "feed": {
      "href": "/v1/players/testplayer/feed"
    },
    "player": {
      "href": "/v1/players/testplayer"
    },
    "self": {
      "href": "/v1/players"
    }
  },
  "status": {
    "status_code": 201,
    "message": "player testplayer created successfully."
  },
  "data": {
    "ID": 0,
    "Name": "testplayer"
  },
  "_embedded": null
}
```


## Duplicate Player

- [ method and url ]
```
$ POST http://0.0.0.0:8080/v1/players
```

- [ example payload ]
```
{
	"name":"testplayer"
}
```

- [ example response ]
```
{
  "_links": {
    "player": {
      "href": "/v1/players/testplayer"
    },
    "self": {
      "href": "/v1/players"
    }
  },
  "status": {
    "status_code": 409,
    "message": "player testplayer already existed"
  },
  "data": {
    "ID": 0,
    "Name": "testplayer"
  },
  "_embedded": null
}
```


# VerifyGetPlayers


## All Players

- [ method and url ]
```
$ GET http://0.0.0.0:8080/v1/players
```

- [ example response ]
```
{
  "_links": null,
  "status": {
    "status_code": 200,
    "message": "players listed successfully."
  },
  "data": [
    {
      "ID": 1,
      "Name": "RainrainWu",
      "Score": 8
    },
    {
      "ID": 2,
      "Name": "!(\u0026GQ(WDUHQ(",
      "Score": 6
    },
    {
      "ID": 3,
      "Name": "  ",
      "Score": 4
    },
    {
      "ID": 4,
      "Name": "__",
      "Score": 3
    }
  ],
  "_embedded": null
}
```


# VerifyGetPlayerFeed


## Finished Player

- [ method and url ]
```
$ GET http://0.0.0.0:8080/v1/players/RainrainWu/feed
```

- [ example response ]
```
{
  "_links": {
    "list": {
      "href": "/players"
    },
    "player": {
      "href": "/players/RainrainWu"
    },
    "self": {
      "href": "/players/RainrainWu/feed"
    }
  },
  "status": {
    "status_code": 200,
    "message": "no quiz left"
  },
  "data": null,
  "_embedded": null
}
```


## Ongoing Player

- [ method and url ]
```
$ GET http://0.0.0.0:8080/v1/players/testplayer/feed
```

- [ example response ]
```
{
  "_links": {
    "list": {
      "href": "/players"
    },
    "player": {
      "href": "/players/testplayer"
    },
    "self": {
      "href": "/players/testplayer/feed"
    }
  },
  "status": {
    "status_code": 200,
    "message": "player testplayer fed successfully."
  },
  "data": {
    "ID": 3,
    "Number": 3,
    "Description": "quiz 3 description.",
    "Score": 1,
    "OptionA": "3.A",
    "OptionB": "3.B",
    "OptionC": "3.C",
    "OptionD": "3.D",
    "Answer": "A"
  },
  "_embedded": [
    {
      "_links": null,
      "status": {
        "status_code": 200,
        "message": "tags listed successfully."
      },
      "data": [
        {
          "id": 3,
          "name": "Security"
        },
        {
          "id": 6,
          "name": "Game"
        },
        {
          "id": 9,
          "name": "CCNS"
        }
      ],
      "_embedded": null
    }
  ]
}
```


# VerifyPostQuizzes


## New Quiz

- [ method and url ]
```
$ POST http://0.0.0.0:8080/v1/quizzes
```

- [ example payload ]
```
{
	"number":999,
	"Description":"test description.",
	"Score":3,
	"OptionA":"test A",
	"OptionB":"test B",
	"OptionC":"test C",
	"OptionD":"test D",
	"Answer":"A"
}
```

- [ example response ]
```
{
  "_links": {
    "quiz": {
      "href": "/v1/quizzes/999"
    },
    "self": {
      "href": "/v1/quizzes"
    },
    "tags": {
      "href": "/v1/quizzes/999/tags"
    }
  },
  "status": {
    "status_code": 201,
    "message": "quiz number 999 created successfully."
  },
  "data": {
    "ID": 0,
    "Number": 999,
    "Description": "test description.",
    "Score": 3,
    "OptionA": "test A",
    "OptionB": "test B",
    "OptionC": "test C",
    "OptionD": "test D",
    "Answer": "A"
  },
  "_embedded": null
}
```


## Duplicate Quiz

- [ method and url ]
```
$ POST http://0.0.0.0:8080/v1/quizzes
```

- [ example payload ]
```
{
	"number":999,
	"Description":"test description.",
	"Score":3,
	"OptionA":"test A",
	"OptionB":"test B",
	"OptionC":"test C",
	"OptionD":"test D",
	"Answer":"A"
}
```

- [ example response ]
```
{
  "_links": {
    "self": {
      "href": "/v1/quizzes"
    }
  },
  "status": {
    "status_code": 409,
    "message": "quiz number 999 already existed"
  },
  "data": {
    "ID": 0,
    "Number": 999,
    "Description": "test description.",
    "Score": 3,
    "OptionA": "test A",
    "OptionB": "test B",
    "OptionC": "test C",
    "OptionD": "test D",
    "Answer": "A"
  },
  "_embedded": null
}
```


# VerifyGetQuizzes


## All Quizzes

- [ method and url ]
```
$ GET http://0.0.0.0:8080/v1/quizzes
```

- [ example response ]
```
{
  "_links": {
    "self": {
      "href": "/v1/quizzes"
    }
  },
  "status": {
    "status_code": 200,
    "message": "quizzes listed successfully."
  },
  "data": [
    {
      "ID": 1,
      "Number": 1,
      "Description": "quiz 1 description.",
      "Score": 1,
      "OptionA": "1.A",
      "OptionB": "1.B",
      "OptionC": "1.C",
      "OptionD": "1.D",
      "Answer": "A"
    },
    {
      "ID": 2,
      "Number": 2,
      "Description": "quiz 2 description.",
      "Score": 1,
      "OptionA": "2.A",
      "OptionB": "2.B",
      "OptionC": "2.C",
      "OptionD": "2.D",
      "Answer": "C"
    },
    {
      "ID": 3,
      "Number": 3,
      "Description": "quiz 3 description.",
      "Score": 1,
      "OptionA": "3.A",
      "OptionB": "3.B",
      "OptionC": "3.C",
      "OptionD": "3.D",
      "Answer": "A"
    },
    {
      "ID": 4,
      "Number": 4,
      "Description": "quiz 4 description.",
      "Score": 2,
      "OptionA": "4.A",
      "OptionB": "4.B",
      "OptionC": "4.C",
      "OptionD": "4.D",
      "Answer": "D"
    },
    {
      "ID": 5,
      "Number": 5,
      "Description": "quiz 5 description.",
      "Score": 3,
      "OptionA": "5.A",
      "OptionB": "5.B",
      "OptionC": "5.C",
      "OptionD": "5.D",
      "Answer": "B"
    },
    {
      "ID": 8,
      "Number": 999,
      "Description": "test description.",
      "Score": 3,
      "OptionA": "test A",
      "OptionB": "test B",
      "OptionC": "test C",
      "OptionD": "test D",
      "Answer": "A"
    }
  ],
  "_embedded": null
}
```


## Query Quizzes by Tag

- [ method and url ]
```
$ GET http://0.0.0.0:8080/v1/quizzes?tag=Security
```

- [ example response ]
```
{
  "_links": {
    "self": {
      "href": "/v1/quizzes"
    }
  },
  "status": {
    "status_code": 200,
    "message": "quizzes listed successfully."
  },
  "data": [
    {
      "ID": 2,
      "Number": 2,
      "Description": "quiz 2 description.",
      "Score": 1,
      "OptionA": "2.A",
      "OptionB": "2.B",
      "OptionC": "2.C",
      "OptionD": "2.D",
      "Answer": "C"
    },
    {
      "ID": 3,
      "Number": 3,
      "Description": "quiz 3 description.",
      "Score": 1,
      "OptionA": "3.A",
      "OptionB": "3.B",
      "OptionC": "3.C",
      "OptionD": "3.D",
      "Answer": "A"
    }
  ],
  "_embedded": null
}
```


# VerifyGetQuiz


## Particular Quiz

- [ method and url ]
```
$ GET http://0.0.0.0:8080/v1/quizzes/999
```

- [ example response ]
```
{
  "_links": {
    "answers": {
      "href": "/v1/answers?quiz=999"
    },
    "list": {
      "href": "/v1/quizzes"
    },
    "self": {
      "href": "/v1/quizzes/999"
    },
    "tags": {
      "href": "/v1/quizzes/999/tags"
    }
  },
  "status": {
    "status_code": 200,
    "message": "quiz number 999 accessed successfully."
  },
  "data": {
    "ID": 8,
    "Number": 999,
    "Description": "test description.",
    "Score": 3,
    "OptionA": "test A",
    "OptionB": "test B",
    "OptionC": "test C",
    "OptionD": "test D",
    "Answer": "A"
  },
  "_embedded": [
    {
      "_links": null,
      "status": {
        "status_code": 200,
        "message": "tags listed successfully."
      },
      "data": null,
      "_embedded": null
    }
  ]
}
```


# VerifyPostQuizTags


## New Quiz Tag

- [ method and url ]
```
$ POST http://0.0.0.0:8080/v1/quizzes/999/tags
```

- [ example payload ]
```
{
	"name":"Engineering"
}
```

- [ example response ]
```
{
  "_links": null,
  "status": {
    "status_code": 201,
    "message": "quiz number 999 registered with tag Engineering successfully."
  },
  "data": {
    "id": 0,
    "name": "Engineering"
  },
  "_embedded": null
}
```


## Duplicate Quiz Tag

- [ method and url ]
```
$ POST http://0.0.0.0:8080/v1/quizzes/999/tags
```

- [ example payload ]
```
{
	"name":"Engineering"
}
```

- [ example response ]
```
{
  "_links": null,
  "status": {
    "status_code": 201,
    "message": "quiz number 999 registered with tag Engineering successfully."
  },
  "data": {
    "id": 0,
    "name": "Engineering"
  },
  "_embedded": null
}
```


# VerifyGetQuizTags


## All Quiz Tags

- [ method and url ]
```
$ GET http://0.0.0.0:8080/v1/quizzes/999/tags
```

- [ example response ]
```
{
  "_links": {
    "answers": {
      "href": "/v1/answers?quiz=999"
    },
    "quiz": {
      "href": "/v1/quizzes/999"
    },
    "self": {
      "href": "/v1/quizzes/999/tags"
    }
  },
  "status": {
    "status_code": 200,
    "message": "tags of quiz number 999 listed successfully."
  },
  "data": [
    {
      "id": 10,
      "name": "Engineering"
    }
  ],
  "_embedded": null
}
```


# VerifyPostAnswers


## New Answer

- [ method and url ]
```
$ POST http://0.0.0.0:8080/v1/answers
```

- [ example payload ]
```
{
	"player_name":"testplayer",
	"quiz_number":999,
	"correct":true
}
```

- [ example response ]
```
{
  "_links": null,
  "status": {
    "status_code": 201,
    "message": "answer created successfully."
  },
  "data": {
    "player_name": "testplayer",
    "quiz_number": 999,
    "correct": true
  },
  "_embedded": null
}
```


## Duplicate Answer

- [ method and url ]
```
$ POST http://0.0.0.0:8080/v1/answers
```

- [ example payload ]
```
{
	"player_name":"testplayer",
	"quiz_number":999,
	"correct":true
}
```
$ POST http://0.0.0.0:8080/v1/answers

- [ example response ]
```
{
  "_links": null,
  "status": {
    "status_code": 409,
    "message": "answer from player testplayer to quiz number 999 already existed"
  },
  "data": null,
  "_embedded": null
}
```


# VerifyGetAnswers


## All Answers

- [ method and url ]
```
$ GET http://0.0.0.0:8080/v1/answers
```

- [ example response ]
```
{
  "_links": {
    "self": {
      "href": "/v1/answers"
    }
  },
  "status": {
    "status_code": 200,
    "message": "answers listed successfully."
  },
  "data": [
    {
      "ID": 1,
      "PlayerID": 1,
      "QuizID": 1,
      "Correct": true
    },
    {
      "ID": 2,
      "PlayerID": 1,
      "QuizID": 2,
      "Correct": true
    },
    {
      "ID": 3,
      "PlayerID": 1,
      "QuizID": 3,
      "Correct": true
    },
    {
      "ID": 4,
      "PlayerID": 1,
      "QuizID": 4,
      "Correct": true
    },
    {
      "ID": 5,
      "PlayerID": 1,
      "QuizID": 5,
      "Correct": true
    },
    {
      "ID": 6,
      "PlayerID": 2,
      "QuizID": 1,
      "Correct": true
    },
    {
      "ID": 7,
      "PlayerID": 2,
      "QuizID": 4,
      "Correct": true
    },
    {
      "ID": 8,
      "PlayerID": 2,
      "QuizID": 5,
      "Correct": true
    },
    {
      "ID": 9,
      "PlayerID": 2,
      "QuizID": 2,
      "Correct": false
    },
    {
      "ID": 10,
      "PlayerID": 2,
      "QuizID": 3,
      "Correct": false
    },
    {
      "ID": 11,
      "PlayerID": 3,
      "QuizID": 1,
      "Correct": true
    },
    {
      "ID": 12,
      "PlayerID": 3,
      "QuizID": 5,
      "Correct": true
    },
    {
      "ID": 13,
      "PlayerID": 3,
      "QuizID": 2,
      "Correct": false
    },
    {
      "ID": 14,
      "PlayerID": 4,
      "QuizID": 4,
      "Correct": true
    },
    {
      "ID": 15,
      "PlayerID": 4,
      "QuizID": 1,
      "Correct": true
    },
    {
      "ID": 16,
      "PlayerID": 4,
      "QuizID": 5,
      "Correct": false
    },
    {
      "ID": 17,
      "PlayerID": 4,
      "QuizID": 2,
      "Correct": false
    },
    {
      "ID": 20,
      "PlayerID": 7,
      "QuizID": 8,
      "Correct": true
    }
  ],
  "_embedded": null
}
```


## Query Answers by Player

- [ method and url ]
```
$ GET http://0.0.0.0:8080/v1/answers?player=testplayer
```

- [ example response ]
```
{
  "_links": {
    "player": {
      "href": "/v1/players/testplayer"
    },
    "self": {
      "href": "/v1/answers"
    }
  },
  "status": {
    "status_code": 200,
    "message": "answers listed successfully."
  },
  "data": [
    {
      "ID": 20,
      "PlayerID": 0,
      "QuizID": 8,
      "Correct": true
    }
  ],
  "_embedded": null
}
```


## Query Answers by Quiz

- [ method and url ]
```
$ GET http://0.0.0.0:8080/v1/answers?quiz=999
```

- [ example response ]
```
{
  "_links": {
    "quiz": {
      "href": "/v1/quizzes/999"
    },
    "self": {
      "href": "/v1/answers"
    }
  },
  "status": {
    "status_code": 200,
    "message": "answers listed successfully."
  },
  "data": [
    {
      "ID": 20,
      "PlayerID": 7,
      "QuizID": 0,
      "Correct": true
    }
  ],
  "_embedded": null
}
```


# VerifyDeletePlayer


## Existed Quiz

- [ method and url ]
```
$ DELETE http://0.0.0.0:8080/v1/quizzes/999
```

- [ example response ]
```
{
  "_links": {
    "list": {
      "href": "/v1/quizzes"
    }
  },
  "status": {
    "status_code": 200,
    "message": "quiz number 999 deleted successfully."
  },
  "data": {
    "ID": 8,
    "Number": 999,
    "Description": "test description.",
    "Score": 3,
    "OptionA": "test A",
    "OptionB": "test B",
    "OptionC": "test C",
    "OptionD": "test D",
    "Answer": "A"
  },
  "_embedded": null
}
```


## Non-existed Quiz

- [ method and url ]
```
$ DELETE http://0.0.0.0:8080/v1/quizzes/999
```

- [ example response ]
```
{
  "_links": {
    "list": {
      "href": "/v1/quizzes"
    }
  },
  "status": {
    "status_code": 400,
    "message": "quiz number 999 not found"
  },
  "data": null,
  "_embedded": null
}
```


# VerifyDeletePlayers


## Existed Player

- [ method and url ]
```
$ DELETE http://0.0.0.0:8080/v1/players/testplayer
```

- [ example response ]
```
{
  "_links": {
    "list": {
      "href": "/players"
    }
  },
  "status": {
    "status_code": 200,
    "message": "player testplayer deleted successfully."
  },
  "data": {
    "ID": 7,
    "Name": "testplayer"
  },
  "_embedded": null
}
```


## Non-existed Player

- [ method and url ]
```
$ DELETE http://0.0.0.0:8080/v1/players/testplayer
```

- [ example response ]
```
{
  "_links": {
    "list": {
      "href": "/players"
    }
  },
  "status": {
    "status_code": 400,
    "message": "player testplayer not found"
  },
  "data": null,
  "_embedded": null
}
```
