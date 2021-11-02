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
package depth

import (
	"github.com/merext/binancectl/v2/api/stream"
	"github.com/spf13/cobra"
)

var symbol string

// depthCmd represents the depth command
var DepthCmd = &cobra.Command{
	Use:   "depth",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines`,
	Run: func(cmd *cobra.Command, args []string) {
		stream.Depth(symbol)
	},
}

func init() {
}
