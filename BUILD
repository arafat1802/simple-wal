load("@gazelle//:def.bzl", "gazelle")
load("@rules_go//go:def.bzl", "go_binary", "go_library")

gazelle(name = "gazelle")

go_library(
    name = "simple-wal_lib",
    srcs = ["main.go"],
    importpath = "simple-wal",
    visibility = ["//visibility:private"],
    deps = ["//utils"],
)

go_binary(
    name = "simple-wal",
    embed = [":simple-wal_lib"],
    visibility = ["//visibility:public"],
)
