load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/roronoadiogo/pick-game-wallpaper
gazelle(name = "gazelle")

go_library(
    name = "pick-game-wallpaper_lib",
    srcs = ["main.go"],
    importpath = "github.com/roronoadiogo/pick-game-wallpaper",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "pick-game-wallpaper",
    embed = [":pick-game-wallpaper_lib"],
    visibility = ["//visibility:public"],
)
