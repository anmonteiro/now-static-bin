FROM debian:7

RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y git curl sudo file build-essential

RUN curl -sSf https://sh.rustup.rs | sh -s -- -y
ENV PATH=/root/.cargo/bin:$PATH

RUN mkdir /compile
WORKDIR /compile

RUN git clone https://github.com/steveklabnik/simple-server /compile

RUN rustup target add x86_64-unknown-linux-musl

RUN cargo build --target=x86_64-unknown-linux-musl --release --example server
RUN mv /compile/target/x86_64-unknown-linux-musl/release/examples/server /server.exe