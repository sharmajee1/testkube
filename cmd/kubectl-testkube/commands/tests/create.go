package tests

import (
	"fmt"
	"strconv"

	"github.com/robfig/cron"
	"github.com/spf13/cobra"

	"github.com/kubeshop/testkube/cmd/kubectl-testkube/commands/common"
	"github.com/kubeshop/testkube/pkg/api/v1/client"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/crd"
	"github.com/kubeshop/testkube/pkg/ui"
)

// NewCreateTestsCmd is a command tp create new Test Custom Resource
func NewCreateTestsCmd() *cobra.Command {

	var (
		testName                 string
		testContentType          string
		file                     string
		executorType             string
		uri                      string
		gitUri                   string
		gitBranch                string
		gitCommit                string
		gitPath                  string
		gitUsername              string
		gitToken                 string
		sourceName               string
		labels                   map[string]string
		variables                map[string]string
		secretVariables          map[string]string
		schedule                 string
		executorArgs             []string
		executionName            string
		variablesFile            string
		envs                     map[string]string
		secretEnvs               map[string]string
		httpProxy, httpsProxy    string
		gitUsernameSecret        map[string]string
		gitTokenSecret           map[string]string
		secretVariableReferences map[string]string
	)

	cmd := &cobra.Command{
		Use:     "test",
		Aliases: []string{"tests", "t"},
		Short:   "Create new Test",
		Long:    `Create new Test Custom Resource`,
		Run: func(cmd *cobra.Command, args []string) {
			crdOnly, err := strconv.ParseBool(cmd.Flag("crd-only").Value.String())
			ui.ExitOnError("parsing flag value", err)

			if testName == "" {
				ui.Failf("pass valid test name (in '--name' flag)")
			}

			namespace := cmd.Flag("namespace").Value.String()
			var client client.Client
			var testLabels map[string]string
			if !crdOnly {
				client, namespace = common.GetClient(cmd)
				test, _ := client.GetTest(testName)
				testLabels = test.Labels

				if testName == test.Name {
					ui.Failf("Test with name '%s' already exists in namespace %s", testName, namespace)
				}
			}

			err = validateCreateOptions(cmd)
			ui.ExitOnError("validating passed flags", err)

			options, err := NewUpsertTestOptionsFromFlags(cmd, testLabels)
			ui.ExitOnError("getting test options", err)

			if !crdOnly {
				executors, err := client.ListExecutors("")
				ui.ExitOnError("getting available executors", err)

				err = validateExecutorType(options.Type_, executors)
				ui.ExitOnError("validating executor type", err)
			}

			err = validateSchedule(options.Schedule)
			ui.ExitOnError("validating schedule", err)

			if !crdOnly {
				_, err = client.CreateTest(options)
				ui.ExitOnError("creating test "+testName+" in namespace "+namespace, err)

				ui.Success("Test created", namespace, "/", testName)
			} else {
				if options.Content != nil && options.Content.Data != "" {
					options.Content.Data = fmt.Sprintf("%q", options.Content.Data)
				}

				if options.ExecutionRequest != nil && options.ExecutionRequest.VariablesFile != "" {
					options.ExecutionRequest.VariablesFile = fmt.Sprintf("%q", options.ExecutionRequest.VariablesFile)
				}

				data, err := crd.ExecuteTemplate(crd.TemplateTest, options)
				ui.ExitOnError("executing crd template", err)

				ui.Info(data)
			}
		},
	}

	cmd.Flags().StringVarP(&testName, "name", "n", "", "unique test name - mandatory")
	cmd.Flags().StringVarP(&testContentType, "test-content-type", "", "", "content type of test one of string|file-uri|git-file|git-dir")

	cmd.Flags().StringVarP(&executorType, "type", "t", "", "test type (defaults to postman/collection)")

	// create options
	cmd.Flags().StringVarP(&file, "file", "f", "", "test file - will be read from stdin if not specified")
	cmd.Flags().StringVarP(&uri, "uri", "", "", "URI of resource - will be loaded by http GET")
	cmd.Flags().StringVarP(&gitUri, "git-uri", "", "", "Git repository uri")
	cmd.Flags().StringVarP(&gitBranch, "git-branch", "", "", "if uri is git repository we can set additional branch parameter")
	cmd.Flags().StringVarP(&gitCommit, "git-commit", "", "", "if uri is git repository we can use commit id (sha) parameter")
	cmd.Flags().StringVarP(&gitPath, "git-path", "", "", "if repository is big we need to define additional path to directory/file to checkout partially")
	cmd.Flags().StringVarP(&gitUsername, "git-username", "", "", "if git repository is private we can use username as an auth parameter")
	cmd.Flags().StringVarP(&gitToken, "git-token", "", "", "if git repository is private we can use token as an auth parameter")
	cmd.Flags().StringVarP(&sourceName, "source", "", "", "source name - will be used together with content parameters")
	cmd.Flags().StringToStringVarP(&labels, "label", "l", nil, "label key value pair: --label key1=value1")
	cmd.Flags().StringToStringVarP(&variables, "variable", "v", nil, "variable key value pair: --variable key1=value1")
	cmd.Flags().StringToStringVarP(&secretVariables, "secret-variable", "s", nil, "secret variable key value pair: --secret-variable key1=value1")
	cmd.Flags().StringVarP(&schedule, "schedule", "", "", "test schedule in a cronjob form: * * * * *")
	cmd.Flags().StringArrayVarP(&executorArgs, "executor-args", "", []string{}, "executor binary additional arguments")
	cmd.Flags().StringVarP(&executionName, "execution-name", "", "", "execution name, if empty will be autogenerated")
	cmd.Flags().StringVarP(&variablesFile, "variables-file", "", "", "variables file path, e.g. postman env file - will be passed to executor if supported")
	cmd.Flags().StringToStringVarP(&envs, "env", "", map[string]string{}, "envs in a form of name1=val1 passed to executor")
	cmd.Flags().StringToStringVarP(&secretEnvs, "secret-env", "", map[string]string{}, "secret envs in a form of secret_name1=secret_key1 passed to executor")
	cmd.Flags().StringVar(&httpProxy, "http-proxy", "", "http proxy for executor containers")
	cmd.Flags().StringVar(&httpsProxy, "https-proxy", "", "https proxy for executor containers")
	cmd.Flags().StringToStringVarP(&gitUsernameSecret, "git-username-secret", "", map[string]string{}, "git username secret in a form of secret_name1=secret_key1 for private repository")
	cmd.Flags().StringToStringVarP(&gitTokenSecret, "git-token-secret", "", map[string]string{}, "git token secret in a form of secret_name1=secret_key1 for private repository")
	cmd.Flags().StringToStringVarP(&secretVariableReferences, "secret-variable-reference", "", nil, "secret variable references in a form name1=secret_name1=secret_key1")

	return cmd
}

