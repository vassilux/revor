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


Vérifier si l'application est en cours d'execution via console de supervisor : supervisorctl
Si l'application est en court d'exécution arrêter l'application : stop revor
Quitter le console : exit

Crée un lien symbolic ln -s /opt/revor/revor_[version] /opt/revor/current 

> **NOTE:**
>
> - En cas si le répertoire existe /opt/revor/current. Supprimer rm -rf /opt/revor/current

### Configuraiton



### Mise à jour 
Mise à jour est identique à l'installation sans la partie de la configuration.





