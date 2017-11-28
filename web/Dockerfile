FROM alpine
# SOURCE : https://curl.haxx.se/docs/caextract.html
ADD ca-certificates.crt /etc/ssl/certs/
ADD src/node_modules /node_modules
COPY src/main.js src/main.js.map index.html /
ADD web /
CMD ["/web"]
