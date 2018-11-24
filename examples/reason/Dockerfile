FROM node:10.13-alpine as build

ENV TERM=dumb \
		LD_LIBRARY_PATH=/usr/local/lib:/usr/lib:/lib

RUN mkdir /esy
WORKDIR /esy

ENV NPM_CONFIG_PREFIX=/esy
RUN npm install -g --unsafe-perm esy@0.4.3

# now that we have esy installed we need a proper runtime

FROM alpine:3.8 as esy

ENV TERM=dumb \
		LD_LIBRARY_PATH=/usr/local/lib:/usr/lib:/lib

WORKDIR /

COPY --from=build /esy /esy

RUN apk add --no-cache \
		ca-certificates wget \
		bash curl perl-utils \
		git patch gcc g++ musl-dev make m4

RUN wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub
RUN wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.28-r0/glibc-2.28-r0.apk
RUN apk add --no-cache glibc-2.28-r0.apk

ENV PATH=/esy/bin:$PATH

RUN mkdir /app
WORKDIR /app
COPY package.json package.json
COPY esy.lock esy.lock
RUN esy fetch
RUN esy build-package @esy-ocaml/substs@0.0.1
RUN esy build-package @opam/base-unix@opam:base
RUN esy build-package ocaml@4.7.1003
RUN esy build-package @opam/ocamlbuild@opam:0.12.0
RUN esy build-package @opam/conf-m4@opam:1
RUN esy build-package @opam/ocamlfind@opam:1.8.0
RUN esy build-package @opam/dune@opam:1.5.1
RUN esy build-package @opam/jbuilder@opam:transition
RUN esy build-package @opam/result@opam:1.3
RUN esy build-package @opam/topkg@opam:1.0.0
RUN esy build-package @opam/cmdliner@opam:1.0.2
RUN esy build-package @opam/base-bytes@opam:base
RUN esy build-package @opam/base64@opam:2.3.0
RUN esy build-package @opam/sexplib0@opam:v0.11.0
RUN esy build-package @opam/base@opam:v0.11.1
RUN esy build-package @opam/ocaml-migrate-parsetree@opam:1.1.0
RUN esy build-package @opam/ocaml-compiler-libs@opam:v0.11.0
RUN esy build-package @opam/ppx_derivers@opam:1.0
RUN esy build-package @opam/stdio@opam:v0.11.0
RUN esy build-package @opam/ppxlib@opam:0.3.1
RUN esy build-package @opam/fieldslib@opam:v0.11.0
RUN esy build-package @opam/uchar@opam:0.0.2
RUN esy build-package @opam/uutf@opam:1.0.1
RUN esy build-package @opam/jsonm@opam:1.0.1
RUN esy build-package @opam/ppx_fields_conv@opam:v0.11.0
RUN esy build-package @opam/ppx_sexp_conv@opam:v0.11.2
RUN esy build-package @opam/seq@opam:base
RUN esy build-package @opam/re@opam:1.8.0
RUN esy build-package @opam/stringext@opam:1.5.0
RUN esy build-package @opam/uri@opam:2.0.0
RUN esy build-package @opam/cohttp@opam:1.2.0
RUN esy build-package @opam/cppo@opam:1.6.5
RUN esy build-package @opam/lwt@opam:4.1.0
RUN esy build-package @opam/cohttp-lwt@opam:1.2.0
RUN esy build-package @opam/astring@opam:0.8.3
RUN esy build-package @opam/num@opam:1.1
RUN esy build-package @opam/parsexp@opam:v0.11.0
RUN esy build-package @opam/sexplib@opam:v0.11.0
RUN esy build-package @opam/ipaddr@opam:2.8.0
RUN esy build-package @opam/fmt@opam:0.8.5
RUN esy build-package @opam/logs@opam:0.6.2
RUN esy build-package @opam/conduit@opam:1.3.0
RUN esy build-package @opam/conduit-lwt@opam:1.3.0
RUN esy build-package @opam/conduit-lwt-unix@opam:1.3.0
RUN esy build-package @opam/magic-mime@opam:1.1.0
RUN esy build-package @opam/cohttp-lwt-unix@opam:1.2.0
RUN esy build-package @opam/menhir@opam:20181113
RUN esy build-package @opam/merlin-extend@opam:0.3
RUN esy build-package @opam/reason@opam:3.3.7
COPY . .

RUN esy b dune build src/main.exe --profile=docker

RUN mv $(esy bash -c 'echo $cur__target_dir/default/src/main.exe') /app/main.exe
