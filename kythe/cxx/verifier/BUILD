package(default_visibility = ["//kythe:default_visibility"])

load("//:visibility.bzl", "PUBLIC_VISIBILITY")
load("//tools:build_rules/lexyacc.bzl", "genlex", "genyacc")

genyacc(
    name = "parser",
    src = "assertions.yy",
    extra_outs = [
        "location.hh",
        "stack.hh",
        "position.hh",
    ],
    header_out = "parser.yy.hh",
    source_out = "parser.yy.cc",
)

genlex(
    name = "lexer",
    src = "assertions.lex",
    out = "lexer.yy.cc",
    includes = [":parser"],
)

cc_library(
    name = "lexparse",
    srcs = [
        ":lexer",
        ":parser",
    ],
    hdrs = [
        "assertion_ast.h",
        "assertions.h",
        "pretty_printer.h",
        ":parser",
    ],
    copts = [
        "-Wno-non-virtual-dtor",
        "-Wno-deprecated-register",
        "-Wno-implicit-fallthrough",
    ],
    deps = [
        "@com_github_google_glog//:glog",
        "@com_google_absl//absl/memory",
        "@com_googlesource_code_re2//:re2",
    ],
)

cc_library(
    name = "lib",
    srcs = [
        "assertions.cc",
        "parser.yy.hh",
        "pretty_printer.cc",
        "verifier.cc",
    ],
    hdrs = [
        "assertion_ast.h",
        "assertions.h",
        "pretty_printer.h",
        "verifier.h",
    ],
    copts = [
        "-Wno-non-virtual-dtor",
        "-Wno-unused-variable",
        "-Wno-implicit-fallthrough",
    ],
    deps = [
        ":lexparse",
        "//kythe/cxx/common:kythe_uri",
        "//kythe/cxx/common:scope_guard",
        "//kythe/proto:common_cc_proto",
        "//kythe/proto:storage_cc_proto",
        "@com_github_gflags_gflags//:gflags",
        "@com_github_google_glog//:glog",
        "@com_google_absl//absl/container:flat_hash_map",
        "@com_google_protobuf//:protobuf",
        "@com_googlesource_code_re2//:re2",
    ],
)

cc_library(
    name = "cmd_lib",
    srcs = [
        "verifier_main.cc",
    ],
    copts = [
        "-Wno-non-virtual-dtor",
        "-Wno-unused-variable",
        "-Wno-implicit-fallthrough",
    ],
    deps = [
        ":lib",
        "//kythe/proto:storage_cc_proto",
        "@com_github_gflags_gflags//:gflags",
        "@com_github_google_glog//:glog",
        "@com_google_protobuf//:protobuf",
    ],
)

cc_binary(
    name = "verifier",
    visibility = [PUBLIC_VISIBILITY],
    deps = [
        ":cmd_lib",
    ],
)

cc_library(
    name = "testlib",
    testonly = 1,
    srcs = [
        "verifier_unit_test.cc",
    ],
    hdrs = [
        "verifier.h",
    ],
    copts = [
        "-Wno-non-virtual-dtor",
        "-Wno-unused-variable",
        "-Wno-implicit-fallthrough",
    ],
    deps = [
        ":lib",
        "//kythe/proto:common_cc_proto",
        "//kythe/proto:storage_cc_proto",
        "//third_party:gtest",
        "@com_github_google_glog//:glog",
        "@com_google_protobuf//:protobuf",
    ],
)

cc_test(
    name = "test",
    size = "small",
    deps = [
        ":testlib",
    ],
)
