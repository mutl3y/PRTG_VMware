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
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// templateCmd represents the template command
var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(basetemplate)
		b := bytes.NewBufferString(basetemplate)
		err := ioutil.WriteFile("prtgvmware.odt", b.Bytes(), os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// templateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// templateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var basetemplate = `
<?xml version="1.0" encoding="UTF-8"?>
<devicetemplate id="customexexml" name="prtgvmware" priority="1">
<check id="ping" meta="ping"/>
<create id="vm" kind="exexml" meta="customexexmlscan" requires="ping">
<metadata>
<exefile>
prtgvmware.exe
</exefile>
<exeparams>
metascan -U https://%host/sdk -u %windowsuser -p %windowspassword -t PRTG 
</exeparams>
</metadata>
<createdata>
      <priority>4</priority>
      <interval>300</interval>
	  <tags>prtgvmware</tags>
	  <avgmode>0</avgmode>
	  
    </createdata>
</create>
</devicetemplate>`