func validateCreateOptions(cmd *cobra.Command) error {
	gitUri := cmd.Flag("git-uri").Value.String()
	gitBranch := cmd.Flag("git-branch").Value.String()
	gitCommit := cmd.Flag("git-commit").Value.String()
	gitPath := cmd.Flag("git-path").Value.String()
	gitUsername := cmd.Flag("git-username").Value.String()
	gitToken := cmd.Flag("git-token").Value.String()
	gitUsernameSecret, err := cmd.Flags().GetStringToString("git-username-secret")
	if err != nil {
		return err
	}

	gitTokenSecret, err := cmd.Flags().GetStringToString("git-token-secret")
	if err != nil {
		return err
	}

	file := cmd.Flag("file").Value.String()
	uri := cmd.Flag("uri").Value.String()
	sourceName := cmd.Flag("source").Value.String()

	hasGitParams := gitBranch != "" || gitCommit != "" || gitPath != "" || gitUri != "" || gitToken != "" || gitUsername != "" ||
		len(gitUsernameSecret) > 0 || len(gitTokenSecret) > 0

	if hasGitParams && uri != "" {
		return fmt.Errorf("found git params and `--uri` flag, please use `--git-uri` for git based repo or `--uri` without git based params")
	}
	if hasGitParams && file != "" {
		return fmt.Errorf("found git params and `--file` flag, please use `--git-uri` for git based repo or `--file` without git based params")
	}

	if file != "" && uri != "" {
		return fmt.Errorf("please pass only one of `--file` and `--uri`")
	}

	if hasGitParams {
		if gitUri == "" && sourceName == "" {
			return fmt.Errorf("please pass valid `--git-uri` flag")
		}
		if gitBranch != "" && gitCommit != "" {
			return fmt.Errorf("please pass only one of `--git-branch` or `--git-commit`")
		}
	}

	if len(gitUsernameSecret) > 1 {
		return fmt.Errorf("please pass only one secret reference for git username")
	}

	if len(gitTokenSecret) > 1 {
		return fmt.Errorf("please pass only one secret reference for git token")
	}

	if (gitUsername != "" || gitToken != "") && (len(gitUsernameSecret) > 0 || len(gitTokenSecret) > 0) {
		return fmt.Errorf("please pass git credentials either as direct values or as secret references")
	}

	return nil
}

func validateExecutorType(executorType string, executors testkube.ExecutorsDetails) error {
	typeValid := false
	executorTypes := []string{}

	for _, ed := range executors {
		executorTypes = append(executorTypes, ed.Executor.Types...)
		for _, et := range ed.Executor.Types {
			if et == executorType {
				typeValid = true
			}
		}
	}

	if !typeValid {
		return fmt.Errorf("invalid executor type '%s' use one of: %v", executorType, executorTypes)
	}

	return nil
}

func validateSchedule(schedule string) error {
	if schedule == "" {
		return nil
	}

	specParser := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	if _, err := specParser.Parse(schedule); err != nil {
		return err
	}

	return nil
}
