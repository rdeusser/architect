usage="Installs plz"

log::info "Installing $(project::name)"

go install -trimpath -gcflags "all=-trimpath=${GOPATH}" -asmflags "all=-trimpath=${GOPATH}" "$(project::repo)/cmd/plz"
