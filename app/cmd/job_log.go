package cmd

import (
	"fmt"
	"log"

	"github.com/linuxsuren/jenkins-cli/client"
	"github.com/spf13/cobra"
)

type JobLogOption struct {
}

var jobLogOption JobLogOption

func init() {
	jobCmd.AddCommand(jobLogCmd)
}

var jobLogCmd = &cobra.Command{
	Use:   "log",
	Short: "Print the job of your Jenkins",
	Long:  `Print the job of your Jenkins`,
	Run: func(cmd *cobra.Command, args []string) {
		if jobOption.Name == "" {
			log.Fatal("need a name")
		}

		jenkins := getCurrentJenkins()
		jclient := &client.JobClient{}
		jclient.URL = jenkins.URL
		jclient.UserName = jenkins.UserName
		jclient.Token = jenkins.Token
		jclient.Proxy = jenkins.Proxy
		jclient.ProxyAuth = jenkins.ProxyAuth

		printLog(jclient, jobOption.Name, 0)
	},
}

func printLog(jclient *client.JobClient, jobName string, start int64) {
	if status, err := jclient.Log(jobName, start); err == nil {
		fmt.Print(status.Text)
		if status.HasMore {
			printLog(jclient, jobName, status.NextStart)
		}
	} else {
		log.Fatal(err)
	}
}