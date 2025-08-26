### innanzitutto ricorda questo | Che cos'è la ridondanza del testo in chiaro? 
Si riferisce al fatto che non tutte le combinazioni di simboli sono equiprobabili o casuali, ma seguono schemi prevedibili, soprattutto nel linguaggio naturale.

Prendiamo il linguaggio italiano:
- Alcune lettere compaiono più spesso di altre:
    - La E è molto più frequente della Z.
- Alcune combinazioni di lettere sono più probabili:
    - Dopo una Q, c’è quasi sempre una U.
- Ci sono parole comuni che si ripetono spesso:
    - "il", "di", "che", "e", "la", ecc.

Questi elementi non sono casuali, ma strutturati. Questo è ciò che si intende per ridondanza. La ridondanza aiuta gli attaccanti nella crittoanalisi!


### Principi alla base dei cifrari simmetrici moderni
Per dare sicurezza computazionale ad un cifrario non perfetto è necessario che la trasformazione di cifratura generi confusione e diffusione (da teoria dell’informazione di C. Shannon).
- **La confusione** si ottiene imponendo al testo cifrato di dipendere in modo complesso dalla chiave e dal testo in chiaro, al punto che è difficile prevedere che cosa accadrà al cifrato modificando anche un solo simbolo del testo in chiaro.
    - La confusione **nasconde la relazione esistente tra testo in chiaro e testo cifrato** e rende poco efficace lo studio del secondo basato su statistiche e ridondanze del primo.
    - La sostituzione polialfabetica è il mezzo più semplice ed efficace per creare confusione.

- **La diffusione** si ottiene imponendo ad ogni simbolo del testo in chiaro di influire sul valore di molti, se non tutti, i simboli del cifrato, al punto che è difficile prevedere quanti e quali di questi modificano il loro valore se si modifica anche un solo simbolo del testo in chiaro.
- La diffusione nasconde la ridondanza del testo in chiaro spargendola all’interno del testo cifrato.
- La trasposizione è il mezzo più semplice ed efficace per ottenere diffusione.

Shannon ha infine indicato il modello del **cifrario composto** come linea guida per conseguire tali proprietà:
- per determinare confusione il cifrario deve avvalersi di operazioni di sostituzione;
- per generare diffusione, deve invece avvalersi di operazioni di trasposizione (o permutazione).

La struttura di un cifrario composto alterna, di conseguenza, stadi di sostituzione (detti S-box) e di permutazione o trasposizione (detti P-box o T-box) per creare confusione e diffusione. 
- per ottenere confusione e diffusione si combinano sostituzione e trasposizione in maniera iterata
- vedi reti di feistel

### Cifrari simmetrici
Tipicamente sono usati per la difesa della riservatezza, ma sono impiegati anche come meccanismi per
- la generazione di numeri pseudocasuali
- l’autenticazione
- l’identificazione.

I cifrari attualmente in uso sono:
- Robusti 
    - resistono a tutti gli attacchi di crittanalisi finora individuati;
- Veloci 
    - nelle realizzazioni con hardware special-purpose elaborano diversi milioni di bit al secondo;
- Efficaci 
    - gestiscono stringhe binarie di lunghezza arbitraria;
- Efficienti 
    - il testo cifrato è praticamente lungo quanto il testo in chiaro. 



Analizziamo due tipi di cifrario:
- **Il cifrario a flusso**:
    - trasforma, con una regola variabile al progredire del flusso di bit (testo…), uno o pochi bit alla volta del testo da cifrare (o decifrare).
    - Questo cifrario punta alla sicurezza dei singoli bit, conveniente sui flussi di dati in tempo reale e seriale (comunicazioni WEP, GSM…).
- **Il cifrario a blocchi**:
    - suddivide in blocchi di bit il testo da cifrare, e per ogni blocco opera una regola di trasformazione fissa.
    - Questo cifrario è tipicamente usato per la protezione di pacchetti, come da 64 o 128 bit (connessioni IP), di file e strutture di dati (comunicazioni IPSec, SFS…).

Vedremo che, tipicamente, è più facile che un cifrario a flusso sia più veloce di un cifrario a blocchi, ma è più facile sbagliare nella progettazione della sicurezza. Ergo, è più facile progettare un cifrario a blocchi sicuro.
- il cifrario a flusso è più veloce del cifrario a blocchi,
- il cifrario a blocchi è più sicuro del cifrario a flusso. 