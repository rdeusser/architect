usage="Builds the project"

log::info "Building $(project::name)"

go build -trimpath -gcflags "all=-trimpath=${GOPATH}" -asmflags "all=-trimpath=${GOPATH}" -o "build/$(project::name)" "$(project::repo)"
