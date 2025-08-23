Bitcoin and its blockchain are problematic under several aspects:
- Limited scalability: the number of **transactions per second is limited by the block size and by the block time**.
- Low efficiency: At least 1h to safely confirm a transaction (devo aspettare 6 blocchi).
- Data expressiveness: in Bitcoin, data can be only economic transactions.

To overcome above open issues, new blockchain and cryptocurrency implementations have started to appear.

In particular, they introduce some new concepts like:
- **Smart Contracts**.
- Decentralized Applications (dApps).
- Decentralized Autonomous Organizations (DAO).
- Non-fungible Tokens (NFTs).
- Oracles.

### Smart contracts
blockchain è un modo per avere una struttura dati condivisa immutabile tra dei peer che raggiungono un consenso sullo stato di questi dati
- bitcoin si limita a rappresentare transazioni su questa struttura dati

Ma nulla vieta di inserire in questa struttura dati quello che vogliamo. In particolare cosa succede se salviamo dei programmi su una blockchain?

A Smart Contract is a computer program stored and executed on a blockchain-based platform.
- Stored on the blockchain = Each node in the network has a copy of it.
- (contratto nome fuorviante, è un programma)

è un programma scritto in un normale linguaggio di programmazione, **salvato ed eseguito su blockchain**
- lo stato del programma è concordato dalla rete

**Come viene eseguito?**
- Each contract has an address that uniquely identifies it, functions that can be invoked, and an internal state (e.g. the value of a variable at a given
moment in time).
- The first thing to do is to deploy the SC on the blockchain. The SC is thus distributed to all nodes.
- To execute a function it is necessary **to perform a transaction to the address of the SC**, specifying the function name and arguments.
- Each node independently executes the operations defined by the contract on its local copy.
- All of them have to get the same result to reach the consensus.  Otherwise the transaction is not valid.

**Limitazioni**
- **In order to achieve consensus, it is essential that the SCs have a deterministic execution**.
- A SC, contrary to what is often said, cannot self-execute itself. A transaction must be explicitly made to it.
- The immutability of the SC prevents adding functionality and/or correcting errors. To upgrade an SC, it is necessary to deploy a new SC and make sure that everyone uses it.
- Since the source code is public, anyone can easily exploit the vulnerabilities.
- **The only data a SC can use is that already present on blockchain**.

**Cosa ottengo**
- esecuzione distribuita, fault tolerant
    - non ho un server centralizzato su cui esegue il codice
    - tutti lo eseguono
- dati immutabili
- tracing completo
    - ho l'intera storia dell'esecuzione del programma
- disaster recovery
    - in eventualità di crash lo stato precedente è riproducibile in quanto salvato su blockchain sotto forma di transazioni


### Decentralized applications
A decentralized application (dApp) is an application built on a decentralized network that **combines a smart contract and a frontend user interface**.

Abbiamo tre parti:
- Frontend: is the interface that users use to interact with the application, e.g. a Web interface.
- Backend: is the heart of the application, which implements the application logic.
    - Is realised through one or more interacting SCs.
- Database: data storage is fully decentralized thanks to the use of the underlying blockchain.


e poi anche altro ma non interessa ...