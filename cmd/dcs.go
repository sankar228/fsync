/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"encoding/json"
	"log"
	"strings"

	"github.com/sankar228/fsync/model"
	"github.com/spf13/cobra"
)

// dcsCmd represents the dcs command
var dcsCmd = &cobra.Command{
	Use:   "dcs",
	Short: "Data Copy Service",
	Long: `This Service is used for Copy data from source to sync
		Track the status of the copy in the SQLLite db.`,
	Run: DataCopyService,
}

func DataCopyService(cmd *cobra.Command, args []string) {

	a, _ := cmd.Flags().GetString("attached")
	log.Printf("attached: %v\n", a)
	source, _ := cmd.Flags().GetString("source")
	dest, _ := cmd.Flags().GetString("dest")
	log.Println("DCS called: ", source, " : ", dest)

	srcConn := GetConnection(source)
	resp, _ := json.Marshal(*srcConn)
	log.Printf("source: %v\n", string(resp))
	destConn := GetConnection(dest)
	resp, _ = json.Marshal(*destConn)
	log.Printf("destination: %v\n", string(resp))

}

/*
sftp://host:port;[user=****];[pass=****];[passkey=****];[path=****]
*/
func GetConnection(connstr string) *model.Connection {
	connTokens := strings.Split(connstr, "://")

	conType := connTokens[0]
	ConnctStr := connTokens[1]
	conProps := make(map[string]string)
	hoststr := strings.Split(ConnctStr, ";")

	for _, prop := range hoststr {
		props := strings.Split(prop, "=")
		if len(props) >= 2 {
			conProps[strings.Split(prop, "=")[0]] = strings.Split(prop, "=")[1]
		} else {
			conProps[strings.Split(prop, "=")[0]] = ""
		}
	}

	log.Printf("conType: %s , conProps: %v\n", conType, conProps)
	log.Printf("username: %s\n", conProps["user"])

	connection := &model.Connection{
		Type:      conType,
		ConnctStr: ConnctStr,
	}
	if v, ok := conProps["user"]; ok {
		connection.User = v
	}
	if v, ok := conProps["pass"]; ok {
		connection.Password = v
	}
	if v, ok := conProps["passkey"]; ok {
		connection.PassKey = v
	}
	if v, ok := conProps["path"]; ok {
		connection.Path = v
	}
	return connection
}

func init() {
	rootCmd.AddCommand(dcsCmd)

	dcsCmd.Flags().StringP("attached", "a", "", "Run the Copy service in attached mode, Used only for testing single copy")
	dcsCmd.Flags().Lookup("attached").NoOptDefVal = "true"

	dcsCmd.Flags().StringP("background", "b", "", `This is the default behaviour.
		Run the Copy service in ditatched mode, Service will be running in the background.
		clients can use RESTApi to invoke copy task
	`)

	dcsCmd.Flags().StringP("source", "s", "", "source connection string sftp://<user>:<(password/keyfile)>@<host> \n or file:///localpath")
	dcsCmd.MarkFlagRequired("source")

	dcsCmd.Flags().StringP("dest", "d", "", "destination connection string sftp://<user>:<(password/keyfile)>@<host> \n or file:///localpath")
	dcsCmd.MarkFlagRequired("dest")
}
