Vizmod
======

Visualization for Ethereum blockchains.

These sources help to show how different nodes in an Ethereum blockchain network work and reach consensus.

It uses a web page, per node, with colored letters to show whether nodes agree on hashes — i.e. reach consensus. Each colored letter symbolizes a hash. When the same color and letter appears in a box this signifies that two nodes have consensus about a certain value.



Basics
------


This concept has three components:

* a modded version of the Ethereum Go client *geth*
* a mini web server that serves dynamic data updates
* a web page with Javascript to display and pull updates

### Modded geth: vizmod-geth

Geth is one of the most used Ethereum clients. This modification of it has been instrumented for the purpose of visualizing, adding additional fast logging and specifically writing out files with only one current value.

### Mini server: vizmod-server

A minimal web server written in Go. It allows for visualization to run anywhere: on the same computer as the node or another. Both ways, a browser is used to pull a page and data updates through this server. The vizmod-server is a completely generic channel with no knowledge of visualization.

### Visualization web page

This page is dynamically created by Javascript and continually updated, usually every second. Three values are displayed that symbolize the state of a node in big, colored, single letters. These are the nine different values that are observed:

* Last Transaction Broadcast
* Last Transaction Received
* Current Work: Nonce Tried
* Proof of Work: Nonce Found
* Proof Accepted: Verified Nonce
* Last Proposed Block
* Current Root Hash
* Current Number of Peers
* Current Chain Height

See detailed descriptions below.



Use
---


### Test Use with Ethereum Mainnet

The easiest setup is using the Ethereum mainnet. If you use two nodes, e.g. running this setup on two networks, you will see the synchronization happen. But it is mainly useful for getting your toes wet and learn how to set up a small private network later.

#### Run the modded geth

Prerequisite: Go and github installed.

Install, build and run a modded geth from source using https://github.com/claryon/vizmod.

	$ git clone https://github.com/claryon/vizmod  
	$ cd vizmod/vizmod-geth  
	$ make  
	$ build/bin/geth  

Geth will start syncing with the Ethereum mainnet. This can take some hours. The visualization will already work.

To speed things up you might be able to copy or link a data directory from another, normal geth instance in.

This modded geth has extra logging that is needed for the visualization. In a different console, go to the vizmod-geth folder and check for files it writes to its current directory that start on `vizmod-`.

Check general activity

	$ tail -f vizmod-full.log

Check that `vizmod-chain-hash.flush` is logged as expected, e.g. catting it to screen repeatedly. Or using

	$ watch cat vizmod-chain-hash.flush

With vizmod-geth running it should show a hash change every 5 to 20 seconds.

Note that it might not change while vizmod-geth is still starting up. That phase can take a minute. It should work though no matter whether vizmod-geth is fast syncing/catching up with the mainnet or in regular mode.

#### Run the miniature web server

Install, build and run the vizmod web server on the same node, from the same repo https://github.com/claryon/vizmod.

	$ cd  
	$ cd vizmod/vizmod-server  
	$ go run vizmod-server.go  


Check the server gives you a page when browsing `localhost:3000/`

#### Check the data link

Check the softlinks in vizmod-server/docroot. They need to point to the flush files in vizmod-geth’s root. They do not share the exact same name but the links only use a middle part E.g.:

	$ ls chain-hash
	$ chain-hash --> ../../vizmod-geth/vizmod-chain-hash.flush

Check that the link vizmod-server/docroot/chain-hash works, e.g. catting it to screen repeatedly. Or using

	$ watch cat chain-hash

With vizmod-geth running it should show a hash change every 5 to 20 seconds.

If this is not the case, create the links in the vizmod-server’s docroot, e.g.:

	$ ln -s ../../vizmod-geth/vizmod-chain-hash.flush chain-hash

You might use the `link` batch for this with the path as parameter:

	$ ./link ../../vizmod-geth

Otherwise, there are nine of them to take care of:

* `tx-broadcast`
* `tx-received`
* `nonce-tried`
* `nonce-found`
* `nonce-accepted`
* `proposed-block`
* `chain-hash`
* `peer-count`
* `chain-height`

Call the visualization up browsing to `localhost:3000/viz.html`

You should now see the instrumentation visualization as expected.


### Use with Private Network

For a private network, each node should get vizmod installed as above. There is no need to install normal, unmodded Ethereum clients, vizmod-geth only adds logging it does not have anything less.

For example, for a private test network on five Raspberry Pis, install vizmod-geth on each. Set a private network number on start up for all of them.

There are cross compiled binaries for Raspbian Pi 2's (ARM 7) in `bin/`.

The visualization can be displayed by any browser that is pointed to the IP of the respective Pi, port 3000. E.g. Displays can be connected to the Pis t hemselves, each displaying a browser pointing to `http://localhost:3000/`. Or five laptops that are on the same local or wifi network can display the data of one Raspberry Pi each. E.g. http://192.168.1.11:3000/ or whatever the IP of any one Pi might be. The connection delay should be in the milliseconds, lower than the actual polling frequency delay of about a second.



Meaning of Displays
------------------


### Tx Broadcast
Hash of the last transaction that this node has been broadcast to the network, i.e. any of its peers.

