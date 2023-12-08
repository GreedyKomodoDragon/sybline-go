# Sybline Go Client

Go client for the [Sybline Message Broker](https://github.com/GreedyKomodoDragon/Sybline).

<p align="center">
  <img src="https://img.shields.io/badge/license-AGPL-blue.svg" alt="AGPL">
   <img src="https://img.shields.io/badge/Go-00ADD8?style=&logo=go&logoColor=white" alt="AGPL">
</p>

# Table of Contents
- [Features](#Features)
- [Installation](#Installation)
- [Documentation](#Documentation)
- [Contributing](#contributing)
- [License](#license)


# Features

See docs for more information: [Sybline Docs](https://www.sybline.com/en/v031/golang/install)

The list of features are:

* Direct Message Routing   
* Batch Messaging
    * Can batch Submit messages to a routing key
    * Can batch ack and nack messages
* Built-in Consumer Struct
    * Handles polling the sybline cluster, you just handle what to do with the messages!
* Security:
    * Account Management
    * mTLS Communication (Non-Account Level)
    * Role based Managed
        * JSON Based Roles
        * Allow and Deny Actions
            * Denys take precedence
        * Restrict Access to Specific Queues or Route Key

# Installation

Go install command:
```
go get github.com/GreedyKomodoDragon/sybline-go
```

# Documentation

All of the go documentation can be found at the [Sybline Documentation](https://www.sybline.com/en/v031/golang/install) website.


### Proto Command

```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    mq.proto
```

# Contributing

Sybline (or any sybline related projects) are open to contributions whether these are new features or bug-fixes.

Please note, if the feature does not align with the original goal of Sybline it will sadly not be accepted; we don't want the scope of Sybline to become too unmaintainable.

If you are interested in the project but have no/little technical experience, please have a look at the [documentation repo](https://github.com/GreedyKomodoDragon/sybline-docs), it always needs changes or translations!

# License

Sybline has been released under GNU Affero General Public License v3.0. 
* This is a copyleft License