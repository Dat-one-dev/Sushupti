#!/bin/sh

set -e

REPO="Dat-one-dev/Sushupti"
INSTALL_DIR="$HOME/.local/bin"

OS=$(uname -s)
ARCH=$(uname -m)

case "$OS" in
    Linux) OS="linux" ;;
    Darwin) OS="darwin" ;;
    *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

case "$ARCH" in
    x86_64) ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

VERSION=$(curl -s https://api.github.com/repos/$REPO/releases/latest | grep '"tag_name"' | cut -d '"' -f 4 | sed 's/^v//')

URL="https://github.com/$REPO/releases/download/v$VERSION/Sushupti_${VERSION}_${OS}_${ARCH}.tar.gz"

mkdir -p "$INSTALL_DIR"

echo "Downloading Sushupti $VERSION..."

curl -L "$URL" -o /tmp/sushupti.tar.gz

tar -xzf /tmp/sushupti.tar.gz -C /tmp

mv /tmp/Sushupti "$INSTALL_DIR/sushupti"

chmod +x "$INSTALL_DIR/sushupti"

echo ""
echo "Sushupti installed!"
echo "Run it with: sushupti"
