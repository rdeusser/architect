usage="Installs the project"

log::info "Installing $(project::name)..."

go install -trimpath -gcflags "all=-trimpath=${GOPATH}" -asmflags "all=-trimpath=${GOPATH}" "$(project::repo)"