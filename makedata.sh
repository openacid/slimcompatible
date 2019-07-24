#!/bin/sh

die()
{
    echo "$@" >&2
    exit 1
}
while read ver; do
    (
    cd ../slim || die cd to slim
    git checkout $ver || die git checkout $ver

    git log -n1 --color --graph --decorate -M --pretty=oneline --abbrev-commit -M
    )
    vnum=${ver#v}
    go run makedata.go -ver $vnum
done <<-END
v0.5.0
v0.5.1
v0.5.2
v0.5.3
v0.5.4
v0.5.5
v0.5.6
v0.5.7
v0.5.8
v0.5.9
END
