/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"os"
	"github.com/spf13/cobra"
	"bufio"
	"strconv"
	"encoding/json"
	"time"
)
// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
	username, _ := cmd.Flags().GetString("user")
	password, _ := cmd.Flags().GetString("password")
logf,logopenerr:=os.OpenFile("log.txt",os.O_CREATE|os.O_APPEND|os.O_RDWR,0644)
if logopenerr!=nil{
fmt.Println("log file open error",logopenerr)
}
	_,err:=os.Stat("entity/curUser.txt")
	if err==nil {
	f,ierr:=os.Open("entity/curUser.txt")
	if ierr!=nil{
	fmt.Println("input file open error:",ierr)
	return
	}
	r:=bufio.NewReader(f)
	buf,err:=r.ReadBytes('\n')
	if err!=nil{
	fmt.Println("file read error:",err)
	return
	}
	buf=buf[:len(buf)-1]
	savedNum,terr:=strconv.Atoi(string(buf))
	if terr!=nil{
	fmt.Println("transform error:",terr)
	return 
	}
	flag:=0
	for i:=0;i<savedNum;i++{
		readJson,rerr:=r.ReadBytes('\n')
		readJson=readJson[:len(readJson)-1]
		if rerr!=nil{
		fmt.Println("file read error:",rerr)
		return
		}
		var tmpUser User
		uerr:=json.Unmarshal(readJson,&tmpUser)
		if uerr!=nil{
		fmt.Println("transform error:",uerr)
		return
		}
		if tmpUser.Password==password && tmpUser.Name==username{
		flag=1
		fmt.Println("Login success!")
		_,werr_:=logf.WriteString(time.Now().Format("2006-01-02 15:04:05")+" "+tmpUser.Name+" login success!"+"\n")
if werr_!=nil{
fmt.Println("log file write error:",werr_)
return
}
		break
		}
	}
	if flag==0{
	fmt.Println("Login fail!")
		_,werr_:=logf.WriteString(time.Now().Format("2006-01-02 15:04:05")+" login fail!"+"\n")
if werr_!=nil{
fmt.Println("log file write error:",werr_)
return
}
	}
	f.Close()
}else{
fmt.Println("Login fail!")
		_,werr_:=logf.WriteString(time.Now().Format("2006-01-02 15:04:05")+" login fail!"+"\n")
if werr_!=nil{
fmt.Println("log file write error:",werr_)
return
}
}

	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	loginCmd.Flags().StringP("user", "u", "Anonymous", "Help message for username")
	loginCmd.Flags().StringP("password", "p", "N", "Help message for password")
}
