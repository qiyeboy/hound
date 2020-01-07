git clone https://github.com/conan25216/hound.git

cd hound

go mod init  
go mod tidy

vim config.json
```
{
  "dbpath" : "db",
  "repos" : {
      "Git" : {
          "url" : "https://github.com/conan25216/hound.git",
        },
     "Local" : {
         "url" : "/yourlocalpath/",
         "exclude-files":["README.md"],
        }
    }
}
```

go install cmds/houndd/main.go

main
