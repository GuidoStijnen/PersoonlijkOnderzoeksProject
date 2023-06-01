# PersoonlijkOnderzoeksProject
POP voor device managment Fonteyn Vakantie parken 

Instructies
Volg de onderstaande stappen om de code overdraagbaar te maken en uit te voeren op een andere omgeving:

Clone de repository: Start met het clonen van de repository naar uw lokale machine waarop u de code wilt uitvoeren.

Installeer afhankelijkheden: De code heeft een aantal externe afhankelijkheden, zoals de MySQL-driver en de YAML-pakketten. U kunt deze afhankelijkheden installeren met behulp van de volgende opdracht:  go mod download
Hiermee worden alle vereiste afhankelijkheden gedownload en lokaal opgeslagen.

Configureer de database: Open het bestand config.yaml en bewerk de volgende configuratiegegevens om verbinding te maken met uw MySQL-databaseserver:
hostname: <hostname>
username: <gebruikersnaam>
password: <wachtwoord>
database: <database>
port: <poort>
Vervang <hostname>, <gebruikersnaam>, <wachtwoord>, <database> en <poort> door de juiste waarden voor uw MySQL-omgeving.

Voer de code uit: Nadat de configuratie is voltooid, kunt u de code uitvoeren met de volgende opdracht: go run main.go
Dit zal de server starten en de code uitvoeren op het opgegeven poortnummer.

Opmerking: Zorg ervoor dat de MySQL-databaseserver actief en toegankelijk is voordat u de code uitvoert.

Toegang tot de applicatie: Open een webbrowser en navigeer naar http://localhost:<poort> waar <poort> overeenkomt met het poortnummer dat is geconfigureerd in de config.yaml-bestand. U zou de applicatie moeten kunnen gebruiken en communiceren met de MySQL-database.

Aanvullende informatie
Databasestructuur: De code maakt gebruik van een enkele tabel genaamd "device" in de MySQL-database. Het volgende is het SQL-schema voor de tabel:CREATE TABLE device (
  id   VARCHAR(50) PRIMARY KEY,
  name VARCHAR(100),
  os   VARCHAR(50),
  ipv4 VARCHAR(15)
);
  
Zorg ervoor dat de databasestructuur overeenkomt met het bovenstaande schema voordat u de code uitvoert.

HTML-templates: De code maakt gebruik van HTML-templates om de webpagina's te renderen. De templatebestanden moeten zich in dezelfde map bevinden als het main.go-bestand. Zorg ervoor dat u de juiste HTML-templates heeft of bewerk de code om uw eigen templates te gebruiken.

Beveiliging: Houd er rekening mee dat de code zoals deze is, mogelijk niet geschikt is voor productiegebruik, omdat het geen uitgebreide beveiligingsmaatregelen bevat. Voordat u de code in een productieomgeving gebruikt, moeten beveiligingsbest practices worden geïmplementeerd, zoals inputvalidatie, autorisatie en beveiligde communicatie.

Codeaanpassingen: Als u wijzigingen wilt aanbrengen in de code om deze aan te passen aan uw specifieke behoeften, kunt u de broncode bewerken met behulp van een teksteditor of een IDE naar keuze. Zorg ervoor dat u vertrouwd bent met de Go-programmeertaal en de gebruikte pakketten voordat u wijzigingen aanbrengt.

Documentatie: Voor meer gedetailleerde documentatie over de gebruikte pakketten en functies, kunt u de officiële documentatie van Go en de betreffende pakketten raadplegen:

Go-documentatie: https://golang.org/doc/
Go MySQL-driver: https://github.com/go-sql-driver/mysql
YAML-pakket (gopkg.in/yaml.v2): https://pkg.go.dev/gopkg.in/yaml.v2
HTML-templatepakket: https://golang.org/pkg/html/template/ 
  
Contact en Hulp
Als u vragen heeft of hulp nodig heeft met betrekking tot deze codebase, neem dan gerust contact met ons op. We helpen u graag verder.

E-mail: Fontysleerling@example.com
Telefoon: +31-123456789
We raden u aan de volgende informatie te verstrekken wanneer u contact met ons opneemt:

Een gedetailleerde beschrijving van uw probleem of vraag.
De stappen die u tot nu toe heeft ondernomen om het probleem op te lossen.
Eventuele foutmeldingen die u heeft ontvangen.
De versie van de Go-programmeertaal en andere relevante pakketten die u gebruikt.
We streven ernaar om binnen 24 uur te reageren op uw vragen en verzoeken om hulp.
