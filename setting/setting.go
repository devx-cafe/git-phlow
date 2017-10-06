package setting

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/go-ini/ini"
	"github.com/praqma/git-phlow/executor"
	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/plugins"
	"errors"
)

//Config git group.name
const (
	defaultBlock                    = "phlow"
	configFileName                  = ".gitconfig"
	configDefaultBlock              = "phlow"
	configServiceField              = "service"
	configRemoteField               = "remote"
	configIssueURLField             = "issue-url"
	configPipelineField             = "pipeline"
	configIntegrationBranchField    = "integration-branch"
	configDeliveryBranchPrefixField = "delivery-branch-prefix"
)

//Default configuration
const (
	InternalDefaultService              = "github"
	InternalDefaultIntegrationBranch    = "master"
	InternalDefaultRemote               = "origin"
	InternalDefaultURL                  = "https://api.github.com"
	InternalDefaultDeliveryBranchPrefix = "ready"
	InternalDefaultScope                = "internal"
	InternalDefaultFile                 = "none"
	InternalDefaultOrigin               = "none"
)

//ProjectSetting ...
type ProjectSetting struct {
	Service              string `ini:"service"`
	IntegrationBranch    string `ini:"integration-branch"`
	Remote               string `ini:"remote"`
	IssueURL             string `ini:"issue-url"`
	PipelineUrl          string `ini:"pipeline"`
	DeliveryBranchPrefix string `ini:"delivery-branch-prefix"`
	Scope                string
	INIBlock             string
}

//ToString ...
func (setting *ProjectSetting) ToString() string {
	r := reflect.ValueOf(setting).Elem()
	t := r.Type()
	msg := ""
	for i := 0; i < t.NumField(); i++ {
		msg += t.Field(i).Name + ": " + r.Field(i).String() + "\n"
	}
	return msg
}

type ConfigError struct {
	errorMessage string
}

func (ce *ConfigError) Error() string {
	return ce.errorMessage
}

func NewConfigError(field, block string) error {
	return &ConfigError{fmt.Sprintf("Error in configuration\n"+
		"Non-optional field missing: %s \nIn configuration block: %s \n", field, block)}
}

//NewProjectStg ...
//initializes project settings from .phlow INI files
func NewProjectStg(INIBlock string) *ProjectSetting {
	return LoadSettings(INIBlock, githandler.Git{Run: executor.RunGit})
}

func LoadSettings(INIBlock string, git githandler.Git) *ProjectSetting {

	//no params have been given, search for default
	if INIBlock == "" {
		INIBlock = configDefaultBlock
	}

	//Load all configurations using git config
	//Errors result in an empty config string, which is git's way to return empty config
	service, errS := git.Config("--get", fmt.Sprintf("%s.%s", INIBlock, configServiceField))
	serviceURL, errSU := git.Config("--get", INIBlock+"."+configIssueURLField)
	remote, errR := git.Config("--get", INIBlock+"."+configRemoteField)
	deliveryBranch, errDB := git.Config("--get", INIBlock+"."+configDeliveryBranchPrefixField)
	integrationBranch, errIB := git.Config("--get", INIBlock+"."+configIntegrationBranchField)

	loadedSetting := ProjectSetting{
		Service:              service,
		IssueURL:             serviceURL,
		Remote:               remote,
		IntegrationBranch:    integrationBranch,
		DeliveryBranchPrefix: deliveryBranch,
		INIBlock:             INIBlock,
	}

	err := ValidateLoadedSetting(&loadedSetting)
	if err != nil {
		//It is the default config, so we will just go to the internal default
		if INIBlock == configDefaultBlock {
			defaultBranch, err := GetDefaultBranchFromInternalDefault()
			if err != nil || strings.TrimSpace(defaultBranch) == "" {
				defaultBranch = InternalDefaultIntegrationBranch
			}
			return &ProjectSetting{
				Service:              InternalDefaultService,
				IntegrationBranch:    defaultBranch,
				Remote:               InternalDefaultRemote,
				IssueURL:             InternalDefaultURL,
				DeliveryBranchPrefix: InternalDefaultDeliveryBranchPrefix,
				PipelineUrl:          InternalDefaultOrigin,
				Scope:                InternalDefaultScope,
			}
		}
		//If all loads fail, we assume that the group does not exists
		if errS != nil && errSU != nil && errR != nil && errDB != nil && errIB != nil {
			fmt.Printf("Error: '%s' configuration does not seem to exist in you configuration files.\n", INIBlock)
			os.Exit(1)
		}
		//Only one or more fields are missing, so we print the error
		fmt.Println(err)
		os.Exit(1)
	}

	return &loadedSetting
}

//BootstrapPhlowConfig ...
//Creates a new .phlow ini file on given location
func BootstrapPhlowConfig(local, integrationBranch string) error {
	pathSeparator := os.PathSeparator
	cfg := ini.Empty()
	sec, _ := cfg.NewSection("phlow")
	sec.Key(configRemoteField).SetValue(InternalDefaultRemote)
	sec.Key(configServiceField).SetValue(InternalDefaultService)
	sec.Key(configIntegrationBranchField).SetValue(integrationBranch)
	sec.Key(configIssueURLField).SetValue(InternalDefaultURL)
	sec.Key(configDeliveryBranchPrefixField).SetValue(InternalDefaultDeliveryBranchPrefix)

	err := cfg.SaveTo(local + string(pathSeparator) + configFileName)
	if err != nil {
		return err
	}
	fmt.Println("Bootstrapping new .gitconfig file")
	return nil
}

//ValidateLoadedSetting ...
func ValidateLoadedSetting(setting *ProjectSetting) error {
	r := reflect.ValueOf(setting).Elem()
	t := r.Type()

	errMsg := ""
	//Non Optional Field checks..
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Name == "Service" && (r.Field(i).String() == "") {
			errMsg += NewConfigError(configServiceField, setting.INIBlock).Error()
		}

		if t.Field(i).Name == "IssueURL" && r.Field(i).String() == "" {
			errMsg += NewConfigError(configIssueURLField, setting.INIBlock).Error()
		}

		if t.Field(i).Name == "IntegrationBranch" && r.Field(i).String() == "" {
			errMsg += NewConfigError(configIntegrationBranchField, setting.INIBlock).Error()
		}

		if t.Field(i).Name == "Remote" && r.Field(i).String() == "" {
			errMsg += NewConfigError(configRemoteField, setting.INIBlock).Error()
		}

		if t.Field(i).Name == "DeliveryBranchPrefix" && r.Field(i).String() == "" {
			errMsg += NewConfigError(configDeliveryBranchPrefixField, setting.INIBlock).Error()
		}
	}
	if errMsg == "" {
		return nil
	}
	return errors.New(errMsg)
}

//GetDefaultBranchFromInternalDefault ...
//Trying to retrieve the default branch from github
func GetDefaultBranchFromInternalDefault() (string, error) {
	git := githandler.Git{Run: executor.RunGit}

	remote, err := git.LSRemote("--get-url", InternalDefaultRemote)
	if err != nil {
		return "", err
	}
	orgAndRepo := githandler.OrgAndRepo(remote)
	token, err := git.Config("--get", "phlow.token")

	branch, err := plugins.DefaultBranchGitHub(InternalDefaultURL, orgAndRepo.Organisation, orgAndRepo.Repository, token)
	if err != nil {
		return "", err
	}
	return branch, nil
}
