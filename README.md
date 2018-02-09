# ReadMe

## What is mktcapService ?

mktcapService is a service that get data from coinmarketcap.com then do the following jobs

* Save Data to assigned SQLDatabase
  * turn on/off use saveToDB in conf.yml
* Monitor Data change and notify user according to the rules has been used
  * turn on/off use quickMonitor in conf.yml
  * use monitorType : assign to monitor specific coins by coin ID  in monitorCoinIDs
  * use  monitorType : top to monitor top monitorLimitRecords rank coins
* Get Specific coin info
  * user apiServer.port in conf.yml to change listen port


## Api

+ get coin info by coin id
```http://127.0.0.1:8080/coin/bitcoin-cash```
```Bitcoin Cash (btc:0.18842700)(usd:2624.03) (1hChg:0.91)(24hrChg:-0.25)(7dayChg:3.99)(lastupdate:2018-01-12 10:39:12 -0500 EST)```


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
  sqlpwd: mypassword
  sqluser: myuser
  tickertable: mktcap_ticker
enableService: 
  quickMonitor: true
  saveToDB: false
quickMonitorService: 
  monitorCoinIDs: 
    - btc
    - ethereum
    - ripple
  monitorIntervalSec: 120
  monitorLimitRecords: 100
  monitorType: assign
ruleSigDiff: 
  - numObervations: 2
    threadholdPercnt: 0.01
  - numObervations: 10
    threadholdPercnt: 8
saveToDBService: 
  saveToDBLimitRecords: 100
  saveToDBSec: 900
slack: 
  slackChannel: mktnotify
  token:  xoxp-166092624144-1xxxxxxxxxxx
apiServer:
  port: 8080


```

## Coin ID Rank 100
```
bitcoin ethereum ripple bitcoin-cash cardano nem litecoin stellar iota tron dash neo monero eos icon qtum bitcoin-gold lisk raiblocks ethereum-classic verge siacoin omisego bytecoin-bcn bitconnect zcash populous stratis bitshares binance-coin dogecoin ardor kucoin-shares status dentacoin steem tether waves vechain dragonchain digibyte dent augur hshare 0x ark komodo veritaseum golem-network-tokens wax basic-attention-token electroneum salt funfair decred kyber-network
ethos experience-points neblio nexus kin reddcoin medibloc substratum pivx aion factom request-network qash storm power-ledger aelf aeternity bytom monacoin gas nxt rchain cobinhood digitalnote iconomi deepbrain-chain syscoin maidsafecoin byteball enigma-project bitcoindark time-new-bank paccoin chainlink santiment quantstamp tenx zcoin digixdao gamecredits gnosis-gno blockv red-pulse storj
```



##### tag `coinmarketcap.com` `crypto market` `cryptocurrency`
