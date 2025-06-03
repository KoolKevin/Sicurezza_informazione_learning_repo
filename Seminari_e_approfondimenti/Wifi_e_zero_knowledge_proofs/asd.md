...


### the problem with WPA2-PSK
- nel protocollo vengono trasmesse tutte le informazioni necessaria a calcolare il MAC
- il mac viene trasmesso
- io posso calcolare il MAC relativo ad una password a caso e vedere 
- se combacia con quello trasmesso allora ho indovinato la password 


**storia**: "How to explain zero-knowledge protocols to your children"

2 key takeaway:
- convinco il verifier che io conosco la password senza mai impiegarla
- il verifier non può reimpiegare la stessa prova in quanto per i terzi non c'è alcuna garanzia che prover e verifier non si siano messi d'accordo, e quindi che il prover non conosca effettivamente la password