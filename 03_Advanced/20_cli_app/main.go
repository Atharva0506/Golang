package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ============================================================
// CLI App built with Cobra + Viper
//
// Cobra: handles commands, subcommands, flags, and help text.
// Viper: handles configuration — env vars, config files, and flags
//        all unified behind one interface.
//
// Run: go run main.go --help
//      go run main.go signal create --symbol BTC --action buy --price 65000
//      go run main.go signal list --symbol ETH
//      go run main.go version
// ============================================================

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// ============================================================
// Root command — the entry point of the CLI
// ============================================================

var rootCmd = &cobra.Command{
	Use:   "trading-bot",
	Short: "A CLI for managing trade signals",
	Long: `trading-bot is a command-line tool for creating and querying
trade signals in the High-Frequency Trading Bot backend.`,
	// PersistentPreRun runs before every subcommand — perfect for config loading.
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initConfig(cmd)
	},
}

func init() {
	// Persistent flags are inherited by ALL subcommands.
	rootCmd.PersistentFlags().String("config", "", "config file (default: ./config.yaml)")
	rootCmd.PersistentFlags().String("server", "http://localhost:8080", "API server URL")
	rootCmd.PersistentFlags().Bool("debug", false, "enable debug logging")

	// Tell Viper to read these flag values too.
	viper.BindPFlag("server", rootCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))

	// Register subcommands
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(signalCmd)
}

// initConfig loads configuration from file + environment variables.
func initConfig(cmd *cobra.Command) {
	cfgFile, _ := cmd.Root().PersistentFlags().GetString("config")

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.trading-bot")
	}

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.trading-bot")
	}

	// Viper automatically reads environment variables when you set a prefix.
	// e.g., TRADING_SERVER=http://prod:8080 will override the "server" config key.
	viper.SetEnvPrefix("TRADING")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		slog.Debug("config loaded", "file", viper.ConfigFileUsed())
	}

	if viper.GetBool("debug") {
		slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})))
	}
}

// ============================================================
// `version` command
// ============================================================

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("trading-bot v1.0.0")
	},
}

// ============================================================
// `signal` command group
// ============================================================

var signalCmd = &cobra.Command{
	Use:   "signal",
	Short: "Manage trade signals",
}

func init() {
	signalCmd.AddCommand(signalCreateCmd)
	signalCmd.AddCommand(signalListCmd)
}

// `signal create` — POST a new signal
var signalCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new trade signal",
	Example: `  trading-bot signal create --symbol BTC --action buy --price 65000
  trading-bot signal create --symbol ETH --action sell --price 3200`,
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol, _ := cmd.Flags().GetString("symbol")
		action, _ := cmd.Flags().GetString("action")
		price, _ := cmd.Flags().GetInt("price")

		// Validate
		if symbol == "" {
			return fmt.Errorf("--symbol is required")
		}
		if action != "buy" && action != "sell" {
			return fmt.Errorf("--action must be 'buy' or 'sell'")
		}
		if price <= 0 {
			return fmt.Errorf("--price must be a positive integer")
		}

		server := viper.GetString("server")
		slog.Info("creating signal",
			"server", server,
			"symbol", symbol,
			"action", action,
			"price", price,
		)

		// In a real CLI you'd make an HTTP call here:
		// resp, err := http.Post(server+"/api/v1/signals", "application/json", body)
		fmt.Printf("✅ Signal created: %s %s @ %d (sent to %s)\n", action, symbol, price, server)
		return nil
	},
}

func init() {
	signalCreateCmd.Flags().String("symbol", "", "trading symbol (e.g. BTC, ETH)")
	signalCreateCmd.Flags().String("action", "", "trade action: buy or sell")
	signalCreateCmd.Flags().Int("price", 0, "entry price")

	// MarkRequired makes cobra return an error automatically if the flag is absent.
	signalCreateCmd.MarkFlagRequired("symbol")
	signalCreateCmd.MarkFlagRequired("action")
	signalCreateCmd.MarkFlagRequired("price")
}

// `signal list` — GET all signals, optionally filtered by symbol
var signalListCmd = &cobra.Command{
	Use:   "list",
	Short: "List trade signals",
	Example: `  trading-bot signal list
  trading-bot signal list --symbol ETH`,
	RunE: func(cmd *cobra.Command, args []string) error {
		symbol, _ := cmd.Flags().GetString("symbol")
		server := viper.GetString("server")

		url := server + "/api/v1/signals"
		if symbol != "" {
			url += "/" + symbol
		}

		slog.Info("listing signals", "url", url)
		// In a real CLI: resp, err := http.Get(url)
		fmt.Printf("📋 Fetching signals from %s\n", url)
		return nil
	},
}

func init() {
	signalListCmd.Flags().String("symbol", "", "filter by symbol (optional)")
}
