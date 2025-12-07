# Hostflow
TODO opis

# Razvoj

# Branching
| Branch       | Namen                                      | Deploy okolje        | Verzije                                                                       |
| ------------ | ------------------------------------------ | -------------------- | ----------------------------------------------------------------------------- |
| **feature/** | Posamezna funkcionalnost ali bugfix        | Lokalno / dev okolje | Ni posebne verzije; verzija se določi ob merge v `dev`                        |
| **dev**      | Integracija vseh feature-jev za testiranje | Dev deploy           | Merge v `dev` lahko poveča minor ali build številko za test                   |
| **main**     | Staging / pred-produkcijska veja           | Staging okolje       | Merge iz `dev` označi staging verzijo; po potrditvi se izdela release verzija |

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

# Github actions

## Okolja
Trenutno je na voljo le Dev okolje, po potrebi se loči na staging.

## Secrets
Secrets se hranijo glede na okolje in se preko github actions prenesejo v yamle za deployment itd.

## CI/CD
- Akcije se izvedejo ob akciji `psuh`. Izvede se build, testi, push na docker registry in deploy s helm charts - specfično samo tisti servisi, kjer je do spremembe prišlo.
- Produkcijska namestitev se proži ročno.

# Struktura projekta
Go mikrostoritve uporabljajo template https://github.com/alexmodrono/gin-restapi-template
TODO

# Zagon
TODO
