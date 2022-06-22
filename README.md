Run `bazel build //...`

Expected: the build passes.

Actual outcome:

```
% bazel build //...
Loading:
Loading: 0 packages loaded
Analyzing: 3 targets (0 packages loaded, 0 targets configured)
INFO: Analyzed 3 targets (0 packages loaded, 0 targets configured).
INFO: Found 3 targets...
[0 / 2] [Prepa] BazelWorkspaceStatusAction stable-status.txt ... (2 actions, 0 running)
ERROR: /Users/gregorbest/code/tmp/BUILD.bazel:10:11: GoCompilePkg genset.a failed: (Exit 1): builder failed: error executing command bazel-out/darwin_arm64-opt-exec-2B5CBBC6/bin/external/go_sdk/builder compilepkg -sdk external/go_sdk -installsuffix darwin_arm64 -src genset.go -embedroot '' -embedroot ... (remaining 21 arguments skipped)

Use --sandbox_debug to see verbose messages from the sandbox and retain the sandbox build root for debugging
compilepkg: panic: 0: got 0 arguments but 1 type parameters

goroutine 19 [running]:
go/types.(*Checker).validateTArgLen(0x1400018c000?, 0x140001b6000?, 0x1, 0x0)
	GOROOT/src/go/types/instantiate.go:131 +0x14c
go/types.(*Checker).instance(0x0, 0x0, {0x1007bd9c8?, 0x1400011c180?}, {0x10092ec60?, 0x0, 0x0}, 0x1400018c000)
	GOROOT/src/go/types/instantiate.go:89 +0x12c
go/types.Instantiate(0x0?, {0x1007bd9c8?, 0x1400011c180?}, {0x10092ec60?, 0x14000182b10?, 0x30?}, 0x0?)
	GOROOT/src/go/types/instantiate.go:59 +0x2a4
golang.org/x/tools/internal/typeparams.Instantiate(...)
	external/org_golang_x_tools/internal/typeparams/typeparams_go118.go:150
golang.org/x/tools/go/ssa.createInstance(0x140001941a0, {0x10092ec60?, 0x0, 0x0}, 0x14000095950, {0x1007bd3b0?, 0x1400009b440}, 0x14000190060)
	external/org_golang_x_tools/go/ssa/instantiate.go:128 +0xcc
golang.org/x/tools/go/ssa.(*instanceSet).lookupOrCreate(0x1400018c020, {0x10092ec60, 0x0, 0x0}, 0x140000b6ad8?)
	external/org_golang_x_tools/go/ssa/instantiate.go:110 +0xe8
golang.org/x/tools/go/ssa.(*Program).needsInstance(0x100780aa0?, 0x0?, {0x10092ec60, 0x0, 0x0}, 0x140000b6ce8?)
	external/org_golang_x_tools/go/ssa/instantiate.go:85 +0x110
golang.org/x/tools/go/ssa.(*builder).expr0(0x140000b79f0, 0x140001944e0, {0x1007be0f8?, 0x140000c0920?}, {0x7, {0x1007bd9c8, 0x1400011ccc0}, {0x0, 0x0}})
	external/org_golang_x_tools/go/ssa/builder.go:781 +0xac8
golang.org/x/tools/go/ssa.(*builder).expr(0x140000b6ef8?, 0x140001944e0, {0x1007be0f8?, 0x140000c0920?})
	external/org_golang_x_tools/go/ssa/builder.go:610 +0x134
golang.org/x/tools/go/ssa.(*builder).setCallFunc(0x100ae8a68?, 0x80?, 0x14000080480?, 0x14000190140)
	external/org_golang_x_tools/go/ssa/builder.go:1006 +0x2f4
golang.org/x/tools/go/ssa.(*builder).setCall(0x10077fa20?, 0x1400009bad0?, 0x140000b3f40, 0x14000190140)
	external/org_golang_x_tools/go/ssa/builder.go:1085 +0x30
golang.org/x/tools/go/ssa.(*builder).expr0(0x140000b79f0, 0x140001944e0, {0x1007bddf8?, 0x140000b3f40?}, {0x7, {0x1007bd9a0, 0x14000099b40}, {0x0, 0x0}})
	external/org_golang_x_tools/go/ssa/builder.go:686 +0x2354
golang.org/x/tools/go/ssa.(*builder).expr(0x14000080480?, 0x140001944e0, {0x1007bddf8?, 0x140000b3f40?})
	external/org_golang_x_tools/go/ssa/builder.go:610 +0x134
golang.org/x/tools/go/ssa.(*builder).stmt(0x140000b7508?, 0x140001944e0, {0x1007be338?, 0x140000c0960?})
	external/org_golang_x_tools/go/ssa/builder.go:2161 +0x12d8
golang.org/x/tools/go/ssa.(*builder).stmtList(0x14000182ae0?, 0x1?, {0x140000996b0?, 0x1, 0x30?})
	external/org_golang_x_tools/go/ssa/builder.go:911 +0x6c
golang.org/x/tools/go/ssa.(*builder).stmt(0x140001944e0?, 0x140001944e0, {0x1007bdd98?, 0x1400009b740?})
	external/org_golang_x_tools/go/ssa/builder.go:2218 +0x8b8
golang.org/x/tools/go/ssa.(*builder).buildFunctionBody(0x14000194000?, 0x140001944e0)
	external/org_golang_x_tools/go/ssa/builder.go:2327 +0x3a4
golang.org/x/tools/go/ssa.(*builder).buildFunction(0x1006e2bd0?, 0x140001944e0)
	external/org_golang_x_tools/go/ssa/builder.go:2267 +0x34
golang.org/x/tools/go/ssa.(*builder).buildCreated(0x1400010f9f0)
	external/org_golang_x_tools/go/ssa/builder.go:2349 +0x2c
golang.org/x/tools/go/ssa.(*Package).build(0x14000190000)
	external/org_golang_x_tools/go/ssa/builder.go:2529 +0xab8
sync.(*Once).doSlow(0x1400018e000?, 0x140000959a0?)
	GOROOT/src/sync/once.go:68 +0x10c
sync.(*Once).Do(...)
	GOROOT/src/sync/once.go:59
golang.org/x/tools/go/ssa.(*Package).Build(...)
	external/org_golang_x_tools/go/ssa/builder.go:2413
golang.org/x/tools/go/analysis/passes/buildssa.run(0x14000184000)
	external/org_golang_x_tools/go/analysis/passes/buildssa/buildssa.go:73 +0x13c
main.(*action).execOnce(0x140000de510)
	external/io_bazel_rules_go/go/tools/builders/nogo_main.go:301 +0x764
sync.(*Once).doSlow(0x0?, 0x0?)
	GOROOT/src/sync/once.go:68 +0x10c
sync.(*Once).Do(...)
	GOROOT/src/sync/once.go:59
main.(*action).exec(0x0?)
	external/io_bazel_rules_go/go/tools/builders/nogo_main.go:245 +0x44
main.execAll.func1(0x0?)
	external/io_bazel_rules_go/go/tools/builders/nogo_main.go:239 +0x50
created by main.execAll
	external/io_bazel_rules_go/go/tools/builders/nogo_main.go:237 +0x48
INFO: Elapsed time: 0.254s, Critical Path: 0.09s
INFO: 2 processes: 2 internal.
FAILED: Build did NOT complete successfully
FAILED: Build did NOT complete successfully
```

Manually running the `nilness` analyzer pass works:

```
% go run golang.org/x/tools/go/analysis/passes/nilness/cmd/nilness@v0.1.11.0 ./...
%
```
