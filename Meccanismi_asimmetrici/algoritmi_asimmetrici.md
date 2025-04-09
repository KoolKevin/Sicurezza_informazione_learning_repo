noi facciamo solo RSA e non ci preoccupiamo degli algoritmi basati sulle curve ellittiche (algoritmi asimmetrici con le prestazioni migliori)

nota: di fatto DH è un cifrario asimmetrico


**OSS IMPORTANTE**: una cifratura asimmetrica è molto meno efficiente di una cifratura simmetrica?
- perchè? dobbiamo fare delle esponenziazioni modulari!
- nel caso simmetrico devo fare solo trasposizioni e sostituzioni


### Come facciamo ad ottimizzare l'esponenziazione modulare?
Algoritmi di teoria dei numeri non sono da sapere per filo e per segno. Tuttavia è importante sapere quali sono gli algoritmi che rendono più efficiente per l'eseponenziazione modulare. Perchè? Nelle implementazioni i vari provider possono utilizzare algoritmi diversi, è desiderabile utilizzare quello più efficiente.
- nel caso di RSA: il tempo di esponenziazione modulare scala con il cubo della dimensione dei bit (1024 -> 2048 -> 8 volte il tempo)

### Come facciamo a generare dei numeri primi molto grandi?
test deterministici -> esponenziali
test probabilistici -> polimoniali ma vanno ripetuti



### RSA
le modalità di cifratura ricompaiono anche in algoritmi asimmetrici

RSA produce un output deterministico
- non ci piace (**capisci il perchè**)
- vogliamo rendere aleatoria un uscita deterministica, come facciamo?
- padding -> OAEP
- non ci interessa come è fatto, ma dobbiamo usare provider che lo implementano

...

per decifrare non faccio la radice n-esima (che è un problema difficile) piuttosto faccio un'altra esponenziazione modulare con dei parametri scelti bene

**NB**: qua la cosa importante è la frammentazione. Siccome ad un unico testo in chiaro deve corrispondere uno e un solo testo cifrato, le operazioni modulari devono essere fatte su dei numeri minori di n -> frammentazione



NOTA: per la decifrazione esiste un particolare algoritmo (di Gartner) che è 4 volte più veloce rispetto a repeated square and multiply -> da tenere a mente nei provider
- questo però richiede di mantenere in memoria i numeri primi p e q generatori di n che dal punto di vista della robustezza sarebbe meglio eliminare dopo avere ottenuto n 

**conclusione**: insomma, la crittografia asimmetrica è sicuramente meno efficiente computazionalmente rispetto a quella simmetrica (devo generare i numeri primi, esponenziazioni modulari, ecc...)



## Cenni sulla sicurezza di RSA


**conclusione 2**: bisogna controllare che il provider che si sta utilizzando impieghi parametri opportuni per efficienza computazionale e contromisure per vari attacchi ad RSA

....



### Cifrario ibrido
nuovo modello per la distribuzione di una chiave simmetrica che si aggiunge a KDC, DH, ecc.

se il messaggio in chiaro è molto più piccolo del modulo, bisogna estendere il messaggio in chiaro con del padding
- altrimenti attacco di forza bruta sullo spazio dei messaggi in chiaro







## Firma digitale
fino ad ora abbiamo visto RSA solamente impiegata nella cifratura dei dati

con RSA posso sia firmare che cifrare, non tutti gli algoritmi ne sono in grado (e.g. alcuni possono solo firmare)
- revesibilità delle chiavi, al posto di passare ad RSA la chiave pubblica, passo la chiave privata ed il procedimento è specchiato



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

# Capisci bene proprietà moltiplicativa di RSA