package cmd

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"
	"slices"
	"strings"
	"time"

	"github.com/coolapso/tfimport/internal/providers"
	"github.com/fatih/color"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

type TFResource struct {
	Address  string
	Type     string
	Name     string
	ImportID string
}

type settings struct {
	useExistingPlan bool
	planFilePath    string
	tool            string
	dryRun          bool
	ignoreList      []string
	delay           time.Duration
	execImport      bool
}

var (
	Version                 = "dev"
	yellow                  = color.New(color.FgYellow)
	execPlanMessage         = "Executing plan..."
	execParseMessage        = "Parsing tfplan..."
	execComputeMessage      = "Computing import IDs..."
	execImportMessage       = "Importing resources..."
	execGenImportFile       = "Generating import file..."
	infoImported            = " ✓  These resources will be imported:\n"
	infoNotImported         = " ⚠️  Resource IDs that cannot be computed often can be solved by running tfimport again, if persists you may have to handle that resource manually. Check the docs or open an issue if you need help\n"
	cannotBeImportedMessage = " ✗  These resources cannot be imported:\n"

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:     "tfimport",
		Version: Version,
		Short:   "Auto pilot for importing resources into tfstate",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			s := initTFImport(cmd)
			execute(s)
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// init rootCmd flags
func init() {
	if Version == "dev" {
		if info, ok := debug.ReadBuildInfo(); ok {
			if info.Main.Version != "" && info.Main.Version != "(devel)" {
				Version = info.Main.Version
			}
		}
	}
	rootCmd.Flags().Bool("tg", false, "executes plan and imports using terragrunt")
	rootCmd.Flags().Bool("terraform", false, "executes plan and imports using Terraform")
	rootCmd.Flags().Bool("dry-run", false, "Shows import details, only used with --run-import")
	rootCmd.Flags().String("plan-file", "", "path to a existing plan file")
	rootCmd.Flags().StringSlice("ignore", []string{}, "comma separated list of resource addresses to ignore, ex: aws_iam_role.*,aws_iam_role_policy_attachment.*, can be set multiple times")
	rootCmd.Flags().Duration("delay", 2*time.Second, "delay between imports (e.g. '1s', '5s', '2h')")
	rootCmd.Flags().Bool("run-import", false, "do not generate import file and import resources directly (you may hit some rate limits issues)")
}

// Checks if any of the flags were set and returns the tool to be used
func initTFImport(cmd *cobra.Command) *settings {
	pterm.Success.Prefix = pterm.Prefix{
		Text:  "✓",
		Style: pterm.NewStyle(pterm.FgGreen),
	}
	pterm.Error.Prefix = pterm.Prefix{
		Text:  "✗",
		Style: pterm.NewStyle(pterm.FgRed),
	}
	pterm.Warning.Prefix = pterm.Prefix{
		Text:  "⚠️",
		Style: pterm.NewStyle(pterm.FgYellow),
	}

	s := &settings{
		tool: "tofu",
	}

	if cmd.Flags().Changed("tg") {
		s.tool = "terragrunt"
	}

	if cmd.Flags().Changed("terraform") {
		s.tool = "terraform"
	}

	if cmd.Flags().Changed("run-import") {
		s.execImport = true
		s.dryRun = true
	}

	if cmd.Flags().Changed("dry-run") {
		s.dryRun = true
	}

	d, err := cmd.Flags().GetDuration("delay")
	if err != nil {
		log.Fatal("could not get delay flag: ", err)
	}
	s.delay = d

	if cmd.Flags().Changed("ignore") {
		ignoreList, err := cmd.Flags().GetStringSlice("ignore")
		if err != nil {
			log.Fatal("could not get ignore flag: ", err)
		}
		s.ignoreList = ignoreList
	}

	if cmd.Flags().Changed("plan-file") {
		s.useExistingPlan = true
		f, err := cmd.Flags().GetString("plan-file")
		if err != nil {
			log.Fatal("could not get plan-file flag: ", err)
		}
		p, err := filepath.Abs(f)
		if err != nil {
			log.Fatal("failed to resolve absolute path for plan file: ", err)
		}
		s.planFilePath = p
	}

	checkTfimporterDir()

	return s
}

func execute(s *settings) {
	if !s.useExistingPlan {
		spin, _ := spinner().Start(execPlanMessage)
		path, err := execPlan(s.tool)
		if err != nil {
			spin.Fail(execPlanMessage)
			log.Fatal(err)
		}
		s.planFilePath = path
		spin.Success(execPlanMessage)
	}

	spin, _ := spinner().Start(execParseMessage)
	plan, err := getTFJson(s.tool, s.planFilePath)
	if err != nil {
		spin.Fail(execParseMessage)
		log.Fatal(err)
	}
	spin.Success(execParseMessage)

	spin, _ = spinner().Start(execComputeMessage)
	ctx := context.Background()
	providerCtx := providers.NewProviderContext(ctx, &plan)

	var toImport []TFResource
	var notImported []TFResource

	toImport, notImported = extractTFResources(providerCtx, plan, s.ignoreList)
	spin.Success(execComputeMessage)

	if s.dryRun {
		dryRun(toImport, notImported)

		if !s.execImport {
			return
		}

		fmt.Print("\nWant to proceed with import? [y/N]: ")
		reader := bufio.NewReader(os.Stdin)
		response, _ := reader.ReadString('\n')

		if strings.ToLower(strings.TrimSpace(response)) != "y" {
			fmt.Println("Aborted.")
			return
		}
	}

	if s.execImport {
		spin, _ := spinner().Start(execImportMessage)
		withErr := execImport(s.tool, toImport, s.delay, spin)
		if withErr {
			spin.Fail(execImportMessage)
		}
		spin.Success(execImportMessage)
		return
	}

	spin, _ = spinner().Start(execGenImportFile)
	filePath, err := genImport(toImport, notImported)
	if err != nil {
		spin.Fail(execGenImportFile)
		log.Fatal(fmt.Errorf("failed to generate import file, %w", err))
	}
	successMessage := fmt.Sprintf("Import file generated at %s\n", filePath)
	spin.Success(successMessage)
	color.Green("You can now run `%s apply` to import the resources\n", s.tool)

}

func execPlan(tool string) (planFilePath string, err error) {
	cwd, _ := os.Getwd()
	planPath := filepath.Join(cwd, ".tfimport", "tfplan")
	planCmd := exec.Command(tool, "plan", fmt.Sprintf("-out=%s", planPath))
	planCmd.Dir = cwd
	var planCmdStderr bytes.Buffer
	planCmd.Stderr = &planCmdStderr
	if err := planCmd.Run(); err != nil {
		return "", fmt.Errorf("failed to execute plan: %w\n%s", err, planCmdStderr.String())
	}

	absPath, err := filepath.Abs(planPath)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path for plan file: %w", err)
	}

	return absPath, err
}

