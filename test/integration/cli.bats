#!/usr/bin/env bats

load helpers

@test "cli: show info" {
    run certm
    [ "$status" -eq 0  ]
    [[ ${lines[0]} =~ "NAME:"  ]]
}

@test "cli: version" {
    run certm -v
    [ "$status" -eq 0  ]
    [[ ${lines[0]} =~ "version"  ]]
}
