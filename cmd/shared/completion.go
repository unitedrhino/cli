package shared

import (
	"fmt"
	"io"
	"strings"
)

func runCompletion(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "usage: ur completion bash|zsh|fish")
		return 2
	}
	shell := args[0]
	switch shell {
	case "bash":
		fmt.Fprintln(stdout, bashCompletionScript())
	case "zsh":
		fmt.Fprintln(stdout, zshCompletionScript())
	case "fish":
		fmt.Fprintln(stdout, fishCompletionScript())
	default:
		fmt.Fprintf(stderr, "unsupported shell %q, valid: bash, zsh, fish\n", shell)
		return 2
	}
	return 0
}

func bashCompletionScript() string {
	var sb strings.Builder
	sb.WriteString(`_ur_cli_completion() {
    local cur prev opts
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"

    local commands="api check config generate-skills help login model scene schema script setup token"
    local model_cmds="template validate generate-script"
    local scene_cmds="validate template"
    local script_cmds="validate template"
    local api_opts="--body --body-file --header -H --fields --summarize --format --transform --debug"
    local config_opts="--list --use"
    local model_template_opts="--json --yaml --output"
    local model_script_opts="--mode --output"
    local global_opts="--app"

    # 全局选项
    if [[ "$cur" == -* ]]; then
        COMPREPLY=( $(compgen -W "$global_opts" -- "$cur") )
        return 0
    fi

    # 第一个参数：命令
    if [ $COMP_CWORD -eq 1 ]; then
        COMPREPLY=( $(compgen -W "$commands" -- "$cur") )
        return 0
    fi

    local cmd="${COMP_WORDS[1]}"

    case "$cmd" in
        api)
            if [[ "$prev" == "--format" ]]; then
                COMPREPLY=( $(compgen -W "json raw yaml" -- "$cur") )
            elif [[ "$prev" == "--body" || "$prev" == "--body-file" || "$prev" == "--transform" ]]; then
                return 0
            else
                COMPREPLY=( $(compgen -W "$api_opts" -- "$cur") )
            fi
            ;;
        config)
            if [[ "$prev" == "--use" ]]; then
                return 0
            else
                COMPREPLY=( $(compgen -W "$config_opts" -- "$cur") )
            fi
            ;;
        model)
            if [ $COMP_CWORD -eq 2 ]; then
                COMPREPLY=( $(compgen -W "$model_cmds" -- "$cur") )
            else
                local sub="${COMP_WORDS[2]}"
                case "$sub" in
                    template)
                        if [[ "$prev" == "--output" ]]; then
                            return 0
                        else
                            COMPREPLY=( $(compgen -W "property event action full $model_template_opts" -- "$cur") )
                        fi
                        ;;
                    generate-script)
                        if [[ "$prev" == "--output" || "$prev" == "--mode" ]]; then
                            return 0
                        else
                            COMPREPLY=( $(compgen -W "$model_script_opts" -- "$cur") )
                        fi
                        ;;
                esac
            fi
            ;;
        scene)
            if [ $COMP_CWORD -eq 2 ]; then
                COMPREPLY=( $(compgen -W "$scene_cmds" -- "$cur") )
            fi
            ;;
        script)
            if [ $COMP_CWORD -eq 2 ]; then
                COMPREPLY=( $(compgen -W "$script_cmds" -- "$cur") )
            fi
            ;;
        check|help|login|schema|setup|token|generate-skills)
            return 0
            ;;
    esac
}

complete -F _ur_cli_completion ur
`)
	return sb.String()
}

func zshCompletionScript() string {
	var sb strings.Builder
	sb.WriteString(`#compdef ur

_ur() {
    local curcontext="$curcontext" state line
    typeset -A opt_args

    _arguments -C \
        '(--app)--app[Specify app context]:app:(iot platform-manage org-manage org-energy console)' \
        '1: :->command' \
        '*:: :->args'

    case "$state" in
        command)
            _values 'commands' \
                'api[Call API endpoint]' \
                'check[Check auth status]' \
                'config[Manage profiles]' \
                'generate-skills[Generate skill docs]' \
                'help[Show help]' \
                'login[Authenticate]' \
                'model[Thing model tools]' \
                'scene[Scene linkage tools]' \
                'schema[Show API schema]' \
                'script[Protocol script tools]' \
                'setup[Initial setup]' \
                'token[Show/manage token]' \
            ;;
        args)
            case "$line[1]" in
                api)
                    _arguments \
                        '--body[JSON body]:json:' \
                        '--body-file[Body file]:file:_files' \
                        '(-H --header)'{-H,--header}'[Custom header]:header:' \
                        '--fields[Field selectors]:selectors:' \
                        '--summarize[Summarize response]' \
                        '--format[Output format]:format:(json raw yaml)' \
                        '--transform[GJSON transform path]:path:' \
                        '--debug[Enable debug logging]'
                    ;;
                config)
                    _arguments \
                        '--list[List profiles]' \
                        '--use[Switch profile]:profile:'
                    ;;
                model)
                    _arguments '1: :->model_cmd' '*:: :->model_args'
                    case "$line[1]" in
                        template)
                            _arguments \
                                '--json[Output JSON]' \
                                '--yaml[Output YAML]' \
                                '--output[Output file]:file:_files'
                            ;;
                        generate-script)
                            _arguments \
                                '--mode[Script mode]:mode:(up-before up-after down-before down-after)' \
                                '--output[Output file]:file:_files'
                            ;;
                    esac
                    ;;
            esac
            ;;
    esac
}

compdef _ur ur
`)
	return sb.String()
}

