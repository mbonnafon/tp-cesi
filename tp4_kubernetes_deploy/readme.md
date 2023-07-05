# Prérequis:

## 1. Kubernetes:

- Pour installer Kubernetes, veuillez suivre la documentation spécifique à Docker Desktop, car il intègre une version de Kubernetes. Vous pouvez trouver les instructions d'installation détaillées sur le lien suivant : https://docs.docker.com/desktop/kubernetes/.
- Installation de la ligne de commande kubernetes: https://kubernetes.io/docs/tasks/tools/

## 2. Environnement de développement intégré (IDE):
Installez VS Code ainsi que le plugin Kubernetes: https://code.visualstudio.com/docs/azure/kubernetes

## 3. Installation du dashboard Kubernetes:
Installation du dashboard:
https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/

Création d'un utilisateur:
https://github.com/kubernetes/dashboard/blob/master/docs/user/access-control/creating-sample-user.md 
 
Test du dashboard avec cette URL:
http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/login

# Argo rollout / dashboard

Le rôle de Argo Rollouts est de faciliter le déploiement et la gestion des mises à jour d'applications sur Kubernetes. Il offre un ensemble d'outils et de fonctionnalités pour effectuer des déploiements progressifs, des mises à jour canari, des rollbacks et des tests A/B. Argo Rollouts permet de gérer de manière automatisée et contrôlée les déploiements d'applications, en minimisant les interruptions de service et en assurant une meilleure stabilité et fiabilité des applications déployées.

## Installation:

ArgoRollout: 
https://argoproj.github.io/argo-rollouts/#quick-start

Ne pas oublier l’installation d’un plugin: https://argo-rollouts.readthedocs.io/en/stable/installation/#kubectl-plugin-installation 

Pour afficher le dashboard:
```
kubectl argo rollouts dashboard
```

## Repository pour tester:
https://github.com/argoproj/rollouts-demo/tree/master/examples 


