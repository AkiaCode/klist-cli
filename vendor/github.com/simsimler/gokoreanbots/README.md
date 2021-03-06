# gokoreanbots

Go 전용 한국 디스코드 봇 리스트 SDK

서버 수 업데이트와 투표 확인 기능을 지원합니다. (랭킹 불러오기 지원 예정)

[Docs](https://pkg.go.dev/github.com/simsimler/gokoreanbots)
## 설치 방법

```
go get github.com/simsimler/gokoreanbots
```

## 사용 방법

1. 서버 수 업데이트만 할 경우

```go
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/simsimler/gokoreanbots"
)

func main() {
	session, _ := discordgo.New("DISCORD BOT TOKEN")
	err := session.Open()
	if err != nil {
		log.Panicln(err)
	}
	gokoreanbots.NewKoreanbots(session, "KOREANBOTS TOKEN", true)
	fmt.Println(session.State.User.Username + "(으)로 로그인했습니다.")
	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
}
```

* * *

2. 투표 여부 불러오기

```go
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/simsimler/gokoreanbots"
)

var (
	session, _ = discordgo.New("DISCORD BOT TOKEN")
	koreanbots = gokoreanbots.NewKoreanbots(session, "KOREANBOTS TOKEN", false)
)

func main() {
	err := session.Open()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(session.State.User.Username + "(으)로 로그인했습니다.")
	session.AddHandler(messageCreate)
	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if s.State.User.ID == m.Author.ID {
		return
	}
	switch m.Content {
	case "!votecheck":
		vote, _ := koreanbots.GetVote(m.Author.ID)
		if vote.Voted {
			_, _ = s.ChannelMessageSend(m.ChannelID, "하트를 누르셨군요")
		} else {
			_, _ = s.ChannelMessageSend(m.ChannelID, "하트를 누르지 않으셨군요")
		}
	}
}
```