Infrequent or no action: will change whenever a transaction enters the network through this very node, e.g. is sent to it via API.

This value is not expected to be in sync across peer nodes.

`vizmod-tx-broadcast.flush` last value  
`vizmod-tx-broadcast.log` timestamped log


### Tx Received
Hash of the last transaction that this node has received from the network, i.e. any of its peers.

Infrequent action: will change whenever a transaction enters the network, at any other peer.

Nodes are not expected to be in sync on this value as they may receive transactions in different order and this value only shows the last. Also the node that sent the transaction in question will not show it as received.

However, a very slow, private system that processes only few transactions a minute, may turn out to be mostly in sync.

`vizmod-tx-received.flush` last value  
`vizmod-tx-received.log` timestamped log


### Work: Nonce Tried
Hash of the last nonce that this node has tried out to form a block.

Continuous action when this node is mining, none if not. Nonce are tried by the millions per second and the log and display will show some snapshots per second.

Nodes are not expected to ever be in sync on this value.

`vizmod-nonce-tried.flush` last value  
`vizmod-nonce-tried.log` timestamped log


### Proof of Work: Nonce Found
Hash of the last nonce that this node found to successfully form a block. This nonce is a valid proof of work.

Slow action around every 5 to 60 seconds when this node is mining in a small private network, none if not. A high mining difficulty of the network — e.g. the Ethereum main net — and accepting blocks from other nodes will stop the node’s search for proof of work. Only on a small private network will this display ever show anything.

Nodes are not expected to be in sync on this value.

`vizmod-nonce-found.flush` last value  
`vizmod-nonce-found.log` timestamped log


### Proof Accepted: Nonce verified
Hash of accepted proof of work from another node: the last nonce that this node found to match a block as expected, coming from a peer. This nonce has been verified to be a valid proof of work then.

Slow action around every 5 to 20 seconds. More even when running in a network with a lot of nodes. High variance expected in a small, private network. Displays independently of wether a node is mining.

Nodes are not expected to be in sync on this value as different nodes will receive different block proposals in different sequence. In a small network nodes might show mostly in sync.

`vizmod-nonce-accepted.flush` last value  
`vizmod-nonce-accepted.log` timestamped log


### Proposed Block
Hash of the last block that this node proposed to the network.

Slow action around every 5 to 60 seconds when this node is mining in a small private network, none if not. A high mining difficulty of the network — e.g. the Ethereum main net — and accepting blocks from other nodes will make the node unable to find and provide blocks in time. Only on a small private network will this display ever show anything.

Nodes are not expected to be in sync on this value. But they are expected to display this value as Root Hash from time to time and all agree on it.

`vizmod-proposed-block.flush` last value  
`vizmod-proposed-block.log` timestamped log


### Root Hash
Hash of the highest, and last block that this node accepts as top of the blockchain. This is ‘the’ blockchain hash that is chained to the next block and that all the nodes form consensus over.

Reliable, slow action around every 5 to 20 seconds.

All nodes are expected to find consensus over this hash. If not, the network is split and a fork happens. This hash can be used to demonstrate a fork with two different consensus’ forming in two partitioned sub networks.

`vizmod-chain-hash.flush` last value  
`vizmod-chain-hash.log` timestamped log


### Number of Peers
Number of peers that this node has and communicates with. The node does not communicate with all of the peers all of the time.

Slow, infrequent action mostly at node start when the peers are searched and connected to.

The nodes are not expected to sync up on this value. In a small, private network where every node ends up having a connection with every participating node, they will.

`vizmod-peer-count.flush` last value  
`vizmod-peer-count.log` timestamped log


### Chain Hash
Number of highest, and last block that this node accepts as top of the blockchain, as counted in unbroken line from the genesis block. This is ‘the’ height of the chain that (mostly) decides which version of the blockchain prevails when block proposals compete or a partition is re-united.

Reliable, slow action around every 5 to 20 seconds.

All nodes are expected to find consensus over this hash. Even when a network is partitioned, the sub partitions are expected to only slowly drift apart in regards to this value, i.e. within minutes. This can be used to demonstrate a fork, to show how work continues on both sides but slower on the weaker side that will lose its state on re-unification.

`vizmod-chain-height.flush` last value  
`vizmod-chain-height.log` timestamped log



Stability
---------

Successful, hands-off 48 hour test run on Ethereum mainnet, including entering and emerging from system hybernation.

Note, the vizmod-* .log files might run a system out of space when run a very long time. Make them links to /dev/null if that is ever intended. They only serve for debugging and are not needed for operation. The vizmod-* .flush files are and they don’t use up disk space. 


Safety
-----

THIS SETUP IS FOR DEMONSTRATION PURPOSES ONLY AND IS NOT HARDENED. IT IS NOT MADE FOR NODES USED IN A PUBLIC NETWORK. THE WEB SERVER MAY EXPOSE THE ENTIRE NODE. THE MINING POWER OF THE MODIFIED CLIENT MAY BE DEGRADED. LOGS MAY WEIGH ON IO THROUGHPUT AND FILL UP THE NODES HARD DISK.

License
-------

(c) 2017 Henning Diedrich, Claryon UG (haftungsbeschränkt).   
geth (c) Ethereum Foundation.  
To be open-sourced by the European Commission.  
