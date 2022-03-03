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
package cmd

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "github.com/PuerkitoBio/goquery"
	. "github.com/logrusorgru/aurora"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "yah_news",
    Short: "Fetch All yahoo articles",
    Long: `A longer description that spans multiple lines and likely contains
    examples and usage of using your application. For example:

    Cobra is a CLI library for Go that empowers applications.
    This application is a tool to generate the needed files
    to quickly create a Cobra application.`,
    // Uncomment the following line if your bare application
    // has an action associated with it:
    Run: func(cmd *cobra.Command, args []string) {
		titles:= [7]string{"国内","経済","エンタメ","スポーツ","国際","IT","地域"}
        if len(os.Args)==1{
            url := [7]string{"https://news.yahoo.co.jp/categories/domestic",
				"https://news.yahoo.co.jp/categories/business",
				"https://news.yahoo.co.jp/categories/entertainment",
				"https://news.yahoo.co.jp/categories/sports",
				"https://news.yahoo.co.jp/categories/world",
				"https://news.yahoo.co.jp/categories/it",
				"https://news.yahoo.co.jp/categories/local"}

			for i:=0; i < len(url); i++{
				fetch_news(url[i],titles[i])
			}
        }
    },
}

func fetch_news(url string,title string){
    doc, err := goquery.NewDocument(url)
    if err != nil {
        panic(err)
    }
	title_text := Bold(Magenta("=== " + title + " ==="))
    articles := doc.Find("div.sc-hAcydR.Lrbus")
    innerarticles:= articles.Find("a")
	fmt.Println(title_text)
    innerarticles.Each(func(i int, s *goquery.Selection){
		attr,_ := s.Attr("href")
		fmt.Println(s.Text())
        fmt.Println(Bold(Cyan(attr)))
    })
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
    cobra.CheckErr(rootCmd.Execute())
}

func init() {
    cobra.OnInitialize(initConfig)

    // Here you will define your flags and configuration settings.
    // Cobra supports persistent flags, which, if defined here,
    // will be global for your application.

    rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.yah_news.yaml)")

    // Cobra also supports local flags, which will only run
    // when this action is called directly.
    rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
    if cfgFile != "" {
        // Use config file from the flag.
        viper.SetConfigFile(cfgFile)
    } else {
        // Find home directory.
        home, err := os.UserHomeDir()
        cobra.CheckErr(err)

        // Search config in home directory with name ".yah_news" (without extension).
        viper.AddConfigPath(home)
        viper.SetConfigType("yaml")
        viper.SetConfigName(".yah_news")
    }

    viper.AutomaticEnv() // read in environment variables that match

    // If a config file is found, read it in.
    if err := viper.ReadInConfig(); err == nil {
        fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
    }
}