func getTFJson(tool string, planPath string) (plan tfjson.Plan, err error) {
	cwd, _ := os.Getwd()
	stateShowCmd := exec.Command(tool, "show", "-json", planPath)
	stateShowCmd.Dir = cwd
	var stateShowCmdStderr bytes.Buffer
	stateShowCmd.Stderr = &stateShowCmdStderr
	output, err := stateShowCmd.Output()
	if err != nil {
		return tfjson.Plan{}, fmt.Errorf("failed to get plan output: %w\n%s", err, stateShowCmdStderr.String())
	}

	if err := json.Unmarshal(output, &plan); err != nil {
		return tfjson.Plan{}, fmt.Errorf("failed to unmarshal plan output: %w", err)
	}

	return plan, err
}

func extractTFResources(ctx *providers.ProviderContext, plan tfjson.Plan, ignoreList []string) ([]TFResource, []TFResource) {
	var toImport []TFResource
	var notImported []TFResource
	for _, r := range plan.ResourceChanges {
		if slices.Contains(r.Change.Actions, "create") {
			if shouldIgnore(r.Address, ignoreList) {
				continue
			}

			importID := ""

			ctx.CurrentResource = r
			if after, ok := r.Change.After.(map[string]any); ok {
				importID = providers.GetImportID(ctx, r.Type, after)
			}

			if importID == "" || importID == providers.MessageProviderNotSupported {
				notImported = append(notImported, TFResource{
					Address:  r.Address,
					Type:     r.Type,
					Name:     r.Name,
					ImportID: importID,
				})
				continue
			}

			toImport = append(toImport, TFResource{
				Address:  r.Address,
				Type:     r.Type,
				Name:     r.Name,
				ImportID: importID,
			})
		}
	}
	return toImport, notImported
}

func execImport(tool string, toImport []TFResource, delay time.Duration, spin *pterm.SpinnerPrinter) (withErrors bool) {
	withErrors = false
	cwd, _ := os.Getwd()
	for _, r := range toImport {
		planCmd := exec.Command(tool, "import", r.Address, r.ImportID)
		planCmd.Dir = cwd
		var planCmdStderr bytes.Buffer
		var planCmdStdout bytes.Buffer
		planCmd.Stderr = &planCmdStderr
		planCmd.Stdout = &planCmdStdout
		if err := planCmd.Run(); err != nil {
			withErrors = true
			stdErrString := planCmdStderr.String()
			if strings.Contains(strings.ToLower(stdErrString), "non-existent remote object") {
				pterm.Warning.Printf("failed to import %s: remote resource does not exist (likely a new resource), will continue...\n", r.Address)
				if delay > 0 {
					time.Sleep(delay)
				}
				continue

			}

			pterm.Error.Printf("\nfailed to execute import for %s: %v\n%s\n", r.Address, err, planCmdStderr.String())
			pterm.Warning.Printf("\nError encountered. Pausing for 10 seconds. Press CTRL+C to abort remaining imports...\n")
			_ = spin.Stop()
			countDown(10)

			spin, _ = spinner().Start(execImportMessage)

		} else {
			pterm.Success.Printf("Successfully imported %s\n", r.Address)
		}

		if delay > 0 {
			time.Sleep(delay)
		}
	}

	return withErrors
}

