// Copyright © 2016 Peter Teich <mail@pteich.xyz>
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
	"github.com/spf13/cobra"

	"github.com/pteich/fofiwano"

	"log"

	"github.com/spf13/viper"
)

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "starts watching for filesystem changes",
	Long: `Fofiwano is a CLI tool to watch folders or files for modifications like added, deleted or changed files.
It then sends a notification to a specific endpoint, e.g. a Slack channel, an URI (HTTP-Request) or executes a command.
`,
	Run: func(cmd *cobra.Command, args []string) {

		var watches []fofiwano.Watcher

		if err := viper.UnmarshalKey("watching", &watches); err != nil {
			log.Fatal(err)
		}

		fofiwano.Watch(watches)
	},
}

func init() {
	RootCmd.AddCommand(watchCmd)
}
