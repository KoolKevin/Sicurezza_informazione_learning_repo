The recent proliferation of digital services has resulted in an unprecedented increase in the number of **digital identities**
- Digital identity refers to the collection of information, attributes, and credentials that uniquely represent an individual, organization, or entity online

Digital identification has gradually shifted from centralized to **decentralized models**

**Centralized Identity**
- Central authority
    - data breaches and single point of failure
- Combination of username and password
    - many vulnerabilities (vedi atatcchi con dizionario a causa di password poco casuali)
    - user prefer convenience over security
- **As many identities as the number of services**
    - poor usability
- Each service provider must securely store, preserve, and safeguard user data
    - companies can be charged


### Federated Identity
In generale, il concetto di federazione è strettamente legato al concetto di **relazione di fiducia**. In un certo senso, la gerarchia di utenti e CA che stanno sotto ad una root-CA forma una federazione

A group of trusted entities share authentication responsibilities (Federazione di Identity Providers), allowing users to **access multiple services with a single identity**

Federated identity refers to linking and using an electronic identity that the user has across several systems.

An application does not necessarily need to obtain in-store credentials to authenticate users
- L'applicazione non deve per forza registrare lei gli utenti e gestirne le credenziali in maniera sicura 

It can use an identity management system that is already storing a user’s electronic identity (IdP) to authenticate the user as long as the application trusts that identity management system (vedi Sign-in with google, ...)

- Individuals can use the same identity across multiple services
- Users are not required to create different accounts
    - identità creata presso un determinato IdP, vale dappertutto nella federazione
- Reduces the number of centralized entities involved
    - non bisogna più creare un account per ogni servizio


**Chi fa parte della federazione?**
Una federazione è formata da più IdP e SP (detti anche Relying Party) dentro ad un ecosistema coordinato (es. una federazione accademica come eduGAIN, SPID in Italia, o eIDAS a livello europeo).
- se ti autentichi con il tuo IdP, puoi accedere anche a servizi di un’altra università o di editori scientifici che fanno parte della federazione, **senza che ogni servizio debba fare un accordo diretto con ogni IdP**.
- Qui l’identità è “federata”, perché l’intero sistema si basa su regole e standard comuni condivisi da tutte le parti della federazione.

**drawbacks**
- Relies on trusted relationships between different parties to authenticate and authorize users
- The Identity Provider (IDP) knows when and where a user logs in, leading to privacy concerns
    - il mio IdP (es. google) vede tutti i servizi a cui accedo
    - privacy lesa
- Users and organizations may become dependent on a single IdP (e.g., Google, Microsoft) without an easy way to migrate
    - si tende a creare un ecosistema chiuso di dipendenze
        - io tendo a fare il login con google e pochi altri
        - se google domani smette di darmi questa possibilità, io non riesco più ad accedere al servizio
    - simile a vendor lock-in
        - mi diventa difficile staccarmi e a migrare su un altro identity provider dato che ho già tanti servizi con cui accedo con lo stesso IdP






Oauth, OpenID Connect, and FIDO 2 are key protocols to implement the federated identity model

**perchè?**









## Self-sovereign identity (SSI)
...



decentralized identifiers
- non eliminabili in quanto replicati sui peer del sistema decentralizzato

...


non c'è un entità centralizzata prendo did e chiavi pubbliche dal verifiable data registry (blockchain) con cui posso verificare le VP e VC



selective disclosure

