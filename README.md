
# Hostflow - razvoj

## Branching
| Branch        | Namen                                              | Deploy okolje        | Verzije                                                                 |
|---------------|----------------------------------------------------|----------------------|-------------------------------------------------------------------------|
| **feature/**  | Razvoj posamezne funkcionalnosti ali izboljšave    | Lokalno / dev        | Brez verzije; verzija se določi ob merge v `dev`                        |
| **bug/**      | Odpravljanje napak                                 | Lokalno / dev        | Brez verzije; verzija se določi ob merge v `dev`                        |
| **dev**       | Integracija vseh funkcionalnosti za testiranje     | Development okolje   | Ob merge lahko poveča *minor* ali *build* verzijo                       |
| **main**      | Stabilna, produkcijsko pripravljena veja           | Staging / Production | Merge iz `dev` označi *release* verzijo (npr. `v1.2.0`)                 |

## Proces razvoja

### 1. Razvoj nove funkcionalnosti
- Ustvari se `feature/<opis>` vejo iz `dev`.
- Razvijalci implementirajo funkcionalnost.
- Po koncu se ustvari **Pull Request (PR)** v `dev`.
- PR vključuje **code review** in testiranje.

### 2. Prenos v dev
- Merge feature veje v `dev`.
- Na `dev` se izvede **CI/CD build** in deploy na dev okolje.
- Preveri se delovanje vseh novih feature-jev skupaj.

### 3. Priprava staging (main)
- Ko so vse funkcionalnosti testirane na `dev`, se `dev` merge-a v `main`.
- Na `main` se izvede **staging deploy**.
- Staging verzija je pripravljena za testiranje pred produkcijo.
- 
### 4. Produkcija
- Namestitev na produkcijo se potem zaradi občutljivosti zaganja ročno.

## Github actions

### 1. Okolja
Trenutno je na voljo le Dev okolje, po potrebi se loči na staging.

### 2. Secrets
Secrets se hranijo glede na okolje in se preko github actions prenesejo v yamle za deployment itd.

### 3. CI/CD
- Akcije se izvedejo ob akciji `psuh`. Izvede se build, testi, push na docker registry in deploy s helm charts - specfično samo tisti servisi, kjer je do spremembe prišlo.
- Produkcijska namestitev se proži ročno.

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
