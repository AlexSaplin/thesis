FROM msitnikov/selachii-base-image


# Install dependencies
WORKDIR /code

# Build & install selachii
COPY . /code/selachii

RUN cd selachii && cmake . && make && \
    cp bin/selachii /bin/selachii && rm -rf /code/selachii

ENTRYPOINT ["/bin/selachii"]
