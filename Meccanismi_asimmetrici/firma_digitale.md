## Firma digitale
fino ad ora abbiamo visto RSA solamente impiegata nella cifratura dei dati

con RSA posso sia firmare che cifrare, non tutti gli algoritmi ne sono in grado (e.g. alcuni possono solo firmare)

RSA gode dellla **proprietà di revesibilità** delle chiavi
- al posto di passare ad RSA la chiave pubblica, passo la chiave privata ed il procedimento è specchiato
- non tutti gli algoritmi asimmetrici hanno questa proprietà


schema di firma con recupero e schema con appendice




5 requisiti:
- di nuovo i termini identificazione e autenticazione mi risultano ambigui. Qua il primo requisito riguarda autenticazione del firmatario e quindi cio che lo garantisce è una PKI
- in questo caso, senza certificazione, la firma consente a chiunque di identificare univocamente NON il firmatario, piuttosto la chiave segreta utilizzata
- ricorda che identificazione ha un aspetto istantaneo


**firma con recupero è inammissibile e non ha validità legale**
- non valgono i requisiti 3 e 5
- posso fare un collage di blocchi firmati presi da tanti documenti firmati dal malcapitato

**NB**: RSA utilizza nativamente firma con recupero, se voglio i requisiti 3 e 5 devo usare qualcos'altro... or do i?
- se utilizzo RSA su u blocco piccolo allora ottengo praticamente uno schema con appendice, di conseguenza posso usare RSA sopra ad un hash 


# CHIEDI A CHATGPT dettagli su questo algoritmo di firma di un certificato: PKCS #1 SHA-256 con crittografia RSA

**proprietà moltiplicativa di RSA**
in sostanza E(m1\*m2) == E(m1)\*E(m2); più precisamente si dovrebbe dire che sono congrui


### Autenticazione di un messaggio oscurato
possiamo sfruttare la proprietà moltiplicativa in uno schema che si chiama firma cieca.

- autenticazione da parte di una **terza parte** di un messaggio firmato da me senza che la terza parte possa conoscerne il contenuto
- la terza parte mi firma un messaggio che non vede in maniera da autenticarlo al destinatario

siccome la terza parte firma, e quindi si assume una responsabilità, solitamente la terza parte impiega anche un protocollo di identificazione per capire chi è che mi sta chiedendo di firmare della roba. 

Altrimenti vulnerabilità! Vedi attacchi alla firma con proprietà moltiplicativa

...

Regola: quando uso una coppia di chiavi per fare firma cieca è bene che venga utilizzata solo nel contesto applicativo legato alla firma cieca e non anche da altre parti... capisci il perchè...

attacco 2: ...