// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	//"model"
	//"cli"
	"encoding/json"
	"utils"
	"fmt"
	"bytes"
	"os"
)

type UserList struct {
	Message string
	Data    []map[string]string
}

func printUser(user map[string]string) {
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "\t")
	if err != nil {
		fmt.Println(err.Error())
	}
	out.WriteTo(os.Stdout)
	fmt.Println()
}

func main() {
	body := "{\"message\": \"OK\", \"data\":[{\"1\":\"2\", \"3\": \"4\"}, {\"1\": \"3\"}]}"
	r := &UserList{}
	err := json.Unmarshal([]byte(string(body)), &r)
	if err != nil {
		utils.PrintError(err.Error())
	}
	for i := 0; i < len(r.Data); i++ {
		printUser(r.Data[i])
	}

	//model.EnsureAgendaDir()
	//cli.Execute()
}
