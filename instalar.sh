#!/bin/sh

set -euo pipefail

aplicar_estilo() {
    local cor="$1"
    local negrito="$2"
    local texto="$3"

    local estilo=""
    local remover_estilo="\033[0m"

    # Verifica se o terminal suporta true colors
    if [[ -t 1 && $COLORTERM == "truecolor" ]]; then
        case "$cor" in
        "vermelho") cor='\e[38;2;255;0;0m' ;;
        "verde") cor='\e[38;2;0;255;0m' ;;
        "amarelo") cor='\e[38;2;255;255;0m' ;;
        "branco") cor='\e[38;2;255;255;255m' ;;
        *) cor="" ;;
        esac
    else
        case "$cor" in
        "vermelho") cor='\033[0;31m' ;;
        "verde") cor='\033[0;32m' ;;
        "amarelo") cor='\033[0;33m' ;;
        "branco") cor='\033[0;2m' ;;
        *) cor="" ;;
        esac
    fi

    # Verifica se deve adicionar o negrito
    if [[ "$negrito" == "true" ]]; then
        estilo='\033[1m'
    fi

    # Imprime o texto com o estilo aplicado
    printf "%b" "${estilo}${cor}${texto}${remover_estilo}"
}

# Exemplos de uso
# log "INFO" "Esta é uma informação destacada."
# log "SUCESSO" "Operação concluída com sucesso."
# log "DEBUG" "Esta é uma mensagem de depuração."
# log "AVISO" "Mensagem de aviso"
# log "ERRO" "Ocorreu um erro!"
log() {
    local nivel="$1"
    local mensagem="$2"

    local cor=""
    local negrito="true"

    case "$nivel" in
    "ERRO") cor="vermelho" ;;
    "INFO") cor="branco" ;;
    "AVISO") cor="amarelo" ;;
    "SUCESSO") cor="verde" ;;
    *) cor="branco"; negrito="false" ;;
    esac

    echo -n "[ "
    aplicar_estilo "$cor" "$negrito" "$nivel"
    echo -n " ]: "
    aplicar_estilo "$cor" "false" "$mensagem"
    echo

    # Verifica se é erro e dá uma saída
    if [[ "$nivel" == "ERRO" ]]; then
        exit 1
    fi
}

if [[ $# -gt 1 ]]; then
    log "ERRO" "Foi recebido mais argumentos do que o esperado. Caso deseje instalar uma versão específica, use por exemplo: v0.1.0."
fi

if ! command -v curl >/dev/null; then
    log "ERRO" "O comando curl é essencial para o processo, mas ele não pôde ser achado."
fi

if [[ ${OS:-} = Windows_NT ]]; then
    target="Windows_x86_64"
    sufixo=".zip"

    if ! command -v unzip >/dev/null || ! command -v 7z >/dev/null; then
        log "ERRO" "Ao menos um dos seguinte programas é necessário para fazer a instalação: unzip, 7z"
    fi
else
    case $(uname -sm) in
    "Darwin x86_64") target="Darwin_x86_64" ;;
    "Darwin arm64") target="Darwin_arm64" ;;
    "Linux aarch64") target="Linux_arm64" ;;
    *) target="Linux_x86_64" ;;
    esac

    sufixo=".tar.gz"

    if ! command -v tar >/dev/null; then
        log "ERRO" "O comando tar é essencial para o processo, mas ele não pôde ser achado."
    fi
fi

GITHUB=${GITHUB-"https://github.com"}
repo_github="$GITHUB/natanfeitosa/portuscript"

arquivo_compactado="$target$sufixo"

if [[ $# = 0 ]]; then
    portuscript_uri=$repo_github/releases/latest/download/$arquivo_compactado
else
    portuscript_uri=$repo_github/releases/download/$1/$arquivo_compactado
fi

raiz_portuscript="${RAIZ_PORTUSCRIPT:-$HOME/.portuscript}"
diretorio_binario="$raiz_portuscript/bin"
# diretorio_binario="./bin"
executavel="$diretorio_binario/portuscript"

if [[ ! -d $diretorio_binario ]]; then
    mkdir -p "$diretorio_binario" ||
        error "ERRO" "Falha ao criar o diretório de instalação \"$diretorio_binario\""
fi

log "DEBUG" "Iniciando download do arquivo compactado"
curl --fail --location --progress-bar --output "$executavel$sufixo" "$portuscript_uri" ||
    log "ERRO" "Falha ao baixar o Portuscript de \"$portuscript_uri\""

log "DEBUG" "Iniciando descompactação"
case "$sufixo" in
".zip")
    if command -v unzip >/dev/null; then
        unzip -d "$diretorio_binario" -o "$executavel$sufixo"
    else
        7z x -o "$diretorio_binario" -y "$executavel$sufixo"
    fi
    ;;
*)
    tar -xf "$executavel$sufixo" -C "$diretorio_binario"
    ;;
