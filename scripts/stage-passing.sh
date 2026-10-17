#!/bin/sh
# Stage the model outputs of every unit that passes `tr check`, and nothing
# of a unit that fails (an agent may still be writing it). Run from the
# repository root after building bin/tr.
set -e
for w in works/*/; do
	work=$(basename "$w")
	for f in "$w"stage/morph/*.json "$w"stage/translate/*.json "$w"stage/review/*.json "$w"stage/alt/*/translate/*.json "$w"stage/alt/*/review/*.json; do
		[ -e "$f" ] || continue
		unit=$(basename "$f" .json)
		alt=""
		case "$f" in */stage/alt/*) alt="--alt $(echo "$f" | sed 's|.*/stage/alt/\([^/]*\)/.*|\1|')" ;; esac
		if ./bin/tr check "$work" "$unit" $alt >/dev/null 2>&1; then
			git add "$f"
		else
			git reset -q -- "$f" 2>/dev/null || true
		fi
	done
done
