/*
   Copyright 2017 nerdicbynature

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

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func usage() {
	fmt.Println("dc-check <docker-compose.yml>")
	os.Exit(1)
}

func main() {
	type envMap map[string]string
	var values envMap = make(envMap)

	if len(os.Args) < 2 {
		usage()
	}
	dcFile := os.Args[1]

	dcContent, err := ioutil.ReadFile(dcFile)
	check(err)

	for _, e := range os.Environ() {
		envVar := strings.Split(e, "=")
		values[envVar[0]] = envVar[1]
	}

	tmpl, err := template.New(dcFile).Parse(string(dcContent))
	if err != nil {
		panic(err)
	}

	var dummy map[string]envMap = make(map[string]envMap)
	dummy["Values"] = values
	err = tmpl.Execute(os.Stdout, dummy)
	if err != nil {
		panic(err)
	}
}
