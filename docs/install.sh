#!/usr/bin/env bash
set -e

REPO="coolapso/tfimport"
BINARY_NAME="tfimport"
INSTALL_DIR="/usr/local/bin"

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$ARCH" in
  x86_64)  ARCH="amd64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  i386|i686) ARCH="386" ;;
  *)
    echo "Unsupported architecture: $ARCH"
    exit 1
    ;;
esac

case "$OS" in
  linux|darwin) EXT="tar.gz" ;;
  mingw*|msys*|cygwin*|windows*) OS="windows"; EXT="zip" ;;
  *)
    echo "Unsupported OS: $OS"
    exit 1
    ;;
esac

if [ -n "$VERSION" ]; then
  LATEST="$VERSION"
else
  LATEST=$(curl -s "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name"' | sed -E 's/.*"tag_name": *"([^"]+)".*/\1/')
fi

if [ -z "$LATEST" ]; then
  echo "Failed to fetch latest release version."
  exit 1
fi

VERSION_NO_V="${LATEST#v}"
FILENAME="${BINARY_NAME}_${VERSION_NO_V}_${OS}_${ARCH}.${EXT}"
DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${LATEST}/${FILENAME}"

TMP_DIR=$(mktemp -d)
trap "rm -rf $TMP_DIR" EXIT

echo "Downloading ${BINARY_NAME} ${LATEST} for ${OS}/${ARCH}..."
curl -sL "$DOWNLOAD_URL" -o "${TMP_DIR}/${FILENAME}"

if [ "$EXT" = "tar.gz" ]; then
  tar -xzf "${TMP_DIR}/${FILENAME}" -C "$TMP_DIR"
else
  unzip -q "${TMP_DIR}/${FILENAME}" -d "$TMP_DIR"
fi

echo "Installing to ${INSTALL_DIR}/${BINARY_NAME}..."
install -m 755 "${TMP_DIR}/${BINARY_NAME}" "${INSTALL_DIR}/${BINARY_NAME}"

echo ""
echo "✅ ${BINARY_NAME} ${LATEST} installed successfully!"
echo "Run '${BINARY_NAME} configure' to get started." 
