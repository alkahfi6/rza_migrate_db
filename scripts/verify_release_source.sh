#!/bin/sh
set -eu

# Recompute the same deterministic source digest used by build_all.bat.
# This runs in GitLab without Go or PowerShell and proves that the committed
# Linux checker was built from the exact source currently under review.

repo_root=${1:-.}
cd "$repo_root"

expected=$(sed -n 's/^source_sha256=//p' tools/linux-amd64/build-info.txt)
case "$expected" in
    ''|*[!0-9a-f]* )
        echo "[error] invalid or missing source_sha256 in build-info.txt" >&2
        exit 1
        ;;
esac
if [ "${#expected}" -ne 64 ]; then
    echo "[error] source_sha256 must contain exactly 64 hexadecimal characters" >&2
    exit 1
fi

file_list=$(mktemp)
hash_list=$(mktemp)
trap 'rm -f "$file_list" "$hash_list"' EXIT HUP INT TERM

find cmd/rza_migrate_db cmd/rza_migrate_check cmd/rza_migrate_report internal -type f -name '*.go' -print > "$file_list"
for path in \
    go.mod \
    go.sum \
    build_all.bat \
    scripts/release_artifacts.ps1 \
    scripts/verify_release_source.sh \
    .gitlab/rza_migrate_check.gitlab-ci.yml
do
    if [ ! -f "$path" ]; then
        echo "[error] release source input is missing: $path" >&2
        exit 1
    fi
    printf '%s\n' "$path" >> "$file_list"
done

LC_ALL=C sort -u "$file_list" -o "$file_list"
while IFS= read -r path; do
    # Ignore Git's checkout-only CRLF/LF conversion, matching the Windows
    # release builder's normalized source hashing.
    digest=$(sed 's/\r$//' "$path" | sha256sum | awk '{print $1}')
    printf '%s  %s\n' "$digest" "$path"
done < "$file_list" > "$hash_list"

actual=$(sha256sum "$hash_list" | awk '{print $1}')
if [ "$actual" != "$expected" ]; then
    echo "[error] release source digest mismatch" >&2
    echo "[error] expected=$expected" >&2
    echo "[error] actual=$actual" >&2
    echo "[hint] run build_all.bat and commit the complete refreshed release" >&2
    exit 1
fi

echo "[ok] release source digest verified: $actual"
