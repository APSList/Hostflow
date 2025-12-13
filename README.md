
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

## Zagon lokalno
TODO

## Zagon Docker
TODO
