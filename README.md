
# Hostflow - razvoj

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
