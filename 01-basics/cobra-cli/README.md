# Building a CLI tool with Cobra

Cobra is a library for building cli tools with go

https://github.com/spf13/cobra

## Installing cobra

```powershell
go install github.com/spf13/cobra-cli@latest
```

## Initializing a project

```powershell
cobra-cli init toolbox
```

this add a new folder named toolbox. if you already have an existing toolbox folder,
you can move those files out and delete the new folder.


## Adding a cmd pallete

Create a pallete called net

```powershell
cobra-cli add net
```

Move the new cmd into the net folder

I then need to make the net cmd a public package. I do this by changing the n to a N.

```go
// netCmd represents the net command
var NetCmd = &cobra.Command{ // change n -> N to make the cmd public
	Use:   "net",
	Short: "Net is a pallete that contains network commds",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
```

I Then need to add the net cmd to the root cmd.

Import the net cmd package
```go
import (
	"os"
	"toolbox/cmd/net"

	"github.com/spf13/cobra"
)
```

then create a new function for importing the pallates

```go
func addSubCommandPalletes() {
	rootCmd.AddCommand(net.NetCmd)
}
```

and then call the function in the init function:


```go
func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.toolbox.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addSubCommandPalletes()
}
```

## Creating a sub command to a pallete

To create a sub command i will use the cobra cli tool

```powershell
cobra-cli add ping
```

Then move the ping command into the net folder

i now need to add the ping command to the net command pallete
and i do this within the ping command init function

```go
func init() {
	// Here you will define your flags and configuration settings.
	NetCmd.AddCommand(pingCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
```

## Adding flags to a cmd

Inside the ping command i have created a url flag.
i need to initialize a variable for the flag


```go
var (
	urlPath string
)
```

and i then need to add it inside the init function. I will also create the logic
for making the flag required.

```go
func init() {
	pingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "The url to ping") // <-- adding the flag

	if err := pingCmd.MarkFlagRequired("url"); err != nil { // <-- adding required and handling error
		fmt.Println(err)
	}

	// Here you will define your flags and configuration settings.
	NetCmd.AddCommand(pingCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
```

## Viper 

There is a companion library for configuring you cli too named viper

https://github.com/spf13/viper

To setup a cli project with viper you should initilize with 

```powershell
cobra-cli init toolbox --viper
```

