

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
	"name":"discord-testplayer",
	"nickname":"testplayer",
	"platform":"discord"
}
```

- [ example response ]
```
{
  "_links": {
    "feed": {
      "href": "/v1/players/discord-testplayer/feed"
    },
    "player": {
      "href": "/v1/players/discord-testplayer"
    },
    "self": {
      "href": "/v1/players"
    }
  },
  "status": {
    "status_code": 201,
    "message": "player discord-testplayer created successfully."
  },
  "data": {
    "_id": 0,
    "name": "discord-testplayer",
    "nickname": "testplayer",
    "platform": "discord"
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
	"name":"discord-testplayer",
	"nickname":"testplayer",
	"platform":"discord"
}
```

- [ example response ]
```
{
  "_links": {
    "player": {
      "href": "/v1/players/discord-testplayer"
    },
    "self": {
      "href": "/v1/players"
    }
  },
  "status": {
    "status_code": 409,
    "message": "player discord-testplayer already existed"
  },
  "data": {
    "_id": 0,
    "name": "discord-testplayer",
    "nickname": "testplayer",
    "platform": "discord"
  },
  "_embedded": null
}
```


# VerifyGetPlayer


## Particular Player

- [ method and url ]
```
$ GET http://0.0.0.0:8080/v1/players/discord-testplayer
```

- [ example response ]
```
{
  "_links": {
    "answers": {
      "href": "/v1/answers?player=discord-testplayer"
    },
    "list": {
      "href": "/v1/players"
    },
    "self": {
      "href": "/v1/players/discord-testplayer"
    }
  },
  "status": {
    "status_code": 200,
    "message": "player accessed successfully."
  },
  "data": {
    "_id": 5,
    "name": "discord-testplayer",
    "nickname": "testplayer",
    "platform": "discord",
    "rank": 5,
    "score": 0,
    "last": 5
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
      "name": "telegram-RainrainWu",
      "nickname": "RainrainWu",
      "platform": "Telegram"
    },
    {
      "_id": 2,
      "name": "line-!(\u0026GQ(WDUHQ(",
      "nickname": "!(\u0026GQ(WDUHQ(",
      "platform": "line"
    },
    {
      "_id": 3,
      "name": "discord-  ",
      "nickname": "  ",
      "platform": ""
    },
    {
      "_id": 4,
      "name": "facebook-__",
      "nickname": "__",
      "platform": "facebook"
    },
    {
      "_id": 5,
      "name": "discord-testplayer",
      "nickname": "testplayer",
      "platform": "discord"
    }
  ],
  "_embedded": null
}
```


# VerifyGetPlayerFeed


## Finished Player

- [ method and url ]
```
$ GET http://0.0.0.0:8080/v1/players/telegram-RainrainWu/feed
```

- [ example response ]
```
{
  "_links": {
    "list": {
      "href": "/v1/players"
    },
    "player": {
      "href": "/v1/players/telegram-RainrainWu"
    },
    "rand": {
      "href": "/v1/players/telegram-RainrainWu/rand"
    },
    "self": {
      "href": "/v1/players/telegram-RainrainWu/feed"
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
$ GET http://0.0.0.0:8080/v1/players/discord-testplayer/feed
```

- [ example response ]
```
{
  "_links": {
    "list": {
      "href": "/v1/players"
    },
    "player": {
      "href": "/v1/players/discord-testplayer"
    },
    "rand": {
      "href": "/v1/players/discord-testplayer/rand"
    },
    "self": {
      "href": "/v1/players/discord-testplayer/feed"
    }
  },
  "status": {
    "status_code": 200,
    "message": "player fed successfully."
  },
  "data": {
    "_id": 1,
    "number": 1,
    "author": "",
    "description": "quiz 1 description.",
    "hint": "quiz 1 hint",
    "score": 1,
    "options": [
      "1.A",
      "1.B",
      "1.C",
      "1.D"
    ],
    "answer": "A",
    "tags": [
      "Network",
      "Language"
    ]
  },
  "_embedded": null
}
```


# VerifyGetPlayerRand


## Any Player

- [ method and url ]
```
$ GET http://0.0.0.0:8080/v1/players/telegram-RainrainWu/rand
```

- [ example response ]
```
{
  "_links": {
    "feed": {
      "href": "/v1/players/telegram-RainrainWu/feed"
    },
    "list": {
      "href": "/v1/quizzes"
    },
    "player": {
      "href": "/v1/players/telegram-RainrainWu"
    },
    "self": {
      "href": "/v1/players/telegram-RainrainWu/rand"
    }
  },
  "status": {
    "status_code": 200,
    "message": "quiz accessed successfully."
  },
  "data": {
    "_id": 1,
    "number": 1,
    "author": "",
    "description": "quiz 1 description.",
    "hint": "quiz 1 hint",
    "score": 1,
    "options": [
      "1.A",
      "1.B",
      "1.C",
      "1.D"
    ],
    "answer": "A",
    "tags": [
      "Network",
      "Language"
    ]
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
	"author":"test author"
	"description":"test description.",
	"hint":"test hint",
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
    "status_code": 400,
    "message": "invalid character '\"' after object key:value pair"
  },
  "data": null,
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
	"author":"test author"
	"description":"test description.",
	"hint":"test hint",
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
    "status_code": 400,
    "message": "invalid character '\"' after object key:value pair"
  },
  "data": null,
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
      "author": "quiz 1 author",
      "description": "quiz 1 description.",
      "hint": "quiz 1 hint",
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
      "author": "quiz 2 author",
      "description": "quiz 2 description.",
      "hint": "quiz 2 hint",
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
      "author": "quiz 3 author",
      "description": "quiz 3 description.",
      "hint": "quiz 3 hint",
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
      "author": "quiz 4 author",
      "description": "quiz 4 description.",
      "hint": "quiz 4 hint",
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
      "author": "quiz 5 author",
      "description": "quiz 5 description.",
      "hint": "quiz 5 hint",
      "score": 3,
      "option_a": "5.A",
      "option_b": "5.B",
      "option_c": "5.C",
      "option_d": "5.D",
      "answer": "B"
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
      "author": "quiz 2 author",
      "description": "quiz 2 description.",
      "hint": "quiz 2 hint",
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
      "author": "quiz 3 author",
      "description": "quiz 3 description.",
      "hint": "quiz 3 hint",
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
  "_links": null,
  "status": {
    "status_code": 400,
    "message": "quiz number 999 not found"
  },
  "data": null,
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
  "_links": {},
  "status": {
    "status_code": 500,
    "message": "quiz number 999 not found"
  },
  "data": null,
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
  "_links": {},
  "status": {
    "status_code": 500,
    "message": "quiz number 999 not found"
  },
  "data": null,
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
  "data": null,
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
	"player_name":"discord-testplayer",
	"quiz_number":999,
	"correct":true
}
```

- [ example response ]
```
{
  "_links": null,
  "status": {
    "status_code": 500,
    "message": "quiz number 999 not found"
  },
  "data": null,
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
	"player_name":"discord-testplayer",
	"quiz_number":999,
	"correct":true
}
```

- [ example response ]
```
{
  "_links": null,
  "status": {
    "status_code": 500,
    "message": "quiz number 999 not found"
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
    }
  ],
  "_embedded": null
}
```


## Query Answers by Player

- [ method and url ]
```
$ GET http://0.0.0.0:8080/v1/answers?player=discord-testplayer
```

- [ example response ]
```
{
  "_links": {
    "player": {
      "href": "/v1/players/discord-testplayer"
    },
    "self": {
      "href": "/v1/answers"
    }
  },
  "status": {
    "status_code": 200,
    "message": "answers listed successfully."
  },
  "data": null,
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
    "self": {
      "href": "/v1/answers"
    }
  },
  "status": {
    "status_code": 500,
    "message": "quiz number 999 not found"
  },
  "data": null,
  "_embedded": null
}
```


# VerifyGetRank


## Rank of active player

- [ method and url ]
```
$ GET http://0.0.0.0:8080/v1/rank
```

- [ example response ]
```
{
  "_links": null,
  "status": {
    "status_code": 200,
    "message": "rank accessed successfully."
  },
  "data": [
    {
      "_id": 1,
      "name": "telegram-RainrainWu",
      "score": 8
    },
    {
      "_id": 2,
      "name": "line-!(\u0026GQ(WDUHQ(",
      "score": 6
    },
    {
      "_id": 3,
      "name": "discord-  ",
      "score": 4
    },
    {
      "_id": 4,
      "name": "facebook-__",
      "score": 3
    }
  ],
  "_embedded": null
}
```


# VerifyDeleteQuiz


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
    "status_code": 400,
    "message": "quiz number 999 not found"
  },
  "data": null,
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
$ DELETE http://0.0.0.0:8080/v1/players/discord-testplayer
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
    "message": "player deleted successfully."
  },
  "data": {
    "_id": 5,
    "name": "discord-testplayer",
    "nickname": "testplayer",
    "platform": "discord"
  },
  "_embedded": null
}
```


## Non-existed Player

- [ method and url ]
```
$ DELETE http://0.0.0.0:8080/v1/players/discord-testplayer
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
    "message": "player discord-testplayer not found"
  },
  "data": null,
  "_embedded": null
}
```
