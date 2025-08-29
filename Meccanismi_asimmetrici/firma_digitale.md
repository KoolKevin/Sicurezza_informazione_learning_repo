## Firma digitale
Il problema della firma digitale ha trovato la sua soluzione ottimale solo con l’avvento degli algoritmi asimmetrici.
- (è possibile firmare anche con algoritmi simmetrici ma non è l'ideale)

Il **grande pregio del meccanismo asimmetrico di firma è** quello di creare un crittogramma la cui **origine può essere verificata con un dato reso di dominio pubblico** dal firmatario:
- chiunque ha di conseguenza la possibilità di giudicare l’autenticità dei documenti prodotti da chiunque altro. 

L’uso del meccanismo richiede due attività, la prima svolta da un ben preciso utente U con la sua chiave privata SU, la seconda svolta da un qualsiasi altro utente X in possesso di una copia autentica di PU:
1. U: **autentica il documento m** calcolando c = S_SU(m) e rendendolo disponibile a chi è interessato;
2. X: **verifica l’autenticità di m** calcolando V_PU(c), cosa che gli fornisce sia m, sia una risposta di tipo
vero/falso.

Per la sicurezza occorrono due proprietà:
- per ogni SU e per ogni m la risposta di V_PU(S_SU(m)) deve essere “vero”.
- per chiunque non abbia SU e per ogni m di sua invenzione, deve essere computazionalmente infattibile trovare un c tale che la risposta di V_PU(c) è “vero".
    - impossibile per un attaccante fabbricare un firma

L’ultima proprietà è essenziale per poter identificare con sicurezza l’origine di un documento; per ottenerla occorre naturalmente impiegare chiavi sufficientemente lunghe (altrimenti attacchi di forza bruta).

La firma asimmetrica di un documento elettronico crea dunque un attendibile e verificabile **collegamento tra l’informazione ivi contenuta e la persona che l’ha fornita**.

La firma digitale **ha validità legale**. Una terza parte, in possesso di una copia autentica di PU e cui è fornita una coppia <m, c>, può così dirimere due differenti dispute.
1. X sostiene che m è stato prodotto da U: se V_PU(c) = <m, SI>, X ha ragione, in caso contrario ha torto;
2. U nega di aver prodotto m: se V_PU(c) = <m, SI>, U ha torto, in caso contrario ha ragione. 

Se la firma è presente, è una prova inconfutabile dell'autenticità del dato 


Infine, ricordiamo quali sono i requisiti che tale strumento deve possedere. La firma digitale:
1. Deve consentire a chiunque di **identificare univocamente il firmatario**
    - dunque, è necessario ricorrere alla crittografia asimmetrica (ameno che non si voglia coinvolgere una terza entità nella crittografia simmetrica);
2. Non deve poter essere imitata da un impostore;
3. Non deve poter essere trasportata da un documento ad un altro;
4. Non deve poter essere ripudiata dall'autore;
5. Deve rendere inalterabile il documento su cui è apposta.





### Algoritmi di firma con recupero e firma con appendice

#### Recupero
Negli algoritmi con recupero il documento m deve avere un valore numerico inferiore a quello del modulo impiegato nei calcoli; se non è così, il documento deve essere preventivamente suddiviso in blocchi. 
 
Il firmatario trasforma dapprima ogni blocco con una funzione di ridondanza R (la scelta di R è critica dal punto di vista della sicurezza) e poi lo sottopone all’algoritmo di firma. Il verificatore controlla se il risultato appartiene al codominio di R: in caso affermativo, ottiene m tramite R-1.  

... insomma, **cifro con la chiave privata**

#### Appendice
Nel caso di messaggi “lunghi” è molto più efficiente l’impiego dei meccanismi con appendice.
- Questo metodo è anche il più usato, consentendo di autenticare messaggi in chiaro (non devo perforza garantire anche la riservatezza se non mi serve).

Il firmatario sottopone il messaggio m ad una funzione hash sicura, firma l’impronta con la sua chiave privata ed affianca a questa etichetta il messaggio in chiaro ed il certificato della sua chiave pubblica.
- ormai sappiamo che ogniqualvolta si utilizza una chiave pubblica, quest'ultima deve essere autenticata mediante un certificato (altrimetti attacchi MIM)

Dall’altra parte del canale il verificatore calcola l’impronta del messaggio ricevuto ed impiega l’algoritmo di verifica (la chiave è quella pubblica del firmatario, che ha estratto dal certificato) per confrontare le due impronte.
- notare che in RSA la verifica corrisponde ad una decifrazione; la risposta si/no mi viene data dal confronto con l'hash del messaggio ricevuto

Si noti che tutti gli algoritmi con recupero possono essere impiegati anche come algoritmi con appendice, avendo l’hash una dimensione sicuramente inferiore a quella del modulo. Il caso tipico è quello dell’algoritmo RSA


#### Considerazioni
La firma con recupero **è inammissibile e non ha validità legale**
- non valgono i requisiti 3 e 5

Da intrusore posso aver intercettato le firme di vari messaggi con i corrispondenti frammenti di testo in chiaro.

A questo punto posso scegliere delle coppie di frammenti (plaintext, firma) e fare un collage in un qualsiasi documento firmato dalla vittima
- la verifica va a buon fine perchè i frammenti sostituiti sono stati firmati con la chiave privata della vittima e il plaintext corrisponde    
- pensa se mi faccio firmare un documento con il mio nome dentro e comincio a sostituire il mio nome dentro ad altri documenti della vittima
    - ancora più facile se la firma è fatta con un cifrario come RSA che è deterministico (malleabile se non sono stati presi accorgimenti)

**NB**: RSA utilizza nativamente firma con recupero, se voglio i requisiti 3 e 5 devo usare qualcos'altro... or do i?
- se utilizzo RSA su u blocco piccolo allora ottengo praticamente uno schema con appendice, di conseguenza posso usare RSA sopra ad un hash 

**NB**: se il messaggio da firmare ha dimensione minore di un blocco, la firma con recupero va anche bene


## Firma con RSA
L’algoritmo RSA è **reversibile**, cioè funziona anche scambiando la posizione delle chiavi. Vediamo cosa succede se si opera in questo modo.
- Sia m < n un testo in chiaro dotato di un certo “formato”, stabilito da un certo standard.
- L’utente U impiega la sua chiave privata SU = {dU, nU} per calcolare `c = E_SU (m) = m^d_U mod n_U`
    - cifratura con la chiave privata (possibile solo perchè RSA è reversibile)
- L’utente X che ha ricevuto c impiega la chiave pubblica PU di U ed ottiene `D_PU(c) = c^e_U mod n_U = m^(d_U*eU) mod nU = m`
    - decifratura con la chiave pubblica (possibile solo perchè RSA è reversibile)

**NB**: non tutti gli algoritmi asimmetrici hanno la proprietà di reversibilità

La prima conclusione è che non deve essere in gioco la riservatezza di m, dato che **PU è di dominio pubblico**.
- tutti possono decifrare

La seconda conclusione è che **così operando si ottiene l’autenticazione dell’origine di m**:
- dal calcolo di D_PU(c), X può ottenere m, verificarne il “formato” e dedurre da questo se c proviene o no da U (a patto naturalmente che SU sia stata tenuta segreta e non sia stata ripudiata (ricorda CRL)).

**RSA a chiavi invertite è dunque un algoritmo di firma con recupero** 

Il metodo è sicuro (chi vuole falsificare la firma di U deve conoscere la chiave segreta d, oppure deve saper fattorizzare n), ma inefficiente (prevedibile, sto cifrando con RSA).

Per ottenere efficienza è però sufficiente **firmare solo l’impronta di m**.
1. U: calcola c = (H(m))^dU mod nU e trasmette m||c;
2. X: estrae m’ dal messaggio ricevuto, calcola H(m’) ed accetta m’ come un messaggio integro e generato da U se e solo se: H(m’) = c^eU mod nU = H(m).

**Nota**: come per la chiave di sessione in un cifrario ibrido, l'hash ha un numero di bit molto minore del modulo n ed è quindi ritenuto insicuro trasformarla
direttamente. Di nuovo si usa del paddding. 




### Firma cieca
La proprietà moltiplicativa di RSA consente di eseguire la cosiddetta firma “a occhi chiusi”, o “con messaggio oscurato”

Consideriamo un messaggio m = m1 × m2. La sua firma da parte di U è:
- `c = m^dU mod nU = ((m1^dU mod nU) × (m2^dU mod nU)) mod nU`

**La firma del prodotto di due messaggi è dunque uguale al prodotto delle due firme**

Tale tecnica risulta particolarmente utile nella situazione in cui un utente X può **richiedere ad un’Autorità T di autenticargli un messaggio m, di cui, però, non vuole rivelare il contenuto**.

Casi di questo tipo sono tutt’altro che infrequenti:
- Nei **sistemi di voto elettronico**, l’elettore si fa autenticare da un’Autorità il suo voto (in forma non leggibile per difenderne la segretezza) e lo deposita poi nell’“urna” in forma leggibile e verificabile. 
    - La firma dell’Autorità garantisce l’anonimità del votante, il suo diritto ad esprimere un voto ed il fatto che voterà una sola volta.
- Nei sistemi di **commercio elettronico**, l’acquirente si fa autenticare da una Autorità il suo denaro elettronico (per un ammontare non leggibile) prima di consegnarlo al venditore in forma leggibile e verificabile.
    - La firma dell’Autorità garantisce al venditore la copertura finanziaria ed al cliente l’anonimato. 

**NB**: siccome la terza parte firma, e quindi **si assume una responsabilità**, solitamente la terza parte impiega anche un **protocollo di identificazione** per capire chi è che mi sta chiedendo di firmare della roba. 

Per oscurare il suo messaggio m, X sceglie a caso un numero r coprimo con nT (svolge il ruolo di _e_) ed invia a T:
- `c1 = m × r^eT mod nT`
    - eT chiave pubblica dell'autorità
    - **prodotto di due messaggi**
    - mi nasconde m all'autorità
    - l'autorità firmando c1 mi firma anche m senza mai conoscerlo

T firma c1 e restituisce a X il testo cifrato
- `c2 = m^dT × r^(eT*dT) mod nT = m^dT × r mod nT`
    - **r viene decifrato dalla firma**
    - m viene firmato
    - X adesso ha il prodotto delle due firme, gli basta annullare r

X moltiplica c2 per l’inverso moltiplicativo di r (mod nT) ed ottiene m autenticato da T:
- `c3 = (m^dT × r × r -1) mod nT = m^dT mod nT = m firmato con la chiave privata di T` 


Ritorniamo ai due esempi.

**La firma cieca serve a separare la fase di autenticazione** (quando l’autorità verifica che tu hai diritto a fare quell’azione) **dalla fase di pubblicazione/uso delt dato** (quando il messaggio diventa leggibile).

La firma dell’autorità dice:
- Questo messaggio è valido
- Proviene da una persona autorizzata
- Non è stato modificato

**Perché serve oscurare il messaggio nella fase di firma?**

- Voto elettronico
    - Se l’autorità vedesse il contenuto del voto quando lo firmi, **saprebbe per chi voti**
        - violazione della segretezza del voto.
    - Con la firma cieca l’autorità non sa cosa ha firmato
    - Tu rimuovi l’offuscamento e ottieni il voto firmato, pronto da inserire nell’urna
    - **Chiunque controlla la firma sa che è un voto valido, ma non sa chi è il proprietario del voto**

**Conclusione**: se firmassi io, tutti saprebbero per chi ho votato/che spese ho fatto. **La firma cieca mi permette di mantenere il mio anonimato**



### Attacchi "matematici" alla firma con RSA

**attacco ridicolo**
Se si impiega uno schema di firma digitale con recupero, chiunque può
- scegliere un numero a caso r
- prendere la chiave pubblica RSA di chiunque altro
- calcolare y = r^e mod n 
- **ottenere r come “firma” inconfutabile del “messaggio” y**.
    - l'operazione di verifica consiste nel fare: r^e mod n -> ma così si ottiene proprio y

Questo attacco non preoccupa però più di tanto, dato che il messaggio y è un numero a caso ed è veramente difficile che abbia un “significato” impegnativo per il proprietario di d. 


**attacco alla firma basato su proprietà moltiplicativa**

Una più pericolosa vulnerabilità discende dalla proprietà moltiplicativa.

Supponiamo che l’intruso **possa richiedere al firmatario X di firmare qualsiasi messaggio di sua scelta, ad esclusione del messaggio m di suo interesse** (questo scenario è detto chosen message attack).

Per ottenere la firma di m è sufficiente che
- l’intruso generi un numero a caso r
- costruisca m1 = m × r
- calcoli m2 = r^-1
- chieda ad X di firmare entrambi i messaggi
- e moltiplichi infine le firme di m1 e di m2 che gli vengono restituite.
    - Ricordiamo: **La firma del prodotto di due messaggi è uguale al prodotto delle due firme**
    - Abbiamo quindi che S(m1) * S(m2) = S(m1 * m2) = S(m * r * r^-1) = **S(m)**
    - **X ha firmato inconsapevolmente un messaggio arbitrario dell'attacante**

Se la stessa chiave è usata sia:
- per firma cieca (dove tu firmi alla cieca, non controlli il contenuto)
- sia per firma “normale” (dove controlli il contenuto prima di firmare)

l’attaccante può costruire uno dei due pezzi con una firma cieca e l’altro con una firma normale, bypassando qualsiasi controllo di “messaggi proibiti”. Di conseguenza: `la chiave usata per apporre firme ad occhi chiusi non deve essere impiegata in alcun altro contesto`.

**attacco alla riservatezza facendo firmare un cifrato**
Il precedente attacco può essere utilizzato anche per rompere la riservatezza di un testo cifrato con RSA intercettato da un intruso e destinato ad un utente X.

L’intruso oscura c, se lo fa firmare da X (il che corrisponde alla decifrazione con la chiave privata), elimina il numero a caso ed ottiene il testo in chiaro che non doveva poter conoscere.

Vale dunque la seguente regola: `chi impiega RSA per decifrare e per firmare deve utilizzare due differenti coppie di chiavi`
- in questa maniera, firmare un cifrato non lo decifra dato che le coppie di chiavi sono diverse