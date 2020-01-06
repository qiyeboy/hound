git clone /hound/.git
cd hound

go mod init  
go mod tidy

vim config.JSON
```
{
  "dbpath" : "db",
  "repos" : {
     "Local" : {
        "url" : "/Users/conanhu/Desktop/goprojects/src/hound/api/"
      },
    "Git" : {
      "url" : "https://github.com/conan25216/tcf.git",
      "exclude-dot-files": true,
      "exclude-files":["README.md","mytest.go"]
    }
  }
}
```
go install cmd/houndd/main.go

***waiting for index built and server on***

main
