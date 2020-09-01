

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
    "_id": 0,
    "correct": true,
    "message": "test provoke message"
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
    "_id": 0,
    "correct": true,
    "message": "test provoke message"
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
      "_id": 1,
      "correct": true,
      "message": "啊不就好棒棒?"
    },
    {
      "_id": 2,
      "correct": true,
      "message": "這題就當作送分啦!"
    },
    {
      "_id": 3,
      "correct": true,
      "message": "你是不是偷看過我們的 database?"
    },
    {
      "_id": 4,
      "correct": false,
      "message": "你踩到陷阱，損失了 1 點生命值"
    },
    {
      "_id": 5,
      "correct": false,
      "message": "系統管理員使用了影分身，你的命中率降低了"
    },
    {
      "_id": 6,
      "correct": true,
      "message": "test provoke message"
    },
    {
      "_id": 7,
      "correct": true,
      "message": "test provoke message"
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
      "_id": 1,
      "correct": true,
      "message": "啊不就好棒棒?"
    },
    {
      "_id": 2,
      "correct": true,
      "message": "這題就當作送分啦!"
    },
    {
      "_id": 3,
      "correct": true,
      "message": "你是不是偷看過我們的 database?"
    },
    {
      "_id": 6,
      "correct": true,
      "message": "test provoke message"
    },
    {
      "_id": 7,
      "correct": true,
      "message": "test provoke message"
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
    "status_code": 201,
    "message": "tag TestTag created successfully."
  },
  "data": {
    "_id": 0,
    "name": "TestTag"
  },
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
      "_id": 1,
      "name": "Network"
    },
    {
      "_id": 2,
      "name": "Language"
    },
    {
      "_id": 3,
      "name": "Security"
    },
    {
      "_id": 4,
      "name": "Hardware"
    },
    {
      "_id": 5,
      "name": "Animation"
    },
    {
      "_id": 6,
      "name": "Game"
    },
    {
      "_id": 7,
      "name": "SysAdmin"
    },
    {
      "_id": 8,
      "name": "School"
    },
    {
      "_id": 9,
      "name": "CCNS"
    },
    {
      "_id": 10,
      "name": "Engineering"
    },
    {
      "_id": 11,
      "name": "Math"
    },
    {
      "_id": 12,
      "name": "Others"
    },
    {
      "_id": 13,
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
    "_id": 0,
    "name": "testplayer"
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
    "_id": 0,
    "name": "testplayer"
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
      "_id": 1,
      "name": "RainrainWu",
      "score": 8
    },
    {
      "_id": 2,
      "name": "!(\u0026GQ(WDUHQ(",
      "score": 6
    },
    {
      "_id": 3,
      "name": "  ",
      "score": 4
    },
    {
      "_id": 4,
      "name": "__",
      "score": 3
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
    "message": "player fed successfully."
  },
  "data": {
    "_id": 3,
    "number": 3,
    "description": "quiz 3 description.",
    "score": 1,
    "options": [
      "3.A",
      "3.B",
      "3.C",
      "3.D",
      "Engineering",
      "Security"
    ],
    "answer": "A",
    "tags": []
  },
  "_embedded": null
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
	"description":"test description.",
	"score":3,
	"option_a":"test A",
	"option_b":"test B",
	"option_c":"test C",
	"option_d":"test D",
	"answer":"A"
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
    "_id": 0,
    "number": 999,
    "description": "test description.",
    "score": 3,
    "option_a": "test A",
    "option_b": "test B",
    "option_c": "test C",
    "option_d": "test D",
    "answer": "A"
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
	"description":"test description.",
	"score":3,
	"option_a":"test A",
	"option_b":"test B",
	"option_c":"test C",
	"option_d":"test D",
	"answer":"A"
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
    "_id": 0,
    "number": 999,
    "description": "test description.",
    "score": 3,
    "option_a": "test A",
    "option_b": "test B",
    "option_c": "test C",
    "option_d": "test D",
    "answer": "A"
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
      "_id": 1,
      "number": 1,
      "description": "quiz 1 description.",
      "score": 1,
      "option_a": "1.A",
      "option_b": "1.B",
      "option_c": "1.C",
      "option_d": "1.D",
      "answer": "A"
    },
    {
      "_id": 2,
      "number": 2,
      "description": "quiz 2 description.",
      "score": 1,
      "option_a": "2.A",
      "option_b": "2.B",
      "option_c": "2.C",
      "option_d": "2.D",
      "answer": "C"
    },
    {
      "_id": 3,
      "number": 3,
      "description": "quiz 3 description.",
      "score": 1,
      "option_a": "3.A",
      "option_b": "3.B",
      "option_c": "3.C",
      "option_d": "3.D",
      "answer": "A"
    },
    {
      "_id": 4,
      "number": 4,
      "description": "quiz 4 description.",
      "score": 2,
      "option_a": "4.A",
      "option_b": "4.B",
      "option_c": "4.C",
      "option_d": "4.D",
      "answer": "D"
    },
    {
      "_id": 5,
      "number": 5,
      "description": "quiz 5 description.",
      "score": 3,
      "option_a": "5.A",
      "option_b": "5.B",
      "option_c": "5.C",
      "option_d": "5.D",
      "answer": "B"
    },
    {
      "_id": 6,
      "number": 999,
      "description": "test description.",
      "score": 3,
      "option_a": "test A",
      "option_b": "test B",
      "option_c": "test C",
      "option_d": "test D",
      "answer": "A"
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
      "_id": 2,
      "number": 2,
      "description": "quiz 2 description.",
      "score": 1,
      "option_a": "2.A",
      "option_b": "2.B",
      "option_c": "2.C",
      "option_d": "2.D",
      "answer": "C"
    },
    {
      "_id": 3,
      "number": 3,
      "description": "quiz 3 description.",
      "score": 1,
      "option_a": "3.A",
      "option_b": "3.B",
      "option_c": "3.C",
      "option_d": "3.D",
      "answer": "A"
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
    "message": "quiz accessed successfully."
  },
  "data": {
    "_id": 6,
    "number": 999,
    "description": "test description.",
    "score": 3,
    "options": [
      "test A",
      "test B",
      "test C",
      "test D"
    ],
    "answer": "A",
    "tags": []
  },
  "_embedded": null
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
    "_id": 0,
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
    "_id": 0,
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
      "_id": 10,
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
      "_id": 1,
      "player_id": 1,
      "quiz_id": 1,
      "correct": true
    },
    {
      "_id": 2,
      "player_id": 1,
      "quiz_id": 2,
      "correct": true
    },
    {
      "_id": 3,
      "player_id": 1,
      "quiz_id": 3,
      "correct": true
    },
    {
      "_id": 4,
      "player_id": 1,
      "quiz_id": 4,
      "correct": true
    },
    {
      "_id": 5,
      "player_id": 1,
      "quiz_id": 5,
      "correct": true
    },
    {
      "_id": 6,
      "player_id": 2,
      "quiz_id": 1,
      "correct": true
    },
    {
      "_id": 7,
      "player_id": 2,
      "quiz_id": 4,
      "correct": true
    },
    {
      "_id": 8,
      "player_id": 2,
      "quiz_id": 5,
      "correct": true
    },
    {
      "_id": 9,
      "player_id": 2,
      "quiz_id": 2,
      "correct": false
    },
    {
      "_id": 10,
      "player_id": 2,
      "quiz_id": 3,
      "correct": false
    },
    {
      "_id": 11,
      "player_id": 3,
      "quiz_id": 1,
      "correct": true
    },
    {
      "_id": 12,
      "player_id": 3,
      "quiz_id": 5,
      "correct": true
    },
    {
      "_id": 13,
      "player_id": 3,
      "quiz_id": 2,
      "correct": false
    },
    {
      "_id": 14,
      "player_id": 4,
      "quiz_id": 4,
      "correct": true
    },
    {
      "_id": 15,
      "player_id": 4,
      "quiz_id": 1,
      "correct": true
    },
    {
      "_id": 16,
      "player_id": 4,
      "quiz_id": 5,
      "correct": false
    },
    {
      "_id": 17,
      "player_id": 4,
      "quiz_id": 2,
      "correct": false
    },
    {
      "_id": 18,
      "player_id": 5,
      "quiz_id": 6,
      "correct": true
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
      "_id": 18,
      "player_id": 0,
      "quiz_id": 6,
      "correct": true
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
      "_id": 18,
      "player_id": 5,
      "quiz_id": 0,
      "correct": true
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
    "_id": 6,
    "number": 999,
    "description": "test description.",
    "score": 3,
    "options": [
      "test A",
      "test B",
      "test C",
      "test D",
      "Engineering"
    ],
    "answer": "A",
    "tags": []
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
    "_id": 5,
    "name": "testplayer"
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
