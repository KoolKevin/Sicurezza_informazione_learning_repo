### Trust problem | problema che ha motivato la nascita della blockchain
La tecnologia blockchain nasce nel contesto dei sistemi finanziari in cui è presente una **terza parte fidata** che fa da garante

Pensa alle carte di credito
- Most common option for payments. Each buyer and each seller has to **open a credit line in advance with a bank or a broker (terza parte fidata)**.
- Then, for each transaction, **the third party has to confirm** that the buyer has the funds to complete the operation.

This “simple” method of payment has several issues:
- **la terza parte deve essere online**
    - Timings: procedures can depend on the opening hours for banks and offices.
    - se la terza parte è offline io non posso procedere con un pagamento/transazione
- **dipendenza dalla terza parte**
    - Costs: each intermediary can demand commissions for the service it provides, even large ones.
    - Data ownership: the user is not the real owner of his/her data or his/her account; 
        - la terza parte sa cosa sto facendo

Cash money has clear advantages compared to credit cards:
- it guarantees full anonymity.
- it does not need third parties.

However, there are other issues:
- we need to “coin money”
    - i.e. a third party that initially distributes the bills and that can be queried against their validity.
- non è digitale, it works only offline.

cio che vorremmo è ottenere il meglio dei due mondi, ovvero:
- vogliamo poter validare le transazioni
- **eliminando la terza parte fidata**
- mantenendo così totale anonimità

_Is it possible to get rid of **trusted third parties** to guarantee the authenticity of documents (vedi CA e certificato) and the validity of economic transactions between parties that **do not trust each other** (PGP non va bene (vabe, non aveva validità legale in ogni caso))?_





### Background sistemi distribuiti e decentralizzati
We can categorise systems according to two criteria:
- where the computation takes place
- and **who makes the decisions**.

Computation
- Centralized: all operations are performed by a single node.
- Distributed: operations are spread across all nodes.

Governance
- Centralized: one single node decides for all.
- **Decentralized: each node participates in the decision-making process**.

#### P2P network
a network in which the computers of connected users act as both client and server.
- Peers make a portion of their resources directly available to other network participants, **without the need for central coordination**.

#### Failures
A failure is a condition that causes a functional unit to fail to perform its required function.
- Benign Failure: when a node stops performing the operations it is supposed to perform.
    - Crash and omissions.
- __Byzantine Failure__: when a node starts to perform arbitrary operations, not present in the correct flow of operations.


## CONSENSO
Distributed processes that have to agree on a single value (e.g., new status of the system).

Properties
- Agreement:
    - If a correct process decides a value, then all correct processes eventually **decide the same value**
- Termination:
    - Every correct process eventually decides some value

Il consenso è un po' quello che vogliamo ottenere con la blockchain, vogliamo che nodi in un sistema decentralizzato, che non si fidano tra di loro, riescano a mettersi daccordo sullo stato delle transazioni senza la terza parte fidata.

**NB**:
- **There is no DETERMINISTIC protocol that solves consensus** (in an asynchronous system where ...)
- But is it often said that **blockchain solves consensus?**
- We will see that this is not actually true. 





...


Sistemi decentralizzati: sistemi in cui ogni nodo di un sistema distribuito partecipa nel prendere decisioni



internet sistema asincrono bizantino




problema dei due generali
- anche nel caso facile di un sistema con fault benigni non abbiamo la certezza se l'altro ha ricevuto i nostri messaggi o meno
- uno dei due generali (processi) ha sempre delle informazioni in più o in meno rispetto all'altro

```perchè a TCP bastano 3 pacchetti e con i due generali andiamo all'infinito?```
- non è che bastano, semplicemente qua non stiamo decidendo della nostra vita/morte
- dopo tre messaggi sono abbastanza sicuro ma non ho garanzia che il server sia presente
    - se al terzo messaggio il server crasha io me ne accorgo solo quando vado in timeout al quarto 


processi corretti = non in crash

**teorema**: non esiste un algoritmo deterministico che permette di fare consenso nel contesto di internet


**blockchain è una soluzione al problema del consenso nel contesto di internet**







### Bitcoin blockchain
Three problems to solve:
1.Where do we store the bills and the transactions?
2.Who coins new money?
3.Who or what gives value to the new currency? (outside the scope of this lecture)

The blockchain is an attempt to answer these questions

