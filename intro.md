I lucidi non sono materiale di studio...

l'orario delle lezioni cambia spesso, controlla

policy di ateneo sull'uso di AI generativa

## Cybersecurity
Paradigma centralizzato
- perimetro ben definito

Paradigma decentralizzato
- perimetri spariscono


**Morris Worm**
1988, prima di allora non esisteva il concetto di attacco informatico

Primo attacco informatico che fà rendere conto la comunità che bisogna pensare alla sicurezza informatica.
- da luogo alla nascita dei CERT (Computer Emergency Response Team)

Capiamo cosa significa sicurezza:
- che cosa vogliamo proteggere?
    - più qualcosa è prezioso, meglio lo vogliamo proteggere. 
    - chi fa sicurezza informatica deve fare anche un **analisi di valore/rischio** 
- da chi?
    - fondamentale il **modello di minaccia** di un protoccolo/sistema
- perchè?
- come?

Sicurezza informatica è un campo ampissimo:
- sicurezza dell'hardware
- sicurezza del firmware
- sicurezza del sistema operativo
- sicurezza del software
    - la maggior parte delle vulnerabilità sono dovute a software scritto male
- **sicurezza dei dati** (focus del corso)
- **sicurezza delle reti** (focus del corso)

### CIA triangle
Obiettivo per la sicurezza dei dati: ci sono tre proprietà fondamentali da garantire
- confidenzialità
- integrità
    - assenza di alterazione durante l'intero ciclo di vita del dato (generazione, trasmissione, memorizzazione)
- disponibilità

**FONDAMENTALE**: bisogna garantire solamente le proprietà richieste dal problema, niente di più niente di meno. Eventuali proprietà aggiuntive si pagano


### Terminologia:
- **Vulnerabilità**: punto debole di un sistema che può rendere realizzabile una minaccia. Non è necessariamente un difetto, può essere una caratteristica intrinseca del sistema ineliminabile
    - Chiaramente, può anche nascere da un implementazione sbagliata. In questo caso la vulnerabilità è risolubile
- **Minaccia**: 
    - intenzionale
    - differenza rispetto a vulnerabilità; es ponte: crepe == vulnerabilità; minaccia == la mia sicurezza quando cammino sul ponte che può crollare se attaccato
- **Attacco**: sfruttare una vulnerabilità per concretizzare una minaccia
    - vettore di attacco
- **Contromisure**:



**Vulnerabilità**
zerodium: azienda broker di vulnerabilità. Compra le vulnerabilità trovate da ricercatori/hobbyists e le rivende
- crazy 

classificazione ...

Esempi di attacchi:
- sniffing e snooping:
    - due tecniche di intercettazione di dati nella sicurezza informatica, spesso utilizzate per spiare le comunicazioni in una rete.
    - Lo sniffing è il processo di cattura e analisi del traffico di rete. Può essere usato per scopi legittimi, come il monitoraggio delle prestazioni della rete, ma anche per attività malevole, come il furto di credenziali o dati sensibili.
        - Esempi di strumenti: Wireshark, tcpdump.
        - Modalità di attacco: l'attaccante può intercettare il traffico su reti Wi-Fi non protette o eseguire un attacco Man-in-the-Middle (MitM).

    - Lo snooping si riferisce alla sorveglianza o al monitoraggio non autorizzato di dati, spesso a livello di dispositivi o sistemi. Può avvenire su e-mail, file o conversazioni di rete.
        - Esempi: un amministratore di sistema che spia dati degli utenti senza permesso, o un malware che raccoglie informazioni dai file.
        - Può essere fatto attraverso software malevoli, keylogger o accessi non autorizzati.

    - Differenza principale:
        - Lo sniffing è focalizzato sull’intercettazione del traffico di rete.
        - Lo snooping riguarda l'accesso non autorizzato ai dati di un sistema o dispositivo.
    - Entrambe le tecniche possono essere prevenute con crittografia (HTTPS, VPN), firewall e segmentazione della rete. 🚀

