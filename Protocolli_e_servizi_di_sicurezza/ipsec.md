
















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


Modalità di incapsulamento
- trasporto
- tunnel
    - protezione totale del pacchetto
    - anche degli header (non faccio sapere neanche mittente e destinatario)
    - per questo motivo non è una modalità utilizzabile in un canale host-host (non sarebbe instradabile)


...


Slide 22 molto importante:
- a seconda del protocollo della suite IPSEC, e che modaliltà di incapsulamento uso, ottengo proprietà di sicurezza più o meno estesa
- bisogna scegliere la combinazione adeguata ai proprio scopi



...

notiamo che in IPSEC prima cifra e poi HMAC, mentre in SSL il contrario
- ipsec più corretto in fase di ricezione: mi permette di risparmiare la decifrazione se il pacchetto non è integro/autentico


**servizio anti-replica**
i seguenti controllo vengono fatti su un pacchetto **valido** (integro ed autentico)
- se arriva qualcosa a sinistra delle finestra si assume che sia qualcosa che sia gia arrivato e quindi si scarta il pacchetto
- se ricade dentro la finestra si marchia quel numero di sequenza come ricevuto e successivi arrivi vengono scartati
- se arriva qualcosa a destra della finestra si sposta la finestra e si marchia quel numero di sequenza come ricevuto   


### protocollo di negoziazione IKE skippato
molto simile al protocollo di negoziazione per SSL





```Chiedi alla prof cosa significa federazione```
```Prova Pratica con wireshark: parsing di pacchetti ipsec/TLS con eventuale decifratura e verifica della firma```