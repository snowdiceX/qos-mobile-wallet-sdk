package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	sdk "github.com/QOSGroup/qos-mobile-wallet-sdk"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	FlagHome     = "home"
	FlagMnemonic = "mnemonic"
	FlagPass     = "passwd"
)

func main() {
	cobra.EnableCommandSorting = false

	rootCmd := &cobra.Command{
		Use:   "qwalletcli",
		Short: "Command line interface for qwallet-sdk",
	}
	rootCmd.PersistentFlags().String(FlagHome, "$HOME/.qoscli/", "home dir")
	viper.BindPFlag(FlagHome, rootCmd.PersistentFlags().Lookup(FlagHome))

	cmdRecover := &cobra.Command{
		Use:   "recover",
		Short: "Recover account from mnemonic",
		RunE:  doRecover,
	}
	cmdRecover.Flags().StringP(FlagMnemonic, "m", "", "Mnemonic")
	viper.BindPFlag(FlagMnemonic, cmdRecover.Flags().Lookup(FlagMnemonic))
	cmdRecover.Flags().StringP(FlagPass, "p", "", "passworld")
	viper.BindPFlag(FlagPass, cmdRecover.Flags().Lookup(FlagPass))

	cmdList := &cobra.Command{
		Use:   "list",
		Short: "List account",
		RunE:  doList,
	}
	rootCmd.AddCommand(cmdRecover, cmdList)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func doList(_ *cobra.Command, _ []string) error {
	home := os.ExpandEnv(viper.GetString(FlagHome))
	viper.Set(FlagHome, home)

	sdk.InitWallet(home)
	ret := sdk.ListAllAccounts()
	showJsonString(ret)
	return nil
}

func doRecover(_ *cobra.Command, _ []string) error {
	home := os.ExpandEnv(viper.GetString(FlagHome))
	viper.Set(FlagHome, home)
	// qosacc19ee0dmedngya6akhyyc2fllqq8hmrgkn24n62g
	// m := "wage maximum acid car catalog aisle attend rookie outdoor unusual donkey script maximum weather tiger expire negative wine evidence grass lemon forget concert planet"
	// regular trumpet envelope oak jar loop comic turkey forest frozen divide pond identify increase magnet power alarm develop depart manual dry gap coin bubble

	m := viper.GetString(FlagMnemonic)
	if m == "" {
		return fmt.Errorf("please input the mnemonic")
	}
	p := viper.GetString(FlagPass)
	if p == "" {
		return fmt.Errorf("please input the passwd")
	}
	sdk.InitWallet(home)
	ret := sdk.ImportMnemonic(m, p)
	showJsonString(ret)
	return nil
}

func showJsonString(js string) (err error) {
	var out bytes.Buffer
	if err = json.Indent(&out, []byte(js), "", "  "); err == nil {
		fmt.Println(out.String())
	}
	return
}
