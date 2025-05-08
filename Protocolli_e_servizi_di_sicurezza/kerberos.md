nato nel 1980

protocollo inizialmente con 3 obiettivi ma di fatto è **rimasto solo l'obiettivo di autenticazione**

non da sapere a memoria, ma bisogna saperci ragionare sopra
- per quale motivo ci sono i timestamp?
- perchè c'è l'autenticatore
- ecc...

non è un'autenticazione federata, ma fortemente centralizzata





**ragionamenti generali**
- dove posiziono il servizio di autenticazione?
    - sulle workstation?
        - le workstation devono tenere traccia della prova di entità dell'utente
        - ci deve essere una relazione di fiducia tra le workstation e i server 
            - e devono essere /tolte se vengono aggiunte/tolte delle workstation/server
        - ci può stare se ci sono poche macchine e quindi l'aggravio ammistrativo del sistema è basso
    - sui server?
        - uhm, stesso problema in realtà


**Idea**:
non ha senso distribuire su tutte le macchine le funzionalità di autenticare gli utenti, utilizziamo un unico server di autenticazione centralizzato a parte.

relazione di fiducia tra i server ed il server di autenticazione.
- I server accettano tutti i messaggi autenticati dal server di autenticazione






Kerberos fa uso di sola crittografia simmetrica per garantire autenticità



dopo essersi identificato (con l'AS), il cliente ottiene un ticket, ovvero una prova di identità, con cui può protrarre autenticare le sue successive richieste ai servizi
- i clienti precondividono con AS i loro segreti che usano per identificarsi
- i servizi precondividono con AS le loro chiavi con cui verrà cifrato il ticket usato per l'autenticazione





primo problema: protocollo passivo. Un primo miglioramento è quindi rendere attiva l'identificazione



obiettivo: identificazione col fine di autenticazione per l'accesso a dei servizi
- anche più volte, non voglio riidentificarmi ogni volta che voglio accedere ad un servizio
- voglio riutilizzare più volte lo stesso ticket




spezziamo le due funzionalità in due servizi separati
- authentication server: identifica (attivamente)
- ticket-granter server: rilascia ticket agli utenti identificati da AS
    - AS e TGS devono avere una relazione di fiducia quindi
    - cosa significa questo se stiamo utilizzando meccanismi simmetrici? AS e TGS precondividono una chiave



... mi sono stancato, altra roba 

- **AUTENTICATORE, associa il ticket al suo legittimo possessore** 
    - concetto utile da generalizzare
    - ad esempio: non basta presentare un certificato, ma devo anche dimostrare di esserne il legittimo proprietario


...



**NB**: Kerberos NON è una sistema di distribuzione delle chiavi. Per certi versi assomiglia in quanto l'AS è centrale, tuttavia esso autentica gli utenti non distribuisce le chiavi. Le chiavi che entrano in gioco in questo servizio devono essere **predistribuite**.


Lo scenario deve essere esteso anche quando richiedo accesso ad un servizio appartenente ad un altro dominio (detto anche reame in kerberos)
- fino ad adesso un utente ha sempre interagito con il proprio AS e TGS
- il mio TGS a questo punto può darmi un ticket per interagire con un TGS remoto il quale si fida di lui

ricorda: **come si creano delle relazioni di fiducia?**
- o con dei certificati (cross certification)
- o nel caso di kerberos con delle chiavi precondivise



scalabilità
- quante chiavi devono essere predistribuite all'AS se ci sono n-utenti?
    - n in un solo dominio
    - n^2 se voglio più domini 