Bitcoin (cryptocurrency) -> Blockchain ("technology")

l'idea di bitcoin è quella di un contante che può essere utilizzata online
- l'idea non è nuova ma i predecessori di bitcoin si basavano tutti su una terza parte fidata centralizzata
- la peculiarità di un bitcoin è stata quella di essere **decentralizzato**








## Blockchain
in pratica è un distributed ledger (libro mastro) che registra 

catena di blocchi contenenti dati linkata da hash-pointer

rete peer-to-peer, ogni nodo ha la propria copia della blockchain

blockchain tamper-resistant
- se modifico un blocco devo anche modificare tutti i blocchi precedenti per aggiustare gli hashpointer

l'informazione corretta è quella della maggioranza dei nodi!
- non basta che io modifica la mia copia locale della blockchain!
- **i fatti veri sono dettati dalla maggiornaza dei blocchi**
- L'unica assunzione che fa la blockchain è che il 75%+1 dei peers siano benigni





### Nakamoto consensus
è un algorimo usato nella rete di bitcoin che permette di avere consenso sul prossimo blocco da aggiungere nel contesto peggiore possibile
- L'unica assunzione che fa la blockchain è che il 75%+1 dei peers siano benigni 
- trustless consensus

concetto di proof-of-work
- non serve per fare consenso
- piuttosto serve a scoraggiare i nodi bizantini
- per essere eligible nella posibilità di inserire un nuovo blocco devi risolvere un puzzle computazionalmente costoso
    - minare bitcoin significa risolvere il puzzle ed aggiungere un blocco alla catena

Il puzzle consiste nel trovare un nonce il cui hash è minore del target
- hash non sono invertibili... e quindi forza bruta è l'unica possibilità


NB: i client della blockchain non devono aggiungere blocchi e quindi non devono fare delle POW


...

La catena più lunga è quella che contiene la verità
- è quella su cui è stata fatta più lavoro, ogni blocco della catena ha dovuto superare un POW prima di essere stato aggiunto

Attenzione: le fork non mi danno garanzia che la mia transazione vada a buon fine
- posso finire in un blocco invalido e quindi la mia transazione viene annullata
- per questo bitcoin base non si presta bene come moneta di base

...

Il meccanismo con cui bitcoin risolve le fork, risolve anche il double spending
- aspetto 6 blocchi

conclusione:
- è un meccanismo che protegge dai bizantini
- risolve fork e double spending
- ... ma devo aspettare un'ora per comprare qualcosa


**51% attack**
un nodo con il 51% della rete può recuperare offline l'evoluzione della catena portata avanti dal resto della rete; in questa maniera può invalidare una catena portando ad una double spending





Il problema della blockchain è che per avere un agreement deterministico, ovvero tutti i nodi della rete concordano sullo stesso stato della rete, è necessario che continuino ad essere aggiunti dei nodi
- Se la blockchain si ferma, potremmo avere delle fork (stato inconsistente)
- la blockchain non deve fermarsi mai!
    - il sistema è safe perchè è live





### Ruoli dei peer
...


il wallet rappresenta un'identità
- non contiene bitcoin
- contiene le chiavi private


I bitcoin dove stanno?
- mi bastano delle transazione (Marco ha dato 10bitcoin a mario)
- ...
- fino ad arrivare ad una transazione originaria in cui sono stati creati i bitcoin
- tengo traccia della catena di movimenti che dimostrano che io sono il possessore di una determinata valuta

Di nuovo, la blockchain è un registro (ledger) che tiene traccia di transazioni













### Blockchain | oltre bitcoin
blockchain è un modo per avere una struttura dati condivisa tra dei peer che raggiungono un consenso sullo stato di questi dati
- bitcoin si limita a rappresentare transazioni su questa struttura dati


**smart contract**
(contratto nome fuorviante)

 è un programma scritto in un normale linguaggio di programmazione, **salvato ed eseguito su blockchain**
- tutti i peer hanno una copia di questo programma
- lo stato del programma è concordato dalla rete

Cosa ottengo
- esecuzione distribuita, fault tolerant
    - non ho un server centralizzato su cui esegue il codice
- dati immutabili
- tracing completo
    - ho l'intera storia dell'esecuzione del programma
- disaster recovery, in eventualità di crash lo stato precedente è riproducibile in quanto salvato su blockchain sotto forma di transazioni
