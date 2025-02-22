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

### Ciclo di vita di un dato
1. comunicazione
2. memorizzazione
3. elaborazione

### CIA triangle
Obiettivo per la sicurezza dei dati: ci sono tre proprietà fondamentali da garantire
1. Confidenzialità
    - solo chi è autorizzato a farlo, può accedere ai dati (leggere, visualizzare, stampare, semplice conoscenza dell’esistenza)

2. Integrità
    - assenza di alterazione durante l'intero ciclo di vita del dato

3. disponibilità

**FONDAMENTALE**: bisogna garantire solamente le proprietà richieste dal problema, niente di più niente di meno. Eventuali proprietà aggiuntive si pagano


"La sicurezza informatica ha lo scopo di proteggere le risorse da accessi indesiderati (integrità), garantire la riservatezza delle informazioni, assicurare il funzionamento e la disponibilità dei servizi a fronte di eventi imprevedibili (disponibilità)"





## Terminologia:
- **Vulnerabilità**: punto debole di un sistema che può rendere realizzabile una **minaccia**.
    - Non è necessariamente un difetto, può essere una caratteristica intrinseca del sistema **ineliminabile**
    - Chiaramente, può anche nascere da un implementazione sbagliata. In questo caso la vulnerabilità è risolubile

- **Minaccia**: un atto ostile, intenzionale o meno, che ha un qualsiasi effetto negativo sulle risorse o sugli utenti del sistema
    - noi lo considereremo sempre intenzionale, c'è un malintenzionato
    
- **Attacco**: qualsiasi azione che sfrutta una vulnerabilità per concretizzare una minaccia
    - **Superficie di attacco**: rappresenta tutte le possibili vulnerabilità di un sistema, ovvero i punti che un attaccante potrebbe sfruttare.
        - Più è ampia la superficie, maggiore è il rischio.
    - **Vettore di attacco**: la tecnica o il metodo che un attaccante utilizza per sfruttare una specifica vulnerabilità all'interno della superficie di attacco.
        - Ad esempio, il phishing è un vettore di attacco che sfrutta la vulnerabilità umana (ingegneria sociale).

- **Contromisure**: azione, dispositivo, procedura o tecnica che consente di **rimuovere o di ridurre una vulnerabilità**

- es ponte:
    - vulnerabilità == crepe
    - minaccia      == il crollo del ponte quando ci cammino sopra 
    - attacco       == terremoto che fa crollare il ponte instabile a causa delle crepe / uno col bazooka che spara sulle crepe
    - contromisure  == riparare le crepe

**Mercato delle Vulnerabilità**
zerodium: azienda broker di vulnerabilità. Compra le vulnerabilità trovate da ricercatori/hobbyists e le rivende
- crazy 


### Esempi di attacchi:
- **sniffing e snooping**:
    - due tecniche di intercettazione di dati nella sicurezza informatica, spesso utilizzate per spiare le comunicazioni in una rete.
    - **Lo sniffing** è il processo di cattura e analisi del traffico di rete. Può essere usato per scopi legittimi, come il monitoraggio delle prestazioni della rete, ma anche per attività malevole, come il furto di credenziali o dati sensibili.
        - Esempi di strumenti: Wireshark, tcpdump.
        - Modalità di attacco: l'attaccante può intercettare il traffico su reti Wi-Fi non protette o eseguire un attacco Man-in-the-Middle (MitM).
    - **Lo snooping** si riferisce alla sorveglianza o al monitoraggio non autorizzato di dati, spesso a livello di dispositivi o sistemi. Può avvenire su e-mail, file o conversazioni di rete.
        - Esempi: un amministratore di sistema che spia dati degli utenti senza permesso, o un malware che raccoglie informazioni dai file.
        - Può essere fatto attraverso software malevoli, keylogger o accessi non autorizzati.
    - Differenza principale:
        - Lo sniffing è focalizzato sull’intercettazione del traffico di rete.
        - Lo snooping riguarda l'accesso non autorizzato ai dati di un sistema o dispositivo.
    - Entrambe le tecniche possono essere prevenute con crittografia (HTTPS, VPN), firewall e segmentazione della rete.

