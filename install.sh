#!/bin/sh
set -e

REPO="JuD4Mo/api-kit"
BINARY="akit"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"
VERSION="${VERSION:-latest}"

detect_os() {
  os=$(uname -s | tr '[:upper:]' '[:lower:]')
  case "$os" in
    linux)  echo "Linux" ;;
    darwin) echo "Darwin" ;;
    *)      echo "Unsupported OS: $os"; exit 1 ;;
  esac
}

detect_arch() {
  arch=$(uname -m)
  case "$arch" in
    x86_64|amd64) echo "x86_64" ;;
    aarch64|arm64) echo "arm64" ;;
    *)             echo "Unsupported architecture: $arch"; exit 1 ;;
  esac
}

OS=$(detect_os)
ARCH=$(detect_arch)

if [ "$VERSION" = "latest" ]; then
  URL="https://github.com/$REPO/releases/latest/download/api-kit_${OS}_${ARCH}.tar.gz"
else
  URL="https://github.com/$REPO/releases/download/$VERSION/api-kit_${OS}_${ARCH}.tar.gz"
fi

echo "Downloading api-kit ($OS/$ARCH)..."
TMPDIR=$(mktemp -d)
curl -sSL "$URL" | tar -xz -C "$TMPDIR"

echo "Installing to $INSTALL_DIR..."
if [ ! -w "$INSTALL_DIR" ]; then
  sudo mv "$TMPDIR/$BINARY" "$INSTALL_DIR/$BINARY"
else
  mv "$TMPDIR/$BINARY" "$INSTALL_DIR/$BINARY"
fi
rm -rf "$TMPDIR"

echo "Installed! Run 'akit new' to get started."
