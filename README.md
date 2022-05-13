[![SonarCloud](https://sonarcloud.io/images/project_badges/sonarcloud-white.svg)](https://sonarcloud.io/summary/new_code?id=tigerwill90_pwdg)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=tigerwill90_pwdg&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=tigerwill90_pwdg)
# pwdg

A minimal high entropy password generator.

**Install**
````bash
go install github.com/tigerwill90/pwdg@latest
````

### Getting started
````bash
pwdg -n 16 -len 15 -col 4

# Generating 16 passwords of length 15 🚀

# Cd1Rt8Lkf&%R$!z (73.86)	pgmG2W%K59gitK^ (88.41)	pY1%Fa97g9B0$oB (75.98)	fZBWiw^CKz60$l* (85.67)
# 0%anL6MdeO47!QK (79.25)	@3Oyj%3Q8MstdVt (94.30)	NL&$Gh3ZfX0axi3 (80.80)	6K!ugjHK3UzBw!1 (81.84)
# G!tN261gVkGt**0 (80.07)	24WK##y%P3wy&^8 (82.08)	QcRTAMtg6Ru*^65 (87.25)	4xiW^iX8e@vj453 (57.03)
# %pTu4jtIX2*VZ6# (89.73)	@^3!Ll%w0Kd1AmX (71.99)	S9htTpq@g04!@69 (63.80)	M0%Czpm%6%8u$ds (82.97)


# Help
pwdg -h
````
