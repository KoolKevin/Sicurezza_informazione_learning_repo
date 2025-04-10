### BASI (molto importante)
obiettivo:
- avere la certezza che al termine del protocollo chi ha dichiarato una certa identità sia confermato che sia effettivamente chi dice di essere

requisiti:
- ...

proprietà da scegliere se garantire o meno in base all'implementazione
-

### Dimostrazione di conoscenza

**protocolli attivi e passivi**
- passivo (o debole): uso sempre lo stesso segreto come prova di identità (conoscenza)
- attivo (o forte): la prova di conoscenza cambia da sessione di identificazione all'altra
    - chiaramente più robusto rispetto a protocolli passivi


### Identificazione passiva
cringe

in locale anche accettabile

nei sistemi in rete inaccettabile, è sempre possibile l'attacco con replica



### Identificazione attiva
ci concentriamo su questa

a partire dalla prova di identità (intercettabile come la pwd) deve essere difficile risalire al segreto che l'ha generata

**one-time pwd**
...

si applica una funzione F che produce le prove di identità più volte per produrre N prove

l'attacco con replica non è più possibile dato che il verificatore cambia le prove di conoscenza ad ogni identificazione

robusto, ma ad un certo punto le prove finiscono -> necessità di rifare tutto il setup



**sfida/risposta**
...


se il nonce è prevedibile l'attaccante può sosituirsi a B e farsi mandare H(S || RB'). A questo punto l'intrusore può spacciarsi per A con B nella prossima sessione dato che conosce l'hash che verrebbe usato da A alla sfida con RB'. 


### Identificazione mutua
...

esempio non robusto
- tengo due sessioni attive

foto sul telefono





### PARLEREMO DI VPN?



