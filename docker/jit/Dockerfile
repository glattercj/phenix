ARG PHENIX_REPO=sandialabs/sceptre-phenix
ARG PHENIX_TAG=main
ARG NODE_ARCH=x64
ARG BUILDPLATFORM

FROM --platform=$BUILDPLATFORM phenix

# https://docs.docker.com/engine/reference/builder/#understand-how-arg-and-from-interact
ARG PHENIX_REPO
ARG PHENIX_TAG
ARG NODE_ARCH

ENV NODE_VERSION 14.21.3

RUN wget -O node.txz https://nodejs.org/dist/v${NODE_VERSION}/node-v${NODE_VERSION}-linux-$NODE_ARCH.tar.xz \
  && tar -xJf node.txz -C /usr/local --strip-components=1 --no-same-owner && rm node.txz \
  && ln -s /usr/local/bin/node /usr/local/bin/nodejs

RUN curl -o- -L https://yarnpkg.com/install.sh | bash

RUN npm install -g @vue/cli redoc-cli
RUN git clone --branch ${PHENIX_TAG} https://github.com/${PHENIX_REPO}.git /usr/local/src/phenix \
  && mkdir -p /opt/phenix/web \
  && cp -a /usr/local/src/phenix/src/go/web/public /opt/phenix/web

WORKDIR /usr/local/src/phenix/src/js
RUN rm yarn.lock && npm install node-sass@npm:sass && npm install
WORKDIR /opt/phenix/web/public
RUN npx redoc-cli build docs/openapi.yml -o docs/index.html --title 'phenix API'

COPY phenix-ui-entrypoint.sh /phenix-ui-entrypoint.sh
RUN  chmod +x /phenix-ui-entrypoint.sh

WORKDIR /opt/phenix

ENTRYPOINT ["/phenix-ui-entrypoint.sh"]
