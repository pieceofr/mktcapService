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

##### tag `coinmarketcap.com` `crypto market`
