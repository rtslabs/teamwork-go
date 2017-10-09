// Copyright © 2017 Gabriel Duke <gabeduke@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/nanobox-io/golang-scribble"

    "teamworkgo/lib"
)

// buildCacheCmd represents the buildCache command
var buildCacheCmd = &cobra.Command{
	Use:   "buildCache",
	Short: "Build the Projects, Task lists and Tasks to local cache",
	Long: `Long description`,
	Run: func(cmd *cobra.Command, args []string) {

		db, _ := scribble.New("/tmp/teamworkgo/.cache", nil)

		projects := lib.GetAllProjects()
		for _, project := range projects.ProjectBeanList {
			taskLists, _ := lib.GetTaskLists(project.Id)
			fmt.Println(project.Name)
			for _, taskList := range taskLists.ProjectBeanList {
				//tasks := lib.GetTask(taskList.Id)
				fmt.Println(taskList.Name)
			}

		}
		// TODO add task lists
		//taskLists := lib.GetTaskLists("projects.json")

		for _, name := range projects.ProjectBeanList {
			db.Write("teamwork", name.Name, "Placeholder for Task Lists2")
		}

		fmt.Println(projects)
	},
}

func init() {
	RootCmd.AddCommand(buildCacheCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCacheCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCacheCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
