/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>
hello
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
	"encoding/json"
	"strconv"
	"time"
)
type User struct{
	Name string
	Password string 
	Email string
	Phone string 
}
// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
logf,logopenerr:=os.OpenFile("log.txt",os.O_CREATE|os.O_APPEND|os.O_RDWR,0644)
if logopenerr!=nil{
fmt.Println("log file open error",logopenerr)
}
	username, _ := cmd.Flags().GetString("user")
	password, _ := cmd.Flags().GetString("password")
	email, _ := cmd.Flags().GetString("email")
	phone, _ := cmd.Flags().GetString("phone")
	var count int =0
	dynamic:=make([]User,0)
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
		}
		count=savedNum
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
		dynamic=append(dynamic,tmpUser)
		}
		f.Close()
	}	


	eflag:=0	
	for i:=0;i<count;i++{
		if dynamic[i].Name==username{
		fmt.Println("The user name has existed,register fail!")
		_,werr_:=logf.WriteString(time.Now().Format("2006-01-02 15:04:05")+" login fail!"+"\n")
		if werr_!=nil{
		fmt.Println("log file write error:",werr_)
		return
		}
		eflag=1
		break
		}
	}
	if eflag==0{
		count++
		dynamic=append(dynamic,User{username,password,email,phone})
		of,oerr:=os.Create("entity/curUser.txt")
		defer of.Close()
		if oerr!=nil{
			fmt.Println("curUser create err:",oerr)
			return
		}
		_,err_:=of.WriteString(strconv.Itoa(count)+"\n")
		if err_!=nil{
		fmt.Println("file write err:",err_)
		return
		}
		for i:=0;i<count;i++{
		tmpUser:=dynamic[i]
		tmpJson,jerr:=json.Marshal(tmpUser)
			if jerr!=nil{
			fmt.Println("Json transform error:",jerr)
			return
			}
		_,jerr_:=of.WriteString(string(tmpJson)+"\n")
			if jerr_!=nil{
			fmt.Println("file write err:",jerr_)
			return
			}
		}
		fmt.Println("Register success!")
		_,werr_:=logf.WriteString(time.Now().Format("2006-01-02 15:04:05")+" "+username+" login success!"+"\n")
		if werr_!=nil{
		fmt.Println("log file write error:",werr_)
		return
		}
	}


	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("user", "u", "Anonymous", "Help message for username")
	registerCmd.Flags().StringP("password", "p", "None", "Help message for password")
	registerCmd.Flags().StringP("email", "e", "None", "Help message for email")
	registerCmd.Flags().StringP("phone", "f", "None", "Help message for phone")
}
