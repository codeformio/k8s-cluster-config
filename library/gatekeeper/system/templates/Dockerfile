FROM golang:1.13

WORKDIR /work

RUN git clone https://github.com/open-policy-agent/opa --branch v0.17.3
RUN cd ./opa && go install .

COPY ./build ./build
RUN cd ./build && go install .


RUN echo "#!/usr/bin/env bash\n" \
"opa test -v *.rego\n" \
"build\n" \
> /bin/build.sh && chmod +x /bin/build.sh

ENTRYPOINT /bin/build.sh
