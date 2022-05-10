/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"sort"

	"github.com/pkg/sftp"
	"github.com/sankar228/fsync/consts"
	"github.com/sankar228/fsync/utils"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

// mirrorCmd represents the mirror command
var mirrorCmd = &cobra.Command{
	Use:   "mirror",
	Short: "Fast and reliable file sync between the systems",
	Long: `Using this utility we can mirror files between source and destination
	This is meant to sync-up directories fast and reliable.`,
	Run: Mirror,
}

func Mirror(cmd *cobra.Command, args []string) {
	source, _ := cmd.Flags().GetString("source")
	dest, _ := cmd.Flags().GetString("dest")
	fmt.Println("mirror called: ", source, " : ", dest)

	source = "[nedetluser!@#123]:nedetluser@10.77.129.5"

	parts := consts.HostRe.FindStringSubmatch(source)
	pass := parts[1]
	user := parts[2]
	host := parts[3] + ":22"
	//keyfile := ""
	fmt.Printf("user: %s, pass: %s, host: %s\n", user, pass, host)

	List(host, user, pass, "/home/nedetluser/ned2parser")
	//copyFile(user, pass, host)

}

func List(host string, user string, pass string, dir string) {
	conn, err := CreateConnection(host, user, pass)
	if utils.CheckErr(err) {
		panic(err)
	}

	client, err := sftp.NewClient(conn)
	if utils.CheckErr(err) {
		panic(err)
	}

	items, err := client.ReadDir(dir)
	if utils.CheckErr(err) {
		panic(err)
	}

	fmt.Println("listing dir: ", dir)
	sort.Sort(ByModTime(items))
	for _, item := range items {
		isdir := "F"
		if item.IsDir() {
			isdir = "D"
		}
		F := fmt.Sprintf("%s %v %s %s\n", isdir, item.ModTime().Format("2006-01-02 15:04"), utils.ByteCountDecimal(item.Size()), item.Name())
		fmt.Println(F)
	}
}
func CreateConnection(host string, user string, pass string) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", host, config)
	return conn, err
}

func copyFile(user string, pass string, host string) {
	conn, err := CreateConnection(host, user, pass)

	if utils.CheckErr(err) {
		panic(err)
	}
	defer conn.Close()
	sftp, err := sftp.NewClient(conn)

	if utils.CheckErr(err) {
		panic(err)
	}
	defer sftp.Close()

	outF, err := os.Create("C:\\Users\\skataba1\\code\\fsync\\crontab_bkp")
	if utils.CheckErr(err) {
		panic(err)
	}
	inF, err := sftp.Open("/home/nedetluser/ned2parser/crontab_bkp")
	if utils.CheckErr(err) {
		panic(err)
	}
	written, err := io.Copy(outF, inF)
	if utils.CheckErr(err) {
		panic(err)
	}
	fmt.Printf("File copied, number of bytes written: %d\n", written)
}

func init() {
	rootCmd.AddCommand(mirrorCmd)

	mirrorCmd.Flags().StringP("source", "s", "", "source connection string sftp://<user>:<(password/keyfile)>@<host> \n or file:///localpath")
	mirrorCmd.MarkFlagRequired("source")

	mirrorCmd.Flags().StringP("dest", "d", "", "destination connection string sftp://<user>:<(password/keyfile)>@<host> \n or file:///localpath")
	mirrorCmd.MarkFlagRequired("dest")
}

type ByModTime []fs.FileInfo

func (fs ByModTime) Len() int           { return len(fs) }
func (fs ByModTime) Less(i, j int) bool { return fs[i].ModTime().After(fs[j].ModTime()) }
func (fs ByModTime) Swap(i, j int)      { fs[i], fs[j] = fs[j], fs[i] }
