----------
revor 
----------

 
### Installation 
Copier le ficheir revor_[version].tar.gz sur la machine cible dans le repertoire /opt/revor
Exécuter tar xvzf revor_[version].tar.gz

Ce placer dans le repertoire : cd /opt/revor/revor_[version]

> **NOTE:**
>
> - En cas d'une nouvelle installation copier le ficheir revor.supervisor.conf dans le répertoire /etc/supervisor/conf.d
>
> - Copier le fichier /opt/revor/revor_[version]/config.sample.json dans /opt/revor/revor_[version]/config.json
> - Addapter ce fichier en cas de besoin à votre environement cible


Vérifier si l'application est en cours d'execution via console de supervisor : supervisorctl
Si l'application est en court d'exécution arrêter l'application : stop revor
Quitter le console : exit

Crée un lien symbolic ls -s /opt/revor/current /opt/revor/revor_[version]
> **NOTE:**
>
> - En cas si le répertoire existe /opt/revor/current. Supprimer rm -rf /opt/revor/current

### Configuraiton
Copier le fichier config.sample.json en config.json : cp config.sample.json config.json
Addapter ce fichier à la configuraiton de système : le dialplan de iPBX et les besoin du client.


### Mise à jour 
Mise à jour est identique à l'installation sans la partie de la configuration.

Il faut copier le ficheir de la configuration actuel (/opt/revor/current/config.json) dans le repertoire /opt/revor/revor_[version] 




