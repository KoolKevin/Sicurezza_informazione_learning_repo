Alice -> Bob

Tre requisiti di autenticazione:
1. autore del messaggio
2. mittente (che può essere diverso dall'autore) del messaggio
3. host


Innanzitutto possiamo decidere di non fare a livello applicativo
- possiamo delegare a livello alcune cose a livello di rete/trasporto 

l'ip del mittente nell'hmac è arbitrario per come lo abbiamo pensato noi
- **non possiamo scappare da una configurazione ipsec se vogliamo garantire l'autenticità dell'ip**
- dobbiamo ipotizzare quindi che i due host siano configurati per comunicare tramite ipsec
- deleghiamo a livello di rete l'autenticazione dell'ip 



Il resto dei requisiti vanno bene come gli abbiamo pensati noi a livello applicativo



Non finisce qua!
- possiamo delegare anche l'autenticazione del mittente con SSL

La firma va fatta a livello appliativo per forza





**Conclusione**: esempio interessante in quanto mostra come possiamo **stratificare i vari protocolli a livelli diversi sfruttando l'incapsulamento dei pacchetti** per ottenere i nostri obiettivi di sicurezza. 