Gli utenti devono tenere segrete alcune delle elaborazioni che eseguono (la decifrazione, l’attestazione di autenticità, la generazione della prova d’identità).

Per tenere segreta una trasformazione è possibile usare:
- una macchina a funzionamento segreto,
- una macchina nota a tutti, ma dotata di un programma segreto,
- una macchina nota a tutti, un programma noto a tutti ed un parametro segreto da fornire prima che inizi la sua esecuzione. 

Macchine a funzionamento segreto
- manutenibilità? scalabilità? ispezione?

Algoritmi segreti
- come sopra

**Algoritmi noti con parametri segreti**
- Decisamente più sicuro è il principio di usare algoritmi di dominio pubblico e di tenere segreto soltanto un loro **parametro, detto chiave**.
- se viene scoperto il parametro lo cambio e lascio inalterate macchina ed algoritmo

```
d'ora in poi con trasformazioni segrete intenderemento algoritmo noto con parametro segreto (segretezza del parametro della trasformazione)
```

### Cifrari simmetrici e asimmetrici
- simmetrica:   k_s == k_d (o facilmente derivabile)
    - segretezza totale della chiave
    - come le distribuisco? entrambi le devono avere
    - tipico per riservatezza

- asimmetrica:  k_s != k_d 
    - k_d (chiave pubblica; può essere nota)
    - ma non si può derivare k_s da k_d (mentre il contrario è facile)
    - tipico per autenticazione

![alt text](immagini/cifrari_simmetrici_e_asimmetrici.png)

- lo spazio delle chiavi N deve essere molto grande
    - in questo modo attacchi di forza bruta non riescono a trovare la chiave giusta

**proprietà delle chiavi simmetriche**
Nei meccanismi simmetrici si ha di norma k=ks=kd. La chiave k, usualmente detta condivisa dato che deve essere preventivamente e segretamente concordata da ogni coppia di corrispondenti, seleziona la trasformazione che ciascuno di essi deve operare. 

Quattro sono le proprietà di sicurezza che gli utenti devono assicurare alla chiave condivisa: 

1. autenticità: quando si concorda il valore della chiave in modo telematico occorre essere ben certi che il corrispondente è proprio quello voluto;
2. segretezza: l’accordo sulla chiave deve avvenire su un canale a prova di intercettazione; la chiave deve poi essere memorizzata su un supporto a prova di
intrusione;
3. integrità: la comunicazione della chiave e la sua successiva memorizzazione devono rilevare la presenza di errori;
4. robustezza: gli algoritmi di crittanalisi che individuano la chiave devono essere computazionalmente infattibili. 

**proprietà delle chiavi asimmetriche**
Nei meccanismi asimmetrici i due differenti valori di chiave da usare alle due estremità del canale insicuro sono preventivamente **generati da uno solo dei due corrispondenti**. Ogni utente U di un algoritmo asimmetrico ha dunque due chiavi:
- la chiave privata SU che U impiega, a seconda del tipo di algoritmo:
    - o per decifrare un messaggio riservato
    - o per autenticare un suo documento
    - o per provare la sua identità

- la chiave pubblica PU che U mette a disposizione di chiunque debba:
    - o inviargli messaggi riservati
    - o verificare l’autenticità dei suoi messaggi.

**NB**: Per un uso sicuro di ogni algoritmo asimmetrico, la funzione PU = f(SU) deve essere unidirezionale. 

le proprietà di sicurezza della chiave privata SU sono tre:
- segretezza: il valore ed il modo con cui è stato ottenuto devono essere tenuti riservati;
- integrità: la memorizzazione deve essere a prova d’errore;
- robustezza: gli algoritmi in grado di calcolare la funzione SU = f^-1(PU) devono essere come minimo sub-esponenziali e la dimensione della chiave pubblica PU ne deve rendere impossibile l’esecuzione.

Le proprietà di sicurezza della chiave pubblica PU sono due:
- autenticità: tutti devono poter essere certi che PU è proprio di U,
- integrità: la memorizzazione deve essere a prova di errore. 









## Crittanalisi
Un'attaccante può provare a fare tre cose quando non ha una chiave:
1. Indovinare
2. Intercettare
3. Dedurre

### Indovinare
La sicurezza dipende da come è stato scelto il segreto (chiave). Se l’utente, per potersela facilmente ricordare a memoria, ha usato un’informazione che l’intruso può prevedere (ad esempio il suo nome, o il suo anno di nascita, o i dati di un suo familiare, ecc.), il tirare ad indovinare ha una notevole probabilità di successo.

La prima difesa preventiva è quindi quella di usare dati casuali.
- R12: ”I bit della stringa che rappresenta un dato segreto devono essere equiprobabili ed indipendenti ”

Abbiamo già anche commentato il fatto che la complessità computazione non mette in conto il caso più favorevole per l’intruso; utile è dunque la contemporanea adozione di un secondo provvedimento.

- R13: ”un dato segreto deve essere frequentemente modificato”. 
    - per lo stesso motivo, usare tante volte lo stesso segreto per molti messaggi non è raccomandabile


