DIR="$(go list -f '{{.Dir}}' modernc.org/b)"
redo-ifchange "$DIR/btree.go"
make --silent -C "$DIR" generic \
| sed 's/KEY/bench.Key/g; s/VALUE/bench.Value/g; s|^\(package .*\)$|\1\n\nimport bench "modernc.org/benchmark-ordered-map/fixture"|'
