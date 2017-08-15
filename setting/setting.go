package setting

import (
	"os"
	"strings"
	"github.com/praqma/git-phlow/executor"
	"github.com/go-ini/ini"
	"fmt"
	"reflect"
	"github.com/go-errors/errors"
)

//Load internals
const (
	phlow_file_name     = ".phlow"
	git_config_fileName = ".gitconfig"
	default_block       = "default"
)

//Default configuration
const (
	internal_default_service                = "github"
	internal_default_integration_branch     = "master"
	internal_default_remote                 = "origin"
	internal_default_issue_url              = "https://api.github.com"
	internal_default_delivery_branch_prefix = "ready"
	internal_default_scope                  = "internal"
	internal_default_file                   = "none"
	internal_pipeline_url                   = "none"
)

//Uses git config commandline interface
//ToolsSetting ...
type ToolsSetting struct {
	User    string
	Token   string
	AutoAdd bool
}

//ProjectSetting ...
type ProjectSetting struct {
	Service              string `ini:"service"`
	IntegrationBranch    string `ini:"integration_branch"`
	Remote               string `ini:"remote"`
	IssueURL             string `ini:"issue_url"`
	PipelineUrl          string `ini:"pipeline"`
	DeliveryBranchPrefix string `ini:"delivery_branch_prefix"`
	Scope                string
	File                 string
	INIBlock             string
}

//NewProjectStg ...
//initializes project settings from .phlow INI files
func NewProjectStg(INIBlock string) *ProjectSetting {
	r := GetLocal()
	g := GetGlobal()
	return LoadProjectSettings(r, g, INIBlock)
}

//NewToolStg ...
//Initializes Tool settings from .gitconfig
func NewToolStg() *ToolsSetting {
	return LoadToolSettings(executor.ExecuteCommand)
}

//LoadProjectSettings ...
func LoadProjectSettings(local, global string, INIBlock string) *ProjectSetting {

	supportedScopes := []string{local, global}
	supportedConfigFiles := []string{phlow_file_name, git_config_fileName}

	//If params are not set, we use the default config
	if INIBlock == "" {
		INIBlock = default_block
	}

	var configScope, configFile string

	var loadSetting = func() *ini.Section {
		for _, scope := range supportedScopes {
			for _, file := range supportedConfigFiles {

				config, err := ini.LooseLoad(scope + "/" + file)
				if err != nil {
					panic(err)
				}

				if sec, _ := config.GetSection(INIBlock); sec != nil {
					configScope = scope
					configFile = file
					return config.Section(INIBlock)
				}
			}
		}
		return nil
	}

	loadedSetting := loadSetting()
	if loadedSetting == nil {
		if INIBlock != default_block {
			fmt.Printf("Error: '%s' configuration does not exist in you configuration files. Following paths were searched: \n", INIBlock)
			for _, scope := range supportedScopes {
				for _, file := range supportedConfigFiles {
					fmt.Println("paths: " + scope + "/" + file)
				}
			}
			os.Exit(1)
		}
		//return internal default because no other configuration exist and no other is specified by params
		return &ProjectSetting{
			Service:              internal_default_service,
			IntegrationBranch:    internal_default_integration_branch,
			Remote:               internal_default_remote,
			IssueURL:             internal_default_issue_url,
			DeliveryBranchPrefix: internal_default_delivery_branch_prefix,
			PipelineUrl:          internal_pipeline_url,
			Scope:                internal_default_scope,
			File:                 internal_default_file,
		}
	}

	//Map section into object
	conf := new(ProjectSetting)
	loadSetting().MapTo(conf)

	//Add configuration origin
	conf.File = configFile
	conf.Scope = configScope
	conf.INIBlock = INIBlock

	if err := ValidateLoadedSetting(conf); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return conf
}

//ValidateLoadedSetting ...
func ValidateLoadedSetting(setting *ProjectSetting) (error) {
	r := reflect.ValueOf(setting).Elem()
	t := r.Type()

	//Non Optional Field checks..
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Name == "Service" && (r.Field(i).String() == "") {
			return errors.New(fmt.Sprintf("Error in configuration file: %s \n"+
				"Non-optional field missing: %s \n In configuration block: %s \n ", setting.Scope+"/"+setting.File, "service", setting.INIBlock))
		}

		if t.Field(i).Name == "IssueURL" && r.Field(i).String() == "" {
			return errors.New(fmt.Sprintf("Error in configuration file: %s \n"+
				"Non-optional field missing: %s \n In configuration block: %s \n ", setting.Scope+"/"+setting.File, "issue_url", setting.INIBlock))
		}

		if t.Field(i).Name == "IntegrationBranch" && r.Field(i).String() == "" {
			return errors.New(fmt.Sprintf("Error in configuration file: %s \n"+
				"Non-optional field missing: %s \n In configuration block: %s \n ", setting.Scope+"/"+setting.File, "integration_branch", setting.INIBlock))
		}

		if t.Field(i).Name == "Remote" && r.Field(i).String() == "" {
			return errors.New(fmt.Sprintf("Error in configuration file: %s \n"+
				"Non-optional field missing: %s \n In configuration block: %s \n ", setting.Scope+"/"+setting.File, "remote", setting.INIBlock))
		}

		if t.Field(i).Name == "DeliveryBranchPrefix" && r.Field(i).String() == "" {
			return errors.New(fmt.Sprintf("Error in configuration file: %s \n"+
				"Non-optional field missing: %s \n In configuration block: %s \n ", setting.Scope+"/"+setting.File, "delivery_branch_prefix", setting.INIBlock))
		}
	}
	return nil
}

//LoadToolSettings ...
func LoadToolSettings(run executor.Runner) *ToolsSetting {
	var set = ToolsSetting{}

	user, err := run("git", "config", "--get", "phlow.user")
	if err != nil {
		panic(err)
	}
	set.User = strings.Replace(user, "\n", "", -1)

	token, err := run("git", "config", "--global", "phlow.token")
	if err != nil {
		panic(err)
	}
	set.Token = strings.Replace(token, "\n", "", -1)

	return &set
}

//GetGlobal ...
func GetGlobal() string {
	return os.Getenv("HOME")
}

//GetLocal ...
func GetLocal() string {
	absoluteRepoPath, err := executor.ExecuteCommand("git", "rev-parse", "--show-toplevel")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(absoluteRepoPath)
}