### Intercettare | memorizzazione del parametro segreto:
Una lunga stringa di bit aleatori è difficile da imparare a memoria. Ogni utente ha dunque bisogno di una memoria, in cui alloggiare tutti i suoi segreti e da cui poter di volta in volta estrarre quello che serve al meccanismo da mettere in esecuzione. 

Per valutare la sicurezza di questa soluzione, consideriamo il caso di una memoria non sicura M, in cui si deve conservare una chiave segreta s, e di un processore sicuro P, in cui deve essere eseguito un algoritmo Tk che trasforma un dato x nel dato y = Ts(x). 

- R15: ”una chiave segreta deve essere mantenuta in memoria in una forma cifrata stabilita dal proprietario”.
- R16: ”la forma cifrata di s deve essere mantenuta anche nel trasferimento su bus da Memoria a Processore”
- R17: ”P deve poter decifrare s solo se il proprietario gli comunica la chiave che ha scelto per cifrarla; una volta usato il testo in chiaro s, P deve cancellarlo accuratamente dalla sua memoria di lavoro ”

il segreto serve a cifrare, ma anch'esso dev'essere cifrato siccome va memorizzato... come fare?
- si utilizza una chiave di secondo livello *u* per cifrare le chiavi di primo livello
- da un lato *u* deve essere casuale, da un altro lato facilmente imparabile a memoria siccome non viene memorizzata
- Entrambe le proprietà vengono di norma conseguite impiegando l’impronta di una frase, la cosiddetta passphrase, che l’utente impara a memoria e comunica al processore su un canale protetto. 

Output cifrato con presenza di pattern ripetitivi da informazioni sfruttabili dall'intrusore
- per questo motivo è importante che gli algoritmi crittografici producano un'uscita il più possibile aleatoria

In un file system sicuro, un utente ha, infatti, la necessità di gestire una **gerarchia di tre dati segreti**.

- Al primo livello della gerarchia dei segreti c’è la **pass phrase** che l’utente deve imparare a memoria e che impiega per decifrare i segreti di secondo livello (accesso al portachiavi). Il meccanismo chiamato in causa è la funzione hash.

- Al secondo livello ci sono, cifrate dalla pass phrase, le chiavi (in generale sono più di una) che l’utente impiega per accedere alle chiavi di cifratura dei file. Ovvero le chiavi alle chiavi che proteggono la riservatezza dei documenti riservati di terzo livello. La memoria per le chiavi che cifrano chiavi è un file, usualmente detto **portachiavi**
(keyring). In questo servizio sono coinvolti un RNG ed un cifrario simmetrico.

- Al terzo livello vi sono le chiavi di cifrazione dei file. Ogni file è cifrato/decifrato con una chiave del portachiavi scelta di volta in volta dal RNG 


Posizione della memoria contenente la chiave cifrata:
- Hard Disk: poco sicuro perché lo stesso calcolatore può essere usato da più utenti;
- Memory Card: la si inserisce nel sistema solo quando vi è necessità del dato segreto;
- Smart Card: molto sicura, consiste nell’avere un calcolatore portatile con una sua memoria contenente s cifrata, a cui sistema invia file x ed ottiene il file cifrato y, in questo modo s non esce mai dal suo ambiente protetto.


### Dedurre
Per gli attuali meccanismi crittografici l’impossibilità di deduzione è sempre verificata.

Un accorgimento ha un ruolo fondamentale nell’ottenimento di “robustezza” agli attacchi: “il legame tra il testo in chiaro ed il testo cifrato deve essere così ben nascosto da far apparire quest’ultimo all’intruso come una variabile aleatoria che assume con uguale probabilità tutti i suoi possibili valori”.

Con questa ipotesi sembrerebbe lecito concludere che l’intruso abbia a disposizione soltanto la forza bruta, attacco che abbiamo già visto poter essere fronteggiato attribuendo alle variabili aleatorie un’adeguata dimensione. In realtà le cose non stanno proprio cosi. Un intruso che ha a disposizione un numero elevato di testi cifrati può condurre attacchi più efficaci mettendo in conto le proprietà statistiche delle variabili aleatorie (può più o meno indovinare qualcosa). 



### Complessità computazionale
...


Nel contesto della sicurezza informatica devo considerare anche il caso migliore
- bisogna adottare degli accorgimenti per evitare che l'intrusore si trovi di fronte ad istanze facili del problema
- e.g. segreto con delle parti che sia ripetono


**unità di misura di robustezza**:
- Anno MIPS -> obsoleto
- Livello di sicurezza (che nome è?) -> numero di operazioni per un'attacco di forza bruta con n bit

128 bit minimo livello di sicurezza



**livello di sicurezza in un cifrario asimmetrico**:
l’attacco che ha successo in questo caso è quello che, conoscendo la chiave pubblica, **inverta la funzione pub = f(priv)**.

Basterebbero 128 bit per avere una chiave privata computazionalmente sicura, ma **per una chiave pubblica servono almeno 2000 bit per rendere difficile l’inversione della funzione**. Dunque, **per simmetria**, entrambe le chiavi saranno di almeno 2000 bit. 