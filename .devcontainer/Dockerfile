FROM mcr.microsoft.com/devcontainers/go:0-1-bullseye

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# RUN go get -x <your-dependency-or-tool>
# USER root

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1

USER vscode

# Install CompileDaemon to allow us to have a build watcher
# For more information, go here: https://github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon@latest