- spoofing:
    - tecnica di attacco informatico in cui un attaccante si maschera da entità fidata per ingannare un sistema o un utente e ottenere accesso a dati sensibili o causare danni. Il termine significa letteralmente "falsificazione" o "camuffamento".
    - Tipi di Spoofing più comuni
        - IP Spoofing: L'attaccante falsifica il proprio indirizzo IP per sembrare un altro dispositivo, spesso per eludere firewall o lanciare attacchi DDoS.
        - DNS Spoofing:  Modifica le risposte DNS per reindirizzare gli utenti a siti falsi, spesso usati per phishing o diffusione di malware.

**syn flooding**:
esempio interessante perchè fa vedere come un protocollo ben costruito ai fini delle comunicazione, progettato in un mondo fidato privo di attacchi informatici.
- presuppone che il client risponda correttamente (vulnerabiltà intrinseca ineliminabile del protocollo TCP)
- se il client non risponde mai/ cambia il suo ip, il buffer del server può andare in overflow

**minacce**:
classificazione...


**contromisure**:
tradeoff: efficiacia/costo di implementazione

classificazione...

modello che rappresenta su 3 dimensioni le caratteristiche da considerare quando si progetta un sistema sicuro (strumenti == contromisure)
- data una vulnerabilità quale contromisure e quali proprietà vanno garantite?
- vulnerabilità durante trasmissione, memorizzazione ed elaborazione

### Oltre CIA
Autenticità del dato:
- quando la destinazione riceve un dato, deve sapere l'identità di chi l'ha inviato
- NB: chi invia il dato non è necessariamente chi lo ha prodotto
- Autenticazione di chi ha inviato il dato: collegato strettamento a sopra; 

NON ripudio:
- chi riceve un dato non può appropriarsi della paternità
- chi crea/invia un dato non può disconoscerne la paternità
- il senso è che, a posteriori, sono in grado di attribuire con certezza chi ha creato/inviato il dato






## Analisi del rischio



### Fasi 

...

- analisi delle dipendenze di funzionamento: un pen-test può propagarsi sfruttando il grafo di dipendenze
- rbac piuttosto che ibac
- catagolazione degli eventi indesiderati: come faccio a sapere quali sono gli eventi indesiderati? cataloghi ben noti (XIRT)



## IPOTESI FONDAMENTALE
Ci concentriamo sulla sicurezza del dato presupponendo un **calcolatore sicuro**, cosa signifca?
- hardware sicuro
    - bisogna garantire integrità del sistema, controllare gli accessi e rilevare le manomissioni
- SO sicuro


Coprocessore
- con memoria e processore
- in cui vengono elaborate e salvate le informazioni sensibili (e.g. chiavi crittografiche)


















entriamo nel vivo del corso

**Meccanismo vs Servizio**
meccanismo: focalizzato su una contromisura per una vulnerabilità
servizio: integra più meccanismi


**Modello del canale insicuro**
ambiente sicuro lato sorgente e destinazione del dato

il canale di comunicazione è insicuro nel senso che una terza entità (che può essere sempre la sorgente/destinazione)

**generazione, trasmissione, memorizzazione -> come fanno ad essere generazione e memorizzazione non sicure???**

attacco passivo: non altera il flusso dei dati. Osserva solo il traffico dei dati
attaco attivo: può alterare il flusso dei dati.


...


controllo dell'accesso al canale infattibile (non scalabile) se il canale è internet
- cosa significa controllare l'accesso al canale? 

una contromisura preventiva che vale sempre è il controllo dell'accesso 

per attacchi attivi -> rilevazione
per attacchi passivi -> prevenzione

dove si collocano i meccanismi/servizi nella pila protocollare ISO/OSI?
- livello 3: ipsec
- livello 4: SSL
- livello 7: a mano

posizionarsi a livelli diversi fa differenza su trasparenza e personalizzazione della cifratura ad esempio