
## Cifrari a blocco
cifrano blocchi di dimensione predefinita e non singoli bit alla volta
- padding di nuovo presente (e sempre presente con anche un intero blocco di padding)

la chiave è la sempre la stessa per ogni blocco





**rete di Feistel**
garantisce confusione e diffusione mediante vari round che applicano sostituzione e trasposizione



A noi interessa lo standard attuale: **AES**

non adotta il modello di Feistel ma adotto anche lui sostituzione e trasposizione


skip di roba.


## Modalità di cifratura | parte che interessa di più!
un cifrario a blocchi non può essere utilizzato solo come visto fino ad adesso (ECB)

Usare ECB non va bene (in generale):
- **determinismo di ECB**
- **malleabilità di ECB**:
    - se uso sempre la stessa chiave posso sostituire il blocco del destinatario con il blocco in cui sono io il destinatario
- se il messaggio è molto piccolo (sta dentro ad un blocco) ECB va vene
    - utilizzato per cifrare chiavi

Altra considerazione sulle modalità di cifratura: **come si propagano gli errori?**
- con ECB non si propaga (l'errore rimane confinato nel blocco modificato)

Le altre modalità diverse da ECB hanno come obiettivo rendere aleatoria l'uscita della cifrazione, ovvero, dato lo stesso messaggio/blocco in chiaro non voglio ottenere lo stesso messaggio cifrato


**LA MONTANARI CHIEDE GLI SCHEMI DELLE MODALITÀ DI CIFRATURA**

**Cypher Block Chaining**:
- pipeline
- che caratteristiche deve avere il v_0? deve essere usato una e una sola volta altrimenti blocchi uguali appartenenti a messaggi in generale diversi vengono cifrati allo stesso modo. Inoltre deve essere casuale e imprevedibile
    - per questi motivi i vettori di inizializzazione vengono generati con PRNG
    - questo requisito non viene sempre rispettato (TLS/SSL)

- efficienza
    - in cifratura non parallelizzabile
    - in decifatura è parallelizzabile

- propagazione degli errori: 
    - **in decifratura**, se un blocco viene alterato, si modifica solo il blocco seguente (vedi esempio)
    - si propaga a tutti i blocchi successivi




**Cypher Feedback e Output Feedback**:
ricordano i cifrari a flusso (sincrono e autosincronizzante)


nello schema OFB il vettore iniziale deve essere unico per DUE MOTIVI
- uno per aleatorietà del cifrato
- l'altro se no ricadiamo nel caso in cui si usa una chiave due volta (vedi xor dei cifrati) il che mina alla riservatezza del messaggio

c'è propagazione dell'errore


**counter**
molto efficiente dato che è parallelizzabile

non ho propagazione dell'errore


# Beast attack
si basa sulla prevedibilità del vettore di inizializzazione... ma non solo! L'attaccante deve poter:
- intercettare il traffico
- e iniettare dei pacchetti specifici nel flusso dati di una sessione già attiva
- o, equivalentemente, induce la sorgente a cifrare una cosa che vuole lui (non facile)

la vulnerabilità dipende da:
- chiave formata da parte fissa e da parte variabile (vettore di inizializzazione pseudo-casuale)
- ogni messaggio applicativo prevede un chiave diversa
- ma nella frammentazione dei protocolli di trasporto, il singolo pacchetto viene cifrato in sequenza con gli altri utilizzando gli ultimi (8)byte del pacchetto precedente come vettore di inizializzazione per la cifratura del pacchetto corrente
- **il vettore di inizializzazione non è più imprevedibile! Lo vedo...** 

...

il cifrato rimane uguale

proprità dell'xor mi permettono di costruire un messaggio in chiaro opportuno che mi permette di risalire al passo 3 l'output del passso 2
- se i due cifrati coincidono, la mia ipotesi di attaccante è corretta!



A livello applicativo è tutto garantito! Ma i progettisti non hanno considerato la prevedibilità a livello di trasporto...




### Paradosso del compleanno e cifrari a blocchi di 64 bit
un cifrario a blocchi è sicuro se e solo se i blocchi hanno dimensione di 128 bit

La probabilità che due blocchi di testo cifrato siano uguali scala con 2^(n/2)