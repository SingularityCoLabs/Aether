#!/usr/bin/env bash
set -euo pipefail

SQLC_VERSION=1.31.1
AETHER_TOOLS_DIR=.tools/bin

platform="$(uname -s)"
architecture="$(uname -m)"

case "$platform/$architecture" in
  Darwin/arm64)
    printf -v asset "sqlc_%s_darwin_arm64.tar.gz" "$SQLC_VERSION"
    checksum="21602158c99eb1f2bae197a66abfb1941d1e9e50b23125bb193349c6b1acc71e"
    ;;
  Darwin/x86_64)
    printf -v asset "sqlc_%s_darwin_amd64.tar.gz" "$SQLC_VERSION"
    checksum="c5af76772e3785d21663a62697056b383f07629979b1bd25b93872e73dbd519b"
    ;;
  Linux/x86_64)
    printf -v asset "sqlc_%s_linux_amd64.tar.gz" "$SQLC_VERSION"
    checksum="497ae4fcdfa64c5b0c311ffe4c2bd991e43991e82e5367792ed78bc2dca27354"
    ;;
  Linux/aarch64 | Linux/arm64)
    printf -v asset "sqlc_%s_linux_arm64.tar.gz" "$SQLC_VERSION"
    checksum="b7cae247740d0c51a1e657479e5b2d21e6fef428f596682a01bc55bf4ab8a23d"
    ;;
  *)
    echo "unsupported sqlc platform: $platform/$architecture" >&2
    echo "install sqlc $SQLC_VERSION manually and place it on PATH" >&2
    exit 1
    ;;
esac

mkdir -p "$AETHER_TOOLS_DIR"

if test -x "$AETHER_TOOLS_DIR/sqlc"; then
  installed="$("$AETHER_TOOLS_DIR/sqlc" version | tr -d 'v')"
  if test "$installed" = "$SQLC_VERSION"; then
    exit 0
  fi
fi

AETHER_SQLC_TMP="$(mktemp -d /tmp/aether-sqlc.XXXXXX)"
trap 'rm -R -- "$AETHER_SQLC_TMP"' EXIT
archive="$AETHER_SQLC_TMP/$asset"
url="https://github.com/sqlc-dev/sqlc/releases/download/v$SQLC_VERSION/$asset"

curl --fail --location --silent --show-error "$url" --output "$archive"

if command -v shasum >/dev/null 2>&1; then
  actual="$(shasum -a 256 "$archive" | awk '{print $1}')"
else
  actual="$(sha256sum "$archive" | awk '{print $1}')"
fi

if test "$actual" != "$checksum"; then
  echo "sqlc checksum mismatch: expected $checksum, got $actual" >&2
  exit 1
fi

tar -xzf "$archive" -C "$AETHER_TOOLS_DIR" sqlc
chmod +x "$AETHER_TOOLS_DIR/sqlc"
"$AETHER_TOOLS_DIR/sqlc" version
