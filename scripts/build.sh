usage="Builds plz"

log::info "Building $(project::name)"

go build -trimpath -gcflags "all=-trimpath=${GOPATH}" -asmflags "all=-trimpath=${GOPATH}" "$(project::repo)/cmd/plz"