func genImport(toImport []TFResource, notImported []TFResource) (path string, err error) {
	cwd, _ := os.Getwd()
	importFilePath, err := filepath.Abs(filepath.Join(cwd, "import.tf"))
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path for import file: %w", err)
	}

	var sb strings.Builder
	header := "# This file was generated by tfimport\n"
	sb.WriteString(header)
	if len(notImported) > 0 {
		sb.WriteString("#\n")
		text := fmt.Sprintf("# %s", cannotBeImportedMessage)
		sb.WriteString(text)
		lines, _ := genNotImportedLines(notImported)
		for _, l := range lines {
			text := fmt.Sprintf("# %s", l)
			sb.WriteString(text)
		}
		text = fmt.Sprintf("#\n# %s\n\n", infoNotImported)
		sb.WriteString(text)
	}

	if len(toImport) > 0 {
		for _, r := range toImport {
			importedBlock := fmt.Sprintf("import {\n  to = %s\n  id = %q\n}\n\n", r.Address, r.ImportID)
			sb.WriteString(importedBlock)
		}
	}

	if err := os.WriteFile(importFilePath, []byte(sb.String()), 0644); err != nil {
		return "", fmt.Errorf("failed to write to import file: %w", err)
	}

	return importFilePath, nil
}

func dryRun(toImport, notImported []TFResource) {
	color.Green(infoImported)
	for _, r := range toImport {
		fmt.Printf("Address: %s, ImportID: %s\n",
			r.Address,
			r.ImportID,
		)
	}
	fmt.Println()

	if len(notImported) > 0 {
		notImportedMessage(notImported)
	}
}

func notImportedMessage(notImported []TFResource) {
	color.Red(cannotBeImportedMessage)
	_, text := genNotImportedLines(notImported)
	fmt.Println(text)
	_, _ = yellow.Println(infoNotImported)
}

func genNotImportedLines(notImported []TFResource) ([]string, string) {
	var sb strings.Builder
	var notImportedLines []string
	for _, r := range notImported {
		if r.ImportID == providers.MessageProviderNotSupported {
			text := fmt.Sprintf("Resource: %s, Reason: %s\n", r.Address, r.ImportID)
			notImportedLines = append(notImportedLines, text)
			sb.WriteString(text)
		} else {
			text := fmt.Sprintf("Resource: %s, Reason: %s\n", r.Address, "ID Cannot be Computed")
			notImportedLines = append(notImportedLines, text)
			sb.WriteString(text)
		}
	}

	return notImportedLines, sb.String()
}

func checkTfimporterDir() {
	if _, err := os.Stat(".tfimport"); os.IsNotExist(err) {
		err := os.Mkdir(".tfimport", 0755)
		if err != nil {
			log.Fatal("failed to create .tfimport directory: ", err)
		}
	}
}

func shouldIgnore(address string, ignoreList []string) bool {
	for _, pattern := range ignoreList {
		matched, err := filepath.Match(pattern, address)
		if (err == nil && matched) || address == pattern {
			return true
		}
	}

	return false
}

func countDown(seconds int) {
	spin := spinner()
	spin.ShowTimer = false
	spin.MessageStyle = pterm.NewStyle(pterm.FgYellow)
	spin, _ = spin.Start("Continuing in...")
	for i := seconds; i > 0; i-- {
		spin.UpdateText(fmt.Sprintf("Continuing in %ds...", i))
		time.Sleep(1 * time.Second)
	}
	spin.Warning("Resuming imports...")
}

func spinner() *pterm.SpinnerPrinter {
	s := pterm.DefaultSpinner.WithSequence("⣷", "⣯", "⣟", "⡿", "⢿", "⣻", "⣽", "⣾")
	s.Style = pterm.NewStyle(pterm.FgGray)

	s.SuccessPrinter = &pterm.PrefixPrinter{
		MessageStyle: &pterm.ThemeDefault.SuccessMessageStyle,
		Prefix: pterm.Prefix{
			Text: pterm.ThemeDefault.Checkmark.Checked,
			// Use FgGreen (text color only) without a background color
			Style: pterm.NewStyle(pterm.FgGreen),
		},
	}

	s.FailPrinter = &pterm.PrefixPrinter{
		MessageStyle: &pterm.ThemeDefault.ErrorMessageStyle,
		Prefix: pterm.Prefix{
			Text:  pterm.ThemeDefault.Checkmark.Unchecked,
			Style: pterm.NewStyle(pterm.FgRed),
		},
	}

	s.WarningPrinter = &pterm.PrefixPrinter{
		MessageStyle: &pterm.ThemeDefault.WarningMessageStyle,
		Prefix: pterm.Prefix{
			Text:  "⚠️",
			Style: pterm.NewStyle(pterm.FgYellow),
		},
	}

	return s
}