- **spoofing**:
    - tecnica di attacco informatico in cui un attaccante si maschera da entità fidata per ingannare un sistema o un utente e ottenere accesso a dati sensibili o causare danni. Il termine significa letteralmente "falsificazione" o "camuffamento".
    - Tipi di Spoofing più comuni
        - IP Spoofing: L'attaccante falsifica il proprio indirizzo IP per sembrare un altro dispositivo, spesso per eludere firewall o lanciare attacchi DDoS.
        - DNS Spoofing:  Modifica le risposte DNS per reindirizzare gli utenti a siti falsi, spesso usati per phishing o diffusione di malware.

**syn flooding**:
esempio interessante perchè fa vedere come un protocollo ben costruito ai fini delle comunicazione, progettato in un mondo fidato privo di attacchi informatici.
- presuppone che il client risponda correttamente (vulnerabiltà intrinseca ineliminabile del protocollo TCP)
- se il client non risponde mai/ cambia il suo ip, il buffer del server può andare in overflow





### Classificazione vulnerabilità
- hardware
- Software/firmware
- Dati
- Linee di comunicazione

### Classificazione minacce:
- disimpegno
- ingenuità
- disastri
- errori
- attacchi intenzionali (detti anche exploit)

### Classificazione contromisure
- prevenzione
- rilevazione
- reazione

**contromisure**:
Ogni contromisura ha un costo, efficacia ed effetti negativi collaterali (producono altre vulnerabilità)
- tradeoff: efficiacia / costo di implementazione e degli effetti collaterali


Interessante il modello che rappresenta su 3 dimensioni le caratteristiche da considerare quando si progetta un sistema sicuro
- (strumenti == contromisure)
- data una vulnerabilità quale contromisure e quali proprietà vanno garantite?
- vulnerabilità durante trasmissione, memorizzazione ed elaborazione






### Oltre CIA
**Autenticazione dell'utente e del dato** (ricadono nell'unico termine autenticazione):
- quando la destinazione riceve un dato, deve sapere l'identità di chi l'ha inviato (autenticazione del dato)
    - NB: chi invia il dato non è necessariamente chi lo ha prodotto
- Autenticazione di chi ha inviato il dato intesa come meccanismo del controllo dell'accesso al canale di invio
    - strettamente collegato a sopra
    - deve esistere una lista degli utenti (o dei gruppi di utenti) potenziali ed un servizio per l’identificazione sicura degli utenti; 

**NON ripudio**:
- chi riceve un dato non può appropriarsi della paternità / non gli si può attribuire la paternità
- chi crea/invia un dato non può disconoscerne la paternità
- il senso è che, a posteriori, sono in grado di attribuire con certezza chi ha creato/inviato il dato (es. contratto di acquisto di una casa)





## Analisi del rischio
bisogna ...

### Fasi 
...
- analisi delle dipendenze di funzionamento: un pen-test può propagarsi sfruttando il grafo di dipendenze
- rbac piuttosto che ibac
- catagolazione degli eventi indesiderati: come faccio a sapere quali sono gli eventi indesiderati? cataloghi ben noti (XIRT)






## IPOTESI FONDAMENTALE del CORSO
Ci concentriamo sulla sicurezza del dato presupponendo un **calcolatore sicuro**, cosa signifca?

"In un sistema sicuro (trusted) devono dunque essere inclusi fin dall’inizio, od eventualmente aggiunti successivamente, **meccanismi e servizi** in grado di controllare il **corretto funzionamento del HW**, di verificare **l’integrità e l’origine del SW**, di rilevare e bloccare ogni uso potenzialmente pericoloso di queste risorse."

Interessante l'uso di un Coprocessore:
- con memoria e processore
- in cui vengono elaborate e salvate le informazioni sensibili (e.g. chiavi crittografiche)





**Meccanismo vs Servizio**
- Meccanismi di sicurezza:  svolgono azioni protettive per **singole e specifiche vulnerabilità**
    -  focalizzato su una contromisura per una vulnerabilità

- Servizi di sicurezza: che sfruttano in generale più meccanismi per proteggere una o più proprietà critiche dell’informazione
    - integrano più meccanismi