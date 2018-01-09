# ReadMe

## What is mktcapService ?

mktcapService is a service that get data from coinmarketcap.com then do the following jobs

* Save Data to assigned SQLDatabase
  * turn on/off use saveToDB in conf.yml
* Monitor Data change and notify user according to the rules has been used
  * turn on/off use quickMonitor in conf.yml

## Notification

* Send Message to Slack and assigned channel according to rule
  * Config slack in conf.yml

## Rule

* Current Supported Rules
  * ruleSigDiff : this rule caculate change of btc , usd and 1hr change in percent and then determine if they exceed threadhold
    * numObervations : number of observations to diff value. ie : 5 will go back to 5 records and diff with current one
    * threadholePercnt : threadhold in percent

## conf.yml

Due to the security reason the conf.yml is not git, you should generate conf.yml and put them in project
conf.ym in mktcapService folder

```
---
database:
  mktcapdb: cryptomarket
  sqlendpoint: 127.0.0.1
  sqlpwd: hello123
  sqluser: hellouser
  tickertable: mktcap_ticker
enableService:
  quickMonitor: true
  saveToDB: true
quickMonitorService:
  monitorIntervalSec: 120
  monitorLimitRecords: 100
ruleSigDiff:
  numObervations: 2
  threadholePercnt: 1
saveToDBService:
  saveToDBLimitRecords: 100
  saveToDBSec: 900
slack:
  slackChannel: coinmarket
  token: xoxp-166092624144-xxxxxxxxxxxxxxxx-xxxxxxx-xxxxxxxxxxx
```

## Coin ID Rank 100
```
bitcoin ethereum ripple bitcoin-cash cardano nem litecoin stellar iota tron dash neo monero eos icon qtum bitcoin-gold lisk raiblocks ethereum-classic verge siacoin omisego bytecoin-bcn bitconnect zcash populous stratis bitshares binance-coin dogecoin ardor kucoin-shares status dentacoin steem tether waves vechain dragonchain digibyte dent augur hshare 0x ark komodo veritaseum golem-network-tokens wax basic-attention-token electroneum salt funfair decred kyber-network
ethos experience-points neblio nexus kin reddcoin medibloc substratum pivx aion factom request-network qash storm power-ledger aelf aeternity bytom monacoin gas nxt rchain cobinhood digitalnote iconomi deepbrain-chain syscoin maidsafecoin byteball enigma-project bitcoindark time-new-bank paccoin chainlink santiment quantstamp tenx zcoin digixdao gamecredits gnosis-gno blockv red-pulse storj
```



##### tag `coinmarketcap.com` `crypto market`
