load("//tools:build_rules/shims.bzl", "go_library", "go_test")

package(default_visibility = ["//kythe:default_visibility"])

go_library(
    name = "reduce",
    srcs = [
        "reduce.go",
        ":reduce_sort.go",
    ],
    deps = [
        "//kythe/go/util/disksort",
        "//kythe/proto:internal_go_proto",
        "@com_github_golang_protobuf//proto:go_default_library",
    ],
)

go_test(
    name = "reduce_test",
    size = "small",
    srcs = ["reduce_test.go"],
    library = "reduce",
    visibility = ["//visibility:private"],
    deps = ["//kythe/go/test/testutil"],
)
