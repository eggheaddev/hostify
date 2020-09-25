#!/bin/sh
# ? this installation script is based on the deno installation scripts
# Copyright 2019 the Deno authors. All rights reserved. MIT license.

set -e

if [ "$(uname -m)" != "x86_64" ]; then
	echo "Error: Unsupported architecture $(uname -m). Only x64 binaries are available." 1>&2
	exit 1
fi

if ! command -v unzip >/dev/null; then
	echo "Error: unzip is required to install Hostify" 1>&2
	exit 1
fi

if [ "$OS" = "Windows_NT" ]; then
	target="windows"
else
	case $(uname -s) in
	Darwin) target="linux" ;;
	*) target="linux" ;;
	esac
fi

if [ $# -eq 0 ]; then
	hostify_asset_path=$(
		curl -sSf https://github.com/eggheaddev/hostify-cli/releases |
			grep -o "/eggheaddev/hostify-cli/releases/download/.*/xhostify-${target}\\.zip" |
			head -n 1
	)
	if [ ! "$hostify_asset_path" ]; then
		echo "Error: Unable to find latest hostify release on GitHub." 1>&2
		exit 1
	fi
	hostify_uri="https://github.com${hostify_asset_path}"
else
	hostify_uri="https://github.com/eggheaddev/hostify-cli/releases/download/${1}/xhostify-${target}.zip"
fi

hostify_install="${HOSTIFY_INSTALL:-$HOME/.hostify}"
bin_dir="$hostify_install/bin"
exe="$bin_dir/hostify"

if [ ! -d "$bin_dir" ]; then
	mkdir -p "$bin_dir"
fi

curl --fail --location --progress-bar --output "$exe.zip" "$hostify_uri"
unzip -d "$bin_dir" -o "$exe.zip"
chmod +x "$exe"
rm "$exe.zip"

echo "Hostify was installed successfully to $exe"

case $SHELL in
/bin/zsh) shell_profile=".zshrc" ;;
*) shell_profile=".bash_profile" ;;
esac
echo "Manually add the directory to your \$HOME/$shell_profile (or similar)"
echo "  export HOSTIFY_INSTALL=\"$hostify_install\""
echo "  export PATH=\"\$HOSTIFY_INSTALL/bin:\$PATH\""
