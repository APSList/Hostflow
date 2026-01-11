
# Hostflow - razvoj

Hostflow je SAAS platforma za upravljanje z nepremičninami za kratkoročni najem. Omogoča enostavno, ekonomično in centralizirano upravljanje z rezervacijami, strankami in transakcijami tudi lastnikom brez tehničnega predznanja. 

## Arhitektura sistema

Hostflow je zasnovana kot Cloud-Native mikrostoritvena aplikacija. Sestavljena je iz 6 mikrostoritev, vsaka s svojim naborom funkcionalnosti :
* Booking (REST):  
  * Vnos/Spremembe/Preklic/Pregled rezervacij 
* Payment (REST): 
  * Vnos/Spremembe/Brisanje/Pregled plačil 
  * Ročna potrditev in preklic plačila 
  * Obdelava dogodkov iz platforme za plačila 
* Property (REST, graphQL): 
  * Pregled (filtriranje, sortiranje) enot za oddajanje 
  * Vnos/Urejanje/Brisanje enot za oddajanje 
  * Urejanje slik za enoto 
* Communication (gRPC): 
  * Pošiljanje email sporočil lastnikom in strankam 
* Profile (REST): 
  * Omejevanje dostopa glede na uporabnika in njegovo vlogo, vnos/pregled/preklic uporabnikov organizacije (omogočeno le lastniku (“owner”) organizacije 
* Customer (gRPC): 
  * Pregled strank, vnos/posodabljanje/vodenje podatkov o strankah
 
Podrobnejša dokumentacija vsake mikrostoritve se nahaja v readme datoteki znotraj svojega repozitorija. Prav tako je tam opisan dostop do dokumentacije API vsake storitve.
 
Topologija aplikacija je zasnovana kot na spodnjem diagramu:
<img width="1037" height="771" alt="hostflow drawio (4)" src="https://github.com/user-attachments/assets/4ad273c4-4255-4050-9d96-b02fe6900ae7" />

Vsaka mikrostoritev ima svojo podatkovno bazo na **Supabase** platformi. Poleg tega sistem je odvisen od naslednjih komponent:
* **Stripe API**
  * Plačila so implementirana preko Stripe platforme
* **Supabase Auth**
  * Za avtentikacijo, avtorizacijo in omejevanje dostopa se uporablja JWT žeton izda s strani Supabase Auth
* **MailerService**
  * Za pošiljanje e-poštnih sporočil
* **Kafka message broker**
  * Za asinhrono obdelovanje sporočil
* **Nginx ingress controller**
  * Za izpostavitev aplikacije izven cluster, določitev pravil in zagotavljenja TLS
* **Angular uporabniški vmesnik**

## Namestitev aplikacije
Aplikacija je nameščane na javnem **Azure** oblaka. Uporablja storitev managed Kubernetes servis za Kubernetes gruščo.
 
## Izpostavitev aplikacije
Aplikacija je javno dostopna preko url: [https://hostflow.software/ui](https://hostflow.software/ui )

Za dostop do REST API storitev in uporabniškega vmesnika izven grušča poskrbi **Nginx ingress controller**. Pravila za usmerjanje so naslednja:
* https://hostflow.software/ui => Uporabniški vmesnik 
* https://hostflow.software/booking => Booking storitev 
* https://hostflow.software/profile => Profile storitev 
* https://hostflow.software/payment => Payment storitev 
* https://hostflow.software/property => Property storitev

Mikrostoritve, ki ne uporabljajo REST protokola pa so dostopne samo znotraj grušča.

**Nginx** s  pomočjo cert-manager samodejno zagotavlja Let's Encrypt **TLS** potrdila za varno **HTTPS** povezavo do vseh mikroservisov. 

## Avtentikacija in avtorizacija ##
Avtentikacija in avtorizacija uporabnikov temelji na storitvi **Supabase Auth**.  Celoten proces temelji na industrijskih protokolih OAuth2 (za avtorizacijo) in OpenID Connect (OIDC) (za avtentikacijo). Ob uspešni prijavi sistem izda varno podpisan JWT (JSON Web Token) žeton. Odjemalec (Angular frontend) posreduje žeton zalednim mikrostoritvam v glavi HTTP zahtevka (Authorization: Bearer <token>). V žetonu so tudi vklučeni metapodatki uporabnika o njegovi organizaciji in vlogo, ki so ključni za zagotavljanje večnajemništva.

## Večnajemništvo ##
Aplikacija za zagotavljanje večnajemništva uporabna model deljenje baze z logično izolacije, kjer podatki organizacij ločijo po obveznem polju “organization_id”. Tabele imajo nastavljene Row Level Security pravila, ki na nivoju baze samodejno filtrira podatke glede na identiteto podano v JWT žetonu. 

Uporabniki so ločeni v organizacije (naročnike aplikacije). Vsaka organizacija s strani administratorjev sistema Hostflow prejme uporabnika z vlogo "OWNER", ta uporabnik lahko nato kreira poljubno mnogo uporabnikov z vlogo "MEMBER". Tem uporabnikom je onemogečen pregled nad uporabniki in plačila ter dodajanje novih uporabnikov.

## Create Booking Flow
* Upravitelj oz. Zunanji sistem ustavari novo rezervacijo in stranka izvede plačilo 
* Upravitelj preko uporabniškega vmesnika oz. Zunanji sistem izbere nastanitev in željen termin. 
* Booking preveri ali je enota razpoložljiva za izbran termin. 
* Če je razpoložljiva ustvari rezervacijo in pokliče Payment za inicializacijo sporočila. 
* Kadar Booking prejme odgovor posodobi rezervacijo in jo posavi v status PAYMENT_REQUIRED. 
* Serverless Supabase Edge funkcija booking-status-update se sproži ob posodobitvi statusa rezervaciji in pokliče Communication za inicializacijo pošiljanja e-poštnega spročila za plačilo. 
* Stranka dobi e-poštno sporočilo z url-jem za plačilo in ga plača. 
*Uspešno plačilo: 
  * Payment iz Stripe plačilnega sistema dobi dogodek, da je bilo plačilo uspešno. 
  * Payment preko sporočilne vrste sporoči dogodek o uspešnem plačilu, ki ga Booking prebere in označi rezervacijo kot potrjeno. 
  * Ob posodobitvi status se ponovno sproži Serverless Supabase funkcija, ki pošlje zahtevo na Communication za pošiljanja maila o potrditvi rezervacije 
  * Communication pošlje mail stranki o uspešnem plačilu in potrdilo o potrjeni rezervaciji. 
* Neuspešno plačilo:
  * Payment iz Stripe plačilnega sistema dobi dogodek, da je bilo plačilo neuspešno. 
  * Payment preko sporočilne vrste sporoči dogodek o neuspešnem plačilu, ki ga Booking prebere in označi rezervacijo kot neuspešno. 
  * Ob posodobitvi status se ponovno sproži Serverless Supabase funkcija, ki pošlje zahtevo na Communication za pošiljanja maila o neuspešnem plačilu. 
  * Kliče Communication, ki pošlje mail o neuspešnem plačilu. 

## Branching
| Branch        | Namen                                              | Deploy okolje        | Verzije                                                                 |
|---------------|----------------------------------------------------|----------------------|-------------------------------------------------------------------------|
| **feature/**  | Razvoj posamezne funkcionalnosti ali izboljšave    | Lokalno        | Brez verzije                      |
| **dev**       | Integracija vseh funkcionalnosti za testiranje     | Development   | Hash zadnjega commita                       |
| **main**      | Stabilna, produkcijsko pripravljena veja           | Production | Release please PR merge samodejno poveca verzijo (npr. `v1.2.0`)                 |

## Proces razvoja

### 1. Razvoj nove funkcionalnosti
- Ustvari se `feature/<opis>` vejo iz `main`.
- Razvijalci implementirajo funkcionalnost.
- Po koncu se ustvari **Pull Request (PR)** v `main`.
- PR vključuje **code review** in testiranje.

### 2. Prenos v dev
- Spremembe (commite) se cherry-pick-a v `dev` branch.
- Na `dev` se izvede **CI/CD build** in deploy na dev okolje.
- Preveri se delovanje vseh novih feature-jev skupaj.

### 3. Release (main)
- Ko je funkcionalnosti testirane na `dev`, se PR-ji merge-ajo v `main`.
- Za nov relese se merge-a autorelease PR
- Release Github workflow propagira novo verzijo in jo namesti
- `dev` branch se ročno resetira na `main`

## 1. Okolja
Okolje se ločuje na Dev in Prod. Na Dev se namešča ob vsaki spremembi na `dev` veji, na Prod pa ob releasu.

## 2. Secrets
Secrets se hranijo glede na okolje in se preko github actions prenesejo v yamle za deployment itd.

## Github actions

### 1. CI/CD
Dev CI/CD workflow se sproži ob vsakem `push`-u na vejo `dev`.

### 2. Release please
Release please github action ustvari release PR ob vsakem `merge`-u v `main` vejo. PR doloci novo verzijo in doda povzetek sprememb v opis.

### 3. PR
PR github action preveri veljavnost PR naslova in zazene teste, enake kot pri CI/CD workflow-u. Brez testov naj PR nebi zdruzil v `main` branch.

### 4. Release
Ko se release please PR `merge`-a v main, se ustvari release, ki sprozi github action za gradnjo iz `main` veje in namestitev na `prod` okolje.

## Struktura projekta
Go mikrostoritve uporabljajo template https://github.com/alexmodrono/gin-restapi-template

.NET projekti uporabljajo basic .NET projektno strukturo - razčlenitev na Business, Api in Test logiko

## Centralizirano beleženje dnevnikov
### Dostop do kibane
- URL: `https://hostflow.software/kibana`
- Uporabnik: `elastic`
- Geslo: (Uporabi geslo, ki si ga pridobil v koraku **4**)

### Navodila za namestitev ELK (Elastic Stack) + Fluent Bit

Opisan je celoten postopek namestitve **Elasticsearch**, **Kibana** (preko ECK operatorja) in **Fluent Bit** za zbiranje logov na Kubernetes gruči.

#### 1. Namestitev ECK Operatorja

Najprej namestimo **Elastic Cloud on Kubernetes (ECK)** operator, ki skrbi za upravljanje Elastic Stack resursov.

```bash
# 1. Dodajanje Elastic Helm repozitorija
helm repo add elastic https://helm.elastic.co
helm repo update

# 2. Namestitev operatorja v ločen namespace 'elastic-system'
helm upgrade --install elastic-operator elastic/eck-operator   -n elastic-system   --create-namespace

# 3. Preverjanje
kubectl -n elastic-system get pods
```
---

#### 2. Priprava Namespace-ov

Ustvarimo ločena namespace-a za bazo in za logiranje.

```bash
kubectl create namespace elastic-stack
kubectl create namespace logging
```

---

#### 3. Zagon baze in vmesnika

Uporabimo pripravljene manifest datoteke.

```bash
# Namestitev Elasticsearch
kubectl apply -f elasticsearch.yaml

# Namestitev Kibane
kubectl apply -f kibana.yaml
```
---

#### 4. Pridobivanje gesla in konfiguracija

Ko se Elasticsearch postavi, operator samodejno ustvari uporabnika `elastic`. Pridobiti moramo njegovo geslo in ustvariti `Secret` v `logging` namespace-u, da se bo Fluent Bit lahko povezal.

Pridobi geslo
```powershell
$Secret = kubectl -n elastic-stack get secret tiny-es-elastic-user -o jsonpath="{.data.elastic}"
$ES_PASSWORD = [System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String($Secret))
Write-Host "Geslo je: $ES_PASSWORD"
```
Ustvari `Secret` za Fluent Bit
V spodnjem ukazu se uporabi spremenljivka `$ES_PASSWORD` (če si uporabil zgornje ukaze). Če spremenljivka ni nastavljena, ročno zamenjaj `$ES_PASSWORD` s pravim geslom.

```bash
kubectl -n logging create secret generic es-credentials   --from-literal=ES_HOST=tiny-es-http.elastic-stack.svc   --from-literal=ES_USER=elastic   --from-literal=ES_PASSWORD="$ES_PASSWORD"   --dry-run=client -o yaml | kubectl apply -f -
```

---

#### 5. Zagon Fluent Bit in Ingressa

Sedaj, ko imamo geslo, lahko zaženemo zbiranje logov in odpremo dostop do Kibane.

```bash
# Zagon Fluent Bit (začne pošiljati loge v Elasticsearch)
kubectl apply -f fluent-bit-es.yaml

# Konfiguracija Ingressa (omogoči dostop preko hostflow.software/kibana)
kubectl apply -f kibana-ingress.yaml
```
---
