go_library(
    name = "errgroup",
    srcs = ["errgroup.go"],
)

go_test(
    name = "errgroup_test",
    srcs = ["errgroup_test.go"],
    deps = [
        ":errgroup",
        ":testify",
    ],
)

go_get(
    name = "testify",
    get = "github.com/stretchr/testify",
    install = [
        "assert",
        "require",
        "vendor/github.com/davecgh/go-spew/spew",
        "vendor/github.com/pmezard/go-difflib/difflib",
    ],
    revision = "f390dcf405f7b83c997eac1b06768bb9f44dec18",
    deps = [":spew"],
)

go_get(
    name = "spew",
    get = "github.com/davecgh/go-spew/spew",
    revision = "ecdeabc65495df2dec95d7c4a4c3e021903035e5",
)
