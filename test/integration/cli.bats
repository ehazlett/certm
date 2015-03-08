#!/usr/bin/env bats

load helpers

@test "cli: show info" {
    run certtool
    [ "$status" -eq 1  ]
    [[ ${lines[0]} =~ "NAME:"  ]]
}

@test "cli: version" {
    run certtool -v
    [ "$status" -eq 0  ]
    [[ ${lines[0]} =~ "version"  ]]
}
