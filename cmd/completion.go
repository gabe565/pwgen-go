package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const CompletionFlag = "completion"

func registerCompletionFlag(cmd *cobra.Command) {
	cmd.Flags().String(CompletionFlag, "", "Output command-line completion code for the specified shell. Can be 'bash', 'zsh', 'fish', or 'powershell'.")
	err := cmd.RegisterFlagCompletionFunc(CompletionFlag, completionCompletion)
	if err != nil {
		panic(err)
	}
}

func completionCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return []string{"bash", "zsh", "fish", "powershell"}, cobra.ShellCompDirectiveNoFileComp
}

func completion(cmd *cobra.Command, args []string) error {
	completionFlag, err := cmd.Flags().GetString(CompletionFlag)
	if err != nil {
		panic(err)
	}

	switch completionFlag {
	case "bash":
		if err := cmd.Root().GenBashCompletion(cmd.OutOrStdout()); err != nil {
			return err
		}
	case "zsh":
		if err := cmd.Root().GenZshCompletion(cmd.OutOrStdout()); err != nil {
			return err
		}
	case "fish":
		if err := cmd.Root().GenFishCompletion(cmd.OutOrStdout(), true); err != nil {
			return err
		}
	case "powershell":
		if err := cmd.Root().GenPowerShellCompletionWithDesc(cmd.OutOrStdout()); err != nil {
			return err
		}
	default:
		return fmt.Errorf("%v: invalid shell", completionFlag)
	}
	return nil
}
