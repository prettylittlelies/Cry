#!/usr/bin/env bash
set -euo pipefail
: "${MLSPP:?path to a cisco/mlspp checkout}"
: "${LIBDAVE:?path to a discord/libdave checkout}"
: "${NLOHMANN:?directory containing nlohmann/json.hpp}"
OUT="${OUT:-$(pwd)/daveinterop-build}"
HERE="$(cd "$(dirname "$0")" && pwd)"
D="$LIBDAVE/cpp"
mkdir -p "$OUT/obj" "$OUT/include"
printf '#pragma once\n#define MLS_NAMESPACE mlspp\n' > "$OUT/include/namespace.h"
FLAGS="-std=c++17 -O1 -w -DWITH_OPENSSL3 -DDISABLE_GREASE -DNDEBUG -I$OUT/include -I$NLOHMANN -I$MLSPP/include -I$MLSPP/lib/bytes/include -I$MLSPP/lib/tls_syntax/include -I$MLSPP/lib/hpke/include -I$MLSPP/third_party -I$D/includes -I$D/src"
SOURCES="$(ls "$MLSPP"/lib/bytes/src/*.cpp "$MLSPP"/lib/tls_syntax/src/*.cpp "$MLSPP"/src/*.cpp)
$(ls "$MLSPP"/lib/hpke/src/*.cpp | grep -v -e mlkem -e hybrid_kem)
$(ls "$D"/src/*.cpp | grep -v -e bindings_ -e boringssl_cryptor)
$(ls "$D"/src/mls/*.cpp | grep -v 'persisted_key_pair.cpp')
$(ls "$D"/src/utils/*.cpp)"
for src in $SOURCES; do
  obj="$OUT/obj/$(echo "$src" | tr '/.:' '___').o"
  [ -f "$obj" ] || g++ $FLAGS -c "$src" -o "$obj" &
done
wait
g++ $FLAGS "$HERE/harness.cpp" "$OUT"/obj/*.o -o "$OUT/harness" -lssl -lcrypto ${EXTRA_LIBS:-} -static-libgcc -static-libstdc++
echo "$OUT/harness"
