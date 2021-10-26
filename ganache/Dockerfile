FROM node:8-alpine
 
RUN npm install -g ganache-cli
RUN mkdir -p /var/ganache
 
ENV DOCKER true
ENV DATADIR /var/ganache
# データを永続化する場合、次回以降も同じmnemonicとnetworkIdで立ち上げる必要性があります
# https://github.com/trufflesuite/ganache-cli/issues/407
 
EXPOSE 8545
 
ENTRYPOINT ["sh", "-c", "ganache-cli --host=0.0.0.0 ]