func fishCompletionScript() string {
	var sb strings.Builder
	sb.WriteString(`# Fish completion for ur CLI

# Global options
complete -c ur -l app -d "Specify app context" -a "iot platform-manage org-manage org-energy console"

# Commands
complete -c ur -n "__fish_use_subcommand" -a "api" -d "Call API endpoint"
complete -c ur -n "__fish_use_subcommand" -a "check" -d "Check auth status"
complete -c ur -n "__fish_use_subcommand" -a "config" -d "Manage profiles"
complete -c ur -n "__fish_use_subcommand" -a "generate-skills" -d "Generate skill docs"
complete -c ur -n "__fish_use_subcommand" -a "help" -d "Show help"
complete -c ur -n "__fish_use_subcommand" -a "login" -d "Authenticate"
complete -c ur -n "__fish_use_subcommand" -a "model" -d "Thing model tools"
complete -c ur -n "__fish_use_subcommand" -a "scene" -d "Scene linkage tools"
complete -c ur -n "__fish_use_subcommand" -a "schema" -d "Show API schema"
complete -c ur -n "__fish_use_subcommand" -a "script" -d "Protocol script tools"
complete -c ur -n "__fish_use_subcommand" -a "setup" -d "Initial setup"
complete -c ur -n "__fish_use_subcommand" -a "token" -d "Show/manage token"

# api subcommand options
complete -c ur -n "__fish_seen_subcommand_from api" -l body -d "JSON body"
complete -c ur -n "__fish_seen_subcommand_from api" -l body-file -d "Body file" -r
complete -c ur -n "__fish_seen_subcommand_from api" -s H -l header -d "Custom header"
complete -c ur -n "__fish_seen_subcommand_from api" -l fields -d "Field selectors"
complete -c ur -n "__fish_seen_subcommand_from api" -l summarize -d "Summarize response"
complete -c ur -n "__fish_seen_subcommand_from api" -l format -d "Output format" -a "json raw yaml"
complete -c ur -n "__fish_seen_subcommand_from api" -l transform -d "GJSON transform path"
complete -c ur -n "__fish_seen_subcommand_from api" -l debug -d "Enable debug logging"

# config subcommand options
complete -c ur -n "__fish_seen_subcommand_from config" -l list -d "List profiles"
complete -c ur -n "__fish_seen_subcommand_from config" -l use -d "Switch profile"

# model subcommands
complete -c ur -n "__fish_seen_subcommand_from model; and not __fish_seen_subcommand_from template validate generate-script" -a "template validate generate-script"

# model template options
complete -c ur -n "__fish_seen_subcommand_from model; and __fish_seen_subcommand_from template" -a "property event action full"
complete -c ur -n "__fish_seen_subcommand_from model; and __fish_seen_subcommand_from template" -l json -d "Output JSON"
complete -c ur -n "__fish_seen_subcommand_from model; and __fish_seen_subcommand_from template" -l yaml -d "Output YAML"
complete -c ur -n "__fish_seen_subcommand_from model; and __fish_seen_subcommand_from template" -l output -d "Output file" -r

# model generate-script options
complete -c ur -n "__fish_seen_subcommand_from model; and __fish_seen_subcommand_from generate-script" -l mode -d "Script mode" -a "up-before up-after down-before down-after"
complete -c ur -n "__fish_seen_subcommand_from model; and __fish_seen_subcommand_from generate-script" -l output -d "Output file" -r

# scene subcommands
complete -c ur -n "__fish_seen_subcommand_from scene; and not __fish_seen_subcommand_from validate template" -a "validate template"

# script subcommands
complete -c ur -n "__fish_seen_subcommand_from script; and not __fish_seen_subcommand_from validate template" -a "validate template"
`)
	return sb.String()
}
