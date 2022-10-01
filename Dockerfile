FROM node:18

# create app directory
WORKDIR /usr/src/app

# install pnpm
RUN npm install -g pnpm

# add tini
ENV TINI_VERSION v0.19.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /tini
RUN chmod +x /tini

# add PocketBase
ENV PB_VERSION 0.7.6
ADD https://github.com/pocketbase/pocketbase/releases/download/v${PB_VERSION}/pocketbase_${PB_VERSION}_linux_amd64.zip /tmp/pb.zip
RUN unzip /tmp/pb.zip -d /pb/

# install app dependencies
COPY package.json pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile

# copy src code and build
COPY . .
RUN pnpm svelte-kit sync
RUN pnpm build

EXPOSE 3000
ENTRYPOINT ["/tini", "--", "./entrypoint.sh"]
