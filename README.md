
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

## Logging (ELK)

V mapi `elk/` je na voljo konfiguracija za logiranje v ELK sklad.

Za lokalno testiranje je na voljo tudi `docker-compose.yml` za zagon ELK sklada

```bash
docker compose up --build
```
Pazi kako konfiguriraš klic na logstash znotraj mikrostoritve:
- Če mikrostoritev teče lokalno, potem je naslov logstasha localhost:5044
- Če mikrostoritev teče znotraj docker omrežja, potem je naslov logstasha logstash:5044

TODO: Dodati deploy in konkretno konfiguracija za dev in produkcijo, ki je ločena od lokalnega testiranja.

## Lokalni zagon
Trenutno se lahko vsaka mikrostoritev zažene lokalno znotraj IDE-ja (GoLand, Visual Studio, ...).
Ali preko Docker datoteke.
TODO: Potrebno dodati docker-compose za zagon vseh mikrostoritev skupaj.

## Helm in deploy
TODO:
