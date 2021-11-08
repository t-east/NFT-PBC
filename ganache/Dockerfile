FROM node:8-alpine
 
RUN npm install -g ganache-cli
RUN mkdir -p /var/ganache
 
ENV DOCKER true
ENV DATADIR /var/ganache
 
EXPOSE 8545
 
ENTRYPOINT ["/bin/sh", "-c", "ganache-cli --host=0.0.0.0 --db=${DATADIR} --gasPrice=0 --mnemonic=\"${MNEMONIC}\" "]