esac
log "SUCESSO" "Parece que a descompactação foi um sucesso"

chmod +x "$executavel$sufixo"
rm "$executavel$sufixo"

log "SUCESSO" "Parabéns, agora você tem o Portuscript disponível em \033[1m$executavel\033[0m"

refresh_command=""
if command -v portuscript >/dev/null; then
	log "INFO" "agora você pode usar o comando 'portuscript --help' para ter um guia de comandos"
else
	case $(basename "$SHELL") in
    fish)
        commands=(
            "set --export DIRETORIO_PORTUSCRIPT $raiz_portuscript"
            "set --export PATH \$DIRETORIO_PORTUSCRIPT/bin \$PATH"
        )

        fish_config=$HOME/.config/fish/config.fish

        if [[ -w $fish_config ]]; then
            {
                echo -e '\n# configuraçõs portuscript'

                for command in "${commands[@]}"; do
                    echo "$command"
                done
            } >>"$fish_config"

            log "INFO" "Adicionado o caminho \"$diretorio_binario\" ao \$PATH em \"$fish_config\""

            refresh_command="source $fish_config"
        else
            log "AVISO" "Adicione manualmente os seguinte comandos ao $fish_config (ou similar):"

            for command in "${commands[@]}"; do
                log "INFO" "  $command"
            done
        fi
        ;;
    zsh)
        commands=(
            "export DIRETORIO_PORTUSCRIPT=$raiz_portuscript"
            "export PATH=\"\$DIRETORIO_PORTUSCRIPT/bin:\$PATH\""
        )

        zsh_config=$HOME/.zshrc

        if [[ -w $zsh_config ]]; then
            {
                echo -e '\n# configuraçõs portuscript'

                for command in "${commands[@]}"; do
                    echo "$command"
                done
            } >>"$zsh_config"

            log "INFO" "Adicionado o caminho \"$diretorio_binario\" ao \$PATH em \"$zsh_config\""

            refresh_command="exec $SHELL"
        else
            log "AVISO" "Adicione manualmente os seguinte comandos ao $zsh_config (ou similar):"

            for command in "${commands[@]}"; do
                log "INFO" "  $command"
            done
        fi
        ;;
    bash)
        commands=(
            "export DIRETORIO_PORTUSCRIPT=$raiz_portuscript"
            "export PATH=\$DIRETORIO_PORTUSCRIPT/bin:\$PATH"
        )

        bash_configs=(
            "$HOME/.bashrc"
            "$HOME/.bash_profile"
        )

        if [[ ${XDG_CONFIG_HOME:-} ]]; then
            bash_configs+=(
                "$XDG_CONFIG_HOME/.bash_profile"
                "$XDG_CONFIG_HOME/.bashrc"
                "$XDG_CONFIG_HOME/bash_profile"
                "$XDG_CONFIG_HOME/bashrc"
            )
        fi

        set_manually=true
        for bash_config in "${bash_configs[@]}"; do

            if [[ -w $bash_config ]]; then
                {
                    echo -e '\n# configuraçõs portuscript'

                    for command in "${commands[@]}"; do
                        echo "$command"
                    done
                } >>"$bash_config"

                log "INFO" "Adicionado o caminho \"$diretorio_binario\" ao \$PATH em \"$bash_config\""

                refresh_command="source $bash_config"
                set_manually=false
                break
            fi
        done

        if [[ $set_manually = true ]]; then
            log "AVISO" "Adicione manualmente os seguinte comandos ao $bash_config (ou similar):"

            for command in "${commands[@]}"; do
                log "INFO" "  $command"
            done
        fi
        ;;
    *)
        log "AVISO" 'Adicione manualmente os seguinte comandos ao ~/.bashrc (ou similar):'
        log "INFO" "  export $install_env=$raiz_portuscript"
        log "INFO" "  export PATH=\"$bin_env:\$PATH\""
        ;;
    esac
fi


echo
log "INFO" "Para um bom início, execute:"
echo

if [[ $refresh_command ]]; then
    log "INFO" "  $refresh_command"
fi

log "INFO" "  portuscript --help"

echo

log "SUCESSO" "Finalmente chegamos ao fim"
log "INFO" "Considere também deixar uma estrelinha no nosso repositório $repo_github"
