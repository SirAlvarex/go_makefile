// Copyright ©2016 Kyle Gordon <siralvarex@gmail.com>
//
// 
//     The MIT License (MIT)
// 
//   Permission is hereby granted, free of charge, to any person obtaining a copy
//   of this software and associated documentation files (the "Software"), to deal
//   in the Software without restriction, including without limitation the rights
//   to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//   copies of the Software, and to permit persons to whom the Software is
//   furnished to do so, subject to the following conditions:
// 
//   The above copyright notice and this permission notice shall be included in all
//   copies or substantial portions of the Software.
// 
//   THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//   IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//   FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//   AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//   LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//   OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//   SOFTWARE.
//     
package cmd

import (

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.int.yammer.com/docker/go_makefile/goMakefile"
)

// cobraCmd respresents the cobra command
var cobraCmd = &cobra.Command{
	Use:   "cobra",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a Cli library for Go that empowers applications. This
application is a tool to generate the needed files to quickly create a Cobra
application.`,
	Run: func(cmd *cobra.Command, args []string) {
		makefile.InitCobra(viper.GetString("binary"))
   },
}

func init() {
	RootCmd.AddCommand(cobraCmd)
	cobraCmd.Flags().String("binary", "", "Name of your binary")
	viper.BindPFlag("binary", cobraCmd.Flags().Lookup("binary"))
	viper.SetDefault("binary", "")

	// Here you will define your flags and configuration settings

	// Cobra supports Persistent Flags which will work for this command and all subcommands
  // cobraCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly
	// cobraCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle" )

}
