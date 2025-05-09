TLS standard di SSL

è un insieme di protocolli (suite) non uno unico
- **handshake**
- **record**
- alert
- change cipherspec


non è completamente trasparente
- quando apro una secure socket devo fare delle chiamate specifiche
- devo quindi modificare eventualmente le mie applicazioni

obiettivi:
- confidenzialità
- integrità
- autenticazione
    - di cosa?
    - del server presso il clienti
    - ma anche viceversa


La sessione
- viene creata con il protocollo di handshake
- tiene traccia dei parametri di sicurezza
    - cifrari
    - chiavi
    - ...
- i parametri di sessione possono essere condivise su più connessioni
- una sessione può avere più connessioni (tipicamente però una)


A sessione e connessione sono associate delle strutture dati
- stato di una sessione
    - ...
    - segreto principale, a partire dal quale vengono creati i segreti per la confidenzialità di ogni connessione
- stato di una connessione
    - ...
    - vettore di inizializzazione per eventuali modalità di cifratura che lo richiedono




### Protocollo di handshake (negoziazione)
come al solito non bisogna studiare a memoria il protocollo, piuttosto bisogna capire come mai si sono scelti quei messaggi


1^ fase:

per lo scambio delle chiavi (simmetriche) ho 3 alternative:
- Diffie-Helman
    - ephemeral
    - fixed
    - anonymous anche no
- KDC
- Cifrario ibrido (cifro la chiave segreta che ho generato per conto mio con la chiave pubblica di quello a cui la voglio mandare)

**per l'autenticazione dei pacchetti ho tre alternative**
- HMAC
    - In SSL viene scelta questa strategie in quanto fare un hash è più efficiente rispetto al fare una firma
    - tuttavia, questa alternativa **è ripudiabile** rispetto al   
- firma
- MAC simmetrico (??? boh, lascia stare)

2^ fase:

autenticazione del server e invio di chiavi

3^ fase:

autenticazione del client

4^ fase:

...


fortezza da lasciar stare (roba proprietaria di SSL che nello standard TLS è sparito)

...


durante i messaggi di server key exchange e certificate, dipendentemente da quali algoritmi, metodo di scambio delle chiavi, ecc. ho negoziato dentro al certificato e nei successivi messaggi risiedono parametri diversi


...


il protocollo di negoziazione termina con i messaggi finished, questi si assicurano dell'integrità e dell'autenticità di quanto scambiato e del segreto concordato.


capiamo: quand'è che il client si accorge che è presente la PoP del certificato del server (potrebbe avere mandato il certificato di un altro)?
- i casi sono 3.
    - ho negoziato cifrario ibrido
        - allora me ne accorgo in client key exchange vedendo se il server riesce ad estrarre il segreto
    - ho negoziato ephemeral DH
        - server key exchange
    - ho negoziato fixed DH
        - qua non c'è una prova di possesso esplicita
        - me ne accorgo in fase 4 quando verifico di aver scambiato lo stesso segreto
        - se il server non è chi dice di essere **la fase 4 non si chiude** (per questo è importante)





Se uso solo il protocollo di handshake sono vulnerabile a redirezioni (di cui non mi accorgo)
- il certificato che ricevo è valido e c'è anche la prova di possesso

Nel protocollo di negoziazione non c'è nessuna operazione che verifica che il certificato che è stato inviato appartiene proprio al server con cui voglio comunicare 
- questo controllo lo devo fare a livello applicativo andando a controllare l'URI nel certificato (browser lo fanno da soli), in quanto a livello di protocollo non si fa niente 



### protocollo record
flusso dei dati

prima si autentica e poi si cifra (da confrontare con ipsec)







## IPSEC
garantisce sempre confidenzialità e autenticazione
- stavolta però siamo a livello IP
- stiamo autenticando e cifrando la comunicazione non tra un client e un server (processi) ma tra macchine fisiche
    - payload IP (e non TCP) incapusula i pacchetti dei layer superiori



Sempre una suite protocollare, ma anche qualcosa di più... una vera e propria architettura di sicurezza

qua abbiamo anche una granularità maggiore rispetto a SSL
- possiamo ottenere solo riservatezza o solo autenticazione (SSL mi da sempre tutto)


Quando si usa IPSEC?
- VPN
- connettere LAN in maniera sicura
- routing sicuro



Chi è che realizza una canale ipsec?
- 4 configurazioni con host e/o gateway


