#!/bin/sh

die()
{
    echo "$@" >&2
    exit 1
}
while read ver; do
    if [ "${ver:0:1}" == "#" ]; then
        continue
    fi

    go get github.com/openacid/slim/trie@$ver || die go-get-$ver

    vnum=${ver#v}
    case $vnum in
        0.5.0|0.5.1|0.5.2|0.5.3|0.5.4|0.5.5|0.5.6|0.5.7|0.5.8|0.5.9)
            echo 0.5.0-9
            go run makedata.go -ver $vnum
            ;;
        *)
            # since 0.5.10, there is an Opt for NewSlimTrie
            echo '>=0.5.10'
            go run make510.go -ver $vnum
            ;;
    esac
done <<-END
# v0.5.0
# v0.5.1
# v0.5.2
# v0.5.3
# v0.5.4
# v0.5.5
# v0.5.6
# v0.5.7
# v0.5.8
# v0.5.9
v0.5.10
END
